package testutils

import "github.com/ethereum/go-ethereum/common"

type MockRuntimeConfig struct {
	P2PPropAddress common.Address
}

func (m *MockRuntimeConfig) P2PProposerAddress() common.Address {
	return m.P2PPropAddress
}
