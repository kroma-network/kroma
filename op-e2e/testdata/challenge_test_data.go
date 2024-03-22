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
		hexutil.MustDecode("0x0030220c6ed9d2841133c2029218be8227e5dd0b53b9d7064c45ed0697829646d411b514eaffb2f105f410ffb5e9c21c994970f1172534da0b1c8233a6185ce3ae"),
		hexutil.MustDecode("0x001c9d0f9f253060bc73748c7019b2eeddf03c7aca6cfa49f381430e6b0cbca1450f429664ad53b5c178a72555db864fecb17070d92757845ed9df8d16c321da09"),
		hexutil.MustDecode("0x001561732813f0e977a7c800a0123a7cbdb83af9f152a4c524a88d45856e56a672194776a139fd44f2f6e090cf3c52536a816df6278bfb2aa7c832c86a3e234718"),
		hexutil.MustDecode("0x0002b0de91c35535a813a2a17cee58400ae8a81228babcd201691c7532365e2a522250889aa2466ea66c193efa77c5003f45124aed019dbc6035f560e1bf85c0fb"),
		hexutil.MustDecode("0x002b11bc4ac76ee779a652e7e93a9871fa88e780285d41c6ed2f1646ec8658b74527ac23d456df5fdd5f0a195c660194fad8eed5a93d1b00da54cfbc28cf0c0636"),
		hexutil.MustDecode("0x00244f31e0a770ed9dbdf583a0e0dbe036f5f9476c3cc761e18b2a475a96bcc10814c463c1a53d595a21f22f3c5cafe495329c68087ab2601468f8fa7e11d88bf7"),
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
		Version:                  eth.OutputVersionV0,
		StateRoot:                common.HexToHash("0x06e23d47dea22feb9523e2817c42de9f05cbc9ce1410ef45bce2dcda2aab7bd6"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0x936a872193703be4201cffc531b00fd6e8c71a092c0e7ea51ab0f8ece00838a4"),
		NextBlockHash:            common.HexToHash("0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x06e23d47dea22feb9523e2817c42de9f05cbc9ce1410ef45bce2dcda2aab7bd6")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber - 1,
		Hash:       common.HexToHash("0x936a872193703be4201cffc531b00fd6e8c71a092c0e7ea51ab0f8ece00838a4"),
		ParentHash: common.HexToHash("0xfe22df475b9806b22e1efe68909a8388a1a25c1d0e9dd0c82f60f2868b979f9c"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	withdrawalHash := common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:      common.HexToHash("0x936a872193703be4201cffc531b00fd6e8c71a092c0e7ea51ab0f8ece00838a4"),
		Coinbase:        CoinbaseAddr,
		Difficulty:      common.Big0,
		Root:            common.HexToHash("0x245d2fa3cbb92f5223e06b23a4c53fc1e1b9cd8edea6769e6749e72c8ff2d384"),
		TxHash:          common.HexToHash("0xc9862b8f3157beea2d258b0a36b836794ea58302951468ab3d0b0885e63bccbf"),
		ReceiptHash:     common.HexToHash("0x1ce54d356f9c39a93aac476d492b27c4c23374e203d9902becd4f4b6826b9b35"),
		Bloom:           types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:          big.NewInt(int64(TargetBlockNumber)),
		GasLimit:        hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:         hexutil.MustDecodeUint64("0xc0b1"),
		Time:            hexutil.MustDecodeUint64("0x65f2c15d"),
		Extra:           hexutil.MustDecode("0x"),
		MixDigest:       common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:         big.NewInt(int64(hexutil.MustDecodeUint64("0x1"))),
		WithdrawalsHash: &withdrawalHash,
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x71eca7451c17ebaa0a17f3efbdc726a214c6a2943ba5883b7c8dfdbcfc70276b"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Mint:       nil,
			Value:      common.Big0,
			Gas:        1000000,
			Data:       hexutil.MustDecode("0xefc674eb00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000065f2c133000000000000000000000000000000000000000000000000000000003b9aca0053c8e248a92c83b511a0057d0926d2fd02652e5bcf6c49df450a3d0ad081400800000000000000000000000000000000000000000000000000000000000000150000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000083400000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000001388"),
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
		StateRoot:                common.HexToHash("0x245d2fa3cbb92f5223e06b23a4c53fc1e1b9cd8edea6769e6749e72c8ff2d384"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96"),
		NextBlockHash:            common.HexToHash("0x2afbf92eea49c24a7e50934e7b05bf21420e5e156bae67b82b41a83423b9dd5d"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x245d2fa3cbb92f5223e06b23a4c53fc1e1b9cd8edea6769e6749e72c8ff2d384")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber,
		Hash:       common.HexToHash("0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96"),
		ParentHash: common.HexToHash("0x936a872193703be4201cffc531b00fd6e8c71a092c0e7ea51ab0f8ece00838a4"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0x2afbf92eea49c24a7e50934e7b05bf21420e5e156bae67b82b41a83423b9dd5d")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x014aae1d03d80de4b524dc0c82a9abb10c0a98c24152c2aaabf3f4c04b1a8c56"),
		TxHash:      common.HexToHash("0xd2f948867c819877a219f9bb39c1b6bcc7967886f2da506980308c7f039d6960"),
		ReceiptHash: common.HexToHash("0x56da5bfe76cb7dd3727c3adfed4c812dc08f3281e623194a2384caf45e4a6a7a"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber + 1)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xc0b1"),
		Time:        hexutil.MustDecodeUint64("0x65f2c15f"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x1"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0xddf44f388cb348d3bc6da77f3ef8b9dfa1393a8821c4a292fdf34c8d18237574"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Gas:        1000000,
			Value:      common.Big0,
			Mint:       nil,
			Data:       hexutil.MustDecode("0xefc674eb00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000065f2c133000000000000000000000000000000000000000000000000000000003b9aca0053c8e248a92c83b511a0057d0926d2fd02652e5bcf6c49df450a3d0ad081400800000000000000000000000000000000000000000000000000000000000000160000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000083400000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000001388"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}
