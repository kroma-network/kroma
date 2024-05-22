// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Hashing } from "../../libraries/Hashing.sol";
import { Types } from "../../libraries/Types.sol";
import { RLPWriter } from "../../libraries/rlp/RLPWriter.sol";
import { Colosseum } from "../../L1/Colosseum.sol";

library ColosseumTestData {
    uint256 internal constant INVALID_BLOCK_NUMBER = 21;
    bytes32 internal constant PREV_OUTPUT_ROOT =
        0xa6b4cc150f0c24daf8b5803491addbe102a388cf1ccec74fbe103a2deb5004e6;
    bytes32 internal constant TARGET_OUTPUT_ROOT =
        0x1ce4527b414c9e5d9edd26e488665f5f0f28c8e75dc622328c66776c93901e87;

    function outputRootProof()
        internal
        pure
        returns (Types.OutputRootProof memory, Types.OutputRootProof memory)
    {
        Types.OutputRootProof memory src = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x263975548df46f3ffc739f602b503f32b4c522026c8c93204929ddd5b65ad202,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0x3392758b5bca8b8319df6180c145ca28152f1b6a3af977bc48ec67d2259dbcd2,
            nextBlockHash: 0x4ecf76378ef03e3a417ac169cb052a879424345c59765aca05fe1fb6259375a9
        });

        Types.OutputRootProof memory dst = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x0475b3d38492c9e58190616eaad4ab033942aa55747d49c5a614b9e751998d5e,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0x4ecf76378ef03e3a417ac169cb052a879424345c59765aca05fe1fb6259375a9,
            nextBlockHash: 0x6c4e19b1fc27f6a075c67f35bd15b21c40025a892e32cdb8d9b5f5d5ec60093a
        });

        return (src, dst);
    }

    function publicInput() internal pure returns (Types.PublicInput memory) {
        bytes32[] memory txHashes = new bytes32[](1);
        txHashes[0] = 0x17deafc4c886b90706f3191fa8ecd152f34ce4fbc36ca35e6de899eb0fc7d86d;

        return
            Types.PublicInput({
                blockHash: 0x4ecf76378ef03e3a417ac169cb052a879424345c59765aca05fe1fb6259375a9,
                parentHash: 0x3392758b5bca8b8319df6180c145ca28152f1b6a3af977bc48ec67d2259dbcd2,
                timestamp: 0x66471e21,
                number: 0x15,
                gasLimit: 0x1c9c380,
                baseFee: 0x1,
                transactionsRoot: 0xb01bdd77e9685c8341733f345113daa34c25a63a37cb76b81de492b36b0cc806,
                stateRoot: 0x0475b3d38492c9e58190616eaad4ab033942aa55747d49c5a614b9e751998d5e,
                withdrawalsRoot: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421,
                txHashes: txHashes,
                blobGasUsed: 0x0,
                excessBlobGas: 0x0,
                parentBeaconRoot: 0x3eeb016384502029f0dc9cc6188d4e5ca8b6547f755b7cfa3749d7512f98c41b
            });
    }

    function blockHeaderRLP() internal pure returns (Types.BlockHeaderRLP memory) {
        return
            Types.BlockHeaderRLP({
                uncleHash: RLPWriter.writeBytes(
                    hex"1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"
                ),
                coinbase: RLPWriter.writeAddress(address(0)),
                receiptsRoot: RLPWriter.writeBytes(
                    hex"886c02379eee108cab1ada4055c4f82b048b7e3bbce0d82bf532c95409d8ad81"
                ),
                logsBloom: RLPWriter.writeBytes(
                    hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
                ),
                difficulty: RLPWriter.writeUint(0),
                gasUsed: RLPWriter.writeUint(0xc9f4),
                extraData: RLPWriter.writeBytes(hex""),
                mixHash: RLPWriter.writeBytes(
                    hex"8bb2786563ea29f638e4e9758d9886e8a1af5b4f7688f4ee6622a6b53df87742"
                ),
                nonce: RLPWriter.writeBytes(hex"0000000000000000")
            });
    }

    struct ProofPair {
        uint256[] proof;
        uint256[] pair;
    }

    function proofAndPair() internal pure returns (ProofPair memory pp) {
        pp.proof = new uint256[](145);
        pp.proof[0] = 0x19d9a5e6c202f1fc261b926b4fa397b107186444defa954f3202a78ceadd0c39;
        pp.proof[1] = 0xeef8959b7c33f876d819fc2b21664304406d94cd2dab27566d5ba8a874900a9;
        pp.proof[2] = 0x1041f72a533a0c43e4911cf8ae7a9ed35f9d5aad24a50927dbb252696a806973;
        pp.proof[3] = 0x2ca4170fc8bbd9445b4b719f90f3bebf7cae7600b4cd5f29bcc90e7492bb2266;
        pp.proof[4] = 0x7ac7b04bce7c2eeef565965b6a95dd45245f5e58c36f4210759600839532e88;
        pp.proof[5] = 0xc85a60ee456aadd6d3434b8b285acb04235582d71ade320845de178de818e11;
        pp.proof[6] = 0x2a12dda71d7eb57f0d537990026892f0e9f5dec46d8c79335a904c41e452cb36;
        pp.proof[7] = 0x1c2702aebbfdad9a9f66bb494fc99e71266b91ca917cac651688d7287ea12fe4;
        pp.proof[8] = 0x21d1b6c50295a0f0825cd2f983057018aff60bf4c03a00de8279f8c31585fb4f;
        pp.proof[9] = 0x2bb46c492f5713a48ce6a557b7eb2a92bb5fdfee98887f1cf998efb38395cd4b;
        pp.proof[10] = 0x14e569e64bbe5334bbc5889dd7b4189a7b2d095eadc5874c12770ffa41837edf;
        pp.proof[11] = 0x97e6e85a659e6cccad6ebfa3724684896d9b9a744eba53ef840119c9f0d833a;
        pp.proof[12] = 0xaaa6e75274ebf2218a5ce18a4665dd8f80bf9d7a98bd6cba3f3277ce451fe9a;
        pp.proof[13] = 0x1e9a1fc27904f3063f42a37dae8ed3cb460a21f399c7c025657c68a03bc6b28d;
        pp.proof[14] = 0x2559f02c90ae5ff601253fe9d54bbd29e9b00c7ec42161a82f7d51ad9a430c6f;
        pp.proof[15] = 0x8eb2e619b1045ed63dd3d4c4cfe75ad0071a0da49f171cc90ffccdbd4430287;
        pp.proof[16] = 0x2a4de9277fc6650dd11f93a6d57b1b56d635b599397bd401eb82348c0db61be6;
        pp.proof[17] = 0x4dc8beb7dcbe307fc9c00297ce2498ababf907fc0438d2cbb270da26617af8;
        pp.proof[18] = 0x62e6988cfc69a117c1c8d478f49dad268af7ed1a0319fd3b1a3a049979a7a4e;
        pp.proof[19] = 0x41640deb8474f39e36b89448bb9f8e4d161d9b143d2a9d5735f19b05fe39f86;
        pp.proof[20] = 0x132f0904719277fe27efd7f7ee6a03b3e8701fa4dbaca7de948b54c3916776ff;
        pp.proof[21] = 0x2534ffe7a93db1c66ee96f83bdc8a0db34ec80c371f5a6d596c2e1461e5a33cd;
        pp.proof[22] = 0x1342d47af85bf70259c34cb1b1428f037902b32b0be91c81a686db73b7ab5cfc;
        pp.proof[23] = 0x300caa42572ce202eed48266694281eb8d91f41c4eef5436177edb7b36aeae3a;
        pp.proof[24] = 0x1e971af80f46ef7b2bdafa8f2ee8cae477945332d1df7a3d7ec4e5e9381def7a;
        pp.proof[25] = 0x1786a7cbc7b7a7ca9a07d222ce467e06eeffd9e9130a30c7dc77bdba791366d4;
        pp.proof[26] = 0xa327254b651a6f3b84d11255f94234d91fbf0b4738ab0eb83c7e343e9e61cf9;
        pp.proof[27] = 0xf176c6b0a65bd46d1757b6b9a1a17eda864ae5023bff11c1da42ef07d0f5db3;
        pp.proof[28] = 0x19251cc0d4000bbbcb9b1f69e0471656ad1b25e6ae507f72fd9eccdd05fe9dca;
        pp.proof[29] = 0xd5a11ce7f6d052f4b4033b6dbd3c2dd1da8e852bf24ccb81ee614f5a7e57a8d;
        pp.proof[30] = 0x8530008500b253bdc4c97b72cc03184f37b4bb8fac18408d4efec3951548bdc;
        pp.proof[31] = 0x2712e1db9b893cf074456318e1eb5d32ccd823e0e90932af865d80892e9d147c;
        pp.proof[32] = 0x18b0fd7ea5064b1dd2827b5ba99b68d1a622002c46b6ff1e22edb22cc296643;
        pp.proof[33] = 0x26b4c4b2e4f1c1f5ff0d10d9b68b60710ed456841554c45391be0a91c0b66f1;
        pp.proof[34] = 0x15ac47815a2d3843718979c166dc1576878a27d9b1a8d828838d52d7d7ef5111;
        pp.proof[35] = 0x1ec53d3bad64c2f32eab818cab5a977c941d9268807fa71278b8acdfcc8e24b5;
        pp.proof[36] = 0x2f5d50dcb47854e0a8219ad5726f25a4ddd1dbe8b0ff3a17801447ebf51cd6b6;
        pp.proof[37] = 0x104c7443eb55fb4070479eeeaa13c957bb066718af07d6d0d7a94562e7857638;
        pp.proof[38] = 0x5be11f7a5682003b03cb2153e683a8b3f3dd521ba3925106f188a86fee8fe52;
        pp.proof[39] = 0x219a017f4bff574e19fe7cb62efef91d19a7a8645a1fffb619726b1bb3aa2711;
        pp.proof[40] = 0xc85fdafce413283d8f4e6246d5422afd3210274bc937bc86d842b9697e7ee39;
        pp.proof[41] = 0x29b82c1dae78d95d4c150f7fef285761504b8f7571c298b6f5088ac49f524fe8;
        pp.proof[42] = 0x286415245fc7d9029e24f43b3a198850d9dee241776eadf4800309efd213da4b;
        pp.proof[43] = 0x1b4d33ac75c22c532ce0a3ea814aa2d11c94e37300900983e0eb35bb5b2a081e;
        pp.proof[44] = 0xc8b256efc67d5d458336277c2476f68cf2dcdb6ebc8a54167e51ddb485b7ff7;
        pp.proof[45] = 0x1cfdb939cfa84ae63d7183417513985c8cb42496d2390f5736875afff28fa91b;
        pp.proof[46] = 0x2b8e5e182a4164c8556e1c7263b79062bdf9d0a23bc415cb7c53b787a7bd9aa8;
        pp.proof[47] = 0x146a8ecb0fe654b20efd592ff8b84a57e2cff5595ff7a393d384435450f4b890;
        pp.proof[48] = 0x16aa31b0af7b84eac629c18bd50c2a691cd4615e6018d3a33e7f026d2a884d45;
        pp.proof[49] = 0x79f24bf322a72e808fbd985a0a5762efff95b343aad42d06840729d6655ac99;
        pp.proof[50] = 0x28a0acee55fdf441d690f10d167959031147558c9c231de94e4257104aa7a67d;
        pp.proof[51] = 0x29e028f5cb7ad29c67178b1a82ae44ee326846d6ef6b1f9b2c3e1cb5db7acbc3;
        pp.proof[52] = 0x25ea3908b7f9e628a82f868db9a1ce6157b0366cc7aa8f6bab8a7b2c5340b776;
        pp.proof[53] = 0x267f6604a563eba2157ba4d689c97bbe2863d3411cb8f4f2b2ece5418d3f4b74;
        pp.proof[54] = 0x2ae7bad8c90bbff29d78ce5d4bb0d7f4c868fe005d1999788cc525c310fd553c;
        pp.proof[55] = 0x78e6ac53084999a3d8e2302204a0a1634d5d3e83312d48351653ca4d61b7eb7;
        pp.proof[56] = 0x1;
        pp.proof[57] = 0x2;
        pp.proof[58] = 0x23395a612af624e2d4083d25ec0591756c16cb88e58abddcf51e1c22f160a024;
        pp.proof[59] = 0x166dd2bb8f3a828997ab14769809e7833a27191882316c8a7b0962d1f717864d;
        pp.proof[60] = 0x5a851911e37694f5c5c5e3b98d2ce610033d73d6f3a131f38418158683008d0;
        pp.proof[61] = 0x12f296f0500691219609499f319b6a1f5d5edcbde0db152fb1814422d1de3984;
        pp.proof[62] = 0x291c72799f5a7c3c514cd1160808c6bdf43c0dc2f501e458667f9a4d0e64c057;
        pp.proof[63] = 0xf8c38ae3482288a56055e2259735dd16068ac641dc3966558b936612aa21b32;
        pp.proof[64] = 0xf1832d6607e55c92c4bee994f7daf7137c3aba09c7e382abea07ae0ebd93acb;
        pp.proof[65] = 0x81e2a085b569dabda9d4348508c5251c614e7e97fa3395691b24d98805f0945;
        pp.proof[66] = 0x109694208171a97a2658578f5aa109e44eacf21a4683d7cd4633e0e230845a11;
        pp.proof[67] = 0x140b8d82dbb59adb9f820f202a8bf76d6ed673655e4105e10be862b861290ca9;
        pp.proof[68] = 0xa07f3957dd8af183d4e4a8ebe2af69c79366d4b4241a0136ca93a77226c1083;
        pp.proof[69] = 0x13488d62027d0ea6dbe32260d9f8bbb10d43dc4ced006e60f2578ff2ea2c258b;
        pp.proof[70] = 0x22ad23fa2fa7ff93f04ab4309fe914a7edc26d21b4aa13461ff0ccdf91219bc2;
        pp.proof[71] = 0x2d678138dbbe19123044a87eaffeccfc3549b5d95025a6cdd1ec4d8147448910;
        pp.proof[72] = 0x7678fda073d29b179654c16bd23a8908a717f32f304ffe8a0f27f1c28cb8fee;
        pp.proof[73] = 0x2ace35ce250865a34313f73023ab119b9bcfc861ba689fa2649a85d8e09f31;
        pp.proof[74] = 0x3871885abf4621c5af081983e14ad617350107c7340e426524f0a96953039d8;
        pp.proof[75] = 0x174ccb627d00b18afd4f8bd2b300b15771c5b3e76f6595dc7df0e9b5cc305a30;
        pp.proof[76] = 0xfc41596cfbee851abeea109786b9dc04e6c694cfa073e122f5b58be5c2f4fb9;
        pp.proof[77] = 0x2a56f2ba074da7a91bfd865e597df0076305625f201e049a97db8fcd670f61b0;
        pp.proof[78] = 0x287b2f8a4efc8ca41988df58b45e8204c7c3334e3484ffbe3525f29811bd8ac0;
        pp.proof[79] = 0x2cec763b034f95d1a67238e3b7af9b69fdc8a4197750b2d8bdaf543dda7819db;
        pp.proof[80] = 0x613df18434c677c9cff6dd41523088bed1c71ade2723f7e5769003e6ec35db8;
        pp.proof[81] = 0x21c8d3632aa90eb52fad20cd3e0661edd04233294351eb352396e5a01ff55fa0;
        pp.proof[82] = 0x1bab88658faeb79be26e540b79f4de6dd6c00160d86ed6de26e7788b48c738b9;
        pp.proof[83] = 0x2d05ad44e7c3af890ac00b530d7b6835a1f2623ad102c6474345e8fe2249f492;
        pp.proof[84] = 0x1c01e7ba1c4c5917dedf200ef03c0cc10ab4d487c492cf6f0b1806343aac4f0f;
        pp.proof[85] = 0x2c6c73d584195053def0e5704cef183b475a6b7ac4260c0db03dc3ebe6be8495;
        pp.proof[86] = 0x15c0edb43e07fe0b22f1fbf4b70773005fc811f69708435c38699cc42fa01b23;
        pp.proof[87] = 0x2cc8622aa0b90cef617f18151bad2ccfbd5a68fe13aa9ace2c8fc110d8e0cc26;
        pp.proof[88] = 0x8c50e06ae3657fbfe2300a048ac5e2c4097cdcc223bc3421650844c26b66ecc;
        pp.proof[89] = 0x2c6c73d584195053def0e5704cef183b475a6b7ac4260c0db03dc3ebe6be8495;
        pp.proof[90] = 0x1;
        pp.proof[91] = 0x28c99b15ae04df4dcde160a9959c7187cde205c5de50dc02154d60e18f5163b5;
        pp.proof[92] = 0x24876d6b1c3c0cc92aef321b4294d4b70ba8c60be5ef09c17ecb80bf71a444ca;
        pp.proof[93] = 0x2f4e3bc3c5da7cdce65b35467654c4e54d665130b583c8433ec59d8bb9a71648;
        pp.proof[94] = 0x3380ff3c601d48ffdc362e1cd63ecfdef5b618bb22fb5dbf7c5f5ff298ff347;
        pp.proof[95] = 0xebb4b6dbb89fecfd3e86cd2dc6c773e97e8efd6d7263c82dad78359d03b08c8;
        pp.proof[96] = 0x18d04432f244554cf3bd1129f7c6523082e2f3201ecba41294025e49041d987f;
        pp.proof[97] = 0x2069e78eb6ae9bc854be91a8d88041e6045d374ba101b19f57c0b7dba690dbbc;
        pp.proof[98] = 0x8e591b731bb2dbc44e03033b6a3be5b1106623931d2c814765f3d43770ef65f;
        pp.proof[99] = 0x61650db02d15001cda395152af849d28775f73e59b016bc3ba4e935179cbe60;
        pp.proof[100] = 0x1cfaac553db3ed2cead5eed1ef4c640c6bd8278392e58c3ad8c3580b9ce42785;
        pp.proof[101] = 0x3bf83e3ff40e84a24dd56a91e987d47498e79aa05e0cea29c3b2ab35473af28;
        pp.proof[102] = 0x219faf9652a12301a885292ceff5d8c97a8c180bdc048baa7faa765ffcee49fa;
        pp.proof[103] = 0x718ce340ccea40a710b5c1ea2576e2dbdf42cd26fbb48d226c35b7affb435d2;
        pp.proof[104] = 0x435f21beb1a974587e0228a6091c7eef49e21262199fd10a26025273d8c0182;
        pp.proof[105] = 0x11135b6bf7c3d99dde9a67d43e2042d403e1da61869920d97e9a94d87b1e320a;
        pp.proof[106] = 0x15bc6df543c620446e8f59285904d0d928bcab4c800ddd9de448bfb7f6572c;
        pp.proof[107] = 0x19793dd85fab0890d209f095779f13b381b41162883139a23eff53b544faea52;
        pp.proof[108] = 0x1f1c57d5c347482efbb84a280584f56ab7fd10ec8c0f8bcc9226a2776cf7a720;
        pp.proof[109] = 0x296dd17916b8e05f69c27e575277af2965cb99897f0f03ffb7729d53243c88d5;
        pp.proof[110] = 0x2cb952f3041a58da70ff6b4ee269ad6393007e105daf71d86b62e7bd4315c3d1;
        pp.proof[111] = 0x6a7739a0fc26fb0a70bfc71feeccbe86460d06f2163ef6dd6ab40e3e48462d8;
        pp.proof[112] = 0x7ead49a7fa92c917f8d169eaebc2e17f90e7a5de1a13c8f0a6e317ea1c1e0eb;
        pp.proof[113] = 0x2c7b7e185a927b0df7680df47a8802468a05962368ac1e2b49254b494fbf3416;
        pp.proof[114] = 0xc7a831b9902b41b480d8070ae0ab20c63f504559a1d5089836aa61aac514fe3;
        pp.proof[115] = 0x2e3c492cce5b8892ba69a522e045ff77611383a01a816c3b80e13072540ada61;
        pp.proof[116] = 0x44f3117b64d9b28e034c13fd5503e1b040b6e8cf9c3fa50d0e2771b95e0b974;
        pp.proof[117] = 0x15cf690ecc5dae137d6819e2c77996cad20df57606b53739371060a2bbf8e31a;
        pp.proof[118] = 0x9e9f2c89d30f957abcbb49f05b3b20cbe2aaebabba7d65ca06112ecd2be0fd2;
        pp.proof[119] = 0xdedb12873610bdcccfdde69e50d5cce5f0415d9bf75e0c6850238a8d6efd2d8;
        pp.proof[120] = 0x2f9bef9f4c0da5e840dd1380b16824f65bcb9f5c1a3fc395612ea6fc2cea3619;
        pp.proof[121] = 0xf20b26f5e36d70764c719d22f58217f1689d9e9b72fa8ddedfbb9a1d58f4101;
        pp.proof[122] = 0x267a535ff79f379b789a1191f765dd965933af29411100df6ff867ae0a14a430;
        pp.proof[123] = 0x177e2aaa4c7d5ba3d1fb393b4b153db4cc76f1402ec8b028573a4415aaacfe03;
        pp.proof[124] = 0x1d925a76cd4dd7a2e0333d4218aaeac2eca28b0570f70611a594ca7c8e6b6bb3;
        pp.proof[125] = 0x255cc0a4531fefc7bc6d71e6dd9cbd4430932056df71a89bddb2e92ec7f18b81;
        pp.proof[126] = 0x47d34bf469acb61bfaeaf3ae664f48e760674b4d5e6e84503dc78295c743d4f;
        pp.proof[127] = 0x271e2c093d187d84b4c5d8ac7a0162afb603b6810c633ae282993f20615bc9b4;
        pp.proof[128] = 0x2e5dbc873d85f16659f02af847927779aa417dc47a7ec2849d68121b4bb60877;
        pp.proof[129] = 0xd64fbe8eab56b4486c71c46e86f9c9ddf63565d10d0813a3eb234ac42c1145b;
        pp.proof[130] = 0x104b3be2969d2c202d65f7e5230d8231311a35e3aa43094857b8fa42ffcaf81a;
        pp.proof[131] = 0x217a93e067fbeaee0dfec3d420417df20b45c4127d6ea4c90b0048df79614264;
        pp.proof[132] = 0x2d39dad695ef52da5216345063d0ff37a05dff81a56f15c030d191dc54d72b25;
        pp.proof[133] = 0x2513e65dbb34c4e2e8eaee15e4f59a4bea8358e335751f84dab861b65e04a254;
        pp.proof[134] = 0xda25c78bec332d8ec9f32b2998254d646e899c22e3db0a7392a959a32c3580c;
        pp.proof[135] = 0x17618ca1ba0b11212cd87ea729a46a473ac9d12e6420681a3d6f43bd3804345;
        pp.proof[136] = 0x254b82f0b003a988704be335cd3ecb82471a9bac8b0d2eb06e0f8dd92c2a2456;
        pp.proof[137] = 0x1a9d8f33d203241c2530eabba9b6ddac8202fb77c888c9d7806a0262e4403a8d;
        pp.proof[138] = 0x376cebb1ffd041c1b02d29e3daf2d58906eb1f8becebabc50484ccdc610c082;
        pp.proof[139] = 0x1ea714a78a3d194f50f57309b4e7386106b01202bf3edfa4b7c027049850871a;
        pp.proof[140] = 0x281e9eb123e0b3ce9af4945284fb1888269a00cdae28a58ff197e8dcc2edd20a;
        pp.proof[141] = 0x12fa6f757a101fbf3f56e85fb3c99b968e5bfaa3b60bcd64dc20cc0822d2fd14;
        pp.proof[142] = 0x2bb7a01c7bffbab9142b26a30296dcee2847b6503746a2075b0f723e2e9f9648;
        pp.proof[143] = 0x14043cabc30eb3b74dd631e5417cd252089eb8bf417d80213f5d2973469e7932;
        pp.proof[144] = 0x280a12a7eacbce09fd7d77c638304d67c0176095e876ea82da3d1d949a5cc7ac;

        pp.pair = new uint256[](4);
        pp.pair[0] = 19096695823754423623148123873573015970890176872266371794684661079871098902785;
        pp.pair[1] = 11521281417540910269396101580613615924952361328842417676530411057268609166447;
        pp.pair[2] = 5499706566509864575301708419078377883501797033394273132167316945164613794624;
        pp.pair[3] = 17894798161211220374432305195961202385860664148552353445548468073882682310505;

        return pp;
    }

    struct Account {
        uint64 nonce;
        uint256 balance;
        bytes32 storageRoot;
        bytes32 codeHash;
    }

    function merkleProof() internal pure returns (Account memory, bytes[] memory) {
        Account memory account = Account({
            nonce: 0,
            balance: 0,
            storageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            codeHash: 0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35
        });
        bytes[] memory proof = new bytes[](13);

        proof[
            0
        ] = hex"0027e039ebdf0f9e7c8a1481ebf7448aae44afec16b045969976d37555b364f6132d6654101a6881a48a968d5a257cc8dc8d980c0cc55d58e47833222a24b2230a";
        proof[
            1
        ] = hex"002023945e0a0e2290059ca87427aa00e3d515d2e9144d8d3e69ee13b7c75615482cf819ee96fa5bf76e6ce7712caf6673df7fc3af269cfe193bed443eeb89527a";
        proof[
            2
        ] = hex"0021d5edc847df9d20356a576c7a4b4b1bb992b427aa4e01412eb08864416db55b1c7c9c38ad18cc007a6e522b638fdf6010aef11727fcb7cdcbfb5edb96e3a74a";
        proof[
            3
        ] = hex"0020140f1c792e0ff14d98ca78ed1c085501566b9917cc0792285ed265e6f8cd5b1e99968ab013de794c675a5f614b8a642294aa47612f1bb6173910a5cd7f0b16";
        proof[
            4
        ] = hex"002b11bc4ac76ee779a652e7e93a9871fa88e780285d41c6ed2f1646ec8658b7450c158053a97ff3e327cfae62631e5878accd83e5803f044c90803d4579ae21ba";
        proof[
            5
        ] = hex"00244f31e0a770ed9dbdf583a0e0dbe036f5f9476c3cc761e18b2a475a96bcc10814c463c1a53d595a21f22f3c5cafe495329c68087ab2601468f8fa7e11d88bf7";
        proof[
            6
        ] = hex"0008f993e0df87a04e71f72a3a237c909e2381eacf176b279925d333ee7b1e36ed03c30671a87c81a313a035fdbe052cc592ae0a604bcf87a5cf163d5a43104574";
        proof[
            7
        ] = hex"00202aa398b4bd976d7c165b2c7bfa6e4b695f18ff78e4cb544dbb0fbab8c6537e15c1089ff56ee758ec382a55a21b76c08ebf64d6a78d7e6ff442793536607510";
        proof[
            8
        ] = hex"0000000000000000000000000000000000000000000000000000000000000000002e957b48192277673a8dde0549358d09d3f9ad6e8db14e08f4fed46f96021a74";
        proof[
            9
        ] = hex"002325a334c56feef28306cece9e3867165ee117aab1c831fe04db286a1b4ff2c80000000000000000000000000000000000000000000000000000000000000000";
        proof[
            10
        ] = hex"00218476186a36a2ddf003ef59459478f44e0cea1ac32870dafca118331259b05f23618448c7fab9e44d30c44be6777aa390c25ad138fde11d22bdefd05f43838b";
        proof[
            11
        ] = hex"012de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c304040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc3524f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127204200000000000000000000000000000000000003000000000000000000000000";
        proof[
            12
        ] = hex"5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449";

        return (account, proof);
    }
}
