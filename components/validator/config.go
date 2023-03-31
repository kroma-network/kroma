package validator

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/components/node/sources"
	chal "github.com/wemixkanvas/kanvas/components/validator/challenge"
	"github.com/wemixkanvas/kanvas/components/validator/flags"
	"github.com/wemixkanvas/kanvas/utils"
	kcrypto "github.com/wemixkanvas/kanvas/utils/service/crypto"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	kmetrics "github.com/wemixkanvas/kanvas/utils/service/metrics"
	kpprof "github.com/wemixkanvas/kanvas/utils/service/pprof"
	krpc "github.com/wemixkanvas/kanvas/utils/service/rpc"
	"github.com/wemixkanvas/kanvas/utils/service/txmgr"
	ksigner "github.com/wemixkanvas/kanvas/utils/signer/client"
)

// Config contains the well typed fields that are used to initialize the output submitter.
// It is intended for programmatic use.
type Config struct {
	L2OutputOracleAddr      common.Address
	ColosseumAddr           common.Address
	PollInterval            time.Duration
	TxManagerConfig         txmgr.Config
	L1Client                *ethclient.Client
	RollupClient            *sources.RollupClient
	AllowNonFinalized       bool
	OutputSubmitterDisabled bool
	ChallengerDisabled      bool
	ProofFetcher            ProofFetcher
	From                    common.Address
	SignerFn                kcrypto.SignerFn
}

// CLIConfig is a well typed config that is parsed from the CLI params.
// This also contains config options for auxiliary services.
// It is transformed into a `Config` before the Validator is started.
type CLIConfig struct {
	/* Required Params */

	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// RollupRpc is the HTTP provider URL for the rollup node.
	RollupRpc string

	// L2OOAddress is the L2OutputOracle contract address.
	L2OOAddress string

	// ColosseumAddress is the Colosseum contract address.
	ColosseumAddress string

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// NumConfirmations is the number of confirmations which we will wait after
	// appending new batches.
	NumConfirmations uint64

	// SafeAbortNonceTooLowCount is the number of ErrNonceTooLowObservations
	// required to give up on a tx at a particular nonce without receiving
	// confirmation.
	SafeAbortNonceTooLowCount uint64

	// ResubmissionTimeout is time we will wait before resubmitting a
	// transaction.
	ResubmissionTimeout time.Duration

	// Mnemonic is the HD seed used to derive the wallet private keys for
	// the validator.
	Mnemonic string

	// HDPath is the derivation path used to obtain the private key for
	// the validator.
	HDPath string

	// PrivateKey is the private key used for the validator.
	PrivateKey string

	RPCConfig krpc.CLIConfig

	// ProverGrpc is the URL of prover grpc server.
	ProverGrpc string

	/* Optional Params */

	// AllowNonFinalized can be set to true to submit outputs
	// for L2 blocks derived from non-finalized L1 data.
	AllowNonFinalized bool

	OutputSubmitterDisabled bool

	ChallengerDisabled bool

	FetchingProofTimeout time.Duration

	LogConfig klog.CLIConfig

	MetricsConfig kmetrics.CLIConfig

	PprofConfig kpprof.CLIConfig

	// SignerConfig contains the client config for signer service
	SignerConfig ksigner.CLIConfig
}

func (c CLIConfig) Check() error {
	if err := c.RPCConfig.Check(); err != nil {
		return err
	}
	if err := c.LogConfig.Check(); err != nil {
		return err
	}
	if err := c.MetricsConfig.Check(); err != nil {
		return err
	}
	if err := c.PprofConfig.Check(); err != nil {
		return err
	}
	if err := c.SignerConfig.Check(); err != nil {
		return err
	}
	return nil
}

