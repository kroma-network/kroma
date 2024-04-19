// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/* Testing utilities */
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { Test } from "forge-std/Test.sol";

import { SecurityCouncilToken } from "../governance/SecurityCouncilToken.sol";
import { TimeLock } from "../governance/TimeLock.sol";
import { UpgradeGovernor } from "../governance/UpgradeGovernor.sol";
import { AssetManager } from "../L1/AssetManager.sol";
import { Colosseum } from "../L1/Colosseum.sol";
import { IValidatorManager } from "../L1/interfaces/IValidatorManager.sol";
import { KromaPortal } from "../L1/KromaPortal.sol";
import { L1CrossDomainMessenger } from "../L1/L1CrossDomainMessenger.sol";
import { L1ERC721Bridge } from "../L1/L1ERC721Bridge.sol";
import { L1StandardBridge } from "../L1/L1StandardBridge.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { SecurityCouncil } from "../L1/SecurityCouncil.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { ValidatorManager } from "../L1/ValidatorManager.sol";
import { ResourceMetering } from "../L1/ResourceMetering.sol";
import { SystemConfig } from "../L1/SystemConfig.sol";
import { ZKMerkleTrie } from "../L1/ZKMerkleTrie.sol";
import { ZKVerifier } from "../L1/ZKVerifier.sol";
import { L2CrossDomainMessenger } from "../L2/L2CrossDomainMessenger.sol";
import { L2ERC721Bridge } from "../L2/L2ERC721Bridge.sol";
import { L2StandardBridge } from "../L2/L2StandardBridge.sol";
import { L2ToL1MessagePasser } from "../L2/L2ToL1MessagePasser.sol";
import { CodeDeployer } from "../libraries/CodeDeployer.sol";
import { Constants } from "../libraries/Constants.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { Types } from "../libraries/Types.sol";
import { IKGHManager } from "../universal/IKGHManager.sol";
import { KromaMintableERC20 } from "../universal/KromaMintableERC20.sol";
import { KromaMintableERC20Factory } from "../universal/KromaMintableERC20Factory.sol";
import { Proxy } from "../universal/Proxy.sol";
import { AddressAliasHelper } from "../vendor/AddressAliasHelper.sol";

contract CommonTest is Test {
    address alice = address(128);
    address bob = address(256);
    address multisig = address(512);

    address immutable ZERO_ADDRESS = address(0);
    address immutable NON_ZERO_ADDRESS = address(1);
    uint256 immutable NON_ZERO_VALUE = 100;
    uint256 immutable ZERO_VALUE = 0;
    uint64 immutable NON_ZERO_GASLIMIT = 50000;
    bytes32 nonZeroHash = keccak256(abi.encode("NON_ZERO"));
    bytes NON_ZERO_DATA = hex"0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff0000";

    uint256 constant MAX_OUTPUT_ROOT_PROOF_VERSION = 0;

    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 indexed version,
        bytes opaqueData
    );

    FFIInterface ffi;

    function setUp() public virtual {
        // Give alice and bob some ETH
        vm.deal(alice, 1 << 16);
        vm.deal(bob, 1 << 16);
        vm.deal(multisig, 1 << 16);

        vm.label(alice, "alice");
        vm.label(bob, "bob");
        vm.label(multisig, "multisig");

        // Make sure we have a non-zero base fee
        vm.fee(1000000000);

        ffi = new FFIInterface();
    }

    function emitTransactionDeposited(
        address _from,
        address _to,
        uint256 _mint,
        uint256 _value,
        uint64 _gasLimit,
        bool _isCreation,
        bytes memory _data
    ) internal {
        emit TransactionDeposited(
            _from,
            _to,
            0,
            abi.encodePacked(_mint, _value, _gasLimit, _isCreation, _data)
        );
    }

    //payable proxy
    function toProxy(address target) internal pure returns (Proxy) {
        return Proxy(payable(target));
    }
}

contract UpgradeGovernor_Initializer is CommonTest {
    address owner = makeAddr("owner");

    // Constructor arguments
    uint256 internal initialVotingDelay = 0;
    uint256 internal initialVotingPeriod = 30;
    uint256 internal initialProposalThreshold = 1;
    uint256 internal votesQuorumFraction = 70;
    uint256 internal minDelaySeconds = 3;
    string internal baseUri = "";

    // Test data
    address internal guardian1 = 0x0000000000000000000000000000000000001004;
    address internal guardian2 = 0x0000000000000000000000000000000000001005;
    address internal guardian3 = 0x0000000000000000000000000000000000001006;
    address internal notGuardian = 0x0000000000000000000000000000000000002000;

    address[] timeLockProposers = new address[](1);
    address[] timeLockExecutors = new address[](1);

    SecurityCouncilToken securityCouncilToken = SecurityCouncilToken(address(new Proxy(multisig)));
    TimeLock timeLock = TimeLock(payable(address(new Proxy(multisig))));
    UpgradeGovernor upgradeGovernor = UpgradeGovernor(payable(address(new Proxy(multisig))));

    SecurityCouncilToken securityCouncilTokenImpl = new SecurityCouncilToken();
    TimeLock timeLockImpl = new TimeLock();
    UpgradeGovernor upgradeGovernorImpl = new UpgradeGovernor();

    function setUp() public virtual override {
        super.setUp();
        vm.startPrank(multisig);

        // setup SecurityCouncilToken
        toProxy(address(securityCouncilToken)).upgradeToAndCall(
            address(securityCouncilTokenImpl),
            abi.encodeCall(SecurityCouncilToken.initialize, owner)
        );

        // setup TimeLock & UpgradeGovernor
        timeLockProposers[0] = address(upgradeGovernor);
        timeLockExecutors[0] = address(upgradeGovernor);

        toProxy(address(timeLock)).upgradeToAndCall(
            address(timeLockImpl),
            abi.encodeCall(
                TimeLock.initialize,
                (minDelaySeconds, timeLockProposers, timeLockExecutors, address(upgradeGovernor))
            )
        );

        toProxy(address(upgradeGovernor)).upgradeToAndCall(
            address(upgradeGovernorImpl),
            abi.encodeCall(
                UpgradeGovernor.initialize,
                (
                    address(securityCouncilToken),
                    payable(address(timeLock)),
                    initialVotingDelay,
                    initialVotingPeriod,
                    initialProposalThreshold,
                    votesQuorumFraction
                )
            )
        );

        //change proxy admin to upgradeGovernor
        toProxy(address(securityCouncilToken)).changeAdmin(address(timeLock));
        toProxy(address(timeLock)).changeAdmin(address(timeLock));
        toProxy(address(upgradeGovernor)).changeAdmin(address(timeLock));

        vm.stopPrank();
    }
}

