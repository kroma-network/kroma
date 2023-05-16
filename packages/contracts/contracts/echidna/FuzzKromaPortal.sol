pragma solidity 0.8.15;

import { KromaPortal } from "../L1/KromaPortal.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ResourceMetering } from "../L1/ResourceMetering.sol";
import { SystemConfig } from "../L1/SystemConfig.sol";
import { ZKMerkleTrie } from "../L1/ZKMerkleTrie.sol";
import { Constants } from "../libraries/Constants.sol";
import { AddressAliasHelper } from "../vendor/AddressAliasHelper.sol";

contract EchidnaFuzzKromaPortal {
    KromaPortal internal portal;
    bool internal failedToComplete;

    constructor() {
        ResourceMetering.ResourceConfig memory rcfg = Constants.DEFAULT_RESOURCE_CONFIG();

        SystemConfig systemConfig = new SystemConfig({
            _owner: address(1),
            _overhead: 0,
            _scalar: 10000,
            _batcherHash: bytes32(0),
            _gasLimit: 30_000_000,
            _unsafeBlockSigner: address(0),
            _config: rcfg
        });
        portal = new KromaPortal({
            _l2Oracle: L2OutputOracle(address(0)),
            _guardian: address(0),
            _paused: false,
            _config: systemConfig,
            _zkMerkleTrie: ZKMerkleTrie(address(0))
        });
    }

    // A test intended to identify any unexpected halting conditions
    function testDepositTransactionCompletes(
        address _to,
        uint256 _mint,
        uint256 _value,
        uint64 _gasLimit,
        bool _isCreation,
        bytes memory _data
    ) public payable {
        failedToComplete = true;
        require(!_isCreation || _to == address(0), "EchidnaFuzzKromaPortal: invalid test case.");
        portal.depositTransaction{ value: _mint }(_to, _value, _gasLimit, _isCreation, _data);
        failedToComplete = false;
    }

    /**
     * @custom:invariant Deposits of any value should always succeed unless
     * `_to` = `address(0)` or `_isCreation` = `true`.
     *
     * All deposits, barring creation transactions and transactions sent to `address(0)`,
     * should always succeed.
     */
    function echidna_deposit_completes() public view returns (bool) {
        return !failedToComplete;
    }
}
