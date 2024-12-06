package testdata

import (
	"fmt"
	"math/big"

	opbindings "github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

type ProofType string

const (
	ZkVMType         ProofType = "zkVM"
	ZkEVMType        ProofType = "zkEVM"
	DefaultProofType           = ZkEVMType
)

var ProofTypes = []ProofType{
	ZkVMType,
	ZkEVMType,
}

func ValidProofType(value ProofType) bool {
	for _, k := range ProofTypes {
		if k == value {
			return true
		}
	}
	return false
}

var (
	TargetBlockNumber = uint64(21)

	// for zkVM challenge
	ZkVMVKeyHash     = common.HexToHash("0x6c15e3bb696329c15f6b963e40ac4c3841e726ef6fcaea042daf4b3d056e8d2f")
	ZkVMWitness      = "0x"
	ZkVMProof        = []byte("0x097da0323fa72c64d59682d6991bab9a50c996336d5fd070f0d56ef3ecb9a5d00d489a87499d9b0026b1fecda38a6edd41d290ffd053a1abb0747b0dbb5b66ff1d35d77f6bcf425a06872e78d8c544966b577e1cbd63975a8d371766ec0b642b279a6d15b4d7b8dd3b47da0f02122e475d6b6a8f97422d56f18acc2184f245e50eae81146701d25724a5c2b9baf82ae93ca264d11c4a481dba73886b01c642932557ed9e6845c6e16562ba4521add47de2c5aa811502c73e8735c874c4db4ba90df225349e121ce92fe27aebb8e236ff905b9911df6f67b52c82a34fe8933c9b053cb8567250f35115a32ae989650483221ce6a4490f092f9543f30a89f065c3116de44ce1a7ca5f947b2557373d893c6989c7b5c2bc42be86f8b059dc6a5fa106bd0c0abd63686b42f3706a366f270767b050e1ba3ec5a454b520023fd40cd60fe55fe86edeea81ff5b7d32347b9c086e969db21bc20c635ca3f6462b52d65f0fc8f4b067b8d4ec08fc2bbf1c23bf0160488d3a828b73acf294b4d42cc1cf4112ca129ba4dca38719090f2be26e2af92380b3ddb339a50fd42fb66e9296c7471ee0f694e35863ed363a77c1ca0e0b3fe1c57f566d36542cb2ef47d3b1eb3daf25ebed15fd6a0adaaf2f37c105b5298d9113d5b2ba641d82bfad0ec4cac0827b1a2c814d228825f742cd68bcec14d229025f551a812af3657b69ecae33f360d00000000728afb9bcba40e0fbc85d775ae7ea50aff5f22ee0e77b7655e732e7a2ee068e1d1c3e5dd8edaa6fc4cd247b4589c8b94c3dac94e2f7a72888ca531087666a4b702eaae316ed0d4df9b072fe8c7392814a25eea4b2c4922d51fb0a6cf64d64968e1e25234f6c4639dbdd3d760ee52c7af731ca9faa84edff95112c335e02a4212d26b444f87cd3f4f831496168d34df8f467ba62b9593672ae25edfeb08b366c24192e1dc95648187b63d52615c37bb108358c6c49ce51bf3e2b77bb15b72efbb6283e36d4e0b3a80f18281201ab420a5a81efc7af8e3cdfb64c693b2ddedb94a21a7462dc9ec33d761306721c1f8d5e8f4b517724a01673bba4853d54591c6d8213b2d1456728311ec65f4f13f40c383f8e72f7890954051eccd05bc3c1fe44e6224c2e29bda437a135eba9ed261aa9a09ccb612dcf5c87dad3e60b3cdbb118d9000000011ae99cd49723c91c7c9001f4256a2c378b1896d5bdfb32f08ce2b329a5587ef720d4ab2bb8e82786d89e196bfe3043c9ffd38e1c622e3d41061c3ad7347dece5")
	ZkVMPublicVaules = []byte("0x2000000000000000ee40a952b20e40f8911e64179baf6afe0cac0e6226f18719560efa55ee51e701200000000000000048f9fe8c248855f9da9e8f6e06f3215f57a6a96f684b5fc63232b565bdb98479200000000000000042c0d60066fbd229758f8deaee337afc6cd0a75ddf120896258a4fd846aafbfd")

	zkVMPrevBlockhash             = common.HexToHash("0x")
	zkVMPrevStateRoot             = common.HexToHash("0x")
	zkVMPrevWithdrawalStorageRoot = common.HexToHash("0x")

	zkVMTargetBlockhash             = common.HexToHash("0x")
	zkVMTargetStateRoot             = common.HexToHash("0x")
	zkVMTargetWithdrawalStorageRoot = common.HexToHash("0x")

	ChallengeL1Head = common.HexToHash("0x")

	// for zkEVM challenge
	coinbaseAddr   = common.HexToAddress("0x0000000000000000000000000000000000000000")
	zeroUint64     = hexutil.MustDecodeUint64("0x0")
	zeroBloom      = types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	emptyExtraData = hexutil.MustDecode("0x")

	zkEVMPrevBlockhash             = common.HexToHash("0x3392758b5bca8b8319df6180c145ca28152f1b6a3af977bc48ec67d2259dbcd2")
	zkEVMPrevStateRoot             = common.HexToHash("0x263975548df46f3ffc739f602b503f32b4c522026c8c93204929ddd5b65ad202")
	zkEVMPrevWithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")

	zkEVMTargetBlockhash             = common.HexToHash("0x4ecf76378ef03e3a417ac169cb052a879424345c59765aca05fe1fb6259375a9")
	zkEVMTargetStateRoot             = common.HexToHash("0x0475b3d38492c9e58190616eaad4ab033942aa55747d49c5a614b9e751998d5e")
	zkEVMTargetWithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")

	l2toL1MesasgePasserBalance  = common.Big0
	l2toL1MesasgePasserCodeHash = common.HexToHash("0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35")
	merkleProof                 = []hexutil.Bytes{
		hexutil.MustDecode("0x0027e039ebdf0f9e7c8a1481ebf7448aae44afec16b045969976d37555b364f6132d6654101a6881a48a968d5a257cc8dc8d980c0cc55d58e47833222a24b2230a"),
		hexutil.MustDecode("0x002023945e0a0e2290059ca87427aa00e3d515d2e9144d8d3e69ee13b7c75615482cf819ee96fa5bf76e6ce7712caf6673df7fc3af269cfe193bed443eeb89527a"),
		hexutil.MustDecode("0x0021d5edc847df9d20356a576c7a4b4b1bb992b427aa4e01412eb08864416db55b1c7c9c38ad18cc007a6e522b638fdf6010aef11727fcb7cdcbfb5edb96e3a74a"),
		hexutil.MustDecode("0x0020140f1c792e0ff14d98ca78ed1c085501566b9917cc0792285ed265e6f8cd5b1e99968ab013de794c675a5f614b8a642294aa47612f1bb6173910a5cd7f0b16"),
		hexutil.MustDecode("0x002b11bc4ac76ee779a652e7e93a9871fa88e780285d41c6ed2f1646ec8658b7450c158053a97ff3e327cfae62631e5878accd83e5803f044c90803d4579ae21ba"),
		hexutil.MustDecode("0x00244f31e0a770ed9dbdf583a0e0dbe036f5f9476c3cc761e18b2a475a96bcc10814c463c1a53d595a21f22f3c5cafe495329c68087ab2601468f8fa7e11d88bf7"),
		hexutil.MustDecode("0x0008f993e0df87a04e71f72a3a237c909e2381eacf176b279925d333ee7b1e36ed03c30671a87c81a313a035fdbe052cc592ae0a604bcf87a5cf163d5a43104574"),
		hexutil.MustDecode("0x00202aa398b4bd976d7c165b2c7bfa6e4b695f18ff78e4cb544dbb0fbab8c6537e15c1089ff56ee758ec382a55a21b76c08ebf64d6a78d7e6ff442793536607510"),
		hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000002e957b48192277673a8dde0549358d09d3f9ad6e8db14e08f4fed46f96021a74"),
		hexutil.MustDecode("0x002325a334c56feef28306cece9e3867165ee117aab1c831fe04db286a1b4ff2c80000000000000000000000000000000000000000000000000000000000000000"),
		hexutil.MustDecode("0x00218476186a36a2ddf003ef59459478f44e0cea1ac32870dafca118331259b05f23618448c7fab9e44d30c44be6777aa390c25ad138fde11d22bdefd05f43838b"),
		hexutil.MustDecode("0x012de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c304040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc3524f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127204200000000000000000000000000000000000003000000000000000000000000"),
		hexutil.MustDecode("0x5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449"),
	}
)

func SetPrevOutputResponse(o *eth.OutputResponse, proofType ProofType) error {
	switch proofType {
	case ZkVMType:
		outputRoot, err := rollup.ComputeL2OutputRoot(&opbindings.TypesOutputRootProof{
			Version:                  eth.OutputVersionV0,
			StateRoot:                zkVMPrevStateRoot,
			MessagePasserStorageRoot: zkVMPrevWithdrawalStorageRoot,
			LatestBlockhash:          zkVMTargetBlockhash,
		})
		if err != nil {
			return fmt.Errorf("mocking error: %w", err)
		}

		o.OutputRoot = outputRoot
		o.WithdrawalStorageRoot = zkVMPrevWithdrawalStorageRoot
		o.StateRoot = zkVMPrevStateRoot
		o.BlockRef = eth.L2BlockRef{
			Number:     TargetBlockNumber - 1,
			Hash:       zkVMPrevBlockhash,
			ParentHash: common.HexToHash("0x"),
		}
	case ZkEVMType:
		outputRoot, err := rollup.ComputeKromaL2OutputRoot(&bindings.TypesOutputRootProof{
			Version:                  eth.OutputVersionV0,
			StateRoot:                zkEVMPrevStateRoot,
			MessagePasserStorageRoot: zkEVMPrevWithdrawalStorageRoot,
			LatestBlockhash:          zkEVMPrevBlockhash,
			NextBlockHash:            zkEVMTargetBlockhash,
		})
		if err != nil {
			return fmt.Errorf("mocking error: %w", err)
		}

		o.OutputRoot = outputRoot
		o.WithdrawalStorageRoot = zkEVMPrevWithdrawalStorageRoot
		o.StateRoot = zkEVMPrevStateRoot
		o.BlockRef = eth.L2BlockRef{
			Number:     TargetBlockNumber - 1,
			Hash:       zkEVMPrevBlockhash,
			ParentHash: common.HexToHash("0x6fd96c96f5ca016848315ed6b2358ba125472019106a4f78ca92745b17f84c34"),
		}
		o.NextBlockRef = eth.L2BlockRef{Hash: zkEVMTargetBlockhash}
	default:
		return fmt.Errorf("unexpected challenge proof type: %s", proofType)
	}

	return nil
}

func SetPrevOutputWithProofResponse(o *eth.OutputWithProofResponse) error {
	err := SetPrevOutputResponse(&o.OutputResponse, ZkEVMType)
	if err != nil {
		return err
	}

	o.PublicInputProof = &eth.PublicInputProof{}
	parentBeaconRoot := common.HexToHash("0x3eeb016384502029f0dc9cc6188d4e5ca8b6547f755b7cfa3749d7512f98c41b")
	o.PublicInputProof.NextBlock = &types.Header{
		ParentHash:       zkEVMPrevBlockhash,
		Coinbase:         coinbaseAddr,
		Difficulty:       common.Big0,
		Root:             zkEVMTargetStateRoot,
		TxHash:           common.HexToHash("0xb01bdd77e9685c8341733f345113daa34c25a63a37cb76b81de492b36b0cc806"),
		ReceiptHash:      common.HexToHash("0x886c02379eee108cab1ada4055c4f82b048b7e3bbce0d82bf532c95409d8ad81"),
		Bloom:            zeroBloom,
		Number:           big.NewInt(int64(TargetBlockNumber)),
		GasLimit:         hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:          hexutil.MustDecodeUint64("0xc9f4"),
		Time:             hexutil.MustDecodeUint64("0x66471e21"),
		Extra:            emptyExtraData,
		MixDigest:        common.HexToHash("0x8bb2786563ea29f638e4e9758d9886e8a1af5b4f7688f4ee6622a6b53df87742"),
		BaseFee:          big.NewInt(int64(hexutil.MustDecodeUint64("0x1"))),
		WithdrawalsHash:  &types.EmptyWithdrawalsHash,
		BlobGasUsed:      &zeroUint64,
		ExcessBlobGas:    &zeroUint64,
		ParentBeaconRoot: &parentBeaconRoot,
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	o.PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.KromaDepositTx{
			SourceHash: common.HexToHash("0x81e84a0b340571d1b0ef61008afa413d4dc9b50884003177f02294cb961b7503"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Mint:       nil,
			Value:      common.Big0,
			Gas:        1000000,
			Data:       hexutil.MustDecode("0x440a5e20000f42400000000000000000000000000000000066471e1e000000000000000d000000000000000000000000000000000000000000000000000000000a83a7d000000000000000000000000000000000000000000000000000000000000000011b075cc318f7c80c85e6aee0d9da5d63dfb91d915d7be13aa2f23ec6b0c380500000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc0000000000000000000000000000000000000000000000000000000000001388"),
		}),
	}
	o.PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	o.PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	o.PublicInputProof.MerkleProof = merkleProof

	return nil
}

func SetTargetOutputResponse(o *eth.OutputResponse, proofType ProofType) error {
	switch proofType {
	case ZkVMType:
		outputRoot, err := rollup.ComputeL2OutputRoot(&opbindings.TypesOutputRootProof{
			Version:                  eth.OutputVersionV0,
			StateRoot:                zkVMTargetStateRoot,
			MessagePasserStorageRoot: zkVMTargetWithdrawalStorageRoot,
			LatestBlockhash:          zkVMTargetBlockhash,
		})
		if err != nil {
			return fmt.Errorf("mocking error: %w", err)
		}

		o.OutputRoot = outputRoot
		o.WithdrawalStorageRoot = zkVMTargetWithdrawalStorageRoot
		o.StateRoot = zkVMTargetStateRoot
		o.BlockRef = eth.L2BlockRef{
			Number:     TargetBlockNumber,
			Hash:       zkVMTargetBlockhash,
			ParentHash: zkVMPrevBlockhash,
		}
	case ZkEVMType:
		nextBlockhash := common.HexToHash("0x6c4e19b1fc27f6a075c67f35bd15b21c40025a892e32cdb8d9b5f5d5ec60093a")
		outputRoot, err := rollup.ComputeKromaL2OutputRoot(&bindings.TypesOutputRootProof{
			Version:                  eth.OutputVersionV0,
			StateRoot:                zkEVMTargetStateRoot,
			MessagePasserStorageRoot: zkEVMTargetWithdrawalStorageRoot,
			LatestBlockhash:          zkEVMTargetBlockhash,
			NextBlockHash:            nextBlockhash,
		})
		if err != nil {
			return fmt.Errorf("mocking error: %w", err)
		}

		o.OutputRoot = outputRoot
		o.WithdrawalStorageRoot = zkEVMTargetWithdrawalStorageRoot
		o.StateRoot = zkEVMTargetStateRoot
		o.BlockRef = eth.L2BlockRef{
			Number:     TargetBlockNumber,
			Hash:       zkEVMTargetBlockhash,
			ParentHash: zkEVMPrevBlockhash,
		}
		o.NextBlockRef = eth.L2BlockRef{Hash: nextBlockhash}
	default:
		return fmt.Errorf("unexpected challenge proof type: %s", proofType)
	}

	return nil
}

func SetTargetOutputWithProofResponse(o *eth.OutputWithProofResponse) error {
	err := SetTargetOutputResponse(&o.OutputResponse, ZkEVMType)
	if err != nil {
		return err
	}

	o.PublicInputProof = &eth.PublicInputProof{}
	o.PublicInputProof.NextBlock = &types.Header{
		ParentHash:  zkEVMTargetBlockhash,
		Coinbase:    coinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x1f5234e71e92fda7263874df353f6445195d58af33cb15bcaa6a37b0df779e4f"),
		TxHash:      common.HexToHash("0x51a7beddaa794ab6fec8312267345c5fc694d13a72b509d30765aadc13475cde"),
		ReceiptHash: common.HexToHash("0xc8d04bf464c80f34cb0083628e8d15b89cf93fe4515276a7af8b2b043bd3f6b9"),
		Bloom:       zeroBloom,
		Number:      big.NewInt(int64(TargetBlockNumber + 1)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xb420"),
		Time:        hexutil.MustDecodeUint64("0x66471e23"),
		Extra:       emptyExtraData,
		MixDigest:   common.HexToHash("0x8bb2786563ea29f638e4e9758d9886e8a1af5b4f7688f4ee6622a6b53df87742"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x1"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	o.PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.KromaDepositTx{
			SourceHash: common.HexToHash("0x40bcbd870b6c68f66e727762654cf39e1491b20be579a13d231a6a6d21f204ce"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Gas:        1000000,
			Value:      common.Big0,
			Mint:       nil,
			Data:       hexutil.MustDecode("0x440a5e20000f42400000000000000000000000010000000066471e1e000000000000000d000000000000000000000000000000000000000000000000000000000a83a7d000000000000000000000000000000000000000000000000000000000000000011b075cc318f7c80c85e6aee0d9da5d63dfb91d915d7be13aa2f23ec6b0c380500000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc0000000000000000000000000000000000000000000000000000000000001388"),
		}),
	}
	o.PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	o.PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	o.PublicInputProof.MerkleProof = merkleProof

	return nil
}