contract MockKro is ERC20 {
    constructor() ERC20("Kroma", "KRO") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract MockKgh is ERC721 {
    constructor() ERC721("Test", "TST") {}

    function mint(address to, uint256 tokenId) external {
        _mint(to, tokenId);
    }
}

contract MockKghManager is IKGHManager {
    function totalKroInKgh(uint256 /* tokenId */) external pure override returns (uint128) {
        return 100e18;
    }
}

contract L2OutputOracle_Initializer is UpgradeGovernor_Initializer {
    // Test target
    ValidatorPool pool;
    ValidatorPool poolImpl;
    AssetManager assetMan;
    AssetManager assetManImpl;
    ValidatorManager valMan;
    ValidatorManager valManImpl;
    L2OutputOracle oracle;
    L2OutputOracle oracleImpl;
    Colosseum colosseum;
    SystemConfig systemConfig;
    KromaPortal mockPortal;

    L2ToL1MessagePasser messagePasser =
        L2ToL1MessagePasser(payable(Predeploys.L2_TO_L1_MESSAGE_PASSER));

    // Constructor arguments
    uint256 internal submissionInterval = 1800;
    uint256 internal l2BlockTime = 2;
    uint256 internal startingBlockNumber = 10;
    uint256 internal startingTimestamp = 1000;
    uint256 internal finalizationPeriodSeconds = 7 days;
    address internal guardian = 0x000000000000000000000000000000000000AaaD;

    // ValidatorPool constructor arguments
    address internal trusted = 0x000000000000000000000000000000000000aaaa;
    uint256 internal requiredBondAmount = 0.1 ether;
    uint256 internal maxUnbond = 2;
    uint256 internal roundDuration = (submissionInterval * l2BlockTime) / 2;
    uint256 internal terminateOutputIndex = 300; // just large enough value, set again in upgrade test

    // AssetManager constructor arguments
    MockKro assetToken;
    MockKgh kgh;
    MockKghManager kghManager;
    uint256 internal undelegationPeriod = 7 days;
    uint128 internal slashingRate = 20;
    uint128 internal minSlashingAmount = 1e18;

    // ValidatorManager constructor arguments
    uint128 internal commissionRateMinChangeSeconds = 7 days;
    uint128 internal jailPeriodSeconds = 7 days;
    uint128 internal jailThreshold = 2;
    uint128 internal maxOutputFinalizations = 10;
    uint128 internal baseReward = 20e18;
    uint128 internal minRegisterAmount = 10e18;
    uint128 internal minStartAmount = 100e18;
    IValidatorManager.ConstructorParams constructorParams;

    // Test data
    address internal asserter = 0x000000000000000000000000000000000000aAaB;
    address internal challenger = 0x000000000000000000000000000000000000AAaC;
    uint256 initL1Time;

    event OutputSubmitted(
        bytes32 indexed outputRoot,
        uint256 indexed l2OutputIndex,
        uint256 indexed l2BlockNumber,
        uint256 l1Timestamp
    );

    event OutputReplaced(uint256 indexed outputIndex, bytes32 newOutputRoot);

    // Advance the evm's time to meet the L2OutputOracle's requirements for submitL2Output
    function warpToSubmitTime() public {
        vm.warp(oracle.nextOutputMinL2Timestamp());
    }

    function _registerValidator(address validator, uint128 assets) internal {
        vm.startPrank(validator);
        assetToken.approve(address(assetMan), uint256(assets));
        valMan.registerValidator(assets, 10, 5);
        vm.stopPrank();
    }

    function setUp() public virtual override {
        super.setUp();

        // Set up asset token and give actors some amount
        assetToken = new MockKro();
        assetToken.mint(trusted, minStartAmount * 10);
        assetToken.mint(asserter, minStartAmount * 10);
        assetToken.mint(challenger, minStartAmount * 10);

        // Set up KGH and KGHManager
        kgh = new MockKgh();
        kghManager = new MockKghManager();

        // Give actors some ETH
        vm.deal(trusted, requiredBondAmount * 10);
        vm.deal(asserter, requiredBondAmount * 10);
        vm.deal(challenger, requiredBondAmount * 10);

        // Deploy proxies
        pool = ValidatorPool(address(new Proxy(multisig)));
        vm.label(address(pool), "ValidatorPool");
        assetMan = AssetManager(address(new Proxy(multisig)));
        vm.label(address(assetMan), "AssetManager");
        valMan = ValidatorManager(address(new Proxy(multisig)));
        vm.label(address(valMan), "ValidatorManager");
        oracle = L2OutputOracle(address(new Proxy(multisig)));
        vm.label(address(oracle), "L2OutputOracle");

        ResourceMetering.ResourceConfig memory config = Constants.DEFAULT_RESOURCE_CONFIG();
        systemConfig = new SystemConfig({
            _owner: address(1),
            _overhead: 0,
            _scalar: 10000,
            _batcherHash: bytes32(0),
            _gasLimit: 30_000_000,
            _unsafeBlockSigner: address(0),
            _config: config,
            _validatorRewardScalar: 5000
        });

        // Mock KromaPortal
        mockPortal = new KromaPortal({
            _l2Oracle: oracle,
            _validatorPool: address(pool),
            _guardian: guardian,
            _paused: false,
            _config: systemConfig,
            _zkMerkleTrie: ZKMerkleTrie(address(0))
        });

        // Deploy the ValidatorPool
        poolImpl = new ValidatorPool({
            _l2OutputOracle: oracle,
            _portal: mockPortal,
            _securityCouncil: guardian,
            _trustedValidator: trusted,
            _requiredBondAmount: requiredBondAmount,
            _maxUnbond: maxUnbond,
            _roundDuration: roundDuration,
            _terminateOutputIndex: terminateOutputIndex
        });

        // Deploy the AssetManager
        assetManImpl = new AssetManager({
            _assetToken: IERC20(assetToken),
            _kgh: IERC721(kgh),
            _kghManager: IKGHManager(kghManager),
            _securityCouncil: guardian,
            _validatorManager: valMan,
            _undelegationPeriod: uint128(undelegationPeriod),
            _slashingRate: slashingRate,
            _minSlashingAmount: minSlashingAmount
        });

        // Deploy the ValidatorManager
        constructorParams = IValidatorManager.ConstructorParams({
            _l2Oracle: oracle,
            _assetManager: assetMan,
            _trustedValidator: trusted,
            _commissionRateMinChangeSeconds: commissionRateMinChangeSeconds,
            _roundDurationSeconds: uint128(roundDuration),
            _jailPeriodSeconds: jailPeriodSeconds,
            _jailThreshold: jailThreshold,
            _maxOutputFinalizations: maxOutputFinalizations,
            _baseReward: baseReward,
            _minRegisterAmount: minRegisterAmount,
            _minStartAmount: minStartAmount
        });
        valManImpl = new ValidatorManager({ _constructorParams: constructorParams });

        // By default the first block has timestamp and number zero, which will cause underflows in
        // the tests, so we'll move forward to these block values.
        initL1Time = startingTimestamp + 1;
        vm.warp(initL1Time);
        vm.roll(startingBlockNumber);
        // Deploy the L2OutputOracle
        oracleImpl = new L2OutputOracle({
            _validatorPool: pool,
            _validatorManager: valMan,
            _colosseum: address(colosseum),
            _submissionInterval: submissionInterval,
            _l2BlockTime: l2BlockTime,
            _startingBlockNumber: startingBlockNumber,
            _startingTimestamp: startingTimestamp,
            _finalizationPeriodSeconds: finalizationPeriodSeconds
        });

        vm.prank(multisig);
        Proxy(payable(address(pool))).upgradeToAndCall(
            address(poolImpl),
            abi.encodeWithSelector(ValidatorPool.initialize.selector)
        );

        vm.prank(multisig);
        Proxy(payable(address(assetMan))).upgradeTo(address(assetManImpl));

        vm.prank(multisig);
        Proxy(payable(address(valMan))).upgradeTo(address(valManImpl));

        vm.prank(multisig);
        Proxy(payable(address(oracle))).upgradeToAndCall(
            address(oracleImpl),
            abi.encodeCall(L2OutputOracle.initialize, (startingBlockNumber, startingTimestamp))
        );

        // Set the L2ToL1MessagePasser at the correct address
        vm.etch(Predeploys.L2_TO_L1_MESSAGE_PASSER, address(new L2ToL1MessagePasser()).code);

        vm.label(Predeploys.L2_TO_L1_MESSAGE_PASSER, "L2ToL1MessagePasser");
    }
}

contract L2OutputOracle_ValidatorSystemUpgrade_Initializer is L2OutputOracle_Initializer {
    uint256 finalizationPeriodOutputNum;

    function setUp() public virtual override {
        super.setUp();

        // last output index of ValidatorPool is 3, first output index of ValidatorManager is 4
        terminateOutputIndex = 3;
        finalizationPeriodSeconds = 2 hours;

        oracleImpl = new L2OutputOracle({
            _validatorPool: pool,
            _validatorManager: valMan,
            _colosseum: address(colosseum),
            _submissionInterval: submissionInterval,
            _l2BlockTime: l2BlockTime,
            _startingBlockNumber: startingBlockNumber,
            _startingTimestamp: startingTimestamp,
            _finalizationPeriodSeconds: finalizationPeriodSeconds
        });
        vm.prank(multisig);
        Proxy(payable(address(oracle))).upgradeTo(address(oracleImpl));

        poolImpl = new ValidatorPool({
            _l2OutputOracle: oracle,
            _portal: mockPortal,
            _securityCouncil: guardian,
            _trustedValidator: trusted,
            _requiredBondAmount: requiredBondAmount,
            _maxUnbond: maxUnbond,
            _roundDuration: roundDuration,
            _terminateOutputIndex: terminateOutputIndex
        });
        vm.prank(multisig);
        Proxy(payable(address(pool))).upgradeTo(address(poolImpl));

        finalizationPeriodOutputNum =
            oracle.FINALIZATION_PERIOD_SECONDS() /
            (oracle.SUBMISSION_INTERVAL() * oracle.L2_BLOCK_TIME());
    }
}

contract Poseidon2Deployer {
    bytes constant poseidon2Code =
        //solhint-disable-next-line max-line-length
        hex"38600c60003961260f6000f37c010000000000000000000000000000000000000000000000000000000060003504806329a5f2f6149063299e566014176200003757fe5b7f109b7f411ba0e4c9b2b70caf5c36a7b194be7c11ad24378bfedb68592ba8118b6020527f16ed41e13bb9c0c66ae119424fddbcbc9314dc9fdbdeea55d6c64543dc4903e06040527f2b90bba00fca0589f617e7dcbfe82e0df706ab640ceb247b791a93b74e36736d6060527f2969f27eed31a480b9c36c764379dbca2cc8fdd1415c3dded62940bcde0bd7716080527f2e2419f9ec02ec394c9871c832963dc1b89d743c8c7b964029b2311687b1fe2360a0527f101071f0032379b697315876690f053d148d4e109f5fb065c8aacc55a0f89bfa60c0527f143021ec686a3f330d5f9e654638065ce6cd79e28c5b3753326244ee65a1b1a760e0527f176cc029695ad02582a70eff08a6fd99d057e12e58e7d7b6b16cdfabc8ee2911610100527f19a3fc0a56702bf417ba7fee3802593fa644470307043f7773279cd71d25d5e0610120527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016024356004356000837f0ee9a592ba9a9518d05986d656f40c2114c4993c11bb29938d21d47304cd8e6e82089050837f00f1445235f2148c5986587169fc1bcd887b08d4d00868df5696fff40956e86483089150837f08dff3487e8ac99e1f29a058d0fa80b930c728730b7ab36ce879f3890ecf73f58408925083818180828009800909905083828180828009800909915083838180828009800909925062000249600052620025ba565b837f2f27be690fdaee46c3ce28f7532b13c856c35342c84bda6e20966310fadc01d082089050837f2b2ae1acf68b7b8d2416bebf3d4f6234b763fe04b8043ee48b8327bebca16cf283089150837f0319d062072bef7ecca5eac06f97d4d55952c175ab6b03eae64b44c7dbf11cfa84089250838181808280098009099050838281808280098009099150838381808280098009099250620002ec600052620025ba565b837f28813dcaebaeaa828a376df87af4a63bc8b7bf27ad49c6298ef7b387bf28526d82089050837f2727673b2ccbc903f181bf38e1c1d40d2033865200c352bc150928adddf9cb7883089150837f234ec45ca27727c2e74abd2b2a1494cd6efbd43e340587d6b8fb9e31e65cc632840892508381818082800980090990508382818082800980090991508383818082800980090992506200038f600052620025ba565b837f15b52534031ae18f7f862cb2cf7cf760ab10a8150a337b1ccd99ff6e8797d42882089050837f0dc8fad6d9e4b35f5ed9a3d186b79ce38e0e8a8d1b58b132d701d4eecf68d1f683089150837f1bcd95ffc211fbca600f705fad3fb567ea4eb378f62e1fec97805518a47e4d9c8408925083818180828009800909905083828180828009800909915083838180828009800909925062000432600052620025ba565b837f10520b0ab721cadfe9eff81b016fc34dc76da36c2578937817cb978d069de55982089050837f1f6d48149b8e7f7d9b257d8ed5fbbaf42932498075fed0ace88a9eb81f5627f683089150837f1d9655f652309014d29e00ef35a2089bfff8dc1c816f0dc9ca34bdb5460c870584089250838181808280098009099050620004bd600052620025ba565b837f04df5a56ff95bcafb051f7b1cd43a99ba731ff67e47032058fe3d4185697cc7d82089050837f0672d995f8fff640151b3d290cedaf148690a10a8c8424a7f6ec282b6e4be82883089150837f099952b414884454b21200d7ffafdd5f0c9a9dcc06f2708e9fc1d8209b5c75b98408925083818180828009800909905062000548600052620025ba565b837f052cba2255dfd00c7c483143ba8d469448e43586a9b4cd9183fd0e843a6b9fa682089050837f0b8badee690adb8eb0bd74712b7999af82de55707251ad7716077cb93c464ddc83089150837f119b1590f13307af5a1ee651020c07c749c15d60683a8050b963d0a8e4b2bdd184089250838181808280098009099050620005d3600052620025ba565b837f03150b7cd6d5d17b2529d36be0f67b832c4acfc884ef4ee5ce15be0bfb4a8d0982089050837f2cc6182c5e14546e3cf1951f173912355374efb83d80898abe69cb317c9ea56583089150837f005032551e6378c450cfe129a404b3764218cadedac14e2b92d2cd73111bf0f9840892508381818082800980090990506200065e600052620025ba565b837f233237e3289baa34bb147e972ebcb9516469c399fcc069fb88f9da2cc28276b582089050837f05c8f4f4ebd4a6e3c980d31674bfbe6323037f21b34ae5a4e80c2d4c24d6028083089150837f0a7b1db13042d396ba05d818a319f25252bcf35ef3aeed91ee1f09b2590fc65b84089250838181808280098009099050620006e9600052620025ba565b837f2a73b71f9b210cf5b14296572c9d32dbf156e2b086ff47dc5df542365a404ec082089050837f1ac9b0417abcc9a1935107e9ffc91dc3ec18f2c4dbe7f22976a760bb5c50c46083089150837f12c0339ae08374823fabb076707ef479269f3e4d6cb104349015ee046dc93fc08408925083818180828009800909905062000774600052620025ba565b837f0b7475b102a165ad7f5b18db4e1e704f52900aa3253baac68246682e56e9a28e82089050837f037c2849e191ca3edb1c5e49f6e8b8917c843e379366f2ea32ab3aa88d7f844883089150837f05a6811f8556f014e92674661e217e9bd5206c5c93a07dc145fdb176a716346f84089250838181808280098009099050620007ff600052620025ba565b837f29a795e7d98028946e947b75d54e9f044076e87a7b2883b47b675ef5f38bd66e82089050837f20439a0c84b322eb45a3857afc18f5826e8c7382c8a1585c507be199981fd22f83089150837f2e0ba8d94d9ecf4a94ec2050c7371ff1bb50f27799a84b6d4a2a6f2a0982c887840892508381818082800980090990506200088a600052620025ba565b837f143fd115ce08fb27ca38eb7cce822b4517822cd2109048d2e6d0ddcca17d71c882089050837f0c64cbecb1c734b857968dbbdcf813cdf8611659323dbcbfc84323623be9caf183089150837f028a305847c683f646fca925c163ff5ae74f348d62c2b670f1426cef9403da538408925083818180828009800909905062000915600052620025ba565b837f2e4ef510ff0b6fda5fa940ab4c4380f26a6bcb64d89427b824d6755b5db9e30c82089050837f0081c95bc43384e663d79270c956ce3b8925b4f6d033b078b96384f50579400e83089150837f2ed5f0c91cbd9749187e2fade687e05ee2491b349c039a0bba8a9f4023a0bb3884089250838181808280098009099050620009a0600052620025ba565b837f30509991f88da3504bbf374ed5aae2f03448a22c76234c8c990f01f33a73520682089050837f1c3f20fd55409a53221b7c4d49a356b9f0a1119fb2067b41a7529094424ec6ad83089150837f10b4e7f3ab5df003049514459b6e18eec46bb2213e8e131e170887b47ddcb96c8408925083818180828009800909905062000a2b600052620025ba565b837f2a1982979c3ff7f43ddd543d891c2abddd80f804c077d775039aa3502e43adef82089050837f1c74ee64f15e1db6feddbead56d6d55dba431ebc396c9af95cad0f1315bd5c9183089150837f07533ec850ba7f98eab9303cace01b4b9e4f2e8b82708cfa9c2fe45a0ae146a08408925083818180828009800909905062000ab6600052620025ba565b837f21576b438e500449a151e4eeaf17b154285c68f42d42c1808a11abf3764c075082089050837f2f17c0559b8fe79608ad5ca193d62f10bce8384c815f0906743d6930836d4a9e83089150837f2d477e3862d07708a79e8aae946170bc9775a4201318474ae665b0b1b7e2730e8408925083818180828009800909905062000b41600052620025ba565b837f162f5243967064c390e095577984f291afba2266c38f5abcd89be0f5b2747eab82089050837f2b4cb233ede9ba48264ecd2c8ae50d1ad7a8596a87f29f8a7777a7009239331183089150837f2c8fbcb2dd8573dc1dbaf8f4622854776db2eece6d85c4cf4254e7c35e03b07a8408925083818180828009800909905062000bcc600052620025ba565b837f1d6f347725e4816af2ff453f0cd56b199e1b61e9f601e9ade5e88db870949da982089050837f204b0c397f4ebe71ebc2d8b3df5b913df9e6ac02b68d31324cd49af5c456552983089150837f0c4cb9dc3c4fd8174f1149b3c63c3c2f9ecb827cd7dc25534ff8fb75bc79c5028408925083818180828009800909905062000c57600052620025ba565b837f174ad61a1448c899a25416474f4930301e5c49475279e0639a616ddc45bc7b5482089050837f1a96177bcf4d8d89f759df4ec2f3cde2eaaa28c177cc0fa13a9816d49a38d2ef83089150837f066d04b24331d71cd0ef8054bc60c4ff05202c126a233c1a8242ace360b8a30a8408925083818180828009800909905062000ce2600052620025ba565b837f2a4c4fc6ec0b0cf52195782871c6dd3b381cc65f72e02ad527037a62aa1bd80482089050837f13ab2d136ccf37d447e9f2e14a7cedc95e727f8446f6d9d7e55afc01219fd64983089150837f1121552fca26061619d24d843dc82769c1b04fcec26f55194c2e3e869acc6a9a8408925083818180828009800909905062000d6d600052620025ba565b837f00ef653322b13d6c889bc81715c37d77a6cd267d595c4a8909a5546c7c97cff182089050837f0e25483e45a665208b261d8ba74051e6400c776d652595d9845aca35d8a397d383089150837f29f536dcb9dd7682245264659e15d88e395ac3d4dde92d8c46448db979eeba898408925083818180828009800909905062000df8600052620025ba565b837f2a56ef9f2c53febadfda33575dbdbd885a124e2780bbea170e456baace0fa5be82089050837f1c8361c78eb5cf5decfb7a2d17b5c409f2ae2999a46762e8ee416240a8cb9af183089150837f151aff5f38b20a0fc0473089aaf0206b83e8e68a764507bfd3d0ab4be74319c58408925083818180828009800909905062000e83600052620025ba565b837f04c6187e41ed881dc1b239c88f7f9d43a9f52fc8c8b6cdd1e76e47615b51f10082089050837f13b37bd80f4d27fb10d84331f6fb6d534b81c61ed15776449e801b7ddc9c296783089150837f01a5c536273c2d9df578bfbd32c17b7a2ce3664c2a52032c9321ceb1c4e8a8e48408925083818180828009800909905062000f0e600052620025ba565b837f2ab3561834ca73835ad05f5d7acb950b4a9a2c666b9726da832239065b7c3b0282089050837f1d4d8ec291e720db200fe6d686c0d613acaf6af4e95d3bf69f7ed516a597b64683089150837f041294d2cc484d228f5784fe7919fd2bb925351240a04b711514c9c80b65af1d8408925083818180828009800909905062000f99600052620025ba565b837f154ac98e01708c611c4fa715991f004898f57939d126e392042971dd90e81fc682089050837f0b339d8acca7d4f83eedd84093aef51050b3684c88f8b0b04524563bc6ea4da483089150837f0955e49e6610c94254a4f84cfbab344598f0e71eaff4a7dd81ed95b50839c82e8408925083818180828009800909905062001024600052620025ba565b837f06746a6156eba54426b9e22206f15abca9a6f41e6f535c6f3525401ea065462682089050837f0f18f5a0ecd1423c496f3820c549c27838e5790e2bd0a196ac917c7ff32077fb83089150837f04f6eeca1751f7308ac59eff5beb261e4bb563583ede7bc92a738223d6f76e1384089250838181808280098009099050620010af600052620025ba565b837f2b56973364c4c4f5c1a3ec4da3cdce038811eb116fb3e45bc1768d26fc0b375882089050837f123769dd49d5b054dcd76b89804b1bcb8e1392b385716a5d83feb65d437f29ef83089150837f2147b424fc48c80a88ee52b91169aacea989f6446471150994257b2fb01c63e9840892508381818082800980090990506200113a600052620025ba565b837f0fdc1f58548b85701a6c5505ea332a29647e6f34ad4243c2ea54ad897cebe54d82089050837f12373a8251fea004df68abcf0f7786d4bceff28c5dbbe0c3944f685cc0a0b1f283089150837f21e4f4ea5f35f85bad7ea52ff742c9e8a642756b6af44203dd8a1f35c1a9003584089250838181808280098009099050620011c5600052620025ba565b837f16243916d69d2ca3dfb4722224d4c462b57366492f45e90d8a81934f1bc3b14782089050837f1efbe46dd7a578b4f66f9adbc88b4378abc21566e1a0453ca13a4159cac04ac283089150837f07ea5e8537cf5dd08886020e23a7f387d468d5525be66f853b672cc96a88969a8408925083818180828009800909905062001250600052620025ba565b837f05a8c4f9968b8aa3b7b478a30f9a5b63650f19a75e7ce11ca9fe16c0b76c00bc82089050837f20f057712cc21654fbfe59bd345e8dac3f7818c701b9c7882d9d57b72a32e83f83089150837f04a12ededa9dfd689672f8c67fee31636dcd8e88d01d49019bd90b33eb33db6984089250838181808280098009099050620012db600052620025ba565b837f27e88d8c15f37dcee44f1e5425a51decbd136ce5091a6767e49ec9544ccd101a82089050837f2feed17b84285ed9b8a5c8c5e95a41f66e096619a7703223176c41ee433de4d183089150837f1ed7cc76edf45c7c404241420f729cf394e5942911312a0d6972b8bd53aff2b88408925083818180828009800909905062001366600052620025ba565b837f15742e99b9bfa323157ff8c586f5660eac6783476144cdcadf2874be45466b1a82089050837f1aac285387f65e82c895fc6887ddf40577107454c6ec0317284f033f27d0c78583089150837f25851c3c845d4790f9ddadbdb6057357832e2e7a49775f71ec75a96554d67c7784089250838181808280098009099050620013f1600052620025ba565b837f15a5821565cc2ec2ce78457db197edf353b7ebba2c5523370ddccc3d9f146a6782089050837f2411d57a4813b9980efa7e31a1db5966dcf64f36044277502f15485f28c7172783089150837f002e6f8d6520cd4713e335b8c0b6d2e647e9a98e12f4cd2558828b5ef6cb4c9b840892508381818082800980090990506200147c600052620025ba565b837f2ff7bc8f4380cde997da00b616b0fcd1af8f0e91e2fe1ed7398834609e0315d282089050837f00b9831b948525595ee02724471bcd182e9521f6b7bb68f1e93be4febb0d3cbe83089150837f0a2f53768b8ebf6a86913b0e57c04e011ca408648a4743a87d77adbf0c9c35128408925083818180828009800909905062001507600052620025ba565b837f00248156142fd0373a479f91ff239e960f599ff7e94be69b7f2a290305e1198d82089050837f171d5620b87bfb1328cf8c02ab3f0c9a397196aa6a542c2350eb512a2b2bcda983089150837f170a4f55536f7dc970087c7c10d6fad760c952172dd54dd99d1045e4ec34a8088408925083818180828009800909905062001592600052620025ba565b837f29aba33f799fe66c2ef3134aea04336ecc37e38c1cd211ba482eca17e2dbfae182089050837f1e9bc179a4fdd758fdd1bb1945088d47e70d114a03f6a0e8b5ba650369e6497383089150837f1dd269799b660fad58f7f4892dfb0b5afeaad869a9c4b44f9c9e1c43bdaf8f09840892508381818082800980090990506200161d600052620025ba565b837f22cdbc8b70117ad1401181d02e15459e7ccd426fe869c7c95d1dd2cb0f24af3882089050837f0ef042e454771c533a9f57a55c503fcefd3150f52ed94a7cd5ba93b9c7dacefd83089150837f11609e06ad6c8fe2f287f3036037e8851318e8b08a0359a03b304ffca62e828484089250838181808280098009099050620016a8600052620025ba565b837f1166d9e554616dba9e753eea427c17b7fecd58c076dfe42708b08f5b783aa9af82089050837f2de52989431a859593413026354413db177fbf4cd2ac0b56f855a888357ee46683089150837f3006eb4ffc7a85819a6da492f3a8ac1df51aee5b17b8e89d74bf01cf5f71e9ad8408925083818180828009800909905062001733600052620025ba565b837f2af41fbb61ba8a80fdcf6fff9e3f6f422993fe8f0a4639f962344c822514508682089050837f119e684de476155fe5a6b41a8ebc85db8718ab27889e85e781b214bace4827c383089150837f1835b786e2e8925e188bea59ae363537b51248c23828f047cff784b97b3fd80084089250838181808280098009099050620017be600052620025ba565b837f28201a34c594dfa34d794996c6433a20d152bac2a7905c926c40e285ab32eeb682089050837f083efd7a27d1751094e80fefaf78b000864c82eb571187724a761f88c22cc4e783089150837f0b6f88a3577199526158e61ceea27be811c16df7774dd8519e079564f61fd13b8408925083818180828009800909905062001849600052620025ba565b837f0ec868e6d15e51d9644f66e1d6471a94589511ca00d29e1014390e6ee4254f5b82089050837f2af33e3f866771271ac0c9b3ed2e1142ecd3e74b939cd40d00d937ab84c9859183089150837f0b520211f904b5e7d09b5d961c6ace7734568c547dd6858b364ce5e47951f17884089250838181808280098009099050620018d4600052620025ba565b837f0b2d722d0919a1aad8db58f10062a92ea0c56ac4270e822cca228620188a1d4082089050837f1f790d4d7f8cf094d980ceb37c2453e957b54a9991ca38bbe0061d1ed6e562d483089150837f0171eb95dfbf7d1eaea97cd385f780150885c16235a2a6a8da92ceb01e504233840892508381818082800980090990506200195f600052620025ba565b837f0c2d0e3b5fd57549329bf6885da66b9b790b40defd2c8650762305381b16887382089050837f1162fb28689c27154e5a8228b4e72b377cbcafa589e283c35d3803054407a18d83089150837f2f1459b65dee441b64ad386a91e8310f282c5a92a89e19921623ef8249711bc084089250838181808280098009099050620019ea600052620025ba565b837f1e6ff3216b688c3d996d74367d5cd4c1bc489d46754eb712c243f70d1b53cfbb82089050837f01ca8be73832b8d0681487d27d157802d741a6f36cdc2a0576881f932647887583089150837f1f7735706ffe9fc586f976d5bdf223dc680286080b10cea00b9b5de315f9650e8408925083818180828009800909905062001a75600052620025ba565b837f2522b60f4ea3307640a0c2dce041fba921ac10a3d5f096ef4745ca838285f01982089050837f23f0bee001b1029d5255075ddc957f833418cad4f52b6c3f8ce16c235572575b83089150837f2bc1ae8b8ddbb81fcaac2d44555ed5685d142633e9df905f66d9401093082d598408925083818180828009800909905062001b00600052620025ba565b837f0f9406b8296564a37304507b8dba3ed162371273a07b1fc98011fcd6ad72205f82089050837f2360a8eb0cc7defa67b72998de90714e17e75b174a52ee4acb126c8cd995f0a883089150837f15871a5cddead976804c803cbaef255eb4815a5e96df8b006dcbbc2767f889488408925083818180828009800909905062001b8b600052620025ba565b837f193a56766998ee9e0a8652dd2f3b1da0362f4f54f72379544f957ccdeefb420f82089050837f2a394a43934f86982f9be56ff4fab1703b2e63c8ad334834e4309805e777ae0f83089150837f1859954cfeb8695f3e8b635dcb345192892cd11223443ba7b4166e8876c0d1428408925083818180828009800909905062001c16600052620025ba565b837f04e1181763050e58013444dbcb99f1902b11bc25d90bbdca408d3819f4fed32b82089050837f0fdb253dee83869d40c335ea64de8c5bb10eb82db08b5e8b1f5e5552bfd05f2383089150837f058cbe8a9a5027bdaa4efb623adead6275f08686f1c08984a9d7c5bae9b4f1c08408925083818180828009800909905062001ca1600052620025ba565b837f1382edce9971e186497eadb1aeb1f52b23b4b83bef023ab0d15228b4cceca59a82089050837f03464990f045c6ee0819ca51fd11b0be7f61b8eb99f14b77e1e6634601d9e8b583089150837f23f7bfc8720dc296fff33b41f98ff83c6fcab4605db2eb5aaa5bc137aeb70a588408925083818180828009800909905062001d2c600052620025ba565b837f0a59a158e3eec2117e6e94e7f0e9decf18c3ffd5e1531a9219636158bbaf62f282089050837f06ec54c80381c052b58bf23b312ffd3ce2c4eba065420af8f4c23ed0075fd07b83089150837f118872dc832e0eb5476b56648e867ec8b09340f7a7bcb1b4962f0ff9ed1f9d018408925083818180828009800909905062001db7600052620025ba565b837f13d69fa127d834165ad5c7cba7ad59ed52e0b0f0e42d7fea95e1906b520921b182089050837f169a177f63ea681270b1c6877a73d21bde143942fb71dc55fd8a49f19f10c77b83089150837f04ef51591c6ead97ef42f287adce40d93abeb032b922f66ffb7e9a5a7450544d8408925083818180828009800909905062001e42600052620025ba565b837f256e175a1dc079390ecd7ca703fb2e3b19ec61805d4f03ced5f45ee6dd0f69ec82089050837f30102d28636abd5fe5f2af412ff6004f75cc360d3205dd2da002813d3e2ceeb283089150837f10998e42dfcd3bbf1c0714bc73eb1bf40443a3fa99bef4a31fd31be182fcc7928408925083818180828009800909905062001ecd600052620025ba565b837f193edd8e9fcf3d7625fa7d24b598a1d89f3362eaf4d582efecad76f879e3686082089050837f18168afd34f2d915d0368ce80b7b3347d1c7a561ce611425f2664d7aa51f0b5d83089150837f29383c01ebd3b6ab0c017656ebe658b6a328ec77bc33626e29e2e95b33ea61118408925083818180828009800909905062001f58600052620025ba565b837f10646d2f2603de39a1f4ae5e7771a64a702db6e86fb76ab600bf573f9010c71182089050837f0beb5e07d1b27145f575f1395a55bf132f90c25b40da7b3864d0242dcb1117fb83089150837f16d685252078c133dc0d3ecad62b5c8830f95bb2e54b59abdffbf018d96fa3368408925083818180828009800909905062001fe3600052620025ba565b837f0a6abd1d833938f33c74154e0404b4b40a555bbbec21ddfafd672dd62047f01a82089050837f1a679f5d36eb7b5c8ea12a4c2dedc8feb12dffeec450317270a6f19b34cf186083089150837f0980fb233bd456c23974d50e0ebfde4726a423eada4e8f6ffbc7592e3f1b93d6840892508381818082800980090990506200206e600052620025ba565b837f161b42232e61b84cbf1810af93a38fc0cece3d5628c9282003ebacb5c312c72b82089050837f0ada10a90c7f0520950f7d47a60d5e6a493f09787f1564e5d09203db47de1a0b83089150837f1a730d372310ba82320345a29ac4238ed3f07a8a2b4e121bb50ddb9af407f45184089250838181808280098009099050620020f9600052620025ba565b837f2c8120f268ef054f817064c369dda7ea908377feaba5c4dffbda10ef58e8c55682089050837f1c7c8824f758753fa57c00789c684217b930e95313bcb73e6e7b8649a4968f7083089150837f2cd9ed31f5f8691c8e39e4077a74faa0f400ad8b491eb3f7b47b27fa3fd1cf778408925083818180828009800909905062002184600052620025ba565b837f23ff4f9d46813457cf60d92f57618399a5e022ac321ca550854ae23918a22eea82089050837f09945a5d147a4f66ceece6405dddd9d0af5a2c5103529407dff1ea58f180426d83089150837f188d9c528025d4c2b67660c6b771b90f7c7da6eaa29d3f268a6dd223ec6fc630840892508381818082800980090990506200220f600052620025ba565b837f3050e37996596b7f81f68311431d8734dba7d926d3633595e0c0d8ddf4f0f47f82089050837f15af1169396830a91600ca8102c35c426ceae5461e3f95d89d829518d30afd7883089150837f1da6d09885432ea9a06d9f37f873d985dae933e351466b2904284da3320d8acc840892508381818082800980090990506200229a600052620025ba565b837f2796ea90d269af29f5f8acf33921124e4e4fad3dbe658945e546ee411ddaa9cb82089050837f202d7dd1da0f6b4b0325c8b3307742f01e15612ec8e9304a7cb0319e01d32d6083089150837f096d6790d05bb759156a952ba263d672a2d7f9c788f4c831a29dace4c0f8be5f8408925083818180828009800909905062002325600052620025ba565b837f054efa1f65b0fce283808965275d877b438da23ce5b13e1963798cb1447d25a482089050837f1b162f83d917e93edb3308c29802deb9d8aa690113b2e14864ccf6e18e4165f183089150837f21e5241e12564dd6fd9f1cdd2a0de39eedfefc1466cc568ec5ceb745a0506edc84089250838181808280098009099050838281808280098009099150838381808280098009099250620023c8600052620025ba565b837f1cfb5662e8cf5ac9226a80ee17b36abecb73ab5f87e161927b4349e10e4bdf0882089050837f0f21177e302a771bbae6d8d1ecb373b62c99af346220ac0129c53f666eb2410083089150837f1671522374606992affb0dd7f71b12bec4236aede6290546bcef7e1f515c2320840892508381818082800980090990508382818082800980090991508383818082800980090992506200246b600052620025ba565b837f0fa3ec5b9488259c2eb4cf24501bfad9be2ec9e42c5cc8ccd419d2a692cad87082089050837f193c0e04e0bd298357cb266c1506080ed36edce85c648cc085e8c57b1ab54bba83089150837f102adf8ef74735a27e9128306dcbc3c99f6f7291cd406578ce14ea2adaba68f8840892508381818082800980090990508382818082800980090991508383818082800980090992506200250e600052620025ba565b837f0fe0af7858e49859e2a54d6f1ad945b1316aa24bfbdd23ae40a6d0cb70c3eab182089050837f216f6717bbc7dedb08536a2220843f4e2da5f1daa9ebdefde8a5ea7344798d2283089150837f1da55cc900f0d21f4a3e694391918a1b3c23b2ac773c6b3ef88e2e422832516184089250838181808280098009099050838281808280098009099150838381808280098009099250620025b1600052620025ba565b60005260206000f35b8360205182098460405184098591088460605185098591088460805183098560a05185098691088560c05186098691088560e0518409866101005186098791088661012051870987910894509250905060005156";

    function deployPoseidon2() public returns (address) {
        return CodeDeployer.deployCode(poseidon2Code);
    }
}

contract Portal_Initializer is L2OutputOracle_Initializer, Poseidon2Deployer {
    ZKMerkleTrie zkMerkleTrie;

    // Test target
    KromaPortal portalImpl;
    KromaPortal portal;

    event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success);
    event WithdrawalProven(
        bytes32 indexed withdrawalHash,
        address indexed from,
        address indexed to
    );

    function setUp() public virtual override {
        super.setUp();

        vm.deal(trusted, requiredBondAmount * 100);
        vm.prank(trusted);
        pool.deposit{ value: trusted.balance }();

        zkMerkleTrie = new ZKMerkleTrie(deployPoseidon2());

        portalImpl = new KromaPortal({
            _l2Oracle: oracle,
            _validatorPool: address(pool),
            _guardian: guardian,
            _paused: true,
            _config: systemConfig,
            _zkMerkleTrie: zkMerkleTrie
        });
        Proxy proxy = new Proxy(multisig);
        vm.prank(multisig);
        proxy.upgradeToAndCall(
            address(portalImpl),
            abi.encodeWithSelector(KromaPortal.initialize.selector, false)
        );
        portal = KromaPortal(payable(address(proxy)));
        vm.label(address(portal), "KromaPortal");
    }
}

