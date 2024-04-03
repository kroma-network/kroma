package predeploys

import "github.com/ethereum/go-ethereum/common"

// Define the information of the deposit contract on the execution layer for the beacon chain.
var (
	BeaconDepositContractAddr     = common.HexToAddress("0x4242424242424242424242424242424242424242")
	BeaconDepositContractCode     = common.FromHex("0x3373fffffffffffffffffffffffffffffffffffffffe14604d57602036146024575f5ffd5b5f35801560495762001fff810690815414603c575f5ffd5b62001fff01545f5260205ff35b5f5ffd5b62001fff42064281555f359062001fff015500")
	BeaconDepositContractCodeHash = common.HexToHash("0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470")
)
