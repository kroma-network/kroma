package validator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/wemixkanvas/kanvas/bindings/bindings"
	"github.com/wemixkanvas/kanvas/components/node/eth"
	chal "github.com/wemixkanvas/kanvas/components/validator/challenge"
	"github.com/wemixkanvas/kanvas/utils"
)

var OutputsPerWeek = big.NewInt(24 * 7)

type ProofFetcher interface {
	FetchProofAndPair(blockRef eth.L2BlockRef) (*chal.ProofAndPair, error)
	Close() error
}

type Challenger struct {
	wg   sync.WaitGroup
	done chan struct{}
	log  log.Logger
	cfg  Config

	ctx      context.Context
	cancel   context.CancelFunc
	callOpts *bind.CallOpts
	txOpts   *bind.TransactOpts

	l2ooContract      *bindings.L2OutputOracle
	colosseumContract *bindings.Colosseum

	fetcher            ProofFetcher
	submissionInterval *big.Int
	checkpoint         *big.Int
}

func NewChallenger(ctx context.Context, cfg Config, l log.Logger) (*Challenger, error) {
	colosseumContract, err := bindings.NewColosseum(cfg.ColosseumAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}
	l2ooContract, err := bindings.NewL2OutputOracle(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	submissionInterval, err := l2ooContract.SUBMISSIONINTERVAL(utils.NewSimpleCallOpts(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get submission interval: %w", err)
	}

	return &Challenger{
		done:     make(chan struct{}),
		log:      l,
		ctx:      ctx,
		cfg:      cfg,
		callOpts: utils.NewCallOptsWithSender(ctx, cfg.From),
		txOpts:   utils.NewSimpleTxOpts(ctx, cfg.From, cfg.SignerFn),

		l2ooContract:      l2ooContract,
		colosseumContract: colosseumContract,

		fetcher:            cfg.ProofFetcher,
		submissionInterval: submissionInterval,
	}, nil
}

func (c *Challenger) IsChallengeInProgress() (bool, error) {
	return c.colosseumContract.IsInProgress(c.callOpts)
}

func (c *Challenger) GetChallengeInProgress() (bindings.TypesChallenge, error) {
	return c.colosseumContract.GetChallengeInProgress(c.callOpts)
}

func (c *Challenger) OutputAtBlockSafe(blockNumber uint64) (*eth.OutputResponse, error) {
	output, err := c.cfg.RollupClient.OutputAtBlock(c.ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	if blockNumber == 0 {
		output.OutputRoot = eth.Bytes32{}
	}

	return output, nil
}

type OutputRange struct {
	OutputIndex *big.Int
	StartBlock  uint64
	EndBlock    uint64
}

func (c *Challenger) GetInvalidOutputRange() (*OutputRange, error) {
	nextOutputIndex, err := c.l2ooContract.NextOutputIndex(c.callOpts)
	if err != nil {
		return nil, err
	}
	if nextOutputIndex.Cmp(common.Big0) == 0 {
		c.log.Info("the output has not been submitted yet.")
		return nil, nil
	}
	latestOutputIndex := new(big.Int).Sub(nextOutputIndex, common.Big1)

	if c.checkpoint == nil {
		if latestOutputIndex.Cmp(OutputsPerWeek) == -1 {
			c.checkpoint = new(big.Int)
		} else {
			c.checkpoint = new(big.Int).Sub(latestOutputIndex, OutputsPerWeek)
		}
	}

	for i := c.checkpoint; i.Cmp(latestOutputIndex) != 1; i.Add(i, common.Big1) {
		output, err := c.l2ooContract.GetL2Output(c.callOpts, i)
		if err != nil {
			return nil, err
		}

		knownOutput, err := c.OutputAtBlockSafe(output.L2BlockNumber.Uint64())
		if err != nil {
			return nil, err
		}

		start := output.L2BlockNumber.Uint64() - c.submissionInterval.Uint64()
		end := output.L2BlockNumber.Uint64()

		if !bytes.Equal(knownOutput.OutputRoot[:], output.OutputRoot[:]) {
			c.checkpoint = i
			c.log.Info(
				"found invalid output",
				"blockNumber", output.L2BlockNumber,
				"outputIndex", i,
				"known", knownOutput.OutputRoot,
				"invalid", common.BytesToHash(output.OutputRoot[:]),
			)
			return &OutputRange{
				OutputIndex: i,
				StartBlock:  start,
				EndBlock:    end,
			}, nil
		} else {
			c.log.Info("confirmed that the output is valid",
				"outputIndex", i,
				"start", start,
				"end", end,
				"outputRoot", common.BytesToHash(output.OutputRoot[:]),
			)
		}
	}

	c.checkpoint = new(big.Int).Add(latestOutputIndex, common.Big1)

	return nil, nil
}

func (c *Challenger) DetermineChallengeTx() (*types.Transaction, error) {
	// Check for a challenge in progress.
	isInProgress, err := c.IsChallengeInProgress()
	if err != nil {
		return nil, fmt.Errorf("unable to get challenge in progress: %w", err)
	}

	if isInProgress {
		isRelated, err := c.IsRelatedChallenge()
		if err != nil {
			return nil, fmt.Errorf("unable to check relationship with challenge: %w", err)
		}

		if !isRelated {
			c.log.Info("another challenge is in progress")
			return nil, nil
		}

		status, err := c.GetStatusInProgress()
		if err != nil {
			return nil, fmt.Errorf("unable to get challenge status: %w", err)
		}

		if !c.cfg.OutputSubmitterDisabled {
			switch status {
			case chal.StatusAsserterTurn:
				return c.Bisect()
			case chal.StatusChallengerTimeout:
				challengeId, err := c.GetChallengeInProgress()
				if err != nil {
					return nil, err
				}

				// TODO(pangssu): Is it necessary to submit challengerTimeout transaction?
				c.log.Info("challenger timed out", "challengeId", challengeId)
				return nil, nil
			}
		}

		if !c.cfg.ChallengerDisabled {
			switch status {
			case chal.StatusChallengerTurn:
				return c.Bisect()
			case chal.StatusAsserterTimeout:
				return c.AsserterTimeout()
			case chal.StatusProveReady:
				return c.ProveFault()
			}
		}

		c.log.Warn("unknown challenge status", "status", status)
	} else if !c.cfg.ChallengerDisabled {
		outputRange, err := c.GetInvalidOutputRange()
		if err != nil {
			return nil, fmt.Errorf("unable to find invalid output: %w", err)
		}

		if outputRange == nil {
			return nil, nil
		}

		return c.CreateChallenge(outputRange)
	}

	return nil, nil
}

func (c *Challenger) IsRelatedChallenge() (bool, error) {
	return c.colosseumContract.IsChallengeRelated(c.callOpts, c.cfg.From)
}

func (c *Challenger) GetStatusInProgress() (uint8, error) {
	return c.colosseumContract.GetStatusInProgress(c.callOpts)
}

func (c *Challenger) LatestChallengeId() (*big.Int, error) {
	return c.colosseumContract.LatestChallengeId(c.callOpts)
}

func (c *Challenger) BuildSegments(turn, segStart, segSize uint64) (*chal.Segments, error) {
	sections, err := c.colosseumContract.GetSegmentsLength(c.callOpts, new(big.Int).SetUint64(turn))
	if err != nil {
		return nil, fmt.Errorf("unable to get segments length of turn %d: %w", turn, err)
	}

	segments := chal.NewEmptySegments(segStart, segSize, sections.Uint64())

	for i, blockNumber := range segments.BlockNumbers() {
		if blockNumber == 0 {
			segments.SetHashValue(0, eth.Bytes32{})
			continue
		}

		output, err := c.OutputAtBlockSafe(blockNumber)
		if err != nil {
			return nil, fmt.Errorf("unable to get output %d: %w", blockNumber, err)
		}

		segments.SetHashValue(i, output.OutputRoot)
	}

	return segments, nil
}

func (c *Challenger) selectFaultPosition(segments *chal.Segments) (*big.Int, error) {
	for i, blockNumber := range segments.BlockNumbers() {
		output, err := c.OutputAtBlockSafe(blockNumber)
		if err != nil {
			return nil, err
		}

		if !bytes.Equal(segments.Hashes[i][:], output.OutputRoot[:]) {
			return big.NewInt(int64(i) - 1), nil
		}
	}

	return nil, errors.New("failed to select fault position")
}

func (c *Challenger) CreateChallenge(outputRange *OutputRange) (*types.Transaction, error) {
	c.log.Info("crafting createChallenge tx",
		"index", outputRange.OutputIndex,
		"start", outputRange.StartBlock,
		"end", outputRange.EndBlock,
	)

	segSize := outputRange.EndBlock - outputRange.StartBlock
	segments, err := c.BuildSegments(1, outputRange.StartBlock, segSize)
	if err != nil {
		return nil, err
	}

	return c.colosseumContract.CreateChallenge(c.txOpts, outputRange.OutputIndex, segments.Hashes)
}

func (c *Challenger) Bisect() (*types.Transaction, error) {
	c.log.Info("crafting bisect tx")

	challenge, err := c.colosseumContract.GetChallengeInProgress(c.callOpts)
	if err != nil {
		return nil, err
	}

	prevSegments := chal.NewSegments(challenge.SegStart.Uint64(), challenge.SegSize.Uint64(), challenge.Segments)
	position, err := c.selectFaultPosition(prevSegments)
	if err != nil {
		return nil, err
	}
	nextTurn := challenge.Turn.Uint64() + 1
	start, size := prevSegments.NextSegmentsRange(position.Uint64())
	nextSegments, err := c.BuildSegments(nextTurn, start, size)
	if err != nil {
		return nil, err
	}

	return c.colosseumContract.Bisect(c.txOpts, position, nextSegments.Hashes)
}

func (c *Challenger) AsserterTimeout() (*types.Transaction, error) {
	c.log.Info("crafting timeout tx")
	return c.colosseumContract.AsserterTimeout(c.txOpts)
}

func (c *Challenger) ChallengerTimeout(challengeId *big.Int) (*types.Transaction, error) {
	c.log.Info("crafting timeout tx")
	return c.colosseumContract.ChallengerTimeout(c.txOpts, challengeId)
}

func (c *Challenger) ProveFault() (*types.Transaction, error) {
	c.log.Info("crafting proveFault tx")

	challenge, err := c.colosseumContract.GetChallengeInProgress(c.callOpts)
	if err != nil {
		return nil, err
	}

	segments := chal.NewSegments(challenge.SegStart.Uint64(), challenge.SegSize.Uint64(), challenge.Segments)
	position, err := c.selectFaultPosition(segments)
	if err != nil {
		return nil, err
	}

	blockNumber := challenge.SegStart.Uint64() + position.Uint64() + 1
	output, err := c.OutputAtBlockSafe(blockNumber)
	if err != nil {
		return nil, err
	}

	fetchResult, err := c.fetcher.FetchProofAndPair(output.BlockRef)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: blockNumber: %d, blockHash: %s",
			err, output.BlockRef.Number, output.BlockRef.Hash,
		)
	}

	return c.colosseumContract.ProveFault(
		c.txOpts,
		position,
		output.OutputRoot,
		fetchResult.Proof,
		fetchResult.Pair,
	)
}