contract Messenger_Initializer is Portal_Initializer {
    L1CrossDomainMessenger L1Messenger;
    L2CrossDomainMessenger L2Messenger =
        L2CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER);

    event SentMessage(
        address indexed target,
        address indexed sender,
        uint256 value,
        bytes message,
        uint256 messageNonce,
        uint256 gasLimit
    );

    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash
    );

    event RelayedMessage(bytes32 indexed msgHash);
    event FailedRelayedMessage(bytes32 indexed msgHash);

    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 mint,
        uint256 value,
        uint64 gasLimit,
        bool isCreation,
        bytes data
    );

    event WhatHappened(bool success, bytes returndata);

    function setUp() public virtual override {
        super.setUp();

        // Setup implementation
        L1CrossDomainMessenger L1MessengerImpl = new L1CrossDomainMessenger(portal);

        // Setup proxy
        Proxy proxy = new Proxy(multisig);
        vm.prank(multisig);
        proxy.upgradeToAndCall(
            address(L1MessengerImpl),
            abi.encodeCall(L1Messenger.initialize, ())
        );
        L1Messenger = L1CrossDomainMessenger(address(proxy));

        vm.etch(
            Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            address(new L2CrossDomainMessenger(address(L1Messenger))).code
        );

        L2Messenger.initialize();

        // Label addresses
        vm.label(address(L1MessengerImpl), "L1CrossDomainMessenger_Impl");
        vm.label(Predeploys.L2_CROSS_DOMAIN_MESSENGER, "L2CrossDomainMessenger");

        vm.label(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)),
            "L1CrossDomainMessenger_aliased"
        );
    }
}

