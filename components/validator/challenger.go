package validator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	chal "github.com/kroma-network/kroma/components/validator/challenge"
	"github.com/kroma-network/kroma/utils"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

type ProofFetcher interface {
	FetchProofAndPair(blockRef uint64) (*chal.ProofAndPair, error)
	Close() error
}

type Challenger struct {
	log    log.Logger
	cfg    Config
	ctx    context.Context
	cancel context.CancelFunc

	l1Client *ethclient.Client

	l2ooContract      *bindings.L2OutputOracle
	l2ooABI           *abi.ABI
	colosseumContract *bindings.Colosseum
	colosseumABI      *abi.ABI

	submissionInterval        *big.Int
	finalizationPeriodSeconds *big.Int
	l2BlockTime               *big.Int
	checkpoint                *big.Int

	l2OutputSub  ethereum.Subscription
	challengeSub ethereum.Subscription

	txCandidatesChan           chan<- txmgr.TxCandidate
	l2OutputSubmittedEventChan chan *bindings.L2OutputOracleOutputSubmitted
	challengeCreatedEventChan  chan *bindings.ColosseumChallengeCreated

	wg sync.WaitGroup
}

func NewChallenger(ctx context.Context, cfg Config, l log.Logger) (*Challenger, error) {
	colosseumContract, err := bindings.NewColosseum(cfg.ColosseumAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	colosseumABI, err := bindings.ColosseumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	l2ooContract, err := bindings.NewL2OutputOracle(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	l2ooABI, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	callOpts := utils.NewSimpleCallOpts(ctx)
	submissionInterval, err := l2ooContract.SUBMISSIONINTERVAL(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission interval: %w", err)
	}
	finalizationPeriodSeconds, err := l2ooContract.FINALIZATIONPERIODSECONDS(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get finalization period seconds: %w", err)
	}
	l2BlockTime, err := l2ooContract.L2BLOCKTIME(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get l2 block time: %w", err)
	}

	return &Challenger{
		log: l,
		cfg: cfg,

		l1Client: cfg.L1Client,

		l2ooContract:      l2ooContract,
		l2ooABI:           l2ooABI,
		colosseumContract: colosseumContract,
		colosseumABI:      colosseumABI,

		submissionInterval:        submissionInterval,
		finalizationPeriodSeconds: finalizationPeriodSeconds,
		l2BlockTime:               l2BlockTime,
	}, nil
}

// initSub initialize subscriptions
func (c *Challenger) initSub(ctx context.Context) {
	opts := &bind.WatchOpts{Context: ctx}

	if !c.cfg.ChallengerDisabled {
		c.l2OutputSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				c.log.Warn("resubscribing after failed L2OutputSubmitted event", "err", err)
			}
			return c.l2ooContract.WatchOutputSubmitted(opts, c.l2OutputSubmittedEventChan, nil, nil, nil)
		})
	}

	c.challengeSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			c.log.Warn("resubscribing after failed ChallengeCreated event", "err", err)
		}
		return c.colosseumContract.WatchChallengeCreated(opts, c.challengeCreatedEventChan, nil, nil, nil)
	})
}

func (c *Challenger) Start(ctx context.Context, txCandidatesChan chan<- txmgr.TxCandidate) error {
	c.ctx, c.cancel = context.WithCancel(ctx)

	c.log.Info("start challenger")

	c.l2OutputSubmittedEventChan = make(chan *bindings.L2OutputOracleOutputSubmitted)
	c.challengeCreatedEventChan = make(chan *bindings.ColosseumChallengeCreated)
	c.txCandidatesChan = txCandidatesChan
	c.initSub(c.ctx)

	// if checkpoint is behind the latest output index, scan the previous outputs from the checkpoint
	nextOutputIndex, err := c.l2ooContract.NextOutputIndex(&bind.CallOpts{Context: c.ctx})
	if err != nil {
		return fmt.Errorf("failed to get the latest output index: %w", err)
	}
	if nextOutputIndex.Cmp(common.Big0) == 0 {
		// if no outputs have been submitted, set checkpoint to 1 because genesis output cannot be challenged.
		c.checkpoint = common.Big1
	} else {
		// set checkpoint to latestOutputIndex (nextOutputIndex - 1)
		c.checkpoint = new(big.Int).Sub(nextOutputIndex, common.Big1)
	}

	if err := c.scanPrevOutputs(c.ctx); err != nil {
		return fmt.Errorf("failed to scan previous outputs: %w", err)
	}

	// if challenge mode on, subscribe L2 output submission events
	if !c.cfg.ChallengerDisabled {
		c.wg.Add(1)
		go c.subscribeL2OutputSubmitted(c.ctx)
	}

	// subscribe challenge creation events
	c.wg.Add(1)
	go c.subscribeChallengeCreated(c.ctx)

	return nil
}

