package validator

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/optsutils"
	"github.com/ethereum-optimism/optimism/op-service/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	chal "github.com/kroma-network/kroma/kroma-validator/challenge"
	"github.com/kroma-network/kroma/kroma-validator/metrics"
)

var deletedOutputRoot = [32]byte{}

type Challenger struct {
	log    log.Logger
	cfg    Config
	ctx    context.Context
	cancel context.CancelFunc
	metr   metrics.Metricer

	l2OOContract      *bindings.L2OutputOracle
	l2OOABI           *abi.ABI
	colosseumContract *bindings.Colosseum
	colosseumABI      *abi.ABI
	valPoolContract   *bindings.ValidatorPoolCaller
	valMgrContract    *bindings.ValidatorManagerCaller
	assetMgrContract  *bindings.AssetManagerCaller

	submissionInterval        *big.Int
	finalizationPeriodSeconds *big.Int
	l2BlockTime               *big.Int
	checkpoint                *big.Int
	requiredBondAmountV1      *big.Int
	requiredBondAmountV2      *big.Int
	valPoolTerminationIndex   *big.Int

	l2OutputSubmittedSub ethereum.Subscription
	challengeCreatedSub  ethereum.Subscription

	l2OutputSubmittedEventChan chan *bindings.L2OutputOracleOutputSubmitted
	challengeCreatedEventChan  chan *bindings.ColosseumChallengeCreated

	wg sync.WaitGroup
}

func NewChallenger(cfg Config, l log.Logger, m metrics.Metricer) (*Challenger, error) {
	colosseumContract, err := bindings.NewColosseum(cfg.ColosseumAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	colosseumABI, err := bindings.ColosseumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	l2OOContract, err := bindings.NewL2OutputOracle(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	l2OOABI, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	valPoolContract, err := bindings.NewValidatorPoolCaller(cfg.ValidatorPoolAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	valMgrContract, err := bindings.NewValidatorManagerCaller(cfg.ValidatorManagerAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	assetMgrContract, err := bindings.NewAssetManagerCaller(cfg.AssetManagerAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	return &Challenger{
		log:  l.New("service", "challenge"),
		cfg:  cfg,
		metr: m,

		l2OOContract:      l2OOContract,
		l2OOABI:           l2OOABI,
		colosseumContract: colosseumContract,
		colosseumABI:      colosseumABI,
		valPoolContract:   valPoolContract,
		valMgrContract:    valMgrContract,
		assetMgrContract:  assetMgrContract,
	}, nil
}

func (c *Challenger) InitConfig(ctx context.Context) error {
	contractWatcher := watcher.NewContractWatcher(ctx, c.cfg.L1Client, c.log)

	err := contractWatcher.WatchUpgraded(c.cfg.L2OutputOracleAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
		defer cCancel()
		submissionInterval, err := c.l2OOContract.SUBMISSIONINTERVAL(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get submission interval: %w", err)
		}
		c.submissionInterval = submissionInterval

		cCtx, cCancel = context.WithTimeout(ctx, c.cfg.NetworkTimeout)
		defer cCancel()
		l2BlockTime, err := c.l2OOContract.L2BLOCKTIME(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get l2 block time: %w", err)
		}
		c.l2BlockTime = l2BlockTime

		cCtx, cCancel = context.WithTimeout(ctx, c.cfg.NetworkTimeout)
		defer cCancel()
		finalizationPeriodSeconds, err := c.l2OOContract.FINALIZATIONPERIODSECONDS(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get finalization period seconds: %w", err)
		}
		c.finalizationPeriodSeconds = finalizationPeriodSeconds

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate l2OO config: %w", err)
	}

	err = contractWatcher.WatchUpgraded(c.cfg.ValidatorPoolAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
		defer cCancel()
		requiredBondAmountV1, err := c.valPoolContract.REQUIREDBONDAMOUNT(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get submission interval: %w", err)
		}
		c.requiredBondAmountV1 = requiredBondAmountV1

		cCtx, cCancel = context.WithTimeout(ctx, c.cfg.NetworkTimeout)
		defer cCancel()
		valPoolTerminationIndex, err := c.valPoolContract.TERMINATEOUTPUTINDEX(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get valPool termination index: %w", err)
		}
		c.valPoolTerminationIndex = valPoolTerminationIndex

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate valPool config: %w", err)
	}

	err = contractWatcher.WatchUpgraded(c.cfg.AssetManagerAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
		defer cCancel()
		requiredBondAmountV2, err := c.assetMgrContract.BONDAMOUNT(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get required bond amount of assetMgr: %w", err)
		}
		c.requiredBondAmountV2 = requiredBondAmountV2

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate assetMgr config: %w", err)
	}

	return nil
}

// initSub initialize subscriptions
func (c *Challenger) initSub() {
	opts := optsutils.NewSimpleWatchOpts(c.ctx)

	if c.cfg.ChallengerEnabled {
		c.l2OutputSubmittedEventChan = make(chan *bindings.L2OutputOracleOutputSubmitted)
		c.l2OutputSubmittedSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				c.log.Warn("resubscribing after failed OutputSubmitted event", "err", err)
			}
			return c.l2OOContract.WatchOutputSubmitted(opts, c.l2OutputSubmittedEventChan, nil, nil, nil)
		})
	}

	c.challengeCreatedEventChan = make(chan *bindings.ColosseumChallengeCreated)
	c.challengeCreatedSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			c.log.Warn("resubscribing after failed ChallengeCreated event", "err", err)
		}
		return c.colosseumContract.WatchChallengeCreated(opts, c.challengeCreatedEventChan, nil, nil, nil)
	})
}