contract Bridge_Initializer is Messenger_Initializer {
    L1StandardBridge L1Bridge;
    L2StandardBridge L2Bridge;
    KromaMintableERC20Factory L2TokenFactory;
    KromaMintableERC20Factory L1TokenFactory;
    ERC20 L1Token;
    ERC20 BadL1Token;
    KromaMintableERC20 L2Token;
    ERC20 NativeL2Token;
    ERC20 BadL2Token;
    KromaMintableERC20 RemoteL1Token;

    event DepositFailed(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes data);

    event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes data);

    event ERC20BridgeInitiated(
        address indexed localToken,
        address indexed remoteToken,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event ERC20BridgeFinalized(
        address indexed localToken,
        address indexed remoteToken,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    function setUp() public virtual override {
        super.setUp();

        vm.label(Predeploys.L2_STANDARD_BRIDGE, "L2StandardBridge");
        vm.label(Predeploys.KROMA_MINTABLE_ERC20_FACTORY, "KromaMintableERC20Factory");

        // Deploy the L1 bridge and initialize it with the address of the
        // L1CrossDomainMessenger
        L1StandardBridge L1Bridge_Impl = new L1StandardBridge(payable(address(L1Messenger)));
        Proxy proxy = new Proxy(multisig);
        vm.prank(multisig);
        proxy.upgradeTo(address(L1Bridge_Impl));

        L1Bridge = L1StandardBridge(payable(address(proxy)));

        vm.label(address(proxy), "L1StandardBridge_Proxy");
        vm.label(address(L1Bridge_Impl), "L1StandardBridge_Impl");

        // Deploy the L2StandardBridge, move it to the correct predeploy
        // address and then initialize it
        L2StandardBridge l2B = new L2StandardBridge(payable(proxy));
        vm.etch(Predeploys.L2_STANDARD_BRIDGE, address(l2B).code);
        L2Bridge = L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE));

        // Set up the L2 mintable token factory
        KromaMintableERC20Factory factory = new KromaMintableERC20Factory(
            Predeploys.L2_STANDARD_BRIDGE
        );
        vm.etch(Predeploys.KROMA_MINTABLE_ERC20_FACTORY, address(factory).code);
        L2TokenFactory = KromaMintableERC20Factory(Predeploys.KROMA_MINTABLE_ERC20_FACTORY);

        L1Token = new ERC20("Native L1 Token", "L1T");

        // Deploy the L2 ERC20 now
        L2Token = KromaMintableERC20(
            L2TokenFactory.createKromaMintableERC20(
                address(L1Token),
                string(abi.encodePacked("L2-", L1Token.name())),
                string(abi.encodePacked("L2-", L1Token.symbol()))
            )
        );

        BadL2Token = KromaMintableERC20(
            L2TokenFactory.createKromaMintableERC20(
                address(1),
                string(abi.encodePacked("L2-", L1Token.name())),
                string(abi.encodePacked("L2-", L1Token.symbol()))
            )
        );

        NativeL2Token = new ERC20("Native L2 Token", "L2T");
        L1TokenFactory = new KromaMintableERC20Factory(address(L1Bridge));

        RemoteL1Token = KromaMintableERC20(
            L1TokenFactory.createKromaMintableERC20(
                address(NativeL2Token),
                string(abi.encodePacked("L1-", NativeL2Token.name())),
                string(abi.encodePacked("L1-", NativeL2Token.symbol()))
            )
        );

        BadL1Token = KromaMintableERC20(
            L1TokenFactory.createKromaMintableERC20(
                address(1),
                string(abi.encodePacked("L1-", NativeL2Token.name())),
                string(abi.encodePacked("L1-", NativeL2Token.symbol()))
            )
        );
    }
}