func (c *Challenger) Stop() error {
	c.log.Info("stop challenger")

	if c.l2OutputSub != nil {
		c.l2OutputSub.Unsubscribe()
	}

	if c.challengeSub != nil {
		c.challengeSub.Unsubscribe()
	}

	c.cancel()
	c.wg.Wait()

	close(c.l2OutputSubmittedEventChan)
	close(c.challengeCreatedEventChan)

	return nil
}

// scanPrevOutputs scans all the previous outputs since the checkpoint within the finalization window.
// If there are invalid outputs, create challenge.
// If there are challenges in progress, keep handling them.
func (c *Challenger) scanPrevOutputs(ctx context.Context) error {
	status, err := c.cfg.RollupClient.SyncStatus(ctx)
	if err != nil {
		return fmt.Errorf("failed to get sync status: %w", err)
	}

	toBlock := new(big.Int).SetUint64(status.CurrentL1.Number)
	// TODO(0xHansLee): add L1BlockTime to rollup config and change to use it
	finalizationStartL1Block := new(big.Int).Sub(toBlock, new(big.Int).Div(c.finalizationPeriodSeconds, big.NewInt(12)))
	// The fromBlock is the maximum value of either genesis block(1) or the first block of the finalization window
	fromBlock := math.BigMax(common.Big1, finalizationStartL1Block)

	outputSubmittedEvent := c.l2ooABI.Events[KeyEventOutputSubmitted]
	challengeCreatedEvent := c.colosseumABI.Events[KeyEventChallengeCreated]

	addresses := []common.Address{c.cfg.ColosseumAddr}
	topics := []common.Hash{challengeCreatedEvent.ID}

	// scan OutputSubmittedEvents only when challenger mode is on
	if !c.cfg.ChallengerDisabled {
		addresses = append(addresses, c.cfg.L2OutputOracleAddr)
		topics = append(topics, outputSubmittedEvent.ID)
	}

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: addresses,
		Topics:    [][]common.Hash{topics},
	}

	logs, err := c.l1Client.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to get event logs related to outputs: %w", err)
	}

	for _, vLog := range logs {
		switch vLog.Address {
		// for OutputSubmitted event
		case c.cfg.L2OutputOracleAddr:
			ev := NewOutputSubmittedEvent(vLog)
			// handle output
			c.wg.Add(1)
			go c.handleOutput(ctx, ev.OutputIndex)
		// for ChallengeCreated event
		case c.cfg.ColosseumAddr:
			ev := NewChallengeCreatedEvent(vLog)
			if ev.OutputIndex.Sign() == 1 && c.isRelatedChallenge(ev.Asserter, ev.Challenger) {
				c.wg.Add(1)
				go c.handleChallenge(ctx, ev.OutputIndex)
			}
		default:
			c.log.Warn("unknown event log", "logs", vLog)
		}
	}

	return nil
}

// subscribeL2OutputSubmitted subscribes the OutputSubmitted event from L2OutputOracle contract.
// If the L2 output root is invalid, create challenge.
// This function should be called only when challenger mode is on.
func (c *Challenger) subscribeL2OutputSubmitted(ctx context.Context) {
	defer c.wg.Done()

	for {
		select {
		case ev := <-c.l2OutputSubmittedEventChan:
			c.log.Info("validating output", "l2BlockNumber", ev.L2BlockNumber, "outputRoot", ev.OutputRoot, "outputIndex", ev.L2OutputIndex)
			// validate all outputs in between the checkpoint and the current outputIndex
			for i := c.checkpoint; i.Cmp(ev.L2OutputIndex) != 1; i.Add(i, common.Big1) {
				c.wg.Add(1)
				go c.handleOutput(ctx, i)
			}
			c.checkpoint = ev.L2OutputIndex
		case <-ctx.Done():
			return
		}
	}
}