// NewCLIConfig parses the CLIConfig from the provided flags or environment variables.
func NewCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		// Required Flags
		L1EthRpc:                  ctx.GlobalString(flags.L1EthRpcFlag.Name),
		RollupRpc:                 ctx.GlobalString(flags.RollupRpcFlag.Name),
		L2OOAddress:               ctx.GlobalString(flags.L2OOAddressFlag.Name),
		ColosseumAddress:          ctx.GlobalString(flags.ColosseumAddressFlag.Name),
		PollInterval:              ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		NumConfirmations:          ctx.GlobalUint64(flags.NumConfirmationsFlag.Name),
		SafeAbortNonceTooLowCount: ctx.GlobalUint64(flags.SafeAbortNonceTooLowCountFlag.Name),
		ResubmissionTimeout:       ctx.GlobalDuration(flags.ResubmissionTimeoutFlag.Name),
		Mnemonic:                  ctx.GlobalString(flags.MnemonicFlag.Name),
		HDPath:                    ctx.GlobalString(flags.HDPathFlag.Name),
		PrivateKey:                ctx.GlobalString(flags.PrivateKeyFlag.Name),
		ProverGrpc:                ctx.GlobalString(flags.ProverGrpcFlag.Name),
		// Optional Flags
		AllowNonFinalized:       ctx.GlobalBool(flags.AllowNonFinalizedFlag.Name),
		OutputSubmitterDisabled: ctx.GlobalBool(flags.OutputSubmitterDisabledFlag.Name),
		ChallengerDisabled:      ctx.GlobalBool(flags.ChallengerDisabledFlag.Name),
		FetchingProofTimeout:    ctx.GlobalDuration(flags.FetchingProofTimeoutFlag.Name),
		RPCConfig:               krpc.ReadCLIConfig(ctx),
		LogConfig:               klog.ReadCLIConfig(ctx),
		MetricsConfig:           kmetrics.ReadCLIConfig(ctx),
		PprofConfig:             kpprof.ReadCLIConfig(ctx),
		SignerConfig:            ksigner.ReadCLIConfig(ctx),
	}
}

// NewValidatorConfig creates a validator config with given the CLIConfig
func NewValidatorConfig(cfg CLIConfig, l log.Logger) (*Config, error) {
	l2ooAddress, err := utils.ParseAddress(cfg.L2OOAddress)
	if err != nil {
		return nil, err
	}

	colosseumAddress, err := utils.ParseAddress(cfg.ColosseumAddress)
	if err != nil {
		return nil, err
	}

	signer, fromAddress, err := kcrypto.SignerFactoryFromConfig(l, cfg.PrivateKey, cfg.Mnemonic, cfg.HDPath, cfg.SignerConfig)
	if err != nil {
		return nil, err
	}

	var fetcher ProofFetcher
	if len(cfg.ProverGrpc) > 0 {
		fetcher, err = chal.NewFetcher(cfg.ProverGrpc, cfg.FetchingProofTimeout, l)
		if err != nil {
			return nil, err
		}
	}

	// Connect to L1 and L2 providers. Perform these last since they are the most expensive.
	ctx := context.Background()
	l1Client, err := utils.DialEthClientWithTimeout(ctx, cfg.L1EthRpc)
	if err != nil {
		return nil, err
	}

	rollupClient, err := utils.DialRollupClientWithTimeout(ctx, cfg.RollupRpc)
	if err != nil {
		return nil, err
	}

	chainID, err := l1Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	txMgrCfg := txmgr.Config{
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
		From:                      fromAddress,
		Signer:                    signer(chainID),
	}

	validatorCfg := &Config{
		L2OutputOracleAddr:      l2ooAddress,
		ColosseumAddr:           colosseumAddress,
		PollInterval:            cfg.PollInterval,
		TxManagerConfig:         txMgrCfg,
		L1Client:                l1Client,
		RollupClient:            rollupClient,
		AllowNonFinalized:       cfg.AllowNonFinalized,
		OutputSubmitterDisabled: cfg.OutputSubmitterDisabled,
		ChallengerDisabled:      cfg.ChallengerDisabled,
		ProofFetcher:            fetcher,
		From:                    fromAddress,
		SignerFn:                signer(chainID),
	}

	return validatorCfg, nil
}