contract ERC721Bridge_Initializer is Messenger_Initializer {
    L1ERC721Bridge L1Bridge;
    L2ERC721Bridge L2Bridge;

    function setUp() public virtual override {
        super.setUp();

        // Deploy the L1ERC721Bridge.
        L1Bridge = new L1ERC721Bridge(address(L1Messenger), Predeploys.L2_ERC721_BRIDGE);

        // Deploy the implementation for the L2ERC721Bridge and etch it into the predeploy address.
        vm.etch(
            Predeploys.L2_ERC721_BRIDGE,
            address(new L2ERC721Bridge(Predeploys.L2_CROSS_DOMAIN_MESSENGER, address(L1Bridge)))
                .code
        );

        // Set up a reference to the L2ERC721Bridge.
        L2Bridge = L2ERC721Bridge(Predeploys.L2_ERC721_BRIDGE);

        // Label the L1 and L2 bridges.
        vm.label(address(L1Bridge), "L1ERC721Bridge");
        vm.label(address(L2Bridge), "L2ERC721Bridge");
    }
}

contract Colosseum_Initializer is Portal_Initializer {
    uint256 immutable CHAIN_ID = 901;
    bytes32 immutable DUMMY_HASH =
        hex"a1235b834d6f1f78f78bc4db856fbc49302cce2c519921347600693021e087f7";
    uint256 immutable MAX_TXS = 100;

    // Test target
    Colosseum colosseumImpl;

    ZKVerifier zkVerifier;
    ZKVerifier zkVerifierImpl;

    SecurityCouncil securityCouncilImpl;
    SecurityCouncil securityCouncil;

    uint256[] segmentsLengths;

    function setUp() public virtual override {
        // Deploy the ZKVerifier
        // Chain ID 901
        Proxy verifierProxy = new Proxy(multisig);
        zkVerifier = ZKVerifier(payable(address(verifierProxy)));
        zkVerifierImpl = new ZKVerifier({
            _hashScalar: 9621709039943366658724195007914543214191235696182878089308100015953499829622,
            _m56Px: 9806184489513941466540325197817869206227113846804865003333373964174130146183,
            _m56Py: 20099975678253040567150653826633105770489605161161511574764394101710764414042
        });
        vm.prank(multisig);
        verifierProxy.upgradeTo(address(zkVerifierImpl));

        // case - L2OutputOracle submissionInterval == 1800
        segmentsLengths.push(9);
        segmentsLengths.push(6);
        segmentsLengths.push(10);
        segmentsLengths.push(6);

        Proxy proxy = new Proxy(multisig);
        colosseum = Colosseum(payable(address(proxy)));

        // Init L2OutputOracle after Colosseum contract deployment
        super.setUp();

        // Deploy the SecurityCouncil (after Colosseum contract deployment)
        securityCouncil = SecurityCouncil(address(new Proxy(multisig)));
        securityCouncilImpl = new SecurityCouncil(
            address(colosseum),
            payable(address(upgradeGovernor))
        );
        vm.prank(multisig);
        toProxy(address(securityCouncil)).upgradeTo(address(securityCouncilImpl));

        colosseumImpl = new Colosseum({
            _l2Oracle: oracle,
            _zkVerifier: zkVerifier,
            _submissionInterval: submissionInterval,
            _creationPeriodSeconds: 6 days,
            _bisectionTimeout: 30 minutes,
            _provingTimeout: 1 hours,
            _dummyHash: DUMMY_HASH,
            _maxTxs: MAX_TXS,
            _segmentsLengths: segmentsLengths,
            _securityCouncil: address(securityCouncil),
            _zkMerkleTrie: address(zkMerkleTrie)
        });
        vm.prank(multisig);
        proxy.upgradeToAndCall(
            address(colosseumImpl),
            abi.encodeCall(Colosseum.initialize, (segmentsLengths))
        );
    }
}

