package testdata

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

var (
	TargetBlockNumber           = uint64(21)
	CoinbaseAddr                = common.HexToAddress("0x0000000000000000000000000000000000000000")
	l2toL1MesasgePasserBalance  = common.Big0
	l2toL1MesasgePasserCodeHash = common.HexToHash("0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35")
	merkleProof                 = []hexutil.Bytes{
		hexutil.MustDecode("0x002c93216458be4d48afd625db7b1f59df392099a3fbb908f1a665e963a5352e93093f4c7d37c8ed44d34d346f08a39995c238ebcd98fddcc5541d55093271c159"),
		hexutil.MustDecode("0x002a543be0e6a209964727ae31f1a57521dedb2a549fb59e5a190b42c5becbb2a82cf819ee96fa5bf76e6ce7712caf6673df7fc3af269cfe193bed443eeb89527a"),
		hexutil.MustDecode("0x0021d5edc847df9d20356a576c7a4b4b1bb992b427aa4e01412eb08864416db55b1c7c9c38ad18cc007a6e522b638fdf6010aef11727fcb7cdcbfb5edb96e3a74a"),
		hexutil.MustDecode("0x0020140f1c792e0ff14d98ca78ed1c085501566b9917cc0792285ed265e6f8cd5b1e99968ab013de794c675a5f614b8a642294aa47612f1bb6173910a5cd7f0b16"),
		hexutil.MustDecode("0x002b11bc4ac76ee779a652e7e93a9871fa88e780285d41c6ed2f1646ec8658b7450c158053a97ff3e327cfae62631e5878accd83e5803f044c90803d4579ae21ba"),
		hexutil.MustDecode("0x00244f31e0a770ed9dbdf583a0e0dbe036f5f9476c3cc761e18b2a475a96bcc10814c463c1a53d595a21f22f3c5cafe495329c68087ab2601468f8fa7e11d88bf7"),
		hexutil.MustDecode("0x0008f993e0df87a04e71f72a3a237c909e2381eacf176b279925d333ee7b1e36ed03c30671a87c81a313a035fdbe052cc592ae0a604bcf87a5cf163d5a43104574"),
		hexutil.MustDecode("0x00202aa398b4bd976d7c165b2c7bfa6e4b695f18ff78e4cb544dbb0fbab8c6537e15c1089ff56ee758ec382a55a21b76c08ebf64d6a78d7e6ff442793536607510"),
		hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000002e957b48192277673a8dde0549358d09d3f9ad6e8db14e08f4fed46f96021a74"),
		hexutil.MustDecode("0x002325a334c56feef28306cece9e3867165ee117aab1c831fe04db286a1b4ff2c80000000000000000000000000000000000000000000000000000000000000000"),
		hexutil.MustDecode("0x00218476186a36a2ddf003ef59459478f44e0cea1ac32870dafca118331259b05f23618448c7fab9e44d30c44be6777aa390c25ad138fde11d22bdefd05f43838b"),
		hexutil.MustDecode("0x012de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c304040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc3524f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127202de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c3"),
		hexutil.MustDecode("0x5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449"),
	}

	zeroUint64       = hexutil.MustDecodeUint64("0x0")
	parentBeaconRoot = common.HexToHash("0xf4db98f24ef4e2ce756ee5c926671ee2d68fba50c72ae29e25f45a1a8ec5966f")

	NextHeader = &types.Header{
		ParentHash:       common.HexToHash("0xc5b85aa0fcef96933ee0e2629c1cd30dd7d14083c03f08b1ab95f1e9f3757b46"),
		Coinbase:         CoinbaseAddr,
		Difficulty:       common.Big0,
		Root:             common.HexToHash("0x2ff9f8349c59d025d078de07d9f7175c3422109c9a11a4075d383ca62a685716"),
		TxHash:           common.HexToHash("0x91c1194e5d35a6f40e1aa4913664dacbcfbb36e8246e9b60ef1c00184f6ff05b"),
		ReceiptHash:      common.HexToHash("0xd6016faa5d7979ca763d5cbc6a228df98194cd3ecbd74e0b22833d55bd9028bf"),
		Bloom:            types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:           big.NewInt(int64(TargetBlockNumber)),
		GasLimit:         hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:          hexutil.MustDecodeUint64("0xdb30"),
		Time:             hexutil.MustDecodeUint64("0x660e9c50"),
		Extra:            hexutil.MustDecode("0x"),
		MixDigest:        common.HexToHash("0x363cf43f812d151c2e9d64de797403c7c73c973ec31b5de66db208de4a5c9d1b"),
		BaseFee:          big.NewInt(int64(hexutil.MustDecodeUint64("0x1"))),
		WithdrawalsHash:  &types.EmptyWithdrawalsHash,
		BlobGasUsed:      &zeroUint64,
		ExcessBlobGas:    &zeroUint64,
		ParentBeaconRoot: &parentBeaconRoot,
	}
)

func SetPrevOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
		Version:                  eth.OutputVersionV0,
		StateRoot:                common.HexToHash("0x0b1fd0c9beef9cb4ac073bc7ac702f237060926b59ad9b58c34225c3c2883042"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0xc5b85aa0fcef96933ee0e2629c1cd30dd7d14083c03f08b1ab95f1e9f3757b46"),
		NextBlockHash:            common.HexToHash("0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x0b1fd0c9beef9cb4ac073bc7ac702f237060926b59ad9b58c34225c3c2883042")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber - 1,
		Hash:       common.HexToHash("0xc5b85aa0fcef96933ee0e2629c1cd30dd7d14083c03f08b1ab95f1e9f3757b46"),
		ParentHash: common.HexToHash("0xf50c945d1e04b20196f23325d78ef3898c3d939e5b032da25c5b3f19fa857b2d"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = NextHeader
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x63d293eeca86da190bd6d3cbd660fa7de2bd0bab18566f3a6899f9933c079ade"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Mint:       nil,
			Value:      common.Big0,
			Gas:        1000000,
			Data:       hexutil.MustDecode("0x440a5e20000f424000000000000000000000000000000000660e9c4d000000000000000d000000000000000000000000000000000000000000000000000000000a83a7d000000000000000000000000000000000000000000000000000000000000000015f9c07c104765a40014bd34a890807136fdce9fc2a7a76fa75b1ac82f19c96030000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc0000000000000000000000000000000000000000000000000000000000001388"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}

func SetTargetOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
		Version:                  eth.OutputVersionV0,
		StateRoot:                common.HexToHash("0x2ff9f8349c59d025d078de07d9f7175c3422109c9a11a4075d383ca62a685716"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b"),
		NextBlockHash:            common.HexToHash("0x7b8aa6641024ecc150b38d9710d40aa6d77f800d2fd1788660914f190840af20"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x2ff9f8349c59d025d078de07d9f7175c3422109c9a11a4075d383ca62a685716")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber,
		Hash:       common.HexToHash("0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b"),
		ParentHash: common.HexToHash("0xc5b85aa0fcef96933ee0e2629c1cd30dd7d14083c03f08b1ab95f1e9f3757b46"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x29fce6172387f4b2ef7f0dcb9b9eb37b59c38acd6c98d1af014d4452b847fb22"),
		TxHash:      common.HexToHash("0xffb80c0722b4a3cd86827c831061530a056cd16b4576925d3b100d4581f85304"),
		ReceiptHash: common.HexToHash("0x55eed20b1c2b45fa5b76072ba6a64e38d625725c883212d3c0357d27a968ea0a"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber + 1)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xc55c"),
		Time:        hexutil.MustDecodeUint64("0x660e9c52"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x363cf43f812d151c2e9d64de797403c7c73c973ec31b5de66db208de4a5c9d1b"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x1"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x4e87ab1f0eda5108be11c44060c190b4a49e2309742e54be65d115de9aeeb4e7"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Gas:        1000000,
			Value:      common.Big0,
			Mint:       nil,
			Data:       hexutil.MustDecode("0x440a5e20000f424000000000000000000000000100000000660e9c4d000000000000000d000000000000000000000000000000000000000000000000000000000a83a7d000000000000000000000000000000000000000000000000000000000000000015f9c07c104765a40014bd34a890807136fdce9fc2a7a76fa75b1ac82f19c96030000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc0000000000000000000000000000000000000000000000000000000000001388"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}
