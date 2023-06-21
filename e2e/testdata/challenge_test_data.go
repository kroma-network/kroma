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
	TargetBlockNumber           = uint64(17)
	CoinbaseAddr                = common.HexToAddress("0x0000000000000000000000000000000000000000")
	l2toL1MesasgePasserBalance  = common.Big0
	l2toL1MesasgePasserCodeHash = common.HexToHash("0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35")
	merkleProof                 = []hexutil.Bytes{
		hexutil.MustDecode("0x000aa0439c4396244e32a290420ad3cc63fcd7492387b6b4b50a151a9c3be187142190905b2b91df9e3921f3cb4a75059a2e7f071b0a036ca3bf5cebca5f1cb420"),
		hexutil.MustDecode("0x0002b8ac403422ca5bf986d532f2a36bf0ed895ac42b0bd739f427d1cdf61745860a428d32b9078aebb23b8e08aed0f5393f8ce5d9050945cdcbde3e9aa754e218"),
		hexutil.MustDecode("0x001fae62daabf99b5ab4efbd7504f308c79180e9bb04448eb51b69951646f4e8e51d0cb347c4624112228de283bda8b87d21bdb2a743cc8d7e2771d1a20afd40d4"),
		hexutil.MustDecode("0x002bf1dc335fbd3c6252fa426d647d20ff86181a1a3ac1f4c53122b5a853436a592a82b9a18c350c5571c5461d90d2e37d3751a17c18bd3c90a701fada42a7446b"),
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
	outputRoot, err := rollup.ComputeL2OutputRootV1(&bindings.TypesOutputRootProof{
		Version:                  rollup.V1,
		StateRoot:                common.HexToHash("0x2fbc5f620a2e28afd9b159b9b7b259ebcca31cea10e3424994a39c1d6d551c18"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
		NextBlockHash:            common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x2fbc5f620a2e28afd9b159b9b7b259ebcca31cea10e3424994a39c1d6d551c18")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber - 1,
		Hash:       common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
		ParentHash: common.HexToHash("0xe6685c9bfcfb882a0353d2949208db791d60de33fa145d461df870a6a85a2353"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x008bc0b55af382fe6508389a13c011b873eb5d7dcbcd3f8ba8f9a2001c196205"),
		TxHash:      common.HexToHash("0x984e84df9b4f0573f6381e05f5445f0e7d82cd95d3a2b916a58369898b82eef3"),
		ReceiptHash: common.HexToHash("0x7298f243c2bd5472fb20330a6ba763aaf9dd07d8b0f818b27277c6b0fd3b85a9"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xc0b1"),
		Time:        hexutil.MustDecodeUint64("0x6489d44b"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x634d0f4"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0xcf6f3fc7fa423808b976c4593f2ada10add979575142cc6a8529870b33097568"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Mint:       nil,
			Value:      common.Big0,
			Gas:        1000000,
			Data:       hexutil.MustDecode("0xefc674eb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006489d429000000000000000000000000000000000000000000000000000000003b9aca00bd1dca8d75e193d89633a0030d5aabd082ff3537d8c3caa57b077c7df1a354b000000000000000000000000000000000000000000000000000000000000000110000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000083400000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000007d0"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}

func SetTargetOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRootV1(&bindings.TypesOutputRootProof{
		Version:                  rollup.V1,
		StateRoot:                common.HexToHash("0x008bc0b55af382fe6508389a13c011b873eb5d7dcbcd3f8ba8f9a2001c196205"),
		MessagePasserStorageRoot: common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127"),
		BlockHash:                common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
		NextBlockHash:            common.HexToHash("0x76da1f020b58638d4f0199a96c46a4c493366c61bbb1dbd3a4ad60166006d6cd"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127")
	(*s).StateRoot = common.HexToHash("0x008bc0b55af382fe6508389a13c011b873eb5d7dcbcd3f8ba8f9a2001c196205")
	(*s).BlockRef = eth.L2BlockRef{
		Number:     TargetBlockNumber,
		Hash:       common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
		ParentHash: common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0x76da1f020b58638d4f0199a96c46a4c493366c61bbb1dbd3a4ad60166006d6cd")}
	(*s).PublicInputProof = &eth.PublicInputProof{}
	(*s).PublicInputProof.NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x141a7dd8d9e4daf81d1ebd599a306bd302e67f58cf92610525c95f13be8409be"),
		TxHash:      common.HexToHash("0xac4aa619e9e046033d4a727ac0748fedb62b58bfcd682c3b0c136220dcf8b251"),
		ReceiptHash: common.HexToHash("0x7298f243c2bd5472fb20330a6ba763aaf9dd07d8b0f818b27277c6b0fd3b85a9"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(int64(TargetBlockNumber + 1)),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xc0b1"),
		Time:        hexutil.MustDecodeUint64("0x6489d44d"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x56ede09"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).PublicInputProof.NextTransactions = types.Transactions{
		types.NewTx(&types.DepositTx{
			SourceHash: common.HexToHash("0x26954d5f47e27b59dcab1e0758eacba5cd5220b48ded3c6c0d68252f13479d85"),
			From:       common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001"),
			To:         &toAddr,
			Gas:        1000000,
			Value:      common.Big0,
			Mint:       nil,
			Data:       hexutil.MustDecode("0xefc674eb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006489d429000000000000000000000000000000000000000000000000000000003b9aca00bd1dca8d75e193d89633a0030d5aabd082ff3537d8c3caa57b077c7df1a354b000000000000000000000000000000000000000000000000000000000000000120000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000083400000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000007d0"),
		}),
	}
	(*s).PublicInputProof.L2ToL1MessagePasserBalance = l2toL1MesasgePasserBalance
	(*s).PublicInputProof.L2ToL1MessagePasserCodeHash = l2toL1MesasgePasserCodeHash
	(*s).PublicInputProof.MerkleProof = merkleProof

	return nil
}
