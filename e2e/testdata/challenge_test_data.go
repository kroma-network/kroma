package testdata

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup"
)

var (
	TargetBlockNumber           = uint64(21)
	CoinbaseAddr                = common.HexToAddress("0x0000000000000000000000000000000000000000")
	l2toL1MesasgePasserBalance  = common.Big0
	l2toL1MesasgePasserCodeHash = common.HexToHash("0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35")
	merkleProof                 = []hexutil.Bytes{
		hexutil.MustDecode("0x0011608e4c1635552de6d1ce4ae41799e32c2d3baa13e05017029c8db503db4e6b0a4478c7f4040b98b10cd90a9e4a6c90f951dc8b65b73845471c49835b55415b"),
		hexutil.MustDecode("0x0002d35b1fb18e43d4cb3883a6a64cb882312db9ee0868c8e425919c2acc67b6fe2862cefc25230c03bcbd86bdb58b6f7662506e5d9af265239f2570df170089a9"),
		hexutil.MustDecode("0x002b189562f03c7653f51509a7840ba909c7bdbbab60778457f01c66649034f87f14c3c9034672c8e4adc74d672f746d96f24d703ab201f95b16c18e120caeb722"),
		hexutil.MustDecode("0x002bf1dc335fbd3c6252fa426d647d20ff86181a1a3ac1f4c53122b5a853436a592f737a81aa1525e0e68bd3e34dbc70f09386352308f59fc5c557ad6f0d389663"),
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
		StateRoot:                common.HexToHash("0x07c36d720614ac2c4b29f4d3f60d9c2ae0c0b38dada03ae98bbdcbe90ecae32c"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0x016fb906ed7974bb75ac39fafb45bf76b14732dfb3a41d12ac04d417c348b376"),
		NextBlockHash:            common.HexToHash("0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x07c36d720614ac2c4b29f4d3f60d9c2ae0c0b38dada03ae98bbdcbe90ecae32c")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber - 1,
		Hash:       common.HexToHash("0x016fb906ed7974bb75ac39fafb45bf76b14732dfb3a41d12ac04d417c348b376"),
		ParentHash: common.HexToHash("0xc1ec84c7d6336bbf8991d62520e4af9beb3071b51e11d582398879972d24a650"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x016fb906ed7974bb75ac39fafb45bf76b14732dfb3a41d12ac04d417c348b376"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x0127ca9f2fc205f1d96450f463e2d9765bb01a1232191773a4e642bebf7fc25e"),
		TxHash:      common.HexToHash("0x90572ca3650e693acd8ec001d84c9c8c3122de052c64e0915cbac2a2f2727a78"),
		ReceiptHash: common.HexToHash("0xcffb7b9369b08f160e112c1d468ecdc8d718f339b38e03cc9f755580945badd6"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0x21ebd"),
		Time:        hexutil.MustDecodeUint64("0x64b7980d"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x3b28afc"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	toAddr2 := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	toAddr3 := common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8")
	toAddr4 := common.HexToAddress("0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc")
	toAddr5 := common.HexToAddress("0x90f79bf6eb2c4f870365e785982e1f101e93b906")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x16007f53726f3adaf12fd069faba06a078deb9b7e35b2a61316cf86523851504"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Mint:       nil,
			Value:      common.Big0,
			Gas:        1000000,
			Data:       hexutil.MustDecode("0xefc674eb000000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000064b7980b000000000000000000000000000000000000000000000000000000000dce13ff4d71108833a096ef91cad4f118d8ca020476be6632931b4447db782e65a657e800000000000000000000000000000000000000000000000000000000000000000000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000083400000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000007d0"),
		}),
		types.NewTx(&types.LegacyTx{
			Nonce:    3,
			GasPrice: hexutil.MustDecodeBig("0x40600f1f"),
			Gas:      21000,
			To:       &toAddr2,
			Value:    hexutil.MustDecodeBig("0x1bc16d674ec80000"),
			Data:     hexutil.MustDecode("0x"),
			V:        hexutil.MustDecodeBig("0x72e"),
			R:        hexutil.MustDecodeBig("0xadf38a3b16751737ba9752760969bb63be6c49eec1024c16b0ce8e7fb44f164f"),
			S:        hexutil.MustDecodeBig("0x4c25ae5e06e5025982cd5b737c0dd14947891d3d7fbb63b9a71abe687ab6772c"),
		}),
		types.NewTx(&types.LegacyTx{
			Nonce:    4,
			GasPrice: hexutil.MustDecodeBig("0x40600f1f"),
			Gas:      21000,
			To:       &toAddr3,
			Value:    hexutil.MustDecodeBig("0x1bc16d674ec80000"),
			Data:     hexutil.MustDecode("0x"),
			V:        hexutil.MustDecodeBig("0x72e"),
			R:        hexutil.MustDecodeBig("0x5c6ed9e164ad6b5878298a3d1759e45a9277a1f22154a64db137714d996b4865"),
			S:        hexutil.MustDecodeBig("0x2760633637cc25dbd988b5be0a2533704c51a1960821ec009d9c2fef03dbc9da"),
		}),
		types.NewTx(&types.LegacyTx{
			Nonce:    5,
			GasPrice: hexutil.MustDecodeBig("0x40600f1f"),
			Gas:      21000,
			To:       &toAddr4,
			Value:    hexutil.MustDecodeBig("0x1bc16d674ec80000"),
			Data:     hexutil.MustDecode("0x"),
			V:        hexutil.MustDecodeBig("0x72e"),
			R:        hexutil.MustDecodeBig("0x34dad3b704cf7caf64dbc0efef8ca8f370108acd19f882c3f9a6823d07deab8"),
			S:        hexutil.MustDecodeBig("0x7390e43b6faef69a91d28de658d87ac0d5ebd26ce46bd25e3111b9bf8a0ebbdc"),
		}),
		types.NewTx(&types.LegacyTx{
			Nonce:    6,
			GasPrice: hexutil.MustDecodeBig("0x40600f1f"),
			Gas:      21000,
			To:       &toAddr5,
			Value:    hexutil.MustDecodeBig("0x1bc16d674ec80000"),
			Data:     hexutil.MustDecode("0x"),
			V:        hexutil.MustDecodeBig("0x72e"),
			R:        hexutil.MustDecodeBig("0x73b0925e53c2bdf785ccd9cb730c8f93e5bda5fa197c5f20fc16adc1f80e60d2"),
			S:        hexutil.MustDecodeBig("0x4f8316e12b331d00f7f97e3a010863a3ae6961d4b7de58ceb3c7d3fe033c1809"),
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
		StateRoot:                common.HexToHash("0x0127ca9f2fc205f1d96450f463e2d9765bb01a1232191773a4e642bebf7fc25e"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55"),
		NextBlockHash:            common.HexToHash("0xd6fa98f354b0976938c8809e922115f9eaff9282045a56dd389fc3dbadda2ff6"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x0127ca9f2fc205f1d96450f463e2d9765bb01a1232191773a4e642bebf7fc25e")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber,
		Hash:       common.HexToHash("0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55"),
		ParentHash: common.HexToHash("0x016fb906ed7974bb75ac39fafb45bf76b14732dfb3a41d12ac04d417c348b376"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0xd6fa98f354b0976938c8809e922115f9eaff9282045a56dd389fc3dbadda2ff6")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x26e8cf51cb4f6029783af42e8b8ca0f8fdd3b3e8d110721b3926b1066782b6cd"),
		TxHash:      common.HexToHash("0xe738e1e439dc406c46c60dd0246a3fdfd13551cfcb0f6c6dd9b5b94511bb29df"),
		ReceiptHash: common.HexToHash("0x4eec2b7d51682df851cce0d4c3d37a6bd3fc141559b878096803f62098c627ca"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber + 1)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0x10395"),
		Time:        hexutil.MustDecodeUint64("0x64b7980f"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x33d522d"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x42e68c7b27556b36e6876700131d6e01d85e33143d22334312fc5ed121382f02"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Gas:        1000000,
			Value:      common.Big0,
			Mint:       nil,
			Data:       hexutil.MustDecode("0xefc674eb000000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000064b7980b000000000000000000000000000000000000000000000000000000000dce13ff4d71108833a096ef91cad4f118d8ca020476be6632931b4447db782e65a657e800000000000000000000000000000000000000000000000000000000000000010000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000083400000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000007d0"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}