contract SecurityCouncil_Initializer is UpgradeGovernor_Initializer {
    address colosseum = makeAddr("colosseum");
    SecurityCouncil securityCouncil;
    SecurityCouncil securityCouncilImpl;

    function setUp() public virtual override {
        super.setUp();
        securityCouncil = SecurityCouncil(address(new Proxy(multisig)));
        securityCouncilImpl = new SecurityCouncil(colosseum, payable(address(upgradeGovernor)));
        vm.prank(multisig);
        toProxy(address(securityCouncil)).upgradeTo(address(securityCouncilImpl));
    }
}

contract FFIInterface is Test {
    function getProveWithdrawalTransactionInputs(
        Types.WithdrawalTransaction memory _tx
    ) external returns (bytes32, bytes32, bytes32, bytes32, bytes[] memory) {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "getProveWithdrawalTransactionInputs";
        cmds[3] = vm.toString(_tx.nonce);
        cmds[4] = vm.toString(_tx.sender);
        cmds[5] = vm.toString(_tx.target);
        cmds[6] = vm.toString(_tx.value);
        cmds[7] = vm.toString(_tx.gasLimit);
        cmds[8] = vm.toString(_tx.data);

        bytes memory result = vm.ffi(cmds);
        (
            bytes32 stateRoot,
            bytes32 storageRoot,
            bytes32 outputRoot,
            bytes32 withdrawalHash,
            bytes[] memory withdrawalProof
        ) = abi.decode(result, (bytes32, bytes32, bytes32, bytes32, bytes[]));

        return (stateRoot, storageRoot, outputRoot, withdrawalHash, withdrawalProof);
    }

    function hashCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external returns (bytes32) {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashCrossDomainMessage";
        cmds[3] = vm.toString(_nonce);
        cmds[4] = vm.toString(_sender);
        cmds[5] = vm.toString(_target);
        cmds[6] = vm.toString(_value);
        cmds[7] = vm.toString(_gasLimit);
        cmds[8] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashWithdrawal(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external returns (bytes32) {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashWithdrawal";
        cmds[3] = vm.toString(_nonce);
        cmds[4] = vm.toString(_sender);
        cmds[5] = vm.toString(_target);
        cmds[6] = vm.toString(_value);
        cmds[7] = vm.toString(_gasLimit);
        cmds[8] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashOutputRootProof(
        bytes32 _version,
        bytes32 _stateRoot,
        bytes32 _messagePasserStorageRoot,
        bytes32 _blockhash,
        bytes32 _nextBlockhash
    ) external returns (bytes32) {
        string[] memory cmds = new string[](8);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashOutputRootProof";
        cmds[3] = Strings.toHexString(uint256(_version));
        cmds[4] = Strings.toHexString(uint256(_stateRoot));
        cmds[5] = Strings.toHexString(uint256(_messagePasserStorageRoot));
        cmds[6] = Strings.toHexString(uint256(_blockhash));
        cmds[7] = Strings.toHexString(uint256(_nextBlockhash));

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashDepositTransaction(
        address _from,
        address _to,
        uint256 _mint,
        uint256 _value,
        uint64 _gas,
        bytes memory _data,
        uint64 _logIndex
    ) external returns (bytes32) {
        string[] memory cmds = new string[](11);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashDepositTransaction";
        cmds[3] = "0x0000000000000000000000000000000000000000000000000000000000000000";
        cmds[4] = vm.toString(_logIndex);
        cmds[5] = vm.toString(_from);
        cmds[6] = vm.toString(_to);
        cmds[7] = vm.toString(_mint);
        cmds[8] = vm.toString(_value);
        cmds[9] = vm.toString(_gas);
        cmds[10] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function encodeDepositTransaction(
        Types.UserDepositTransaction calldata txn
    ) external returns (bytes memory) {
        string[] memory cmds = new string[](12);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "encodeDepositTransaction";
        cmds[3] = vm.toString(txn.from);
        cmds[4] = vm.toString(txn.to);
        cmds[5] = vm.toString(txn.value);
        cmds[6] = vm.toString(txn.mint);
        cmds[7] = vm.toString(txn.gasLimit);
        cmds[8] = vm.toString(txn.isCreation);
        cmds[9] = vm.toString(txn.data);
        cmds[10] = vm.toString(txn.l1BlockHash);
        cmds[11] = vm.toString(txn.logIndex);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes));
    }

    function encodeCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external returns (bytes memory) {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "encodeCrossDomainMessage";
        cmds[3] = vm.toString(_nonce);
        cmds[4] = vm.toString(_sender);
        cmds[5] = vm.toString(_target);
        cmds[6] = vm.toString(_value);
        cmds[7] = vm.toString(_gasLimit);
        cmds[8] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes));
    }

    function decodeVersionedNonce(uint256 nonce) external returns (uint256, uint256) {
        string[] memory cmds = new string[](4);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "decodeVersionedNonce";
        cmds[3] = vm.toString(nonce);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (uint256, uint256));
    }

    function getMerkleTrieFuzzCase(
        string memory variant
    ) external returns (bytes32, bytes memory, bytes memory, bytes[] memory) {
        string[] memory cmds = new string[](3);
        cmds[0] = "./scripts/go-ffi/go-ffi";
        cmds[1] = "trie";
        cmds[2] = variant;

        return abi.decode(vm.ffi(cmds), (bytes32, bytes, bytes, bytes[]));
    }
}