func (c *Challenger) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)

	if err := c.InitConfig(c.ctx); err != nil {
		return err
	}
	c.initSub()

	c.wg.Add(1)
	go c.loop()

	return nil
}

func (c *Challenger) Stop() error {
	if c.l2OutputSubmittedSub != nil {
		c.l2OutputSubmittedSub.Unsubscribe()
	}
	c.challengeCreatedSub.Unsubscribe()

	c.cancel()
	c.wg.Wait()

	if c.l2OutputSubmittedEventChan != nil {
		close(c.l2OutputSubmittedEventChan)
	}
	close(c.challengeCreatedEventChan)

	return nil
}

func (c *Challenger) loop() {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		select {
		case <-c.ctx.Done():
			return
		default:
			if c.cfg.ChallengerEnabled {
				if err := c.updateCheckpoint(); err != nil {
					c.log.Error(err.Error())
					continue
				}
			}

			if err := c.scanPrevOutputs(); err != nil {
				c.log.Error("failed to scan previous outputs", "err", err)
				continue
			}

			// if challenge mode on, subscribe L2 output submission events
			if c.cfg.ChallengerEnabled {
				c.wg.Add(1)
				go c.subscribeL2OutputSubmitted()
			}

			// subscribe challenge creation events
			c.wg.Add(1)
			go c.subscribeChallengeCreated()

			return
		}
	}
}

// updateCheckpoint updates checkpoint which is the last checked output index, so the next output handling starts after
// this point. If checkpoint is behind the latest output index, handle the previous outputs from the checkpoint.
func (c *Challenger) updateCheckpoint() error {
	cCtx, cCancel := context.WithTimeout(c.ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	nextOutputIndex, err := c.l2OOContract.NextOutputIndex(optsutils.NewSimpleCallOpts(cCtx))
	if err != nil {
		return fmt.Errorf("failed to get the latest output index: %w", err)
	}
	if nextOutputIndex.Cmp(common.Big0) == 0 {
		// if no outputs have been submitted, set checkpoint to 1 because genesis output cannot be challenged
		c.checkpoint = common.Big1
	} else {
		// set checkpoint to latestOutputIndex (nextOutputIndex - 1)
		c.checkpoint = new(big.Int).Sub(nextOutputIndex, common.Big1)
	}
	c.metr.RecordChallengeCheckpoint(c.checkpoint)
	return nil
}

// scanPrevOutputs scans all the previous outputs before current L1 block within the finalization window.
// If there are invalid outputs, create challenge.
// If there are challenges in progress, keep handling them.
func (c *Challenger) scanPrevOutputs() error {
	status, err := c.cfg.RollupClient.SyncStatus(c.ctx)
	if err != nil {
		return fmt.Errorf("failed to get sync status: %w", err)
	}

	toBlock := new(big.Int).SetUint64(status.HeadL1.Number)
	// TODO(0xHansLee): add L1BlockTime to rollup config and change to use it
	finalizationStartL1Block := new(big.Int).Sub(toBlock, new(big.Int).Div(c.finalizationPeriodSeconds, big.NewInt(12)))
	// The fromBlock is the maximum value of either genesis block(1) or the first block of the finalization window
	fromBlock := math.BigMax(common.Big1, finalizationStartL1Block)

	outputSubmittedEvent := c.l2OOABI.Events[KeyEventOutputSubmitted]
	challengeCreatedEvent := c.colosseumABI.Events[KeyEventChallengeCreated]

	addresses := []common.Address{c.cfg.ColosseumAddr}
	topics := []common.Hash{challengeCreatedEvent.ID}

	// scan OutputSubmittedEvents only when challenger mode is on
	if c.cfg.ChallengerEnabled {
		addresses = append(addresses, c.cfg.L2OutputOracleAddr)
		topics = append(topics, outputSubmittedEvent.ID)
	}

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: addresses,
		Topics:    [][]common.Hash{topics},
	}

	logs, err := c.cfg.L1Client.FilterLogs(c.ctx, query)
	if err != nil {
		return fmt.Errorf("failed to get event logs related to outputs: %w", err)
	}

	for _, vLog := range logs {
		switch vLog.Address {
		// for OutputSubmitted event
		case c.cfg.L2OutputOracleAddr:
			ev, err := NewOutputSubmittedEvent(vLog)
			if err != nil {
				c.log.Error("failed to parse output submitted event", "err", err)
				continue
			}
			// handle output
			c.wg.Add(1)
			go c.handleOutput(ev.L2OutputIndex)
		// for ChallengeCreated event
		case c.cfg.ColosseumAddr:
			ev, err := NewChallengeCreatedEvent(vLog)
			if err != nil {
				c.log.Error("failed to parse challenge created event", "err", err)
				continue
			}
			if ev.OutputIndex.Sign() == 1 && c.isRelatedChallenge(ev.Asserter, ev.Challenger) {
				c.wg.Add(1)
				go c.handleChallenge(ev.OutputIndex, ev.Asserter, ev.Challenger)
			}
		default:
			c.log.Warn("unknown event log", "logs", vLog)
		}
	}

	return nil
}

