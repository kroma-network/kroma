package testdata

import (
	"fmt"
	"math/big"

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
	ZkVMVKeyHash     = common.HexToHash("0x00ad7cc293beb89796e8a9713ab7189b33faf6975325da0707d2d9cb129db055")
	ZkVMProof        = hexutil.MustDecode("0x54bdcae3256eaae3a90569ad01460800c08e2e06c3855ced450de49dd34e32854eb67a8816b72b8294ea24733d2b47253edf6796321dc105f65d4509ee27dcdf849c453327152d4769901e03502fb829cfb6712c6f6e1088feb007a2cea4d8ae37a055671a40948624b98a91069134ea8ea6226e2f6a1457abec73c7d54c3f84eae953650fff44bd6eb127f188b42cf0b4552d9332ab154323b704fce64a396882f54dab173beedccb24459be8ba08df8b8682b90bed088ca80581e6280ab0299e26a5440c0dea757e11484d1bfb9cc54cb5b999b3d9b82ac4106ad0beb705d2a40467061ec853e02449daa488a94aae8c9d1147065857a08c80ced9a04decd9d5f96d64001c6675f4f317af95829e40e3274101e9baa1e539111a0654c3628da41cc4610dd0dde84ff63c37fd213cb6ae5a86e22c0a088249aa3752ee4beb023777e02802629c8c5bfb2197357eb1566a77f4169177c826ba62d784693318795cea293e048bebd2ac5e88b24ff21e8547a3edf369173a0c3106ee2fef2d807103f9a1c91d7ef6307a47781f8ec796db6f3e1e6256b14aefb0be95d15818eaea4fa6b8ab2b2605cc47f8010abd47fabc7be65ea31096f26cbdc6ac2fd7fd8af7be18fa9e0a3500fb95305139353357d88322ec0d64026fe5a9b5c5f8a124de743207eb2a1dcf72d0168eae8a898ce3fec30a266d95ab8a8359547e7c56894f89f940d90e097cfe99058f651935fbe199211af862819d46cbb1722643447f9e5100714ddf2c3805f779b66f864138e7596ec5a224816bb75cb6a269387cb10bf269229d301fff5a2f5ef5191d95b89be192e5295bb20eb7f2b81bb48513cc2e1f389686da0eb207a4d13446c0b1482f1b1d1ff61fde440f4d194b2f1b9a4ea1bc2c394bc72be7e54c5ff0fff46f22939d5bb4f5e8e12683cbf26915f030efa53541e61583081eef0c4eef42367c3f341f6283fe706def1fd0b918c440144a12cb30fea3620db53fda424496de0e44526cb992e2071e8a4d7c3c427dfe45a1ad2ed61fac6b1d353e45fb0bb3371366931164f97d439a05a188e38673f2c9aba54e4677272c085ab877f90baf7e56a92c2ab2447d75b662fc04cccf3f9716dfaf575662e2a003d77395c3c8a90706526d8ed1dbfde71dde5739670a8bf30fecc98f6b2929fa1876db545698cfc3233f1101212ebe78c6db2460071fcce10504c35ff472738e")
	ZkVMPublicValues = hexutil.MustDecode("0x20000000000000008906045a3f098db0d7a5f7d9400e1299718fc9563b30417e887cf0c91a375fea200000000000000053ae7b0c4ef070efb9bec7ae130edce19797b2283f020e085a69796ad43b6c512000000000000000c5f36b4618de22a58113263406f398342a7624d832fda0ccd05ab61b0ceb5352")
	ZkVMWitness      = "0x"

	zkVMPrevOutput   = common.HexToHash("0x8906045a3f098db0d7a5f7d9400e1299718fc9563b30417e887cf0c91a375fea")
	zkVMTargetOutput = common.HexToHash("0x53ae7b0c4ef070efb9bec7ae130edce19797b2283f020e085a69796ad43b6c51")

	ChallengeL1Head = common.HexToHash("0xc5f36b4618de22a58113263406f398342a7624d832fda0ccd05ab61b0ceb5352")

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
		o.OutputRoot = eth.Bytes32(zkVMPrevOutput)
		// Other values than output root won't be used in zkVM challenge, so we don't have to mock them
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
		o.OutputRoot = eth.Bytes32(zkVMTargetOutput)
		// Other values than output root won't be used in zkVM challenge, so we don't have to mock them
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