// subscribeChallengeCreated subscribes the ChallengeCreated event from Colosseum contract and handle challenge.
func (c *Challenger) subscribeChallengeCreated(ctx context.Context) {
	defer c.wg.Done()

	for {
		select {
		case ev := <-c.challengeCreatedEventChan:
			// when challenge created, handle it
			if ev.OutputIndex.Sign() == 1 && c.isRelatedChallenge(ev.Asserter, ev.Challenger) {
				c.wg.Add(1)
				go c.handleChallenge(ctx, ev.OutputIndex)
			}
		case <-ctx.Done():
			return
		}
	}
}

// handleOutput handles output when output submitted
func (c *Challenger) handleOutput(ctx context.Context, outputIndex *big.Int) {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
	Loop:
		select {
		case <-ticker.C:
			outputRange, err := c.ValidateOutput(ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to validate output", "err", err, "outputIndex", outputIndex)
				break Loop
			}

			// if output is valid, terminate handling
			if outputRange == nil {
				c.log.Info("output is validated", "outputIndex", outputIndex)
				return
			}

			// check if the challenge is in progress already.
			isInProgress, err := c.IsChallengeInProgress(ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to get the status of challenge", "err", err, "outputIndex", outputIndex)
				break Loop
			}

			// if challenge is in progress, terminate handling
			if isInProgress {
				c.log.Info("found invalid output, but is already in progress", "outputIndex", outputIndex)
				return
			}

			// if there is no challenge on invalid output, create a new challenge
			tx, err := c.CreateChallenge(ctx, outputRange)
			if err != nil {
				c.log.Error("failed to create createChallenge tx", "err", err, "outputIndex", outputIndex)
				break Loop
			}

			c.submitChallengeTx(tx)
			return
		case <-ctx.Done():
			return
		}
	}
}