// subscribeL2OutputSubmitted subscribes the OutputSubmitted event from L2OutputOracle contract.
// It handles all the outputs between the checkpoint output index and the output index from the watched event.
// If the L2 output root is invalid, create challenge.
// This function should be called only when challenger mode is on.
func (c *Challenger) subscribeL2OutputSubmitted() {
	defer c.wg.Done()

	for {
		select {
		case ev := <-c.l2OutputSubmittedEventChan:
			c.log.Info("watched output submitted event", "l2BlockNumber", ev.L2BlockNumber, "outputRoot", ev.OutputRoot, "outputIndex", ev.L2OutputIndex)
			// if the emitted output index is less than or equal to the checkpoint, it is considered reorg occurred.
			if ev.L2OutputIndex.Cmp(c.checkpoint) <= 0 {
				c.wg.Add(1)
				go c.handleOutput(new(big.Int).Set(ev.L2OutputIndex))
			} else {
				// validate all outputs between the checkpoint and the current outputIndex
				for i := new(big.Int).Add(c.checkpoint, common.Big1); i.Cmp(ev.L2OutputIndex) != 1; i.Add(i, common.Big1) {
					c.wg.Add(1)
					go c.handleOutput(new(big.Int).Set(i))
				}
			}
			c.checkpoint = ev.L2OutputIndex
			c.metr.RecordChallengeCheckpoint(c.checkpoint)
		case <-c.ctx.Done():
			return
		}
	}
}

// subscribeChallengeCreated subscribes the ChallengeCreated event from Colosseum contract and handle challenge.
func (c *Challenger) subscribeChallengeCreated() {
	defer c.wg.Done()

	for {
		select {
		case ev := <-c.challengeCreatedEventChan:
			c.log.Info("watched challenge created event", "outputIndex", ev.OutputIndex, "challenger", ev.Challenger)
			// when challenge created, handle it
			if ev.OutputIndex.Sign() == 1 && c.isRelatedChallenge(ev.Asserter, ev.Challenger) {
				c.wg.Add(1)
				go c.handleChallenge(ev.OutputIndex, ev.Asserter, ev.Challenger)
			}
		case <-c.ctx.Done():
			return
		}
	}
}

// handleOutput handles output when output submitted, creates challenge if the output is invalid.
// This function should be called only when challenger mode is on.
func (c *Challenger) handleOutput(outputIndex *big.Int) {
	c.log.Info("handling output to detect invalid output", "outputIndex", outputIndex)
	defer c.wg.Done()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		select {
		case <-c.ctx.Done():
			return
		default:
			// check if challenge creation period is not past
			isInCreationPeriod, err := c.IsInChallengeCreationPeriod(c.ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to get if challenge creation period is not past", "err", err, "outputIndex", outputIndex)
				continue
			}
			// if challenge creation period is past, terminate handling
			if !isInCreationPeriod {
				c.log.Info("challenge creation period is already past", "outputIndex", outputIndex)
				return
			}

			outputs, err := c.OutputsAtIndex(c.ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to get outputs when handling output", "err", err, "outputIndex", outputIndex)
				continue
			}

			outputRange := c.ValidateOutput(outputIndex, outputs)
			// if output is valid, terminate handling
			if outputRange == nil {
				c.log.Info("output is validated", "outputIndex", outputIndex)
				return
			}

			// if challenge from another challenger is already proven and output is deleted, terminate handling
			if IsOutputDeleted(outputs.RemoteOutput.OutputRoot) {
				c.log.Info("found invalid output, but output is already deleted", "outputIndex", outputIndex)
				return
			}

			// check the status of my challenge
			status, err := c.GetChallengeStatus(c.ctx, outputIndex, c.cfg.TxManager.From())
			if err != nil {
				c.log.Error("unable to get challenge status", "err", err, "outputIndex", outputIndex)
				continue
			}
			// if challenge is already in progress, terminate handing
			if status != chal.StatusNone && status != chal.StatusChallengerTimeout {
				c.log.Info("found invalid output, but challenge is already in progress", "outputIndex", outputIndex)
				return
			}

			canCreateChallenge, err := c.CanCreateChallenge(c.ctx, outputIndex)
			if err != nil {
				c.log.Error(err.Error())
				continue
			}
			if !canCreateChallenge {
				continue
			}

			// if all of the above conditions are satisfied, create a new challenge
			tx, err := c.CreateChallenge(c.ctx, outputRange)
			if err != nil {
				c.log.Error("failed to create createChallenge tx", "err", err, "outputIndex", outputIndex)
				continue
			}

			if err := c.submitChallengeTx(tx); err != nil {
				c.log.Error("failed to submit create challenge tx", "err", err, "outputIndex", outputIndex)
				continue
			}

			c.log.Info("submit create challenge tx", "outputIndex", outputIndex)
			return
		}
	}
}

