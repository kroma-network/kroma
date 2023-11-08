package testdata

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

var (
	TargetBlockNumber           = uint64(21)
	CoinbaseAddr                = common.HexToAddress("0x0000000000000000000000000000000000000000")
	l2toL1MesasgePasserBalance  = common.Big0
	l2toL1MesasgePasserCodeHash = common.HexToHash("0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35")
	merkleProof                 = []hexutil.Bytes{
		hexutil.MustDecode("0x001d70e33820638e04e2a9edff096ef8a372067f654795d98810db23d4aa05650d0a7cae21b7fd43b77b12d45602e15180298ba59195186ad8b84685bd4df131c6"),
		hexutil.MustDecode("0x0009c507e2a822ac5a2e95760cbd5348f34c04148c5e2bc10f064632d424b53a1a29fc543090f1b44ab0ba6f939c9e8c9d329c614de53e2f38ec7457842b980ff2"),
		hexutil.MustDecode("0x0021e148200d68427f4cd748b40d577a3ebc24f29dba6a51690cb56625632735f5047b12ada42bc39c3f53a187bdd7cb98db773a6c0475cb7cf813861296273c51"),
		hexutil.MustDecode("0x002bf1dc335fbd3c6252fa426d647d20ff86181a1a3ac1f4c53122b5a853436a5900d62fc393b0150bfe7c840be535455215e921f99f673f440d5ddba5cc55ac07"),
		hexutil.MustDecode("0x00188f4a9b605caa246413179613bffe6c4efcd717d4d68a2dbc531123c8d2a4531d12fb4a76dc568a4655d2e360e9bdcd9f86542d6f7d327e3140477c92aa2920"),
		hexutil.MustDecode("0x00244f31e0a770ed9dbdf583a0e0dbe036f5f9476c3cc761e18b2a475a96bcc10809387d2d4be643d4d8df653dcb8495eebc36eba91f993f78fe843a799a6d0fb9"),
		hexutil.MustDecode("0x0008f993e0df87a04e71f72a3a237c909e2381eacf176b279925d333ee7b1e36ed03c30671a87c81a313a035fdbe052cc592ae0a604bcf87a5cf163d5a43104574"),
		hexutil.MustDecode("0x00202aa398b4bd976d7c165b2c7bfa6e4b695f18ff78e4cb544dbb0fbab8c6537e15c1089ff56ee758ec382a55a21b76c08ebf64d6a78d7e6ff442793536607510"),
		hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000002e957b48192277673a8dde0549358d09d3f9ad6e8db14e08f4fed46f96021a74"),
		hexutil.MustDecode("0x002325a334c56feef28306cece9e3867165ee117aab1c831fe04db286a1b4ff2c80000000000000000000000000000000000000000000000000000000000000000"),
		hexutil.MustDecode("0x00218476186a36a2ddf003ef59459478f44e0cea1ac32870dafca118331259b05f23618448c7fab9e44d30c44be6777aa390c25ad138fde11d22bdefd05f43838b"),
		hexutil.MustDecode("0x012de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c304040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc3524f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127202de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c3"),
		hexutil.MustDecode("0x5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449"),
	}
)

func SetPrevOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
		Version:                  rollup.V0,
		StateRoot:                common.HexToHash("0x2ecc9f95421c4f8c6acfd73a9983b021e79b381c9e80991b9b45da927c926c4f"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0x2d8d7264743ac0648b2b0fae0137cb0f77b2a952f5583a2cc6abf0c72f4f1b80"),
		NextBlockHash:            common.HexToHash("0x5cd3ba48964223516867ee8036fb0121c095d93a4301084f3fa37d811655d1e8"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x2ecc9f95421c4f8c6acfd73a9983b021e79b381c9e80991b9b45da927c926c4f")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber - 1,
		Hash:       common.HexToHash("0x2d8d7264743ac0648b2b0fae0137cb0f77b2a952f5583a2cc6abf0c72f4f1b80"),
		ParentHash: common.HexToHash("0x8c86435a69f9411db54e8ece1b55e389c6f68a4f93e1c477a409ff24bb916749"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0x5cd3ba48964223516867ee8036fb0121c095d93a4301084f3fa37d811655d1e8")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x2d8d7264743ac0648b2b0fae0137cb0f77b2a952f5583a2cc6abf0c72f4f1b80"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x1370c09d12e3aefefbe29bcecaa9a1adac759ee1b6657065f8e103f56b364037"),
		TxHash:      common.HexToHash("0x2bd4ac406e80d8401dcd8a3770eac5630ff2883931d1acc6fdd86ece3a23c4d9"),
		ReceiptHash: common.HexToHash("0xf75fc90d6167f310f60db885364953174fcd99a498b8629ea197074022a0eb67"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0x10371"),
		Time:        hexutil.MustDecodeUint64("0x64e4b763"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x3a61b96"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x98a9de84b8d737c493a2a3a8fc8713fbb5e950ed8b890ca912b1d10399a77070"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Mint:       nil,
			Value:      common.Big0,
			Gas:        1000000,
			Data:       hexutil.MustDecode("0xefc674eb000000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000064e4b75f000000000000000000000000000000000000000000000000000000000dcbf333756f13a3a20d49dcdc6dc48fc618c9d79b476d00693744946e85882a9df2e21600000000000000000000000000000000000000000000000000000000000000010000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc0000000000000000000000000000000000000000000000000000000000000834000000000000000000000000000000000000000000000000000000000016e3600000000000000000000000000000000000000000000000000000000000000000"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}

func SetTargetOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
		Version:                  rollup.V0,
		StateRoot:                common.HexToHash("0x1370c09d12e3aefefbe29bcecaa9a1adac759ee1b6657065f8e103f56b364037"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0x5cd3ba48964223516867ee8036fb0121c095d93a4301084f3fa37d811655d1e8"),
		NextBlockHash:            common.HexToHash("0xf7cda4b3b93fbdd37eac444fd3341ac1d8627ba861d706f272e947579121fb53"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x1370c09d12e3aefefbe29bcecaa9a1adac759ee1b6657065f8e103f56b364037")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber,
		Hash:       common.HexToHash("0x5cd3ba48964223516867ee8036fb0121c095d93a4301084f3fa37d811655d1e8"),
		ParentHash: common.HexToHash("0x2d8d7264743ac0648b2b0fae0137cb0f77b2a952f5583a2cc6abf0c72f4f1b80"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0xf7cda4b3b93fbdd37eac444fd3341ac1d8627ba861d706f272e947579121fb53")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x5cd3ba48964223516867ee8036fb0121c095d93a4301084f3fa37d811655d1e8"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x2cd3f8d15616fcbf9d24f46e5a87506ef1b95abce40f94400b2adcbf2fc0c121"),
		TxHash:      common.HexToHash("0x6312e7568e45027e8916c174bb9b09b9b89778491e4edaf2e47753c1aa191d25"),
		ReceiptHash: common.HexToHash("0x6328a94ff7ae439b334202f0bd6726d0cc24abdbf74d93f750fd29ebdfbaa882"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber + 1)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xceb5"),
		Time:        hexutil.MustDecodeUint64("0x64e4b765"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x331dc7e"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x2bab73de95ac558e82da0ff8537d1d0a289a26debc46b86a49c64bda50bf152d"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Gas:        1000000,
			Value:      common.Big0,
			Mint:       nil,
			Data:       hexutil.MustDecode("0xefc674eb000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000064e4b762000000000000000000000000000000000000000000000000000000000c143523254c69ffda05296015136a3ee1e3fb142dec7e3c295f0e4a2a5ef584b0762d3d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc0000000000000000000000000000000000000000000000000000000000000834000000000000000000000000000000000000000000000000000000000016e3600000000000000000000000000000000000000000000000000000000000000000"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}