// handleChallenge handles challenge according to its status and role.
func (c *Challenger) handleChallenge(ctx context.Context, outputIndex *big.Int) {
	defer c.wg.Done()

	ticker := time.NewTicker(c.cfg.ChallengerPollInterval)
	defer ticker.Stop()

	for {
	Loop:
		select {
		case <-ticker.C:
			challenge, err := c.GetChallenge(ctx, outputIndex)
			if err != nil {
				c.log.Error("failed to get challenge", "err", err, "outputIndex", outputIndex)
				break Loop
			}

			isAsserter := challenge.Asserter == c.cfg.TxManager.From()
			isChallenger := challenge.Challenger == c.cfg.TxManager.From()

			// check the status of challenge
			status, err := c.GetChallengeStatus(ctx, outputIndex)
			if err != nil {
				c.log.Error("unable to get challenge status", "err", err, "outputIndex", outputIndex)
				break Loop
			}

			// if the challenge is inactivated, terminate handling
			if isInactivated(status) {
				c.log.Error("challenge is not in progress", "challengeStatus", status)
				return
			}

			// if asserter
			if isAsserter && !c.cfg.OutputSubmitterDisabled {
				if status == chal.StatusAsserterTurn {
					tx, err := c.Bisect(ctx, outputIndex)
					if err != nil {
						c.log.Error("asserter: failed to create bisect tx", "err", err, "outputIndex", outputIndex)
						break Loop
					}
					c.submitChallengeTx(tx)
				}
			}

			// if challenger
			if isChallenger && !c.cfg.ChallengerDisabled {
				switch status {
				case chal.StatusChallengerTurn:
					tx, err := c.Bisect(ctx, outputIndex)
					if err != nil {
						c.log.Error("challenger: failed to create bisect tx", "err", err, "outputIndex", outputIndex)
						break Loop
					}
					c.submitChallengeTx(tx)
				case chal.StatusAsserterTimeout, chal.StatusReadyToProve:
					skipSelectPosition := (status == chal.StatusAsserterTimeout)
					tx, err := c.ProveFault(ctx, outputIndex, skipSelectPosition)
					if err != nil {
						c.log.Error("failed to create prove fault tx", "err", err, "outputIndex", outputIndex)
						break Loop
					}
					c.submitChallengeTx(tx)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *Challenger) submitChallengeTx(tx *types.Transaction) {
	c.txCandidatesChan <- txmgr.TxCandidate{
		TxData:   tx.Data(),
		To:       tx.To(),
		GasLimit: 0,
	}
}

func (c *Challenger) IsChallengeInProgress(ctx context.Context, outputIndex *big.Int) (bool, error) {
	return c.colosseumContract.IsInProgress(&bind.CallOpts{Context: ctx}, outputIndex)
}

func (c *Challenger) GetChallenge(ctx context.Context, outputIndex *big.Int) (bindings.TypesChallenge, error) {
	return c.colosseumContract.GetChallenge(&bind.CallOpts{Context: ctx}, outputIndex)
}

func (c *Challenger) OutputAtBlockSafe(ctx context.Context, blockNumber uint64) (*eth.OutputResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cancel()
	return c.cfg.RollupClient.OutputAtBlock(ctx, blockNumber)
}

func (c *Challenger) OutputWithProofAtBlockSafe(ctx context.Context, blockNumber uint64) (*eth.OutputResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.cfg.NetworkTimeout)
	defer cancel()
	return c.cfg.RollupClient.OutputWithProofAtBlock(ctx, blockNumber)
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
	copy(balance[:], p.L2ToL1MessagePasserBalance.Bytes()[:])

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
	remoteOutput bindings.TypesCheckpointOutput
	localOutput  *eth.OutputResponse
}

func (c *Challenger) outputsAtIndex(ctx context.Context, outputIndex *big.Int) (*Outputs, error) {
	remoteOutput, err := c.l2ooContract.GetL2Output(&bind.CallOpts{Context: ctx}, outputIndex)
	if err != nil {
		return nil, err
	}

	localOutput, err := c.OutputAtBlockSafe(ctx, remoteOutput.L2BlockNumber.Uint64())
	if err != nil {
		return nil, err
	}

	return &Outputs{remoteOutput, localOutput}, nil
}

type OutputRange struct {
	OutputIndex *big.Int
	StartBlock  uint64
	EndBlock    uint64
}

// ValidateOutput validates the output given the outputIndex
func (c *Challenger) ValidateOutput(ctx context.Context, outputIndex *big.Int) (*OutputRange, error) {
	outputs, err := c.outputsAtIndex(ctx, outputIndex)
	if err != nil {
		return nil, err
	}

	start := outputs.remoteOutput.L2BlockNumber.Uint64() - c.submissionInterval.Uint64()
	end := outputs.remoteOutput.L2BlockNumber.Uint64()

	if !bytes.Equal(outputs.localOutput.OutputRoot[:], outputs.remoteOutput.OutputRoot[:]) {
		c.log.Info(
			"found invalid output",
			"blockNumber", outputs.remoteOutput.L2BlockNumber,
			"outputIndex", outputIndex,
			"local", outputs.localOutput.OutputRoot,
			"invalid", common.BytesToHash(outputs.remoteOutput.OutputRoot[:]),
		)
		return &OutputRange{
			OutputIndex: outputIndex,
			StartBlock:  start,
			EndBlock:    end,
		}, nil
	} else {
		c.log.Info("confirmed that the output is valid",
			"outputIndex", outputIndex,
			"start", start,
			"end", end,
			"outputRoot", common.BytesToHash(outputs.remoteOutput.OutputRoot[:]),
		)
		return nil, nil
	}
}

func (c *Challenger) isRelatedChallenge(asserter common.Address, challenger common.Address) bool {
	return c.cfg.TxManager.From() == asserter || c.cfg.TxManager.From() == challenger
}

func (c *Challenger) GetChallengeStatus(ctx context.Context, outputIndex *big.Int) (uint8, error) {
	return c.colosseumContract.GetStatus(&bind.CallOpts{Context: ctx}, outputIndex)
}

func (c *Challenger) BuildSegments(ctx context.Context, turn uint8, segStart, segSize uint64) (*chal.Segments, error) {
	sections, err := c.colosseumContract.GetSegmentsLength(&bind.CallOpts{Context: ctx}, turn)
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
	c.log.Info("crafting createChallenge tx",
		"index", outputRange.OutputIndex,
		"start", outputRange.StartBlock,
		"end", outputRange.EndBlock,
	)

	segSize := outputRange.EndBlock - outputRange.StartBlock
	segments, err := c.BuildSegments(ctx, 1, outputRange.StartBlock, segSize)
	if err != nil {
		return nil, err
	}

	txOpts := utils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.CreateChallenge(txOpts, outputRange.OutputIndex, segments.Hashes)
}

func (c *Challenger) Bisect(ctx context.Context, outputIndex *big.Int) (*types.Transaction, error) {
	c.log.Info("crafting bisect tx")

	challenge, err := c.colosseumContract.GetChallenge(&bind.CallOpts{Context: ctx}, outputIndex)
	if err != nil {
		return nil, err
	}

	prevSegments := chal.NewSegments(challenge.SegStart.Uint64(), challenge.SegSize.Uint64(), challenge.Segments)
	position, err := c.selectFaultPosition(ctx, prevSegments)
	if err != nil {
		return nil, err
	}
	nextTurn := challenge.Turn + 1
	start, size := prevSegments.NextSegmentsRange(position.Uint64())
	nextSegments, err := c.BuildSegments(ctx, nextTurn, start, size)
	if err != nil {
		return nil, err
	}

	txOpts := utils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.Bisect(txOpts, outputIndex, position, nextSegments.Hashes)
}

func (c *Challenger) ChallengerTimeout(ctx context.Context, outputIndex *big.Int) (*types.Transaction, error) {
	c.log.Info("crafting timeout tx")
	txOpts := utils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.ChallengerTimeout(txOpts, outputIndex)
}

// ProveFault creates proveFault transaction for invalid output root
// TODO: ProveFault will take long time, so that we may have to handle it carefully
func (c *Challenger) ProveFault(ctx context.Context, outputIndex *big.Int, skipSelectPosition bool) (*types.Transaction, error) {
	c.log.Info("crafting proveFault tx")

	outputs, err := c.outputsAtIndex(ctx, outputIndex)
	if err != nil {
		return nil, err
	}

	challenge, err := c.colosseumContract.GetChallenge(&bind.CallOpts{Context: ctx}, outputIndex)
	if err != nil {
		return nil, err
	}

	// When asserter timeout, skip finding fault position since the same segments have been stored in colosseum.
	position := common.Big0
	blockNumber := challenge.SegStart.Uint64()
	if !skipSelectPosition {
		segments := chal.NewSegments(challenge.SegStart.Uint64(), challenge.SegSize.Uint64(), challenge.Segments)
		position, err = c.selectFaultPosition(ctx, segments)
		if err != nil {
			return nil, err
		}

		blockNumber = challenge.SegStart.Uint64() + position.Uint64()
	}

	fetchResult, err := c.cfg.ProofFetcher.FetchProofAndPair(blockNumber)
	if err != nil {
		return nil, fmt.Errorf("%w: blockNumber: %d", err, blockNumber)
	}

	proof, err := c.PublicInputProof(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	txOpts := utils.NewSimpleTxOpts(ctx, c.cfg.TxManager.From(), c.cfg.TxManager.Signer)
	return c.colosseumContract.ProveFault(
		txOpts,
		outputIndex,
		outputs.localOutput.OutputRoot,
		position,
		proof,
		fetchResult.Proof,
		// NOTE(0xHansLee): the hash of public input (pair[4], pair[5]) is not needed in proving fault.
		// It can be calculated using public input sent to colosseum contract.
		fetchResult.Pair[:4],
	)
}

// isInactivated checks if the challenge is inactivated.
func isInactivated(status uint8) bool {
	return status == chal.StatusNone ||
		status == chal.StatusChallengerTimeout ||
		status == chal.StatusProven ||
		status == chal.StatusApproved
}