// handleChallenge handles related challenge according to its status and role when challenge created.
func (c *Challenger) handleChallenge(outputIndex *big.Int, asserter common.Address, challenger common.Address) {
	c.log.Info("handling related challenge", "outputIndex", outputIndex, "asserter", asserter, "challenger", challenger)
	defer c.wg.Done()

	isAsserter := asserter == c.cfg.TxManager.From()
	isChallenger := challenger == c.cfg.TxManager.From()

	ticker := time.NewTicker(c.cfg.ChallengePollInterval)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		select {
		case <-c.ctx.Done():
			return
		default:
			// check the status of challenge
			status, err := c.GetChallengeStatus(c.ctx, outputIndex, challenger)
			if err != nil {
				c.log.Error("unable to get challenge status", "err", err, "outputIndex", outputIndex, "challenger", challenger)
				continue
			}
			// if challenge is not in progress, terminate handling
			if status == chal.StatusNone {
				c.log.Info("challenge is not in progress", "outputIndex", outputIndex, "challenger", challenger)
				return
			}

			output, err := c.GetL2Output(c.ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to get output when handling challenge", "err", err, "outputIndex", outputIndex)
				continue
			}
			isOutputDeleted := IsOutputDeleted(output.OutputRoot)

			isOutputFinalized, err := c.IsOutputFinalized(c.ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to get if output is finalized when handling challenge", "err", err, "outputIndex", outputIndex)
				continue
			}

			// if asserter
			if isAsserter {
				// if output is already deleted, asserter has no incentives to handle challenge any further
				if isOutputDeleted {
					c.log.Info("do nothing because output is already deleted", "outputIndex", outputIndex, "challenger", challenger)
					return
				}
				// if output is already finalized and not `ChallengerTimeout` status, terminate handling
				if isOutputFinalized && status != chal.StatusChallengerTimeout {
					c.log.Info("output is already finalized when handling challenge", "outputIndex", outputIndex, "challenger", challenger)
					return
				}
				switch status {
				case chal.StatusAsserterTurn:
					tx, err := c.Bisect(c.ctx, outputIndex, challenger)
					if err != nil {
						c.log.Error("failed to create bisect tx", "err", err, "outputIndex", outputIndex, "challenger", challenger)
						continue
					}
					if err := c.submitChallengeTx(tx); err != nil {
						c.log.Error("failed to submit bisect tx", "err", err, "outputIndex", outputIndex, "challenger", challenger)
						continue
					}
				case chal.StatusChallengerTimeout:
					// call challenger timeout to take challenger's bond away
					tx, err := c.ChallengerTimeout(c.ctx, outputIndex, challenger)
					if err != nil {
						c.log.Error("failed to create challenger timeout tx", "err", err, "outputIndex", outputIndex, "challenger", challenger)
						continue
					}
					if err := c.submitChallengeTx(tx); err != nil {
						c.log.Error("failed to submit challenger timeout tx", "err", err, "outputIndex", outputIndex, "challenger", challenger)
						continue
					}
				default:
					continue
				}
			}

			// if challenger
			if isChallenger && c.cfg.ChallengerEnabled {
				// if output has been already deleted, cancel challenge to refund challenger's bond
				if isOutputDeleted && status != chal.StatusChallengerTimeout {
					tx, err := c.CancelChallenge(c.ctx, outputIndex)
					if err != nil {
						c.log.Error("failed to create cancel challenge tx", "err", err, "outputIndex", outputIndex)
						continue
					}
					if err := c.submitChallengeTx(tx); err != nil {
						c.log.Error("failed to submit cancel challenge tx", "err", err, "outputIndex", outputIndex)
						continue
					}
				}

				// if output is already finalized, terminate handling
				if isOutputFinalized {
					c.log.Info("output is already finalized when handling challenge", "outputIndex", outputIndex)
					return
				}

				// Challenger doesn't need to check if output is already deleted or not. Because when trying to bisect or prove fault with deleted output index,
				// the contract automatically cancels the challenge.
				switch status {
				case chal.StatusChallengerTurn:
					tx, err := c.Bisect(c.ctx, outputIndex, challenger)
					if err != nil {
						c.log.Error("failed to create bisect tx", "err", err, "outputIndex", outputIndex)
						continue
					}
					if err := c.submitChallengeTx(tx); err != nil {
						c.log.Error("failed to submit bisect tx", "err", err, "outputIndex", outputIndex)
						continue
					}
				case chal.StatusAsserterTimeout, chal.StatusReadyToProve:
					skipSelectFaultPosition := status == chal.StatusAsserterTimeout
					tx, retry, err := c.ProveFault(c.ctx, outputIndex, challenger, skipSelectFaultPosition)
					if err != nil {
						c.log.Error("failed to create prove fault tx", "err", err, "outputIndex", outputIndex)
						continue
					}
					if retry {
						continue
					}
					if err := c.submitChallengeTx(tx); err != nil {
						c.log.Error("failed to submit prove fault tx", "err", err, "outputIndex", outputIndex)
						continue
					}
				default:
					continue
				}
			}
		}
	}
}

