// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Hashing } from "../../libraries/Hashing.sol";
import { Types } from "../../libraries/Types.sol";
import { RLPWriter } from "../../libraries/rlp/RLPWriter.sol";
import { Colosseum } from "../../L1/Colosseum.sol";

library ColosseumTestData {
    uint256 internal constant INVALID_BLOCK_NUMBER = 3000;
    bytes32 internal constant PREV_OUTPUT_ROOT =
        0xccac4c06a5ea0106ec15bb374df8c97084492ffd527764eae88a794200432c9e;
    bytes32 internal constant TARGET_OUTPUT_ROOT =
        0xb0617c5bfa546a0b6cc576e5d270731c866c542f694664636648b9469135d353;

    function outputRootProof()
        internal
        pure
        returns (Types.OutputRootProof memory, Types.OutputRootProof memory)
    {
        Types.OutputRootProof memory src = Types.OutputRootProof({
            version: bytes32(uint256(1)),
            stateRoot: 0x0d987e48a7951caba0bcbb4f7d6049ecb41f98f68acdd724219f76a8ca5b720e,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0xfcf55897ef0e68aa8b3b39b1c8d6650a16663b967ad9d01b5d25d1c2bf39e144,
            nextBlockHash: 0xf37a1d0486fe80ccfa63a519512d30c240ade862fa3faca49ede025209b39471
        });

        Types.OutputRootProof memory dst = Types.OutputRootProof({
            version: bytes32(uint256(1)),
            stateRoot: 0x1931d7bf7e9da61441a4145ec18807680141d6a37d68890f9c786e31a364d1ae,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0xf37a1d0486fe80ccfa63a519512d30c240ade862fa3faca49ede025209b39471,
            nextBlockHash: bytes32(abi.encode())
        });

        return (src, dst);
    }

    function publicInput() internal pure returns (Types.PublicInput memory) {
        bytes32[] memory txHashes = new bytes32[](1);
        txHashes[0] = 0xa856617fcd81b386d003c195fc82ae03a79c6d7538e1763324b0eb7d207207d0;

        return
            Types.PublicInput({
                coinbase: 0x0000000000000000000000000000000000000000,
                timestamp: 0x645db07e,
                number: 0xbb8,
                difficulty: 0x0,
                gasLimit: 0xe4e1c0,
                baseFee: 0x9,
                chainId: 901,
                transactionsRoot: 0x8eb182b0a61fed8cc5f190d4c1d87dbe5e2df9a346cb36d26b8d71639ca92b37,
                stateRoot: 0x1931d7bf7e9da61441a4145ec18807680141d6a37d68890f9c786e31a364d1ae,
                txHashes: txHashes
            });
    }

    function publicInputHash(Colosseum colosseum) internal view returns (bytes32) {
        (Types.OutputRootProof memory srcOutputRootProof, ) = outputRootProof();
        Types.PublicInput memory pi = publicInput();

        uint256 maxTxs = colosseum.MAX_TXS();

        bytes32[] memory dummyHashes;
        if (pi.txHashes.length < maxTxs) {
            dummyHashes = Hashing.generateDummyHashes(colosseum.DUMMY_HASH(), maxTxs - pi.txHashes.length);
        }

        return
            Hashing.hashPublicInput(
                srcOutputRootProof.stateRoot,
                pi,
                colosseum.CHAIN_ID(),
                dummyHashes
            );
    }

    function blockHeaderRLP() internal pure returns (Types.BlockHeaderRLP memory) {
        return
            Types.BlockHeaderRLP({
                parentHash: RLPWriter.writeBytes(
                    hex"fcf55897ef0e68aa8b3b39b1c8d6650a16663b967ad9d01b5d25d1c2bf39e144"
                ),
                uncleHash: RLPWriter.writeBytes(
                    hex"1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"
                ),
                receiptsRoot: RLPWriter.writeBytes(
                    hex"ff07e961a994197df0f9406d373ef5d124e6febecf85a74ca0bcd436b2c5ad0e"
                ),
                logsBloom: RLPWriter.writeBytes(
                    hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
                ),
                gasUsed: RLPWriter.writeUint(0xf4240),
                extraData: RLPWriter.writeBytes(hex""),
                mixHash: RLPWriter.writeBytes(
                    hex"0000000000000000000000000000000000000000000000000000000000000000"
                ),
                nonce: RLPWriter.writeBytes(hex"0000000000000000"),
                withdrawalsRoot: hex""
            });
    }

    struct ProofPair {
        uint256[] proof;
        uint256[] pair;
    }

    function proofAndPair(bytes32 _publicInputHash) internal pure returns (ProofPair memory pp) {
        pp.proof = new uint256[](145);
        pp.proof[0] = 0x2db84b4be2adf33a9442a48e35273b442afb13491b53b5ddda9d28afa4bcc94f;
        pp.proof[1] = 0x250d0813a637b6894bcfda982e940ace442939d02d7d308bd269da5bc0d90c75;
        pp.proof[2] = 0x2490d94e7a9603f18a9a56bac4a4b9e8bade53765f36c8f47538723c99327c4a;
        pp.proof[3] = 0x2a9abdbc434023d970e442cc2b7faa601e2ab1056fd84d436e0f52fcbd166780;
        pp.proof[4] = 0x2cc4cf0d5d6183bb6e5d08abfe3f8ec7657979e9f36c63bc11bdbdc12f20f37c;
        pp.proof[5] = 0x13815ac8dd7a6a9fa4d8fc7ec4a6a873673edf90aade9e0dbe04ce14c42dda;
        pp.proof[6] = 0x04aa40d93c70c51f3bbf328d859be6a26fa47d42e20b90d588327735848403ad;
        pp.proof[7] = 0x0df93b34368a431946abd12d067ce6d53931b53bac92d60694caf306e72d2a50;
        pp.proof[8] = 0x11bab1ba91e7fc49ce58a1289a6a2d7fa4d93cf04cfc371edcb7b11c4a2b4bfc;
        pp.proof[9] = 0x248829f768e0b1cd3cb41dd2b1482a9a05d58e458d3e025f86cc84c5d5da3fb4;
        pp.proof[10] = 0x184b5e7a6fb160ac14f2a53ce08b28954b43b223e976f46b64c8995f84def8c3;
        pp.proof[11] = 0x20b15c5b8b8b44f49235ed5c6e1f8666188d11351b77d46b2856c4a33b48fec0;
        pp.proof[12] = 0x226cd7be86dfc4bca11d05ba1e9ba1ebef9d1616d58a17f6d3fec001d0067c30;
        pp.proof[13] = 0x0a261d4c6019b229b5294470107816b2f79cb244b424e15c1778a257f9cd4790;
        pp.proof[14] = 0x1852f911c1ea051ea2a50f9e5f39dee0852d4badbad834ac148b8b369846e907;
        pp.proof[15] = 0x2288c4537d56b6a28fa0a2cfe09766e93e90f2eb92c81688ea2cdeef0afc72ba;
        pp.proof[16] = 0x0d81a02023adaf01abe3590c13673e7c581aef8bd581a80aee86761f71f90b21;
        pp.proof[17] = 0x10d44a001fee7e0dbc06b6b57e4d8393437ed0412a449a4801f22058b67b0bd1;
        pp.proof[18] = 0x19332e828f4096178fe74522eabce81df92b68210b1b16d09778c3f9e687b1b6;
        pp.proof[19] = 0x13bf83223bc5e1073739b33ec41a99929db83dc3a00ab5129f35fb93c63f4f0c;
        pp.proof[20] = 0x1e382c6a607aa3d0a86b3c7378a6bec5e83e322efc33d87ba95b8edc656c2ffe;
        pp.proof[21] = 0x1d4e41c3f10623e542a1a554c1a0599d921ea0583a80c5f5801a6be0b36ff769;
        pp.proof[22] = 0x14798c2f163d17e71bd1b7f9748a544729c7e5dc1f8b69b845b7e02aede0e411;
        pp.proof[23] = 0x022c93db9a989e57489a594ea5a1c304e29e163f2adcad0a5010425370adb624;
        pp.proof[24] = 0x28fe543b1afdb31116b0c5e53e91a010dfa016703b8e9b04316029db1a11510c;
        pp.proof[25] = 0x124c6ee5db1ca2bdd3f962f2e3cce58a10650ac6b967829530e687f000c214a2;
        pp.proof[26] = 0x2d60c72ba9a3a0d187491b06c4df0c41e130a5d2df33920cdab35fe93f578d4e;
        pp.proof[27] = 0x2dc1271c3caa3a5864fe162ba3522ed39cd7aef06064fd03df0cfd54461999d0;
        pp.proof[28] = 0xfa4fef5681a8cb9644672210a01bc2a58ba899993de49f0a77134349ea9761;
        pp.proof[29] = 0x1960979499d99f991de05360b62b03ded924e108efaa4801cc79dbd4f85278f7;
        pp.proof[30] = 0x115a6687310ea2d9c49a9f85ca3a0b99d1a8ec6a18330743bbca6e0347664871;
        pp.proof[31] = 0x2eb8b744498591ffdc3e48fca13343accb0031e75767b8eb9c5ba01a862c413a;
        pp.proof[32] = 0x0114854f677f1947572248f21e615e4bbdbfee6ce2fd929d59e9e08d8c923d77;
        pp.proof[33] = 0x131046b319403852d7c020cb46caa55f725f050065c9188be78f29994a92281f;
        pp.proof[34] = 0x0188fee4217392d9bb7706d6a69f2e74351495b76fd73b77ec3c869ac9c66924;
        pp.proof[35] = 0x1cc8b335c07d157f2d55d7c2d38345bc4d73da18b5b6ca02b5d4a31fe353d11e;
        pp.proof[36] = 0x14f3a9e0b502d8ac7f64ef0a7d05429f67b6ebadd76bc5b61c8ab52332dacf99;
        pp.proof[37] = 0x19945e1f83400748b639727337349beefead70716cdedd9478b78ad228043cd9;
        pp.proof[38] = 0x12744893971ccde3ace7c85cfc2f239875f2a8c146f90e85b690d5b404beb474;
        pp.proof[39] = 0x04ef6e29fd15a5abe467245d1fe9a2d026fb545e41ccc16ac8b177d0421bd6b3;
        pp.proof[40] = 0x23d6063a0ced33b3b6425a328459edc1b65a84a03c017197f811f8739420ece7;
        pp.proof[41] = 0x1b3a63370c5c8870929204d5ae96fa6c94370c8c2b801090d87edd82c739ebc2;
        pp.proof[42] = 0x0260fc633ac72379d6e4114feaf38fb3fa3384f39a8e0cad4c825d2f5ebd8dc5;
        pp.proof[43] = 0x1a1e87a3b3337abc1d56a9f0260106ae1033689397b9e25b3847e2fafb69f649;
        pp.proof[44] = 0x21caaab53935e64ca343bf4aae9e01bb59cf8523d41757af3241a396be50f7df;
        pp.proof[45] = 0x01989a630f1dacc2e93e790600027e77f71b0c4cf09bb2c3a2407fab32951850;
        pp.proof[46] = 0x2b6b6beda62c3b6aa93194d39ab4f0f8709f7324496b4f4cbab666c6f15f0482;
        pp.proof[47] = 0x2d245ac58fb7feb64de4d83398fae025dc249eae59ef4969cce90e68803e68e1;
        pp.proof[48] = 0x1dc774f206d611c88cecfd81459c0e0e2778b1ea9e23b48d5a21a260a85b882c;
        pp.proof[49] = 0x29c93dd2e91809ccecf63d2dba8b69e3f3d4f2d1846590bdec28055c150d5665;
        pp.proof[50] = 0x1ba848ee49f53210dc1503f36bf92982628d9760d3c06fae021f50e99403b943;
        pp.proof[51] = 0x11d2a769e44495989e1ec075249c7cb6d8e26d12b834355b895017a7f92ffb38;
        pp.proof[52] = 0x0e3874c7ce21a8c229097fe511fcfb145e8f982f9a7a89efcc658a744ae9667a;
        pp.proof[53] = 0x227b6c05d9ddf0e45d8072e83105ca6054afa7f9343aa384ee993871be72a63b;
        pp.proof[54] = 0x1655bc2672198f8eed1944b9c5c10cc83da835fa6b3b65027539b69cc152dc6a;
        pp.proof[55] = 0x5dd989cf88d741e5cbeb6c7406ef33dad7667d9acbac2ecede2a546993ae4b;
        pp.proof[56] = 0x00;
        pp.proof[57] = 0x00;
        pp.proof[58] = 0x2d76c2e02aa6d2aaeb5bca6079c4701825ac73336b66ed69e508bae79f46c4f1;
        pp.proof[59] = 0x201a372c82f860bd5691cf0062560f6cbddf0ef1ee7e0d1762590ccfcc7c8597;
        pp.proof[60] = 0x298980f45d7a9d8f7de17d8b2992a1eb866111fbcc5308ad1f42ee97f666a8da;
        pp.proof[61] = 0x1112ccfc43917a7fadf94a93e1fd20671316e3abb44f1116170feb413ec7d104;
        pp.proof[62] = 0x05e2d7cd66f8c1ca5336b26f2e98f183a627b198147e8085022c3a940dd3196b;
        pp.proof[63] = 0x215766098fb23c74696d3706393ddb9a396ca17758bd63cdbd472157fd9f29df;
        pp.proof[64] = 0x05748fdf0f3353b25462c3365e0d7d5e5ef3fc42760e74c9bdb1ce33af18766d;
        pp.proof[65] = 0x026a7759962fa812bfbadcf5b0f57dcb70d6b693e87faf9078e640a0f58d0c73;
        pp.proof[66] = 0x0df6fc918932db1f885601a8bcd8765f2ae47b3ab8e6057b8032394c97ab2a0a;
        pp.proof[67] = 0x01f7c49650d7ad562aa1f40bc068d463d51eb92664221e3bea5d37df79de23bd;
        pp.proof[68] = 0x17f4510bfc9f8272f0920ab367148da96589621ad49087fe462c67c5000424c5;
        pp.proof[69] = 0x2e09167ff264977e20b815fa6af5074d501132a36e7828d60027a287589adddd;
        pp.proof[70] = 0x2163af6bc43129a82f8ef845174977866bfc62fe289bcff07e65601820d87547;
        pp.proof[71] = 0x1b176a40533b9244b19a5022763eb6716fe57be2f19ca57bd5b3075924e66e2c;
        pp.proof[72] = 0x0b0faa8423dfbeab0f2a8c56330e0b6de0101241f1a0bc25d44b5f48d73c60d3;
        pp.proof[73] = 0x1994b2e08132822991c92222b71d7425402d4aa38114ea43aed804cc8fa5850f;
        pp.proof[74] = 0x02e8759280468e1fc2b831734154a63238b9e8766b67c57d34e63d692963e94e;
        pp.proof[75] = 0x07b5df97dff873ffd28a1aac69d0cf0eff3d2d327a19a97908919133b9102538;
        pp.proof[76] = 0x2e05d794ef0fa654f4319bdf071bc2e18b7c84ab71a5526f7027035183f5080c;
        pp.proof[77] = 0x210ecfd49532a3beeda06210ca026b1558dcce83e4468cfd174ffe100858259e;
        pp.proof[78] = 0x29c9ac5471e1d9b9ef46ddb5eede61f52684ce4cb82886fa478f45696d9de7d7;
        pp.proof[79] = 0x15680777ad270e04d48d69b00c93d59b9456d69b6378a49c107535dd0f66b5d2;
        pp.proof[80] = 0x2cd8ada9d3bfa05bcd4ac63ecbd16a2d0f6e8f3041a4477d6d9b35f37edd4a37;
        pp.proof[81] = 0x0aaf4bc683e3967d07c96343a25648b07d237ae33d1da5971b77ae9dae2955df;
        pp.proof[82] = 0x2cc473cead4472dbea390fa73bb61766c94090c6e34d8da1db57dee735a3efd8;
        pp.proof[83] = 0x194fb396d45beefe932befe4f67e23431f5f529f2005b49ae553f30495ef9353;
        pp.proof[84] = 0x25c2bf9c4d955b4a71bba35c46c114a2fa8e9ed171918e3dce6cf80410f87940;
        pp.proof[85] = 0x0b06ab766e7a8824c4506cfa82e33fc4c9255161d10d8caee641bebe1c146ec8;
        pp.proof[86] = 0x109d6860031accac7b6a3de0281ce548411a183e255489e2db47cec180d46222;
        pp.proof[87] = 0x15183c73e1a213e84c08ef462f386d0e8ed43b68aee9d3e60af90b4f25a740e1;
        pp.proof[88] = 0x21e05690a731c64c3de54ffc544f5fdf0f1983cc6c648438a9d549da56ec79bc;
        pp.proof[89] = 0x0b06ab766e7a8824c4506cfa82e33fc4c9255161d10d8caee641bebe1c146ec8;
        pp.proof[90] = 0x00;
        pp.proof[91] = 0x0b2c8caabaa2caf038b83bbfc7739f991d635de720c3e5cc895801b241933c23;
        pp.proof[92] = 0x2e93bb3a5dfbdcb8cbad6fd1bc85002769831bdf102db42140b79cd8f5ae851a;
        pp.proof[93] = 0x2af618448e46a8cc8ba91c7211c8bc37c8e1e8c5ec123a71b3ea684229a92061;
        pp.proof[94] = 0x19da634c22e60424bf6a9dd10bb28d6ec8d22aceceb1a21259b7ad512b80d7d8;
        pp.proof[95] = 0x19f0c7ab77bc981eeeb03defc3300bb3eb73eeb116368e353b37df4d1026bdd9;
        pp.proof[96] = 0x231bb3c25d644facb72a230652a00e3eb8c69d7e97d918a35c99fbeb96b283ae;
        pp.proof[97] = 0x25ce5cab38dd6fe45eb58a54db8965b229533b9026e629d01a0ff5e93fb36613;
        pp.proof[98] = 0x101f0fe2780f5df9659ceef59699edc843c63ea21d06ccb9de952896bae3367d;
        pp.proof[99] = 0x16186ed2512656f776c97498ce2576f28ebb54e8c86a0da7f41835264c4629ed;
        pp.proof[100] = 0x04e96ea6dc24ed0cec658845c1f4bcdbea3de9a796d69df9f59a2612a1ab9607;
        pp.proof[101] = 0x2a54185888e01ef4dafc8517bed9edabaafec3fc81124b505ae21ac43001319f;
        pp.proof[102] = 0x2a3c6af9332905821d87edbadaa97c3fd17db5d4cbe9757210d9735f0ceec752;
        pp.proof[103] = 0x05c5a92af02e95a6a0a69eb25694d0fb74462f315fa1ad1818e5a85f33e491a2;
        pp.proof[104] = 0x1bc3abe88a59e2c92f4e1280305854058f0021c760e64203bb5fbcb66f896126;
        pp.proof[105] = 0x25956f67e3a354fbd7ff67e032190f0d158d3f874b7b5bad3c38afb6e903a642;
        pp.proof[106] = 0x105ede25ccaee90ac72f6708c8a38bf2850815dc69ea60a7736bafa1f1f2ac86;
        pp.proof[107] = 0x11df505e995d9610be41329fb0ca1c01e5b2bed405b41054816fce6669d6295c;
        pp.proof[108] = 0x0848561d29d0c2d5258586217ac15ef0ce03fcf01e605af9c515e3570dc425fe;
        pp.proof[109] = 0x2980e3871aa6bfc828920c1ee3783c2e17f1ef598b4a584255ed947944b074cf;
        pp.proof[110] = 0x0a105490d794934fdf7736e205dc95a97edac64c22e1dd5e291ee3183188701c;
        pp.proof[111] = 0x082a220ee89fdf4b6a30e72f5a1b82d5aec1a87662232483cf68d6c3c2a1253e;
        pp.proof[112] = 0x0f87b14b92855e790ed930a65bbd5d46ffd07cd2988ee6bc3539179d85f670d6;
        pp.proof[113] = 0x08bf58730a1f7c62121dae23d291b9cb02733d2e7af22a04e66b92bf1f54fdac;
        pp.proof[114] = 0x0cf23d10dca096f4e5fced9df398b59d9c0824344847874364d689eeecf12710;
        pp.proof[115] = 0x254abac420c9ef64fc68d874129ba838624ba2f3311282825632e8cf47821e2a;
        pp.proof[116] = 0x06127cea34afb05ce8ec3195449fc48ffdc8657ecb11bc7a9fc8db1e4dc3603e;
        pp.proof[117] = 0x30621754c15ee309248b234cc075f50ed8877d15f85d77fbff1c6de28c23dacf;
        pp.proof[118] = 0x13dff59922f93949189c27f30415ca9d54accb449dd119b9475b2d15cad9a6f3;
        pp.proof[119] = 0x05031a89082237448f24a66ffd905666246c2d85a49f7b4eb0da3875fa2660be;
        pp.proof[120] = 0x10f5edd2e9fa1feebe23a308d315aab10d6974f2e3d552a2981755d403dd2736;
        pp.proof[121] = 0x01e26f56d99ccbfaf01af471bffd470ce8a3e27bc679e2720985b9a1c09449d5;
        pp.proof[122] = 0x0b5a2123fa330a2e336e3442d7397e4987fef041d1d97837a552d3ad24e6fb89;
        pp.proof[123] = 0x1817d114ac9114e78db222a7c74b02fc30764bd11cdad32754a7aa72666ab623;
        pp.proof[124] = 0x2a00f0c2cee59c02158e8db512c73385f28ab1eb70b7ef73e86341228e808c2c;
        pp.proof[125] = 0x0dcc90c63560947d908d34924d5d2514f7bb1be784b1c010ae2bb6719586b38a;
        pp.proof[126] = 0x281af6ca14665d0794140820893ebb480e8d1ef8cdf0fe37094fb71f34612efc;
        pp.proof[127] = 0x07a67b4ab73ddd87281b6b021a4b2f19eb88367695fb6066abfacef5fa02c825;
        pp.proof[128] = 0x255568cf81305745af3dcb913e6f4077e2d48aa352f60c6e707869c6cfb73e3c;
        pp.proof[129] = 0x1229b93581b80ff61a1ad68c537f8af5d883fc11a140f76c62629d58a7d2c845;
        pp.proof[130] = 0x147ba8fab506e63208261ab6c22276ec35844aabef134aba290c6590f571d3e6;
        pp.proof[131] = 0x0edd6a1fe0fad061c6a45e8c0aa4d262a31f0904350ed34337243d2d276962f8;
        pp.proof[132] = 0x11477afad1a93575ef7bfa68c0e04484a9108e79ef63d616ca614b05baa7682e;
        pp.proof[133] = 0x0146480b406ef84e790dfea079f512525dcbf361341faa069531fbd44aef0425;
        pp.proof[134] = 0x0365d0de11cfad24fbd45f19d074d47e08fd3061115ab25de33aae7593db8275;
        pp.proof[135] = 0x29c6c9fefedb97686bcc4a71342dabe64c0c263ba7136cc4354177418370d819;
        pp.proof[136] = 0x0942c5e91535d5709f74e17107b787fdaf245b0711e2c1309340a47daef6d935;
        pp.proof[137] = 0x0ee6d8ea0859a1b88b041585d306351a67fb45f200dd966533a3da400091ca2f;
        pp.proof[138] = 0x2faa206c54f151b6421a4377d3a90068f0b488e27de457b468be4cd13bd66137;
        pp.proof[139] = 0x22a59373bcc665291647afc23f16af49e218a499c70fdfbe6c3d32220ba949a9;
        pp.proof[140] = 0x2667d6f5bf28ebfc754e1bc3214234bf14853c097473bbe5574e7c91b09a08ee;
        pp.proof[141] = 0x117425b1586f174cbae9dae1325482e2c239e15d5354f8e952674cf847cd97bb;
        pp.proof[142] = 0x242881b0dacfde429de1a253a9866681e0d2a02b0a84b7f21b2cfd195210ab12;
        pp.proof[143] = 0x1e15ec1297adf2bd19b2098779edcfc324e5a140c3f43b225a0337177acd7d66;
        pp.proof[144] = 0x10735949561e38cbf5801fc6f4005a165325ab8b08b7cd40d7e9d76729b5999d;

        pp.pair = new uint256[](5);
        pp.pair[0] = 11043928860274208474100222903659720401105736444942783912777777731178123111009;
        pp.pair[1] = 4733049234560598175814075307329400052685708929235536361180057682884349522951;
        pp.pair[2] = 12927143393367435583780384041605583199602086804808450831263369829150907626360;
        pp.pair[3] = 4387398449564854858399845335905273397295727913637629556341886693575315457608;
        pp.pair[4] = uint256(_publicInputHash);

        return pp;
    }
}
