package e2eutils

import (
	"crypto/ecdsa"
	"fmt"

	hdwallet "github.com/ethereum-optimism/go-ethereum-hdwallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// DefaultMnemonicConfig is the default mnemonic used in testing.
// We prefer a mnemonic rather than direct private keys to make it easier
// to export all testing keys in external tooling for use during debugging.
// If these values are changed, it is subject to breaking tests. They
// must be in sync with the values in the DeployConfig used to create the system.
var DefaultMnemonicConfig = &MnemonicConfig{
	Mnemonic:         "test test test test test test test test test test test junk",
	CliqueSigner:     "m/44'/60'/0'/0/0",
	TrustedValidator: "m/44'/60'/0'/0/1", // 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
	Batcher:          "m/44'/60'/0'/0/2",
	Deployer:         "m/44'/60'/0'/0/3",
	Alice:            "m/44'/60'/0'/0/4",
	SequencerP2P:     "m/44'/60'/0'/0/5",
	Bob:              "m/44'/60'/0'/0/7",
	Mallory:          "m/44'/60'/0'/0/8", // 0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f
	SysCfgOwner:      "m/44'/60'/0'/0/0",

	// [Kroma: START],
	Challenger1: "m/44'/60'/0'/0/11", // 0x71bE63f3384f5fb98995898A86B02Fb2426c5788
	Challenger2: "m/44'/60'/0'/0/12", // 0xFABB0ac9d68B0B445fB7357272Ff202C5651694a
	Guardian:    "m/44'/60'/0'/0/13", // 0x1CBd3b2770909D4e10f157cABC84C7264073C9Ec
	// [Kroma: END]
}

// MnemonicConfig configures the private keys for the hive testnet.
// It's json-serializable, so we can ship it to e.g. the hardhat script client.
type MnemonicConfig struct {
	Mnemonic string

	CliqueSigner string
	Deployer     string
	SysCfgOwner  string

	// rollup actors
	TrustedValidator string
	Batcher          string
	SequencerP2P     string

	// prefunded L1/L2 accounts for testing
	Alice   string
	Bob     string
	Mallory string

	// [Kroma: START]
	Challenger1 string
	Challenger2 string
	Guardian    string
	// [Kroma: END
}

// Secrets computes the private keys for all mnemonic paths,
// which can then be kept around for fast precomputed private key access.
func (m *MnemonicConfig) Secrets() (*Secrets, error) {
	wallet, err := hdwallet.NewFromMnemonic(m.Mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}
	account := func(path string) accounts.Account {
		return accounts.Account{URL: accounts.URL{Path: path}}
	}

	deployer, err := wallet.PrivateKey(account(m.Deployer))
	if err != nil {
		return nil, err
	}
	cliqueSigner, err := wallet.PrivateKey(account(m.CliqueSigner))
	if err != nil {
		return nil, err
	}
	sysCfgOwner, err := wallet.PrivateKey(account(m.SysCfgOwner))
	if err != nil {
		return nil, err
	}
	trustedValidator, err := wallet.PrivateKey(account(m.TrustedValidator))
	if err != nil {
		return nil, err
	}
	batcher, err := wallet.PrivateKey(account(m.Batcher))
	if err != nil {
		return nil, err
	}
	sequencerP2P, err := wallet.PrivateKey(account(m.SequencerP2P))
	if err != nil {
		return nil, err
	}
	alice, err := wallet.PrivateKey(account(m.Alice))
	if err != nil {
		return nil, err
	}
	bob, err := wallet.PrivateKey(account(m.Bob))
	if err != nil {
		return nil, err
	}
	mallory, err := wallet.PrivateKey(account(m.Mallory))
	if err != nil {
		return nil, err
	}

	// [Kroma: START]
	challenger1, err := wallet.PrivateKey(account(m.Challenger1))
	if err != nil {
		return nil, err
	}
	challenger2, err := wallet.PrivateKey(account(m.Challenger2))
	if err != nil {
		return nil, err
	}
	guardian, err := wallet.PrivateKey(account(m.Guardian))
	if err != nil {
		return nil, err
	}
	//	[Kroma: END]

	return &Secrets{
		Deployer:         deployer,
		SysCfgOwner:      sysCfgOwner,
		CliqueSigner:     cliqueSigner,
		TrustedValidator: trustedValidator,
		Batcher:          batcher,
		SequencerP2P:     sequencerP2P,
		Alice:            alice,
		Bob:              bob,
		Mallory:          mallory,
		Wallet:           wallet,

		// [Kroma :START]
		Challenger1: challenger1,
		Challenger2: challenger2,
		Guardian:    guardian,
		//	[Kroma: END]
	}, nil
}

// Secrets bundles secp256k1 private keys for all common rollup actors for testing purposes.
type Secrets struct {
	Deployer     *ecdsa.PrivateKey
	CliqueSigner *ecdsa.PrivateKey
	SysCfgOwner  *ecdsa.PrivateKey

	// rollup actors
	TrustedValidator *ecdsa.PrivateKey
	Batcher          *ecdsa.PrivateKey
	SequencerP2P     *ecdsa.PrivateKey

	// prefunded L1/L2 accounts for testing
	Alice   *ecdsa.PrivateKey
	Bob     *ecdsa.PrivateKey
	Mallory *ecdsa.PrivateKey

	// Share the wallet to be able to generate more accounts
	Wallet *hdwallet.Wallet

	// [Kroma: START]
	Challenger1 *ecdsa.PrivateKey
	Challenger2 *ecdsa.PrivateKey
	Guardian    *ecdsa.PrivateKey
	// [Kroma: END]
}

// EncodePrivKey encodes the given private key in 32 bytes
func EncodePrivKey(priv *ecdsa.PrivateKey) hexutil.Bytes {
	privkey := make([]byte, 32)
	blob := priv.D.Bytes()
	copy(privkey[32-len(blob):], blob)
	return privkey
}

func EncodePrivKeyToString(priv *ecdsa.PrivateKey) string {
	return hexutil.Encode(EncodePrivKey(priv))
}

// Addresses computes the ethereum address of each account,
// which can then be kept around for fast precomputed address access.
func (s *Secrets) Addresses() *Addresses {
	return &Addresses{
		Deployer:         crypto.PubkeyToAddress(s.Deployer.PublicKey),
		CliqueSigner:     crypto.PubkeyToAddress(s.CliqueSigner.PublicKey),
		SysCfgOwner:      crypto.PubkeyToAddress(s.SysCfgOwner.PublicKey),
		TrustedValidator: crypto.PubkeyToAddress(s.TrustedValidator.PublicKey),
		Batcher:          crypto.PubkeyToAddress(s.Batcher.PublicKey),
		SequencerP2P:     crypto.PubkeyToAddress(s.SequencerP2P.PublicKey),
		Alice:            crypto.PubkeyToAddress(s.Alice.PublicKey),
		Bob:              crypto.PubkeyToAddress(s.Bob.PublicKey),
		Mallory:          crypto.PubkeyToAddress(s.Mallory.PublicKey),

		// [Kroma: START]
		Challenger1: crypto.PubkeyToAddress(s.Challenger1.PublicKey),
		Challenger2: crypto.PubkeyToAddress(s.Challenger2.PublicKey),
		Guardian:    crypto.PubkeyToAddress(s.Guardian.PublicKey),
		// [Kroma: END]
	}
}

// Addresses bundles the addresses for all common rollup addresses for testing purposes.
type Addresses struct {
	Deployer     common.Address
	CliqueSigner common.Address
	SysCfgOwner  common.Address

	// rollup actors
	TrustedValidator common.Address
	Batcher          common.Address
	SequencerP2P     common.Address

	// prefunded L1/L2 accounts for testing
	Alice   common.Address
	Bob     common.Address
	Mallory common.Address

	// [Kroma: START]
	Challenger1 common.Address
	Challenger2 common.Address
	Guardian    common.Address
	// [Kroma: END]
}

func (a *Addresses) All() []common.Address {
	return []common.Address{
		a.Deployer,
		a.CliqueSigner,
		a.SysCfgOwner,
		a.TrustedValidator,
		a.Batcher,
		a.SequencerP2P,
		a.Alice,
		a.Bob,
		a.Mallory,

		// [Kroma: START]
		a.Challenger1,
		a.Challenger2,
		a.Guardian,
		// [Kroma: END]
	}
}