func (c *Challenger) submitChallengeTx(tx *types.Transaction) error {
	return c.cfg.TxManager.SendTransaction(c.ctx, tx).Err
}

// CanCreateChallenge checks if challenger is in the status that can create challenge.
func (c *Challenger) CanCreateChallenge(ctx context.Context, outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	from := c.cfg.TxManager.From()

	var balance, requiredBondAmount *big.Int
	if c.IsValPoolTerminated(outputIndex) {
		if isInJail, err := c.isInJail(ctx); err != nil {
			return false, err
		} else if isInJail {
			c.log.Warn("validator is in jail")
			return false, nil
		}

		validatorStatus, err := c.valMgrContract.GetStatus(optsutils.NewSimpleCallOpts(cCtx), from)
		if err != nil {
			return false, fmt.Errorf("failed to fetch the validator status: %w", err)
		}
		c.metr.RecordValidatorStatus(validatorStatus)

		if validatorStatus != StatusActive {
			c.log.Warn("validator is not in the status that can create a challenge", "status", validatorStatus)
			return false, nil
		}

		balance, err = c.assetMgrContract.TotalValidatorKroNotBonded(optsutils.NewSimpleCallOpts(cCtx), from)
		if err != nil {
			return false, fmt.Errorf("failed to fetch balance: %w", err)
		}
		requiredBondAmount = c.requiredBondAmountV2
	} else {
		var err error
		balance, err = c.valPoolContract.BalanceOf(optsutils.NewSimpleCallOpts(cCtx), from)
		if err != nil {
			return false, fmt.Errorf("failed to fetch deposit amount: %w", err)
		}
		requiredBondAmount = c.requiredBondAmountV1
	}

	c.metr.RecordUnbondedDepositAmount(balance)

	// Check if the unbonded deposit amount is less than the required bond amount
	if balance.Cmp(requiredBondAmount) == -1 {
		c.log.Warn(
			"unbonded deposit is less than bond attempt amount",
			"requiredBondAmount", requiredBondAmount,
			"unbonded_deposit", balance,
		)
		return false, nil
	}

	c.log.Info("unbonded deposit amount and required bond amount",
		"unbonded_deposit", balance, "required_bond", requiredBondAmount)

	return true, nil
}

func (c *Challenger) IsValPoolTerminated(outputIndex *big.Int) bool {
	return c.valPoolTerminationIndex.Cmp(outputIndex) < 0
}

func (c *Challenger) isInJail(ctx context.Context) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	from := c.cfg.TxManager.From()
	isInJail, err := c.valMgrContract.InJail(optsutils.NewSimpleCallOpts(cCtx), from)
	if err != nil {
		return false, fmt.Errorf("failed to fetch the jail status: %w", err)
	}

	return isInJail, nil
}

func (c *Challenger) IsInChallengeCreationPeriod(ctx context.Context, outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.colosseumContract.IsInCreationPeriod(optsutils.NewSimpleCallOpts(cCtx), outputIndex)
}

func (c *Challenger) IsOutputFinalized(ctx context.Context, outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.l2OOContract.IsFinalized(optsutils.NewSimpleCallOpts(cCtx), outputIndex)
}

func (c *Challenger) GetChallenge(ctx context.Context, outputIndex *big.Int, challenger common.Address) (bindings.TypesChallenge, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.colosseumContract.GetChallenge(optsutils.NewSimpleCallOpts(cCtx), outputIndex, challenger)
}

func (c *Challenger) OutputAtBlockSafe(ctx context.Context, blockNumber uint64) (*eth.OutputResponse, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.cfg.RollupClient.OutputAtBlock(cCtx, blockNumber)
}

func (c *Challenger) OutputWithProofAtBlockSafe(ctx context.Context, blockNumber uint64) (*eth.OutputWithProofResponse, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.cfg.RollupClient.OutputWithProofAtBlock(cCtx, blockNumber)
}

