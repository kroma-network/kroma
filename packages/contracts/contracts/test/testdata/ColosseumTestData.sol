// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Hashing } from "../../libraries/Hashing.sol";
import { Types } from "../../libraries/Types.sol";
import { RLPWriter } from "../../libraries/rlp/RLPWriter.sol";
import { Colosseum } from "../../L1/Colosseum.sol";

library ColosseumTestData {
    uint256 internal constant INVALID_BLOCK_NUMBER = 21;
    bytes32 internal constant PREV_OUTPUT_ROOT =
        0x6f3194bc2f916583747a794e2857627fd1d29fbc68c6c330745d1269d26fc67e;
    bytes32 internal constant TARGET_OUTPUT_ROOT =
        0x2a11f436327b23ee68aaf4e468b0b77199416d32dd18e3b116099abdfec29619;

    function outputRootProof()
        internal
        pure
        returns (Types.OutputRootProof memory, Types.OutputRootProof memory)
    {
        Types.OutputRootProof memory src = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x06e23d47dea22feb9523e2817c42de9f05cbc9ce1410ef45bce2dcda2aab7bd6,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0x936a872193703be4201cffc531b00fd6e8c71a092c0e7ea51ab0f8ece00838a4,
            nextBlockHash: 0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96
        });

        Types.OutputRootProof memory dst = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x245d2fa3cbb92f5223e06b23a4c53fc1e1b9cd8edea6769e6749e72c8ff2d384,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96,
            nextBlockHash: 0x2afbf92eea49c24a7e50934e7b05bf21420e5e156bae67b82b41a83423b9dd5d
        });

        return (src, dst);
    }

    function publicInput() internal pure returns (Types.PublicInput memory) {
        bytes32[] memory txHashes = new bytes32[](1);
        txHashes[0] = 0xa56648e0f3071be0f185b83a97c2e3cf083e1cd1808c960fc97a0e16e1ae29cc;

        return
            Types.PublicInput({
                blockHash: 0xf4f75cec53957dc5b44a5aaff8d778da4d23b79e278f9d4dc8965ac147d7fc96,
                parentHash: 0x936a872193703be4201cffc531b00fd6e8c71a092c0e7ea51ab0f8ece00838a4,
                timestamp: 0x65f2c15d,
                number: 0x15,
                gasLimit: 0x1c9c380,
                baseFee: 0x1,
                transactionsRoot: 0xc9862b8f3157beea2d258b0a36b836794ea58302951468ab3d0b0885e63bccbf,
                stateRoot: 0x245d2fa3cbb92f5223e06b23a4c53fc1e1b9cd8edea6769e6749e72c8ff2d384,
                withdrawalsRoot: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421,
                txHashes: txHashes
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
                    hex"1ce54d356f9c39a93aac476d492b27c4c23374e203d9902becd4f4b6826b9b35"
                ),
                logsBloom: RLPWriter.writeBytes(
                    hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
                ),
                difficulty: RLPWriter.writeUint(0),
                gasUsed: RLPWriter.writeUint(0xc0b1),
                extraData: RLPWriter.writeBytes(hex""),
                mixHash: RLPWriter.writeBytes(
                    hex"0000000000000000000000000000000000000000000000000000000000000000"
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
        pp.proof[0] = 0x1a4fa7303ffcad780726423d94b1b3464f06e29d7dbebcf9d6e43f4e7aacb089;
        pp.proof[1] = 0x193576e16ebc2fad8b4fe09a30cea72b4a8000f137cce050b796150c2718e871;
        pp.proof[2] = 0x1412c51128048c6f6b2c4f24e714b44c4b56a63134fdcf0330a060ffcf62e6c;
        pp.proof[3] = 0xa7a14d581d685f27bd0da828cbd1af7164f3e09aa8ea3d1cb00aa1a5dabdefd;
        pp.proof[4] = 0x82e795582ccb3343a6c55864bb142b0c945ac5280c77b37e76f0caa486d98a3;
        pp.proof[5] = 0x1233418401e1fd4686489d61c5462ea7c789ba68d33376fa476176803fcceac2;
        pp.proof[6] = 0x10d6360e1b1fac6f91978ae8a382361caa17dc75a03a47d47491274528cac974;
        pp.proof[7] = 0x25219565a1662ad218e357d0dc333e42e520e8bb776ce9a751cbfc29a0940d41;
        pp.proof[8] = 0x200d9a52a6e687bda64f5ef1b6bdeeddfb83674d96b19f7b3eb5767bb4d1a57b;
        pp.proof[9] = 0x1a3b0540202a34956f83e3651ffc8b26f9788a3074ca0a61fc1805ed2353e6cc;
        pp.proof[10] = 0x18e4b80bd9483c4f424a44678c9d41fc4961dfb59c409e4f950197f8660042b7;
        pp.proof[11] = 0x4cb69ca854fd0050f2ddf810297c516e563e7b407776c2401ed49ada6cdfdd2;
        pp.proof[12] = 0x64ba1c035548086568ad8d8bcec60521a51d69c96d8823d2439237a0ed3c8fa;
        pp.proof[13] = 0x2e570dc0085d912427ac930adac23067f820feb87bc0976b7a94b63af060fb5e;
        pp.proof[14] = 0x27baee2dfaffd68bf61190f0e7dd94e4eeb7f1b62b02913653fe2b0b80ecf4f4;
        pp.proof[15] = 0x2d9403794f554ee2369ba20684da441e93051f9463050c63397b7ad1ffb86276;
        pp.proof[16] = 0x93d986b88ac136ebdcb74fd420958433be4f8e86d5a098cfb8dd5ac3bbfbf55;
        pp.proof[17] = 0x2f846fc7f9eac3cb5b2ff948ac8106aac401c5d0264cac825f87bcf276eb2eb9;
        pp.proof[18] = 0x16aa2b0a6d9d3bcfc405025f5f9a06f19effed83b9055116404f9b62974e41b3;
        pp.proof[19] = 0x17d81e09e2177633bf41fe6cff73fe89099c169f926e77be4dbd2217544d6570;
        pp.proof[20] = 0x7f29b16da8e1b5398566325a6fed474470404d6094da7bccbb3fc229958f79f;
        pp.proof[21] = 0x2e6fb0ebf603290c327c8a79d04f94f188931cb19de486b2bfcf55b6e3c9a900;
        pp.proof[22] = 0x25da8f322896e3d03cb241f417c66cc3148175bf7914183ff6b0d709d00fc4fc;
        pp.proof[23] = 0x1be897cd230733d81bd4ff7f65614ed868a7139e37aa28e69482cd38517d0bf9;
        pp.proof[24] = 0x13b65f32a8e50e654582fa904547d589687e904043a05cabd6798ccb72f92dcd;
        pp.proof[25] = 0x12e05da93b8fea3589bdff889955e5b1d8bc00ba25eb77a3e1cf1b4a5bf34fa0;
        pp.proof[26] = 0xfeb533eeab500979968cff96aa2d884e923fbc933cacc2ade144fee8aadd31a;
        pp.proof[27] = 0x1634c88c050640ecc4895e5c866267d9d53375916b3a9e5e660a3569690e87f;
        pp.proof[28] = 0x445d7dd3b43a39773657af03114f104863827b86c2a3d504665bc5b40e4da3a;
        pp.proof[29] = 0x264ca4944add7a5f1527b92d6b0aa7cadf5fe88ed8bbb53b7bd2b191378e07a3;
        pp.proof[30] = 0x8530008500b253bdc4c97b72cc03184f37b4bb8fac18408d4efec3951548bdc;
        pp.proof[31] = 0x2712e1db9b893cf074456318e1eb5d32ccd823e0e90932af865d80892e9d147c;
        pp.proof[32] = 0x18b0fd7ea5064b1dd2827b5ba99b68d1a622002c46b6ff1e22edb22cc296643;
        pp.proof[33] = 0x26b4c4b2e4f1c1f5ff0d10d9b68b60710ed456841554c45391be0a91c0b66f1;
        pp.proof[34] = 0x1e57d4e134f1daedf32559b4975bcfbb45f4fd8eaca37c56278b5575c568509;
        pp.proof[35] = 0x14072e5222c5a43d3c45f3be844d9a445fc5cec4c7243f3f51855d68c0b1a37f;
        pp.proof[36] = 0xfc3c13bac8237b88e734fe9e7a10508cdbb20e2aba3c82542d14c4f19438a1b;
        pp.proof[37] = 0x1e6c8129cff5125ca3ff7cf0e4471569a23767a82e6eca4444e536f7781733be;
        pp.proof[38] = 0x1833e4a4367c1305fb606e881c03d0639a9aa85303fec0af8c0dbf389b7739b4;
        pp.proof[39] = 0x6c7b2f24534c150a46fbb1985b0a8d13145b75c14fc4540d4c9fe113e7c6988;
        pp.proof[40] = 0x1f9b07f5146b8c4b13edbcbd7c8ca0792f508f215315b205327789a97eb11972;
        pp.proof[41] = 0x6148ecc6e98660f9dbdbc8f5b24b2b9cd6f1282d7f6c89a74dd44d7e558268a;
        pp.proof[42] = 0xc9d7a8fa4dd5e0bb8a5ccb48d186ea25c75faa2e811ccc0688b4f28040fa2b8;
        pp.proof[43] = 0x38e390411f0279cdf6f035c0fc2f8daef165901ee11a2f28adcf8c97c7970c2;
        pp.proof[44] = 0x2095833ea202a19d38dcdf3cf38c2ac647792571fc45036b8f5f32490078eb56;
        pp.proof[45] = 0xeb23aa693738c36cc58c301e40c5ea228724865222c596e3da681f6a5759757;
        pp.proof[46] = 0x132b7e449d7e9c5178354fe5a99798e9adcfe88d38aed3fa18b535e9d5faf0a3;
        pp.proof[47] = 0x11009bc48ed9b304b2b88fa6a45d93bd536c240dd8e7f03208ce1bbde43ba74;
        pp.proof[48] = 0x1b4ee695787b24b669eabc594564771f5d34ae2091b641b6d05495f96ec27dda;
        pp.proof[49] = 0x1254fd98c0da49fc15dbcd04ecba8b6b67100f9b93086abcbc393d5819c00a0f;
        pp.proof[50] = 0xe0a728cb22fa4633cce0e891aa6b039b27eb644c61f2a6fbe8ae6ee7cec6ea7;
        pp.proof[51] = 0xdc39efcdb5fb1f8bff907c8170fd612e7f4290e7feb3e94d12dbf4ea7a90d19;
        pp.proof[52] = 0xf5e1faf705a5a6b90011d6c8485e592dfe17da44715ce41fb4a0559384f4f7a;
        pp.proof[53] = 0x1ac6a8210be7ba46cd2c42d621484caeb753ba06e6b3b7a78f2d1b6f3c29f71;
        pp.proof[54] = 0x14d7ab23c37d991433a192c0d1b012117a7ee61d813f6714ab0010e0d0f58e60;
        pp.proof[55] = 0x1d72b6fbbda25169348838cb88666814ba728ffe9e3f2a1b94a911f51f7ddd84;
        pp.proof[56] = 0x01;
        pp.proof[57] = 0x02;
        pp.proof[58] = 0x12f61d3c7ae852ab83c4b0f1e2e181fbb140ab16bf18cc8f5c164b8283242c20;
        pp.proof[59] = 0xd289d8ff6b45f1470a27ab925a1de2176386e3af0c7eea57c7124617f20e0f1;
        pp.proof[60] = 0x2b2be0bc31cad577322c51a63e3fdff68035974c4bb9cb27e85cc23377462ba8;
        pp.proof[61] = 0x2769aed974dea2329f49d1c0bc9454da24c4d37a8872b760bdeccc02eaa317c9;
        pp.proof[62] = 0x5b1f9e762011a3ef147208c7406e0acb9cdf9292cd7b5a6e9345666a09c474e;
        pp.proof[63] = 0x4eedf2a4b5930689e5cc5aa09e645e1e38ce70d39f25413c2545be1791e246b;
        pp.proof[64] = 0x2a5258ef877e779bb60f5316e3ccbfb5cb668e1c9f654575bc7f1e95ef51a562;
        pp.proof[65] = 0xf1b5f50f4e5f21783a46294c5b760a8b98d579260a62cdec2ac309b4bbdd898;
        pp.proof[66] = 0x780967041a3ecec1ab95a889aa27b3e205ff30c4201aaa10fa8e224c6f9ba0e;
        pp.proof[67] = 0xebf6f809e99a3133007532e82c2fd787d637bb5307fa48880b3b981c9415d6;
        pp.proof[68] = 0x2a01fe691c17c2fe992254794ab963c26a116026fb3645df49fced7ee2e8cb5d;
        pp.proof[69] = 0xae1d648b82de2fd5ce5b07d6bf795f72649ee23e91c3857076ab33f68c3780;
        pp.proof[70] = 0x2391551c2875a7401263ddbcc430a1640f19a55289eb9e2f604d3c8d21f94693;
        pp.proof[71] = 0x24cd354b9a605bde6fc47b61c56f2fef50861375e08990dc80d4061e8e7ecb1f;
        pp.proof[72] = 0x5293fc5baa97522c104b818fb30eb78fd32dde691a7b19b015337b8678ba8bb;
        pp.proof[73] = 0x1258772f43393835524592dad121627aab93ec206a500610a189f7da38d23bff;
        pp.proof[74] = 0x19b1e89ac0cea94512f3357285534ec7a50aecda889897fe8987fbcb3ede0b23;
        pp.proof[75] = 0x86be6e663daaee8fa6c0ae3d2d5258ee97945abcad95873d5fe073b78e77ea;
        pp.proof[76] = 0x260f1bd628ddfefd2dcf01b8e3228bcb489de9f5c0aa2db389a946a67d62531b;
        pp.proof[77] = 0x171ea09df61954173c034e5a369a03bedc1fb7d1164b93bc1495daab4253d69e;
        pp.proof[78] = 0x7c274bd65e03009fa50367bab8674c9024d9d24e4682ec77691168c54a94fbe;
        pp.proof[79] = 0x574d50f1652b81e8fa244ac5f68f54e3ec4e824ad8785e67bba05ce3aad16ed;
        pp.proof[80] = 0x1ea3968ef186dd4d44df9a035c0de712202a8d3ee29a3450891e85715e9380a7;
        pp.proof[81] = 0x121130c76b649f458ba41a13b728f52414b8dec346a13339a351658ef66eeb19;
        pp.proof[82] = 0x158ab79396c82a7715f83c7eef66804a70acfce74162386c1bc22b181d348117;
        pp.proof[83] = 0x1d60e7eb2f9add7efe046308a588915dea3c6ba0c9e15f43600e503964d263ea;
        pp.proof[84] = 0xfffd842a6816beb8ee1d5af4c89ab0b91e307f90047562eb6676400ec0d8574;
        pp.proof[85] = 0x1bc0cb9eb6dc60385a38054d5e83923be16a6028f6f1f5ec89dce155efee52f1;
        pp.proof[86] = 0x3a7e9c868091cd0fcd338591933f1bb0badc49e035b711b4e012cb732a36534;
        pp.proof[87] = 0x19c9be7c3f52bbdf45056d381ee1536f33622f3e361ae6cfef944c8a5c9817ab;
        pp.proof[88] = 0x22e4c08d12e9607ce3eeb6ce40d9610b32a8b8f041c5a230cac45304f690ec3a;
        pp.proof[89] = 0x1bc0cb9eb6dc60385a38054d5e83923be16a6028f6f1f5ec89dce155efee52f1;
        pp.proof[90] = 0x01;
        pp.proof[91] = 0xc9432aae2acb0d35b8bca43e8c13435207dc74e27f5ad3fcd702a9dfd399906;
        pp.proof[92] = 0x177e6aa59ab8ba9f936b8ea1c45abf5e275b009aa65214da58c603ec00cb15ff;
        pp.proof[93] = 0x2849aef51ba59269880846bd17e7864605414bae22052a69ce8ca2e94d03707b;
        pp.proof[94] = 0x2315be81231c28034363998aa0c7b489de1a2585167e9236d4b8eae211183d24;
        pp.proof[95] = 0x281ffeca6ac7d0fec2d34b60edcb9835cb03099ffbfe49dd59991b54c22f53ce;
        pp.proof[96] = 0x1fd858ab503175b5c961b6f342e7d1b56cf1f04cdd15bb36337028dc7f1ad702;
        pp.proof[97] = 0xa94d124938bdd42888f9a0292823b6f81f7c6bfd369d9bcfedbd9757fd9b17b;
        pp.proof[98] = 0x18b7b8d3bcde1be50ae63909fc36fbe9c501e0ff75888274771311816ab2f77;
        pp.proof[99] = 0x2a450a5c811b75a3aa85e4574b241d5ede3b7c325c91fbf7de68100b32062f13;
        pp.proof[100] = 0x1745ae7db8ef875597de68cee65d6aadb645ce50b753b3b8fb5092169808d057;
        pp.proof[101] = 0x470343d98474b44397282e6358ed9c14824d677cf7e5b6ff35d4bd4e8a0f514;
        pp.proof[102] = 0x2a9ae76d9b5aad6a69ff0918926a46b85a3999626353fd21ea3a253848ed6fcd;
        pp.proof[103] = 0x27e631f53a3e8b8366ede51882990606e7a005c2e087d694bb0001590e90428f;
        pp.proof[104] = 0x25c13b2e160d0323e8d2ac63969f9252e5bd7db81677bbdd5dcbd2806f693705;
        pp.proof[105] = 0x13ad8b06c5b42e33bf9f3b99f7ad32922c211c1e7a8249aa6465686aac51f6ea;
        pp.proof[106] = 0x164daddd4e585ce4f8c323dab639d8ead4d0d480e38e4c3a6fc1559226bb52b0;
        pp.proof[107] = 0x15e57e55a3daab69febe89b97a6ac304209b62b0f51cd3289d16070454c4a8d6;
        pp.proof[108] = 0x130ea863da9c690790516140f75e8b1059895e15edecfa1c29c9dc849263ae3;
        pp.proof[109] = 0x17af902e17d23bd7f20a5323f452d0f2210188b160c038e7dc712df373ea644e;
        pp.proof[110] = 0x191331073adf36d73c913debf6ccd711da1b9a299773f2bcbd4067421ecc1e44;
        pp.proof[111] = 0x155f58bb7d8295284651945a40f70b4ab58ff839853d320d4ec5103512db007;
        pp.proof[112] = 0x12e4ad2012f6b66c0dc5c5290e3967618ba196e609c9326235252fc2c8d44077;
        pp.proof[113] = 0x33154f7ca94a4ff572a9fbebca768f0e598cfc359ec21d2b1760280d7b27a6b;
        pp.proof[114] = 0x2c07021cba60778677c4d40a926464fbbc257e21b96ef779883fe88f9b7359bf;
        pp.proof[115] = 0x8c6f207f58e03ddcd3c3c2d8a3b6b76c29e7c9fd3cd7c526b233cad8ae667fa;
        pp.proof[116] = 0x2c33519fdbc357e54f79a96a8dc7368ee47bf77e69c654f4c6b73cd7c9a4131a;
        pp.proof[117] = 0x2fa54b58dd598b4255b8464b922d83a3e7aa3d5a4596a1709c1f99cd1664e8f2;
        pp.proof[118] = 0x2daee11805e2df3f4d21887da1cc6bb25d2dbaa6d80a31b0d3a430217e388c15;
        pp.proof[119] = 0x209af82723d78a0bfab870ba13a0784df424e8adea58c2ca67f1243879966feb;
        pp.proof[120] = 0xf1c4fa98392d1408982cba8fe1616f76d885bb7648b43a15b56fa25cb4b6ebe;
        pp.proof[121] = 0x2996f85947896aa17780ac61fcdbd7d5a557aa503d0039d605784a8c4e316df2;
        pp.proof[122] = 0x8fc08066bf541f6852c6a1b1851306232f8da1455a833ad2844b0c9bfdbb8fa;
        pp.proof[123] = 0x2e6737ee7bc8b98df7fa4a4d83c6f5357f382d61be0142f3140c34c85b994067;
        pp.proof[124] = 0x191b24f8edd7cb70cbda333940203c3145a8ad7b35638e92c21c5193a002a1ce;
        pp.proof[125] = 0x1a9b206665649dcc6b1efc3f7be3e5e9acfe507423e73322cbd122edf56075cb;
        pp.proof[126] = 0x1e1eb6b8d1b07dee23279f82c80bd08604de00c91475b3d6de5aa718408a4bae;
        pp.proof[127] = 0xd404704de14d77813613a738913aa07c92fecd7858327bf62f2706249364a04;
        pp.proof[128] = 0x14e72a9472dc054123e3fbbf5fdb00500bd95a54b70240f4165f513cb77cd3e3;
        pp.proof[129] = 0x13408dc706efe1ef6e149b01c38e9addedc33b0cc73d562a812bedb00a8824f4;
        pp.proof[130] = 0x195a75cd5698187ab317a5c6610928283311072bf7b7e1d030546b82f5c376f3;
        pp.proof[131] = 0xe7aa767f62c8b31b33922c25bc79aa5ec97945187c92fc7fcd783017ad53f7;
        pp.proof[132] = 0x1be7928bc83ee4443106a074de688874172185f56ec6973ba38f0ce45dbf2e4c;
        pp.proof[133] = 0x16ad5325772e3cb7b305a38a2983eb66b66ad9f64b7f9046bff773394b100f20;
        pp.proof[134] = 0x265ab9018bfd4f8d1ecf4566034ed0514e4be351773c67dd7eda929adacd0a3d;
        pp.proof[135] = 0x268e7bccc1611939ac40cb1dad2b4ae09f0ecff22cbaaed17ad12f1a8e5981e9;
        pp.proof[136] = 0x26d200a45fa2a812272b3d27583f575fbd33469a5fc77f6c8d542268df434d41;
        pp.proof[137] = 0xc8cf0ef533bdec13fe8c3afb4c727fdf40d7640531d3b895c43cba33728fd;
        pp.proof[138] = 0x206bfefbfeadd3ca651cc6a4cacd3a6bd9a11138a168ced4c9a2fef580ebfcab;
        pp.proof[139] = 0x18c302726ec28cb7e1c520593971089f8a6b82e037a0c46a66574f30935e5a76;
        pp.proof[140] = 0x10258208b9a13e532eeec948e370525df6f9691701957434614c0e63cb02e3ac;
        pp.proof[141] = 0x1328aacf49d422e63e905e5ae0d73fe7b7c343d0a7f2cec84983eea722124ab7;
        pp.proof[142] = 0x2af6fa2a48b313d824d0e880ff9f7cb95759b7734c6f7e278ecc65654617f7e7;
        pp.proof[143] = 0x2520908f994fb0b0af7cb9558c347e2b3e89bb7329d571d93ebe746315c28692;
        pp.proof[144] = 0xf848c6ef8c3a3e39a87751054427da1dbfbfe772c128cc48877a1331ecbe710;

        pp.pair = new uint256[](4);
        pp.pair[0] = 20609418911151138324450803130776494232294221285099929470349628441649047881358;
        pp.pair[1] = 20913628764449516285815319081544515959636715479736357390778761120727213427486;
        pp.pair[2] = 18839768183412911405495259650022226421487305205049218581564251173897383055455;
        pp.pair[3] = 1021326530201674480459694577141291720298692740014895280268019946662538854609;

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
        ] = hex"0030220c6ed9d2841133c2029218be8227e5dd0b53b9d7064c45ed0697829646d411b514eaffb2f105f410ffb5e9c21c994970f1172534da0b1c8233a6185ce3ae";
        proof[
            1
        ] = hex"001c9d0f9f253060bc73748c7019b2eeddf03c7aca6cfa49f381430e6b0cbca1450f429664ad53b5c178a72555db864fecb17070d92757845ed9df8d16c321da09";
        proof[
            2
        ] = hex"001561732813f0e977a7c800a0123a7cbdb83af9f152a4c524a88d45856e56a672194776a139fd44f2f6e090cf3c52536a816df6278bfb2aa7c832c86a3e234718";
        proof[
            3
        ] = hex"0002b0de91c35535a813a2a17cee58400ae8a81228babcd201691c7532365e2a522250889aa2466ea66c193efa77c5003f45124aed019dbc6035f560e1bf85c0fb";
        proof[
            4
        ] = hex"002b11bc4ac76ee779a652e7e93a9871fa88e780285d41c6ed2f1646ec8658b74527ac23d456df5fdd5f0a195c660194fad8eed5a93d1b00da54cfbc28cf0c0636";
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
        ] = hex"012de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c304040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc3524f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127202de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c3";
        proof[
            12
        ] = hex"5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449";

        return (account, proof);
    }
}
