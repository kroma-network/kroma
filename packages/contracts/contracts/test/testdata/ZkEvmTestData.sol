// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Hashing } from "../../libraries/Hashing.sol";
import { Types } from "../../libraries/Types.sol";
import { RLPWriter } from "../../libraries/rlp/RLPWriter.sol";
import { Colosseum } from "../../L1/Colosseum.sol";

library ZkEvmTestData {
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
            latestBlockhash: 0x3392758b5bca8b8319df6180c145ca28152f1b6a3af977bc48ec67d2259dbcd2,
            nextBlockHash: 0x4ecf76378ef03e3a417ac169cb052a879424345c59765aca05fe1fb6259375a9
        });

        Types.OutputRootProof memory dst = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x0475b3d38492c9e58190616eaad4ab033942aa55747d49c5a614b9e751998d5e,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            latestBlockhash: 0x4ecf76378ef03e3a417ac169cb052a879424345c59765aca05fe1fb6259375a9,
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
        pp.proof[0] = 0x2fdc701e2fc2653e7cde260b104d358633ea35479f3566f37236e75706d54f33;
        pp.proof[1] = 0x20e163bada0a37e8fa3deb9351bd0650666c4d647e53fe32b922ae1c9d101e51;
        pp.proof[2] = 0x2ff90e2299ff8b41bd46c28886f942b13b3403ba00ce4e3508315bb2f0817d73;
        pp.proof[3] = 0x1319ef5829fa99343eabdf8ec2c71ab7c90d1bc9ae4fc5b11997804ca774c094;
        pp.proof[4] = 0x141c2b685b2f50ec83d52a949a12cdb4f68c16256c0a702d89c1213f7963ba83;
        pp.proof[5] = 0x2f5ba8155cc1b933bf562e927360806745e091357004fbaf4ce48af303237dab;
        pp.proof[6] = 0x1a9db5518083f55faaf5cf4a39ef6f8e807e2a552ed2497f625dc86cecfa9d96;
        pp.proof[7] = 0x4381419865b3d502c7948d6e3386051cd3ca8ac2babb61114226cb656aadc40;
        pp.proof[8] = 0x26bdb01d53248faa5f6be5f84261843bdf98d32663daea2a208368722e82041f;
        pp.proof[9] = 0x10b113efa8cde6008d954837725d669f7fd871323630cd2215ab7f1e3f328dc6;
        pp.proof[10] = 0x1060475bdb157d3408970d16783d31ade2c5e9a743083210edeeaf304139ec48;
        pp.proof[11] = 0x2dfb8e2a8f6d25affde058ef2328e054049b1d2ae766f17ac776206132c5f4;
        pp.proof[12] = 0x1f608debd4be61572a326a19889bc5bb262af977f71cc8c32da457d3d9c1dbb2;
        pp.proof[13] = 0x241fcc502979e0fb4139840b856eb1a44b510b335c0150e9af78cdfacc0fa09b;
        pp.proof[14] = 0x2b1d57c633994b351bc7801ec2195577ca30b6f98f7c77df377bbd18266483fd;
        pp.proof[15] = 0xc0d2689f57f47d90dcb2e45d381469e7c6f80780591f2c2c43d59f6451ac010;
        pp.proof[16] = 0x7bf445b712435740eec5b3efcc35f62c917e7b925d00542de540e046c91f8c6;
        pp.proof[17] = 0x12c0b193bc80ea8e404ae09938fc5c5da4ee9dcb6798bae4e014c85fbab14258;
        pp.proof[18] = 0x227783555f99aa8e7908858ab92eb31ebe282a172a84daf72fc1f30d4ade3d72;
        pp.proof[19] = 0x22edfc50c8356b43de0aec9afca9f51f67dc8a971208a2a274ea507753b64e83;
        pp.proof[20] = 0x2b63bdc3622cbaa0f3f3e863a3360176b3cf3eb5ccdff40fa7acfced95527199;
        pp.proof[21] = 0x2f56d4df46eff6bdbe0f25410a5220b772a0a1894afae426d1ebe859bcd9e4e;
        pp.proof[22] = 0x17781b72fe5e7f7dafaedf62df1f94623dd8abe939ce877d9e400a1d10fb3e40;
        pp.proof[23] = 0x203019ca6ca965d10ec5abccbefd6abb484c7fc1ee42c921497f4e1c53718794;
        pp.proof[24] = 0x25cae74ac841cf4235753c0fd94f88c672c102e4e75cea7c1606f5e58223cd2b;
        pp.proof[25] = 0x24fad6f1bbbb053ec3a295cc29db23428ba6a33f2255beff698701ad0fb117b0;
        pp.proof[26] = 0x456b2f41d1ec50c0683e5f98bebefe05cd44f916524ab960a0000fc208524e3;
        pp.proof[27] = 0x216aa9d63e33e43130b2abc0a829e2e0fabf844d317856e2eb9005a7ff1d5849;
        pp.proof[28] = 0x5d7f3f092ce02ada2e2c6d77726c1d720cc05030990ed1ecb62266c02cc0819;
        pp.proof[29] = 0xc51b37fbab31153457b435bf059d8635692aa09ff4f97679176e141b5a82833;
        pp.proof[30] = 0x8530008500b253bdc4c97b72cc03184f37b4bb8fac18408d4efec3951548bdc;
        pp.proof[31] = 0x2712e1db9b893cf074456318e1eb5d32ccd823e0e90932af865d80892e9d147c;
        pp.proof[32] = 0x18b0fd7ea5064b1dd2827b5ba99b68d1a622002c46b6ff1e22edb22cc296643;
        pp.proof[33] = 0x26b4c4b2e4f1c1f5ff0d10d9b68b60710ed456841554c45391be0a91c0b66f1;
        pp.proof[34] = 0x733daf05b53f092db2567b03f0cf6df2c195c258ca8893618719a2f5395e09e;
        pp.proof[35] = 0x298fcccf5cb83a796f43906c6d8cd2202e8040b63cec6f66d4d9aec401461f3b;
        pp.proof[36] = 0x2973feb364253d4217e068d4e3657b187fd6ec0009a665fa80295355addd23d7;
        pp.proof[37] = 0x1ce601db03a8e98f2fa26bee2377d6f00ab229f8606dc97db0619567d19e1b6a;
        pp.proof[38] = 0x169e40d328ea51375e431384cde7e81130730c0814a9e283ac32b09810ae374;
        pp.proof[39] = 0x8a3453625b283e4a321a24fd24bc2b9b03bfc6dbf00a1f4741cb4ccc6fe2273;
        pp.proof[40] = 0x1cf1ee4719170cc23d0a1287cc384c9be5fa566de5861d24cd6f4202f6a0d67e;
        pp.proof[41] = 0x1af131d0b74c4a6b36b6d339edd46b184ef0d09c625f18c981c7cce975973c7b;
        pp.proof[42] = 0x104c0df3e11629330e9dde22eca39f1b82c96dc18bab150c83bfc96de824613f;
        pp.proof[43] = 0x6a5dca3e358382ff75fa88c8a383293d2b8dba8bff01e2b6df2fde3c88724f5;
        pp.proof[44] = 0x26cf560f8fef9766dfaa33fa525ea92f77181903ae9464e3367c97c7574366a2;
        pp.proof[45] = 0x3d4f9c9990c2045d8b84edf3076ec4cda1f63f5badb298a022079a934025012;
        pp.proof[46] = 0x1ceb5ec3210e5e239cd464fedef910c22e0e7fb1d25b1c18f79eaa5ebd59eed9;
        pp.proof[47] = 0x2b5011a3752039e5225664bb6c06d8ebef4a3672fb116561b52f817856e41cff;
        pp.proof[48] = 0x156abfab2abf6cac084b854d79889247d305d4b1f63721006d6f6db717a23c40;
        pp.proof[49] = 0x427ff2078978fbf3b042ba5f0f5b42a77ca1e519f240fa954102b430eda6a1b;
        pp.proof[50] = 0x2e69577fc75810032736717b82449d7b0e0634b7a7ef0e68c5e67546f19882cf;
        pp.proof[51] = 0x2af6acdde4a3f844f382cec882bed33276e30da371a9bcec34b9c56665c3eb49;
        pp.proof[52] = 0x91c725d1a56e7c71473370e410f0233e42cc01e4919e126849bfdc7869015b6;
        pp.proof[53] = 0x1afc040523ffdc20c41e9b0086f6df66881c5362fb762ff606d7bbef67dba411;
        pp.proof[54] = 0x32912137bd17d268da4c0cf3e63809ff5dbfb2d50c51fc3a8193dbb3c32c9fd;
        pp.proof[55] = 0x2633048356c7474beb41e70683501398ad590f3bb2a775c1396e8d87bdf4707b;
        pp.proof[56] = 0x1;
        pp.proof[57] = 0x2;
        pp.proof[58] = 0x19298d85b22c5b64d151af4eec207d3118151bbd114f9e30895826c3e9470dc7;
        pp.proof[59] = 0x2e9211c354feb0844ce9c26e3821565b4df0600692c897244b6fe435e509bf62;
        pp.proof[60] = 0x15127a0bcac2120d31681d94be04b1aab00f275c7a26d9d6401515483597fd7a;
        pp.proof[61] = 0x15e1889388505f334245ad0d8da8012d4b6526e3340acc49db00925d38fdb235;
        pp.proof[62] = 0x2774408029a42a6d71d4c85900ee16d2331e17d7ce421cbdca566a7effe1edbd;
        pp.proof[63] = 0x135916e1edf74f527c2eeb6930db2ff0203373615c5d1d425316ae1ac4101b4f;
        pp.proof[64] = 0x167266dcde80b5d1c40449502c52f735df68fcab9234e5b6e7d1103c6556088a;
        pp.proof[65] = 0xe362ea0a8a858e86087a9b82ad1f3fea26a6d1cedf664cb0b10a88a168650e1;
        pp.proof[66] = 0x1abd7878d55ce08f61618d3e946ee81f00365aaf9e3b31a0945b4866a146958a;
        pp.proof[67] = 0x1117c3b387d6a82ba670f0c75f0ea6dd970dd7bcf6d09f9abe6cec3f57699363;
        pp.proof[68] = 0x2ae4acf945ce936cf5db0aef46a5e84d2d239bee157f3d71c8d9ad629dacbc6c;
        pp.proof[69] = 0x2e3ddb4048d19c973d02097eb52ee61c8f3c0a66e23cd51b96df90fc30e48f67;
        pp.proof[70] = 0x171b8ec0e608f6ee8c60e9ed6846b0131873183b0255041a4c32e4c072f0c607;
        pp.proof[71] = 0x6b2f713fe8dbf616f7250105434d2bd1d7668801a9f0758f0286de930e74755;
        pp.proof[72] = 0x1e8ff17e7e7685971522c73910148843c4b1e193e499976a98e42fcdfe9a66c7;
        pp.proof[73] = 0x28130ff5a8302b3be6f879c29c4b3a911dfb591559e065f7863d7f2d4e20a06d;
        pp.proof[74] = 0x203119f6880c100978e57a16928ceb499a338100280b77304fc2a683b875acf6;
        pp.proof[75] = 0x2b6d77c422e3a09cf1195a9e06d72a42fa0b8756f3f7da1e3d6d73eda861ab21;
        pp.proof[76] = 0x2995100f21ebbb0ed55238beff3fb1589fae216956799fb95c7d5c5bd8ad1062;
        pp.proof[77] = 0x3392ab37e6cb61bcce3ef65baee748d182f82d416017bd1919da44e6cfa06c2;
        pp.proof[78] = 0x21b7fb3ec98d689f8dbb33b8dc997fe519e6cf42842becb41bdd6207fbf31a15;
        pp.proof[79] = 0x640ac8f555cfbd8187de1798619fb582166a4e2289c385b93458338345be1c9;
        pp.proof[80] = 0x2934c513819aef27feab1675716a0f6778dbe068954910871502a053172a604a;
        pp.proof[81] = 0x1328a5dccc432984b4ae5008071399613afc0b3b70bb69955e6102f6bb6dc755;
        pp.proof[82] = 0x2a2d8b64cd5715da065355649dc8799193da4a80ba050f288b4fcf7e9d84ee55;
        pp.proof[83] = 0xdf39890daa9d5f31323e8e5196aed99738d0b5befbb2d18b62a9321034475f6;
        pp.proof[84] = 0x256544ba5be9d8b2baef20431320c6b6f7e350fda0251be39e3328c5252bbef7;
        pp.proof[85] = 0x1ac46e3cfb0c48fe121e100c267fe5cd5bed9e325d068d80f0e7df372b5a07d3;
        pp.proof[86] = 0x11e3e0c9ab14421c63ec40458035f91c0237f8ca73c71adce2a004a510a2c800;
        pp.proof[87] = 0x1b9afdd0c11454697064ef251d763ec64b82661259c928dc8f9a24722ea2dcf7;
        pp.proof[88] = 0x49196e14f0ec7e36775e83a47ff44f3ec5c78573d284b4ed37f281f3dcd68d6;
        pp.proof[89] = 0x1ac46e3cfb0c48fe121e100c267fe5cd5bed9e325d068d80f0e7df372b5a07d3;
        pp.proof[90] = 0x1;
        pp.proof[91] = 0x283e4999955f4b48df712a2599b9773a0dc72af92ecca15a786b865aa81322c7;
        pp.proof[92] = 0x293a14832e8589b63b2f9a7c90ab26731d2f564eaec101636312b6e008496681;
        pp.proof[93] = 0xe34c528162fc030c699d9c2b92d66c1c3c8880f921ce4892a8ffacd1abc7736;
        pp.proof[94] = 0x1a1637ca3fe047d061493a6086752183e0984950d56e96e9349652ccff8ba6f6;
        pp.proof[95] = 0x21a7ccf6542418b4c3137090a176e8391ecd21c200765ce065275ca0823a6164;
        pp.proof[96] = 0x125f2a84c0cdf4bbca74acebf16341d1e3bc69a3ff54cf27e391c9b895125ca4;
        pp.proof[97] = 0x41e681ea154045a7344f8ac8186e2ebb97885a71ad4fef8738dc9780e1429c6;
        pp.proof[98] = 0x214af2832141f344916986a9a37f2bbb64a53ab300270c27ac91748b7072aa7a;
        pp.proof[99] = 0x1320806a17d2aece661ac2b64d3ecf15585b6f2feb8d4f812baa4a853273dcb1;
        pp.proof[100] = 0xaa7605552a9438dc66fef0c8bdb400f22e97abfa0679c6657f16aab29648d80;
        pp.proof[101] = 0xe40996061a05ca08f0bbdf22998bba471962fe8035631507266f354604438f6;
        pp.proof[102] = 0xeaa81d869d44a2d7d889e4a2c17bf56b17a5bbbfa8ee5792328b45041727184;
        pp.proof[103] = 0x1b7a626a3da725089f01592d17d933bef6798c800b1491d139c8bb084a6fa3c6;
        pp.proof[104] = 0xe71f70e8fc50e98793f0e99edd7c98f66d2a00d8d152bcdd57d9d0a6639c84;
        pp.proof[105] = 0x198f538b66a078bc1b6c28bb3276c529898fe3d662da97e90e528ea6ec357e47;
        pp.proof[106] = 0x2afcd51a00412f766e15a29313c8038fdfe6e74b954a2ea99240710c7fd3edd8;
        pp.proof[107] = 0x125dbe2f4583d7c90feb258e89ebd78a9cbe830e6443f77250dba013b349491a;
        pp.proof[108] = 0x224f7469c7af9c46159d189820101084685d7c93b474eab074dfa299e0440364;
        pp.proof[109] = 0x1deb5a29a5cb81d10c78d887370496b0af77ffc7b7ab8903884e63c10a464d80;
        pp.proof[110] = 0x2942084a484b2a4cff51cf5a8142364fae513e6837541adccdf15df4969d2d94;
        pp.proof[111] = 0x20bfb117e72a5f3bfbf444901d138f1cb9043883856d0030424f7a0863700783;
        pp.proof[112] = 0x29e39f869cbf67f7fbeb6dc90f14b665e4f45983cd239f8ea6e28a1bc0c50e73;
        pp.proof[113] = 0x1c75b6f9c84c01b98fbcad94d60857b1ee15e26292ff129f24af84ade45d57b6;
        pp.proof[114] = 0x159a258bb6c9d0e374ab4210ea2f532792bce6a9a4644983dd1bd1f77e73944c;
        pp.proof[115] = 0x46d0b1289d0bc8e29455d7f45723ad84fe4587eb80ae461f93bf06d30f4285e;
        pp.proof[116] = 0x209eb6f96cfb673ead9408263d879bc946a7498615fe0cb34627ea7e3195b307;
        pp.proof[117] = 0x7a58e80342271db0159527464587829f6ebe978de540ca1b7437b8641d34004;
        pp.proof[118] = 0x221a16a594bd8b074d794a99a09d7595f0d697a86954c71db0760375339b05ea;
        pp.proof[119] = 0x1b45b979933e8943040531bbf351faf94bdbd273be56b8e75a5fd2bd4afee36;
        pp.proof[120] = 0xd64d1118820fe4e71ca22c358f1385094b6591f62d92654762f1a084a339e37;
        pp.proof[121] = 0x113150b95a3f8e33fda003bddd14b8abe423af5ac52f48859ba5dfaed3b60af4;
        pp.proof[122] = 0x14e6cabeaa32518e598b1ddf0111bd4d7a17535fa8cd3f4ee696b547f3caf133;
        pp.proof[123] = 0x12e66ba628954b85a54b3355d3e8554c2b1bc7cd49cdec8c9b4661826de567d7;
        pp.proof[124] = 0x287eecc8ae49ac791762ef603eb26f9fc3faf8bc63f93c36c129dc066f5e44cd;
        pp.proof[125] = 0x2084548b3e2a9bb378150c271bba041244aa3fcb98a5bf08b4108843dd35abc6;
        pp.proof[126] = 0x223671c02f7e06e05a348a48855f29dc994774a95d53170e92f08913c3289732;
        pp.proof[127] = 0x2926305890a7b4e58dc6c95f62ed50af5a238b821c133c57cceb40ac68147073;
        pp.proof[128] = 0xe06553ad4c2977e78505070e3bed99ba2c18e244dfa6d54acbd329846915387;
        pp.proof[129] = 0x1fcee863cf9716df6598e069f7e3fa1ff5380756889606e4d4a58324c390b69d;
        pp.proof[130] = 0x2f1e6c1f198dd99bf24ecbcddf54b5f2b524cf6965b90d908c3e3ec6c3e09cf2;
        pp.proof[131] = 0x16344e152261742616dbcf2f1fb5d29a1b1e5ed651917f618a4f8733c96c5504;
        pp.proof[132] = 0x23e956b4bde5235788ab1bd89ed6a4b9b87389c617516e929cdf13881078d57;
        pp.proof[133] = 0x16586d1c98b070745c3627c826301bc09f8c0b8d45adc7a2a4cb752e0200e03;
        pp.proof[134] = 0x2c8251b2ff135efe2751351d0efcd6fcba7ec7f386e0ced0f154421d3c16bcff;
        pp.proof[135] = 0x241b9dd867478d9ada0e8dabc97236698f4fcb18ca172321f8306f8e2e6f5e14;
        pp.proof[136] = 0x1bf0174f39de19422c6e6f5d828e9df5f6fabf91989fa438093bdb4e2906a84c;
        pp.proof[137] = 0x719dc859db44b1be16f6320dc5de993b84a5ba229eff0492624446f1cc60bcd;
        pp.proof[138] = 0x2ecaa202ea0326c1daba798c09c367cd30140d0a68d6012e9bfd938366c0043a;
        pp.proof[139] = 0x1ee958c600a49cb9692e58acfeeb373b45f3768ec907c05c8dd44dd2af948065;
        pp.proof[140] = 0xc8ba6f0b100a87339b98243fc084db0ecf2035d1ff45602eea0b44755c1ba8e;
        pp.proof[141] = 0x112e334ea1529f733727c4969afcdd102f40571bff41d9b8fe4e006b89e8112d;
        pp.proof[142] = 0x22d028a5168fc31ac986bd84ab578fb2a4d32e04a178ecf113953d4a840c57e1;
        pp.proof[143] = 0x2677596fc15bb717a6c679e8a4be04fd4f7e75fca48f8904b9eb5bef67f673d3;
        pp.proof[144] = 0x1e5af52b9b8b81cf029cff490d6d0b8596ce779a1b879de5736286cea386eabd;

        pp.pair = new uint256[](4);
        pp.pair[0] = 1439937077562332533928559358929951268199674421421171954406633361314021765471;
        pp.pair[1] = 525991771894247758857595702512857049103382678372340463500401100215936874823;
        pp.pair[2] = 16610762691051063341183748375589157509446274807887837367684010822518324280674;
        pp.pair[3] = 2418211688952629122600292090656011756990119463711732129096727527695014522543;

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