func (c *Challenger) PublicInputProof(ctx context.Context, blockNumber uint64) (bindings.TypesPublicInputProof, error) {
	srcOutput, err := c.OutputWithProofAtBlockSafe(ctx, blockNumber)
	if err != nil {
		return bindings.TypesPublicInputProof{}, err
	}

	dstOutput, err := c.OutputWithProofAtBlockSafe(ctx, blockNumber+1)
	if err != nil {
		return bindings.TypesPublicInputProof{}, err
	}

	publicInput, err := srcOutput.ToPublicInput()
	if err != nil {
		return bindings.TypesPublicInputProof{}, err
	}

	rlp, err := srcOutput.ToBlockHeaderRLP()
	if err != nil {
		return bindings.TypesPublicInputProof{}, err
	}

	p := dstOutput.PublicInputProof

	var balance [32]byte
	copy(balance[:], common.BigToHash(p.L2ToL1MessagePasserBalance).Bytes())

	// TODO(chokobole): Do we need to deep copy of this?
	merkleProof := make([][]byte, len(p.MerkleProof))
	for i, b := range p.MerkleProof {
		merkleProof[i] = b
	}

	return bindings.TypesPublicInputProof{
		SrcOutputRootProof:          srcOutput.ToOutputRootProof(),
		DstOutputRootProof:          dstOutput.ToOutputRootProof(),
		PublicInput:                 publicInput,
		Rlps:                        rlp,
		L2ToL1MessagePasserBalance:  balance,
		L2ToL1MessagePasserCodeHash: p.L2ToL1MessagePasserCodeHash,
		MerkleProof:                 merkleProof,
	}, nil
}

type Outputs struct {
	RemoteOutput *bindings.TypesCheckpointOutput
	LocalOutput  *eth.OutputResponse
}

func (c *Challenger) OutputsAtIndex(ctx context.Context, outputIndex *big.Int) (*Outputs, error) {
	remoteOutput, err := c.GetL2Output(ctx, outputIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote output: %w", err)
	}

	localOutput, err := c.OutputAtBlockSafe(ctx, remoteOutput.L2BlockNumber.Uint64())
	if err != nil {
		return nil, fmt.Errorf("failed to get local output: %w", err)
	}

	return &Outputs{&remoteOutput, localOutput}, nil
}

func (c *Challenger) GetL2Output(ctx context.Context, outputIndex *big.Int) (bindings.TypesCheckpointOutput, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.l2OOContract.GetL2Output(optsutils.NewSimpleCallOpts(cCtx), outputIndex)
}

type OutputRange struct {
	OutputIndex *big.Int
	StartBlock  uint64
	EndBlock    uint64
	L1Origin    eth.BlockID
}

// ValidateOutput validates the output for the given outputIndex.
func (c *Challenger) ValidateOutput(outputIndex *big.Int, outputs *Outputs) *OutputRange {
	start := outputs.RemoteOutput.L2BlockNumber.Uint64() - c.submissionInterval.Uint64()
	end := outputs.RemoteOutput.L2BlockNumber.Uint64()

	if !bytes.Equal(outputs.LocalOutput.OutputRoot[:], outputs.RemoteOutput.OutputRoot[:]) {
		c.log.Info(
			"found invalid output",
			"blockNumber", outputs.RemoteOutput.L2BlockNumber,
			"outputIndex", outputIndex,
			"local", outputs.LocalOutput.OutputRoot,
			"invalid", common.BytesToHash(outputs.RemoteOutput.OutputRoot[:]),
		)
		return &OutputRange{
			OutputIndex: outputIndex,
			StartBlock:  start,
			EndBlock:    end,
			L1Origin:    outputs.LocalOutput.BlockRef.L1Origin,
		}
	} else {
		c.log.Info("confirmed that the output is valid",
			"outputIndex", outputIndex,
			"start", start,
			"end", end,
			"outputRoot", common.BytesToHash(outputs.RemoteOutput.OutputRoot[:]),
		)
		return nil
	}
}

func (c *Challenger) isRelatedChallenge(asserter common.Address, challenger common.Address) bool {
	return c.cfg.TxManager.From() == asserter || c.cfg.TxManager.From() == challenger
}

func (c *Challenger) GetChallengeStatus(ctx context.Context, outputIndex *big.Int, challenger common.Address) (uint8, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	return c.colosseumContract.GetStatus(optsutils.NewSimpleCallOpts(cCtx), outputIndex, challenger)
}

func (c *Challenger) BuildSegments(ctx context.Context, turn uint8, segStart, segSize uint64) (*chal.Segments, error) {
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()

	sections, err := c.colosseumContract.SegmentsLengths(optsutils.NewSimpleCallOpts(cCtx), big.NewInt(int64(turn-1)))
	if err != nil {
		return nil, fmt.Errorf("unable to get segments length of turn %d: %w", turn, err)
	}

	segments := chal.NewEmptySegments(segStart, segSize, sections.Uint64())

	for i, blockNumber := range segments.BlockNumbers() {
		output, err := c.OutputAtBlockSafe(ctx, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("unable to get output %d: %w", blockNumber, err)
		}

		segments.SetHashValue(i, output.OutputRoot)
	}

	return segments, nil
}