// Used for testing a future upgrade beyond the current implementations.
// We include some variables so that we can sanity check accessing storage values after an upgrade.
contract NextImpl is Initializable {
    // Initializable occupies the zero-th slot.
    bytes32 slot1;
    bytes32[19] __gap;
    bytes32 slot21;
    bytes32 public constant slot21Init = bytes32(hex"1337");

    function initialize() public reinitializer(2) {
        // Slot21 is unused by an of our upgradeable contracts.
        // This is used to verify that we can access this value after an upgrade.
        slot21 = slot21Init;
    }
}

contract Reverter {
    fallback() external {
        revert();
    }
}

// Useful for testing reentrancy guards
contract CallerCaller {
    event WhatHappened(bool success, bytes returndata);

    fallback() external {
        (bool success, bytes memory returndata) = msg.sender.call(msg.data);
        emit WhatHappened(success, returndata);
        assembly {
            switch success
            case 0 {
                revert(add(returndata, 0x20), mload(returndata))
            }
            default {
                return(add(returndata, 0x20), mload(returndata))
            }
        }
    }
}

// Used for testing the `CrossDomainMessenger`'s per-message reentrancy guard.
contract ConfigurableCaller {
    bool doRevert = true;
    address target;
    bytes payload;

    event WhatHappened(bool success, bytes returndata);

    /**
     * @notice Call the configured target with the configured payload OR revert.
     */
    function call() external {
        if (doRevert) {
            revert("ConfigurableCaller: revert");
        } else {
            (bool success, bytes memory returndata) = address(target).call(payload);
            emit WhatHappened(success, returndata);
            assembly {
                switch success
                case 0 {
                    revert(add(returndata, 0x20), mload(returndata))
                }
                default {
                    return(add(returndata, 0x20), mload(returndata))
                }
            }
        }
    }

    /**
     * @notice Set whether or not to have `call` revert.
     */
    function setDoRevert(bool _doRevert) external {
        doRevert = _doRevert;
    }

    /**
     * @notice Set the target for the call made in `call`.
     */
    function setTarget(address _target) external {
        target = _target;
    }

    /**
     * @notice Set the payload for the call made in `call`.
     */
    function setPayload(bytes calldata _payload) external {
        payload = _payload;
    }

    /**
     * @notice Fallback function that reverts if `doRevert` is true.
     *         Otherwise, it does nothing.
     */
    fallback() external {
        if (doRevert) {
            revert("ConfigurableCaller: revert");
        }
    }
}
