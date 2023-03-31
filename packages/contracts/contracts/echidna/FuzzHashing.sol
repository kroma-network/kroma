pragma solidity 0.8.15;

import { Encoding } from "../libraries/Encoding.sol";
import { Hashing } from "../libraries/Hashing.sol";

contract EchidnaFuzzHashing {
    bool internal failedCrossDomainHashHighVersion;
    bool internal failedCrossDomainHashV0;

    /**
     * @notice Takes the necessary parameters to perform a cross domain hash with a randomly
     * generated version. Only schema version 0 is supported and all others should revert.
     */
    function testHashCrossDomainMessageHighVersion(
        uint16 _version,
        uint240 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) public {
        // generate the versioned nonce
        uint256 encodedNonce = Encoding.encodeVersionedNonce(_nonce, _version);

        // hash the cross domain message. we don't need to store the result since the function
        // validates and should revert if an invalid version (>0) is encoded
        Hashing.hashCrossDomainMessage(encodedNonce, _sender, _target, _value, _gasLimit, _data);

        // check that execution never makes it this far for an invalid version
        if (_version > 0) {
            failedCrossDomainHashHighVersion = true;
        }
    }

    /**
     * @notice Takes the necessary parameters to perform a cross domain hash using the v0 schema
     * and compares the output of a call to the unversioned function to the v0 function directly
     */
    function testHashCrossDomainMessageV0(
        uint240 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) public {
        // generate the versioned nonce with the version set to 0
        uint256 encodedNonce = Encoding.encodeVersionedNonce(_nonce, 0);

        // hash the cross domain message using the unversioned and versioned functions for
        // comparison
        bytes32 sampleHash1 = Hashing.hashCrossDomainMessage(
            encodedNonce,
            _sender,
            _target,
            _value,
            _gasLimit,
            _data
        );
        bytes32 sampleHash2 = Hashing.hashCrossDomainMessageV0(
            encodedNonce,
            _sender,
            _target,
            _value,
            _gasLimit,
            _data
        );

        // check that the output of both functions matches
        if (sampleHash1 != sampleHash2) {
            failedCrossDomainHashV0 = true;
        }
    }

    /**
     * @custom:invariant `hashCrossDomainMessage` reverts if `version` is > `0`.
     *
     * The `hashCrossDomainMessage` function should always revert if the `version` passed is > `0`.
     */
    function echidna_hash_xdomain_msg_high_version() public view returns (bool) {
        // ASSERTION: A call to hashCrossDomainMessage will never succeed for a version > 1
        return !failedCrossDomainHashHighVersion;
    }

    /**
     * @custom:invariant `version` = `0`: `hashCrossDomainMessage` and `hashCrossDomainMessageV0`
     * are equivalent.
     *
     * If the version passed is 0, `hashCrossDomainMessage` and `hashCrossDomainMessageV0` should be
     * equivalent.
     */
    function echidna_hash_xdomain_msg_0() public view returns (bool) {
        // ASSERTION: A call to hashCrossDomainMessage and hashCrossDomainMessageV0
        // should always match when the version passed is 0
        return !failedCrossDomainHashV0;
    }
}