func (c *Challenger) selectFaultPosition(ctx context.Context, segments *chal.Segments) (*big.Int, error) {
	for i, blockNumber := range segments.BlockNumbers() {
		output, err := c.OutputAtBlockSafe(ctx, blockNumber)
		if err != nil {
			return nil, err
		}

		if !bytes.Equal(segments.Hashes[i][:], output.OutputRoot[:]) {
			return big.NewInt(int64(i) - 1), nil
		}
	}

	return nil, errors.New("failed to select fault position")
}

func (c *Challenger) CreateChallenge(ctx context.Context, outputRange *OutputRange) (*types.Transaction, error) {
	outputIndex := outputRange.OutputIndex
	l1BlockHash := outputRange.L1Origin.Hash
	l1BlockNumber := new(big.Int).SetUint64(outputRange.L1Origin.Number)

	c.log.Info("crafting createChallenge tx",
		"index", outputIndex,
		"start", outputRange.StartBlock,
		"end", outputRange.EndBlock,
		"l1BlockHash", l1BlockHash.TerminalString(),
		"l1BlockNumber", l1BlockNumber,
	)

	segSize := outputRange.EndBlock - outputRange.StartBlock
	segments, err := c.BuildSegments(ctx, 1, outputRange.StartBlock, segSize)
	if err != nil {
		return nil, err
	}

	txOpts := optsutils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.CreateChallenge(txOpts, outputIndex, l1BlockHash, l1BlockNumber, segments.Hashes)
}

func (c *Challenger) Bisect(ctx context.Context, outputIndex *big.Int, challenger common.Address) (*types.Transaction, error) {
	c.log.Info("crafting bisect tx", "outputIndex", outputIndex, "challenger", challenger)

	challenge, err := c.GetChallenge(ctx, outputIndex, challenger)
	if err != nil {
		return nil, err
	}

	prevSegments := chal.NewSegments(challenge.SegStart.Uint64(), challenge.SegSize.Uint64(), challenge.Segments)
	position, err := c.selectFaultPosition(ctx, prevSegments)
	if err != nil {
		return nil, err
	}
	// if the first segment is different between challenger and asserter, return error
	if position.Cmp(common.Big0) == -1 {
		return nil, errors.New("the first segment must be matched when bisecting")
	}

	nextTurn := challenge.Turn + 1
	start, size := prevSegments.NextSegmentsRange(position.Uint64())
	nextSegments, err := c.BuildSegments(ctx, nextTurn, start, size)
	if err != nil {
		return nil, err
	}

	txOpts := optsutils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.Bisect(txOpts, outputIndex, challenger, position, nextSegments.Hashes)
}

func (c *Challenger) ChallengerTimeout(ctx context.Context, outputIndex *big.Int, challenger common.Address) (*types.Transaction, error) {
	c.log.Info("crafting challenger timeout tx", "outputIndex", outputIndex, "challenger", challenger)

	txOpts := optsutils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.ChallengerTimeout(txOpts, outputIndex, challenger)
}

func (c *Challenger) CancelChallenge(ctx context.Context, outputIndex *big.Int) (*types.Transaction, error) {
	c.log.Info("crafting cancel challenge tx", "outputIndex", outputIndex)

	txOpts := optsutils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.CancelChallenge(txOpts, outputIndex)
}

// ProveFault creates proveFault transaction for invalid output root.
// TODO: ProveFault will take long time, so that we may have to handle it carefully.
func (c *Challenger) ProveFault(
	ctx context.Context, outputIndex *big.Int, challenger common.Address, skipSelectFaultPosition bool,
) (*types.Transaction, bool, error) {
	challenge, err := c.GetChallenge(ctx, outputIndex, challenger)
	if err != nil {
		return nil, true, err
	}

	// when asserter timeout, skip finding fault position since the same segments have been stored in colosseum
	position := common.Big0
	blockNumber := challenge.SegStart
	if !skipSelectFaultPosition {
		prevSegments := chal.NewSegments(blockNumber.Uint64(), challenge.SegSize.Uint64(), challenge.Segments)
		position, err = c.selectFaultPosition(ctx, prevSegments)
		if err != nil {
			return nil, true, fmt.Errorf("failed to select fault position(outputIndex: %s, challengerAddress: %s): %w", outputIndex.String(), challenger.String(), err)
		}

		blockNumber = new(big.Int).Add(blockNumber, position)
	}

	targetBlockNumber := new(big.Int).Add(blockNumber, common.Big1)
	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	header, err := c.cfg.L2Client.HeaderByNumber(cCtx, targetBlockNumber)
	defer cCancel()
	if err != nil {
		return nil, true, err
	}

	// if the target block time is after Kroma MPT time, generate zkVM proof otherwise zkEVM proof
	if c.cfg.RollupConfig.IsKromaMPT(header.Time) {
		return c.proveFaultWithZkVm(ctx, outputIndex, challenge, targetBlockNumber, header.Hash().Hex(), position)
	} else {
		return c.proveFaultWithZkEvm(ctx, outputIndex, challenger, targetBlockNumber, position)
	}
}

