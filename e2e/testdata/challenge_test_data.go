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
	TargetBlockNumber = uint64(17)
	CoinbaseAddr      = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func SetPrevOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRootV1(&bindings.TypesOutputRootProof{
		Version:                  rollup.V1,
		StateRoot:                common.HexToHash("0x2fbc5f620a2e28afd9b159b9b7b259ebcca31cea10e3424994a39c1d6d551c18"),
		MessagePasserStorageRoot: common.HexToHash("0x"),
		BlockHash:                common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
		NextBlockHash:            common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x")
	(*s).StateRoot = common.HexToHash("0x2fbc5f620a2e28afd9b159b9b7b259ebcca31cea10e3424994a39c1d6d551c18")
	(*s).BlockRef = eth.L2BlockRef{
		Hash:       common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
		ParentHash: common.HexToHash("0xe6685c9bfcfb882a0353d2949208db791d60de33fa145d461df870a6a85a2353"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769")}
	(*s).NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x008bc0b55af382fe6508389a13c011b873eb5d7dcbcd3f8ba8f9a2001c196205"),
		TxHash:      common.HexToHash("0x984e84df9b4f0573f6381e05f5445f0e7d82cd95d3a2b916a58369898b82eef3"),
		ReceiptHash: common.HexToHash("0x7298f243c2bd5472fb20330a6ba763aaf9dd07d8b0f818b27277c6b0fd3b85a9"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(17),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xc0b1"),
		Time:        hexutil.MustDecodeUint64("0x6489d44b"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x634d0f4"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).NextTransactions = types.Transactions{
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

	return nil
}

func SetTargetOutputResponse(s **eth.OutputResponse) error {
	outputRoot, err := rollup.ComputeL2OutputRootV1(&bindings.TypesOutputRootProof{
		Version:                  rollup.V1,
		StateRoot:                common.HexToHash("0x008bc0b55af382fe6508389a13c011b873eb5d7dcbcd3f8ba8f9a2001c196205"),
		MessagePasserStorageRoot: common.HexToHash("0x"),
		BlockHash:                common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
		NextBlockHash:            common.HexToHash("0x76da1f020b58638d4f0199a96c46a4c493366c61bbb1dbd3a4ad60166006d6cd"),
	})
	if err != nil {
		return fmt.Errorf("mocking error: %w", err)
	}

	(*s).OutputRoot = outputRoot
	(*s).WithdrawalStorageRoot = common.HexToHash("0x")
	(*s).StateRoot = common.HexToHash("0x008bc0b55af382fe6508389a13c011b873eb5d7dcbcd3f8ba8f9a2001c196205")
	(*s).BlockRef = eth.L2BlockRef{
		Hash:       common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
		ParentHash: common.HexToHash("0x7296c63860b949715042cc8a60b5dcea924801b62ab7aa9ea40345abd570db40"),
	}
	(*s).NextBlockRef = eth.L2BlockRef{Hash: common.HexToHash("0x76da1f020b58638d4f0199a96c46a4c493366c61bbb1dbd3a4ad60166006d6cd")}
	(*s).NextBlock = &types.Header{
		ParentHash:  common.HexToHash("0x3534f7f01bdd7d6b568f4eb60cb6a4b08131678be85cf84a61cebd2f1ae81769"),
		Coinbase:    CoinbaseAddr,
		Difficulty:  common.Big0,
		Root:        common.HexToHash("0x141a7dd8d9e4daf81d1ebd599a306bd302e67f58cf92610525c95f13be8409be"),
		TxHash:      common.HexToHash("0xac4aa619e9e046033d4a727ac0748fedb62b58bfcd682c3b0c136220dcf8b251"),
		ReceiptHash: common.HexToHash("0x7298f243c2bd5472fb20330a6ba763aaf9dd07d8b0f818b27277c6b0fd3b85a9"),
		Bloom:       types.BytesToBloom(hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
		Number:      big.NewInt(18),
		GasLimit:    hexutil.MustDecodeUint64("0x1c9c380"),
		GasUsed:     hexutil.MustDecodeUint64("0xc0b1"),
		Time:        hexutil.MustDecodeUint64("0x6489d44d"),
		Extra:       hexutil.MustDecode("0x"),
		MixDigest:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BaseFee:     big.NewInt(int64(hexutil.MustDecodeUint64("0x56ede09"))),
	}
	toAddr := common.HexToAddress("0x4200000000000000000000000000000000000002")
	(*s).NextTransactions = types.Transactions{
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

	return nil
}