// proveFaultWithZkVm fetches zkVM witness data and proof to create proveFaultWithZkVm transaction.
// It sends requests to external RPC and immediately returns, then continuously checks the status of request and moves to next step.
func (c *Challenger) proveFaultWithZkVm(
	ctx context.Context, outputIndex *big.Int, challenge bindings.TypesChallenge, targetBlockNumber *big.Int, blockHash string, position *big.Int,
) (*types.Transaction, bool, error) {
	c.log.Info("crafting proveFaultWithZkVm tx", "outputIndex", outputIndex, "challenger", challenge.Challenger)

	l1Head := hex.EncodeToString(challenge.L1Head[:])

	requestStatus, err := c.cfg.WitnessGenerator.RequestWitness(ctx, blockHash, l1Head)
	if err != nil {
		return nil, true, fmt.Errorf("failed to request witness(target block number: %s): %w", targetBlockNumber.String(), err)
	}
	if *requestStatus != chal.RequestCompleted {
		c.log.Info("witness generation is not completed", "targetBlockNumber", targetBlockNumber, "status", *requestStatus)
		return nil, true, nil
	}

	witnessResult, err := c.cfg.WitnessGenerator.GetWitness(ctx, blockHash, l1Head)
	if err != nil {
		return nil, true, fmt.Errorf("failed to get witness(target block number: %s): %w", targetBlockNumber.String(), err)
	}

	requestStatus, err = c.cfg.ZkVMProofFetcher.RequestProve(ctx, blockHash, l1Head, witnessResult.Witness)
	if err != nil {
		return nil, true, fmt.Errorf("failed to request proof(target block number: %s): %w", targetBlockNumber.String(), err)
	}
	if *requestStatus != chal.RequestCompleted {
		c.log.Info("proof generation is not completed", "targetBlockNumber", targetBlockNumber, "status", *requestStatus)
		return nil, true, nil
	}

	proofResult, err := c.cfg.ZkVMProofFetcher.GetProof(ctx, blockHash, l1Head)
	if err != nil {
		return nil, true, fmt.Errorf("failed to get proof(target block number: %s): %w", targetBlockNumber.String(), err)
	}

	txOpts := optsutils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	tx, err := c.colosseumContract.ProveFaultWithZkVm(
		txOpts,
		outputIndex,
		position,
		bindings.TypesZkVmProof{
			ZkVmProgramVKey: proofResult.VKeyHash,
			PublicValues:    proofResult.PublicValues,
			ProofBytes:      proofResult.Proof,
		},
	)
	if err != nil {
		return nil, true, err
	}

	return tx, false, nil
}

// proveFaultWithZkEvm fetches public input and zkEVM proof to create proveFaultWithZkEvm transaction.
// It sends requests to external RPC and waits until responded.
func (c *Challenger) proveFaultWithZkEvm(
	ctx context.Context, outputIndex *big.Int, challenger common.Address, targetBlockNumber *big.Int, position *big.Int,
) (*types.Transaction, bool, error) {
	c.log.Info("crafting proveFaultWithZkEvm tx", "outputIndex", outputIndex, "challenger", challenger)

	srcBlockNumber := new(big.Int).Sub(targetBlockNumber, common.Big1)
	proof, err := c.PublicInputProof(ctx, srcBlockNumber.Uint64())
	if err != nil {
		return nil, true, fmt.Errorf("failed to get public input proof(prev block number: %s): %w", srcBlockNumber.String(), err)
	}

	cCtx, cCancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cCancel()
	trace, err := c.cfg.L2Client.GetBlockTraceByNumber(cCtx, targetBlockNumber)
	if err != nil {
		return nil, true, fmt.Errorf("failed to get block trace(target block number: %s): %w", targetBlockNumber.String(), err)
	}

	traceBz, err := json.Marshal(trace)
	if err != nil {
		return nil, true, fmt.Errorf("failed to marshal block trace(target block number: %s): %w", targetBlockNumber.String(), err)
	}

	fetchResult, err := c.cfg.ZkEVMProofFetcher.FetchProofAndPair(ctx, string(traceBz))
	if err != nil {
		return nil, true, fmt.Errorf("failed to fetch proof and pair(target block number: %s): %w", targetBlockNumber.String(), err)
	}

	txOpts := optsutils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	tx, err := c.colosseumContract.ProveFaultWithZkEvm(
		txOpts,
		outputIndex,
		position,
		bindings.TypesZkEvmProof{
			PublicInputProof: proof,
			Proof:            fetchResult.Proof,
			// NOTE(0xHansLee): the hash of public input (pair[4], pair[5]) is not needed in proving fault.
			// It can be calculated using public input sent to colosseum contract.
			Pair: fetchResult.Pair[:4],
		},
	)
	if err != nil {
		return nil, true, err
	}

	return tx, false, nil
}

// IsOutputDeleted checks if the output is deleted.
func IsOutputDeleted(outputRoot [32]byte) bool {
	return bytes.Equal(outputRoot[:], deletedOutputRoot[:])
}
