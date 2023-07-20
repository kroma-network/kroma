// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Hashing } from "../../libraries/Hashing.sol";
import { Types } from "../../libraries/Types.sol";
import { RLPWriter } from "../../libraries/rlp/RLPWriter.sol";
import { Colosseum } from "../../L1/Colosseum.sol";

library ColosseumTestData {
    uint256 internal constant INVALID_BLOCK_NUMBER = 21;
    bytes32 internal constant PREV_OUTPUT_ROOT =
        0x263c3dff39a9c9fc685a5d07616a19124b5484054ca44656edabe1569bd64a88;
    bytes32 internal constant TARGET_OUTPUT_ROOT =
        0x11facd9abdc2f4c2286f3336b558638a1fbfde64c264233ace144b56f6dda885;

    function outputRootProof()
        internal
        pure
        returns (Types.OutputRootProof memory, Types.OutputRootProof memory)
    {
        Types.OutputRootProof memory src = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x07c36d720614ac2c4b29f4d3f60d9c2ae0c0b38dada03ae98bbdcbe90ecae32c,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0x016fb906ed7974bb75ac39fafb45bf76b14732dfb3a41d12ac04d417c348b376,
            nextBlockHash: 0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55
        });

        Types.OutputRootProof memory dst = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x0127ca9f2fc205f1d96450f463e2d9765bb01a1232191773a4e642bebf7fc25e,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55,
            nextBlockHash: 0xd6fa98f354b0976938c8809e922115f9eaff9282045a56dd389fc3dbadda2ff6
        });

        return (src, dst);
    }

    function publicInput() internal pure returns (Types.PublicInput memory) {
        bytes32[] memory txHashes = new bytes32[](5);
        txHashes[0] = 0x347f8e7947782c8ce40a50e7cb2133dc33db5e2bf42535b82a523d131de07c5c;
        txHashes[1] = 0x11d7498a85370b5ff0c83427e3e1631629db7f6ee37d13187588cd5990d8315c;
        txHashes[2] = 0xaa84dcd2e2bf71ab166a297ff7b57a7a4f200627ffcccaa039a90cc8030b55c8;
        txHashes[3] = 0x97b1830777dd0e122e9f3c202958b3970e6f15902cff51f1d27eabe82aed079b;
        txHashes[4] = 0xdd190c63d32075b707635b6ea9704c4d23952fdd1992a40652f40a159a2d8497;

        return
            Types.PublicInput({
                blockHash: 0xb1a1d1c45920976d42b489c8c80680498ccd41bfc3086ff6d1fb07cdd0870d55,
                parentHash: 0x016fb906ed7974bb75ac39fafb45bf76b14732dfb3a41d12ac04d417c348b376,
                timestamp: 0x64b7980d,
                number: 0x15,
                gasLimit: 0x1c9c380,
                baseFee: 0x3b28afc,
                transactionsRoot: 0x90572ca3650e693acd8ec001d84c9c8c3122de052c64e0915cbac2a2f2727a78,
                stateRoot: 0x0127ca9f2fc205f1d96450f463e2d9765bb01a1232191773a4e642bebf7fc25e,
                withdrawalsRoot: 0x0,
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
                    hex"cffb7b9369b08f160e112c1d468ecdc8d718f339b38e03cc9f755580945badd6"
                ),
                logsBloom: RLPWriter.writeBytes(
                    hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
                ),
                difficulty: RLPWriter.writeUint(0),
                gasUsed: RLPWriter.writeUint(0x21ebd),
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
        pp.proof[0] = 0x14e783919de06914136b530fd57e4013e93986463156040aa9dcee69a95a3f06;
        pp.proof[1] = 0x47112a4b747334e3177c66dda5e863f443e96f5f84d13c4b2943c261cf4959c;
        pp.proof[2] = 0x3576436bc1f19086793791f6bfb1a61ff2c58be9bde6e04bd97792224e1667a;
        pp.proof[3] = 0x2e3074feca3a72b5b76766545f5983dbbd74df7394ec340e3b10e68a348dbfc8;
        pp.proof[4] = 0x11b9df28741cba47a3505b83174047026abb4b91ef62cfb57d09be0f506d3696;
        pp.proof[5] = 0x1615cb789139466edfd611db3e05d6684ba2a5a3e5889175ccc8a6d4fe50ee80;
        pp.proof[6] = 0xdd54fc8600cb83dfc3e45f6d4ccd7e51ffe0b499166392ff4e557b2cce30ee2;
        pp.proof[7] = 0x145f4531281f4a1850930f5fb8e5c7197ba2ff283166caefeb14e657f4e08633;
        pp.proof[8] = 0x25b90ae5ba3dbac34bd69a5e0b9ab8cd5e8fd8338d863ade0653ea2442d7e4f3;
        pp.proof[9] = 0x1aed0be46c8d2ca30a4eb602a4de46407571d0b1e3e92d4cfe5e480b60f3ba24;
        pp.proof[10] = 0xfb57da872830a262859e8cfa7ba195564a7bb92e73fa495e69548d679d9e0a8;
        pp.proof[11] = 0x898bd8f6133ebceef397430a95f324233bfe3ec89d5ba2bf81af2ae38361c7a;
        pp.proof[12] = 0x8f596b198b6eaf5afac7b18ad72d13d04a0944e519cb22cb7754c477adfd0db;
        pp.proof[13] = 0x74f4ce48d4f2b687bce0eacae7ed86aecfd78a38e67e03318bae22aeeca531a;
        pp.proof[14] = 0x25e1fb509885c2408458739e54b9874664d2920e6e1f9666e0516eb92303ac77;
        pp.proof[15] = 0xbd4b50f9638cdf842994133d90ecadad2bb4b8c90b8b1ee7cee197491c899db;
        pp.proof[16] = 0x9027c283d46a4926e52b91a3a261834a4fbe49583b84ed6f529a1e095d03b07;
        pp.proof[17] = 0x21eaa0dfb06ce47b6bdcd879a478a5094be96158a7d58672d1a36c5deb619d16;
        pp.proof[18] = 0x1f1aedc227a78fb5a5e18039d6866ef0f84b8c7a0beb1911954a97fc1fdbd31f;
        pp.proof[19] = 0x2b883e10bd651db05fa10b9913e42a855e672dbabc608c5af760fb8cc0931ffc;
        pp.proof[20] = 0x125f1b4f2b5d4d9271c8e9f70db358d6ba22f1cee1ae62f4f8e4597264f79a6a;
        pp.proof[21] = 0x8ee9d530aa46d22b502a54ba8be27d767f1035e2a1ff5518b3d9838a2fbec94;
        pp.proof[22] = 0x104e6d61d1e61e03dfc42071bfe5e8e173245a73e760fc191aa5bd488d2f8a5c;
        pp.proof[23] = 0x2e3ca1b4ea9b193afb96f7e8525ec299248f9596e168136ed3e801b32a79e915;
        pp.proof[24] = 0x126070385a2b9cec94c239c6078ed6f4f39361ddb4ee96f4d19cd08ade7c97cd;
        pp.proof[25] = 0xeded57b2a7d1ca9a26c595fb77056d490979aea72e28e744532948fa92c1744;
        pp.proof[26] = 0x250a17b2c376c55ba5ec57e0a89965a95712af16debb3a712d47ae0932c45609;
        pp.proof[27] = 0xfd8073119c810e238128542284c6b4fcf5cbeb234573bee7cf3f9a98b5b5692;
        pp.proof[28] = 0x652829df71aa07b8505f2f4d4dac7f67ebcaada17172bc7e31a4998c99806a8;
        pp.proof[29] = 0x14586bdfaa1b67c7252e3f41b65172a2151df8c641972e44e22f3b78851bfc43;
        pp.proof[30] = 0x1ae3c28032c1038e7542362c1982c6688d2278a9e98ab54b5f8a2dc5755981c7;
        pp.proof[31] = 0x12d276bdf2fa5b5845a160dab2b87f7ceec887085e41e77a45e8c6ea42299ebe;
        pp.proof[32] = 0x2e81389b7bc027e7031bbd49f7531f2d39cd8defe65958b7a071984621ad7f06;
        pp.proof[33] = 0x249b935f2f12fda0773e3d1bb8defa99ea3bc8e780188cef2d3b6d9563bdbb5e;
        pp.proof[34] = 0xf3f0670075e17fa81aaa96a0faf38474798ba4e251be7b71834d068faab21e1;
        pp.proof[35] = 0x1391eef7048a66a5ffa537ea26e871ad6a64f60f3221e8d5f2e71b55ecaf4f1b;
        pp.proof[36] = 0x24ec794eb996e4db8179531163beb1a865070aa5314f76e17f0086f9af31ec5c;
        pp.proof[37] = 0xbc7eb1901aefeb499064539549c318c17de4b81d895962180edbf63a170d15a;
        pp.proof[38] = 0x24c0b26ed38fe97dcc902221ddee6d8dfca797e5821f750f2285ff6a24acd9ff;
        pp.proof[39] = 0x1cee7f06a645b4d07133fcc2ab7a515a5c147374a76dacd170073234ef2827db;
        pp.proof[40] = 0x1ee45d0198a015ed64fe754312fd38a66c7f7d19e3c296657645f959fbae2727;
        pp.proof[41] = 0x24b72515c09b5fa8ddc59dafc3592fca63e15b519255545400cf318a83adb9be;
        pp.proof[42] = 0x396b9abb4ee2aea0613bf7d0b9b282431ed921502e5210536c3569174f6c3ea;
        pp.proof[43] = 0xf1831fa1a5d8d48c51b6dda117e09cfa1821f142ba0441bb75e9ea83d390df7;
        pp.proof[44] = 0x238af58ef595450d346dd330464df27d1af4991d7a0ad2b22f6042a1b02e8f15;
        pp.proof[45] = 0x1445be844d54fbe0ba9e3ba7d4419b4c1931233c758d4b265de315443a250290;
        pp.proof[46] = 0x22f2d8cf4c128474edd10d8565ec88a769a663adeb8248fbad719730c8c53778;
        pp.proof[47] = 0x2d8e3735defdbcbe1bc6dab9bd52897431fb887fe0bd061870fb0fe5851375ad;
        pp.proof[48] = 0x110fe9d4efbdfddf414784ddba69f51c379f0b11df17ae095c975c34df661cb;
        pp.proof[49] = 0x51808b6507f74ce9162773d5d298258f1627e820983277d0ee54c029a292663;
        pp.proof[50] = 0x28f189eff19361a1b1f345e55fe052b90b6fc47c9c689bc72b471814e962e9c5;
        pp.proof[51] = 0x2335d63c42f144096e549a96c66b177bae60c38c4ed90c8df18b161dfcb1f8ba;
        pp.proof[52] = 0xd0d7f934b56f78922dbab317457468769a398780d56421957d4496cf579efff;
        pp.proof[53] = 0x23e8626492014698c3ede55502a8a805460825d98f279370d92a92ebfbedc4e;
        pp.proof[54] = 0x2affb42f683b61952e50662ffed95c882e5949ea8e6ae60f8938b7307adeb456;
        pp.proof[55] = 0x1269e108082c9484d5c73e03f3511cc021d37b3911355fa69f77a65cd3421cd5;
        pp.proof[56] = 0x1;
        pp.proof[57] = 0x2;
        pp.proof[58] = 0x2b024a5d05957db019280c931e7e54e683eb881c6ec598126e850dfd29052277;
        pp.proof[59] = 0x62bc18f56e275c3cd448216f9dd21ce549f4be843b9b5e72e2dfcb0e7474268;
        pp.proof[60] = 0x18d270b455b4ffb80fa2ea5eb630635d34c908ed18aa91663e65cd028c11ee7c;
        pp.proof[61] = 0x1e0f552b39daf767510d0d57b578c6916b95663ff9918540d0f389b7bf5f1ab6;
        pp.proof[62] = 0x2738cc863508a42a354cef2f7b75175977f36e70add4facf7ec897dc9d344c42;
        pp.proof[63] = 0x18606eac211b8165e516a8c67f735506de2492644fd0e35448734e7132da9fce;
        pp.proof[64] = 0xb2236576b641a58ea5f96fbecd27740d799496a6d9990cea762c3ee1a46c557;
        pp.proof[65] = 0x2ad9d4dd1321e35a5d9339cc7ec54666fbfeee9361d884d4eba98b5af1823692;
        pp.proof[66] = 0x1afff79f282b7405fc34ace07b333f7eb6a0e23cd1ff75431fb9687cc1fc3405;
        pp.proof[67] = 0xf78e305ff093f3adbc3fe1f383381f04f6614d84a77762628d718ddf2c113fc;
        pp.proof[68] = 0x1b62aca3f1c51de4393832de7ac7cbc69795b3b2d7bd8db6c5c6ce1106fa882c;
        pp.proof[69] = 0x2e615d8f8207f329f7f09436b536e7782e4218fb10eeea750f21efd43f7a60f4;
        pp.proof[70] = 0x12c17e82452dec535ee8ce7620bbbe8e0ecbaf03bcc901999dd3edc1dca48b8d;
        pp.proof[71] = 0x22005b3efd296be9c669a21e0da3ca7714c5193f91d96ca2956601c25d92ec58;
        pp.proof[72] = 0x232774507c4b2c399a397ecc0dd35c40e265305832e1e42e68e1e5f51429f294;
        pp.proof[73] = 0x1488eaa9ce3d1cbdd071fbc9dcf2b2925b4c08ad1af261f52644d9b8d38515fe;
        pp.proof[74] = 0xaf34b96bf90461a50d09e10212d5ee6e1a2218139a33257abb8c3a1d1d95818;
        pp.proof[75] = 0x23d967417306d3b68a235c82b84e3cd258640af8246d1be789ea292ef2664e44;
        pp.proof[76] = 0x20ec4681e0cb3f0a967d517607f80dafc5361485ab4a3e89d66165578d6c36e5;
        pp.proof[77] = 0x48fbc9d1a1ff97e0b26193acdd832c8704288f04571da9a8c1e7c126d73e291;
        pp.proof[78] = 0x1b48dd85161adeb09754886475263885ad58c08916e6e9a5fac132c2e450fa75;
        pp.proof[79] = 0x2868c6c1c723ac3fb5fcb9f5ef4b34ba59955f4697100b16be4a3e70cc3047f5;
        pp.proof[80] = 0xb5e3fa01ecf708f3951ccb1e8bea7e29607a7876a8085634aebad612d243d3;
        pp.proof[81] = 0xba0b82e4cb4cc9f205dc024cebb125755560523281db4193299a7d487cc3841;
        pp.proof[82] = 0x7cff2a41abc2d849413fdb0486df0521584a67c971fa14a0ba58e035c0f8b24;
        pp.proof[83] = 0x2678f3c5c2605f9aee45918d15a8b41160354464b0f8a645b1cbef6c661e3b9;
        pp.proof[84] = 0x2f8746ee48676606468d451ea6c76ba52a6d50bccbd5ea73fedd1a81e5e27f68;
        pp.proof[85] = 0x2196ec84056bf89252d974a928bcfaa6b1688df80af1a5338ca37d0eca8a972c;
        pp.proof[86] = 0x220acb603416f3f745a3e441ce47e6410a1b4807999dec28d131c6c80677f4f8;
        pp.proof[87] = 0x2fd9408a4e11545339ab808180ac60218b21b516de13bf0cdf8ab0b6fa8e56ef;
        pp.proof[88] = 0x2d2117ac2e2fc25cfaee4f8cf418ada83b1ef045c9f81ba644dd7200bc758f0;
        pp.proof[89] = 0x2196ec84056bf89252d974a928bcfaa6b1688df80af1a5338ca37d0eca8a972c;
        pp.proof[90] = 0x1;
        pp.proof[91] = 0x1d8b642fc241c3f1e4728317eaaa9e7a4807e3b66ed4b73af9c97bb6c22bc489;
        pp.proof[92] = 0x8126347ec1e34cf1e36d04de2ad7803562ede631e96c55c084a328e21534a9a;
        pp.proof[93] = 0x14d3b36066212b7f7d4eeed156a3d1b44b2c3922840d791d6dfb488dbbef146c;
        pp.proof[94] = 0x2a1ffa9618152033ac9ac804cdb0a6dc1e45a8f359d7141cd13399ad0968c00;
        pp.proof[95] = 0x88fdc5e27acac404adb19de1876cbadb11392876e91a1fb442e9c1b1ceae4b2;
        pp.proof[96] = 0x1683b9ec5a127f20b6f006e1d91196ac101108f24588621093d1a9f25996332c;
        pp.proof[97] = 0x2b271cda1fdec74fcfdb29f4d9a9ecc43e0ed35cbcdbf8cacbd9f4417709a3c4;
        pp.proof[98] = 0x40f1bfb5941533e6af01185cbaecac591d4c826918d7d184ad700432d32e654;
        pp.proof[99] = 0x2d42491e1039b2ec7e3245b5e1e2e191111818fce6ec17158d808cab4f4540a;
        pp.proof[100] = 0x15843baa6fef3772664746bd8920a811ce683242dec13620e98d531638782b3b;
        pp.proof[101] = 0x5f8cb33869d3e2b51bfe6afa55bcd4b34361c1ab8e56456c166a599db75f0c2;
        pp.proof[102] = 0x2623b442c36913dac485d2d67d089dea111a3cb43228775473a2383b48155dda;
        pp.proof[103] = 0x1398da3821fcc7fdc8f61ee0ba2561a7db2d456df8d9d64c76376ced5e2cc771;
        pp.proof[104] = 0x3603355cbcf57b6fee000a869470e99752558d4e8fcb19afa2641864fe7070e;
        pp.proof[105] = 0x53e277494b7d5003ca1ae909a2a8e794f7ca2a726355f16dd146bf95449ed12;
        pp.proof[106] = 0xe7eb5c93cc1ffde32798c9461c6521d6dfc3b309b980a0380a84a96d9a81b70;
        pp.proof[107] = 0x115c8778644a34d408ef782b32d239cd7ba9ac19a54ea63112df40be655b5d35;
        pp.proof[108] = 0x1157d25dc3a3571504b693d67ba580983ac369a5a7ad5d4578f6b826144bfeec;
        pp.proof[109] = 0x2cea95e3ce25a545ec389ba373fe9bbd6455a017aa3d9891d5830a1daa962a40;
        pp.proof[110] = 0xefc23d06d04ed438d2b850e4102023aef74694861d0c2f79fc1fcbe6b4e4e25;
        pp.proof[111] = 0x149e93044ea6b1c5efe38deeee6d0a7b1b43d8dde1f758d571173d025f240cc6;
        pp.proof[112] = 0x787371c3d83cb5d66cd8bbb7138f7849795c841e91791634ced953a8f559af3;
        pp.proof[113] = 0xb3772a82de5e69383c82ed6f10842e1d689c912a41cb12436fe14556484c934;
        pp.proof[114] = 0x26f57637bf5c91cd1a22f02c74e368d0e4ee06fbfbf22ca2968c0e7605e4201a;
        pp.proof[115] = 0x25297496114bbf7f3d0a365d6338c1ff7660a95ca7b792831ce80da2e06a2774;
        pp.proof[116] = 0xa7f88552e875d584f726e918ddfd2d38327182184b6be6726d4075e4ee17b40;
        pp.proof[117] = 0x1f8c8ce3aa200f5b2c6eca8334849ed8a4e673214c2bdfeb2be31a6a4fa18ed4;
        pp.proof[118] = 0x9a81cc96dbf5d5a28c764b2fa17e7d1bcd4295e062af20b2f49fdafaa9d1fc1;
        pp.proof[119] = 0x10504317c05b7bd1ab65e57c25e3700e2d0d25d01478aa0402e0fd435552c971;
        pp.proof[120] = 0x2dfbc34669d876c60334cf164a0183da2fe5c14fe85b08fd7f27de3aceed4ce5;
        pp.proof[121] = 0xcb7621718d25fd8dab8d67cb9c701148ee62aaa8bf48bc621b2949215068c82;
        pp.proof[122] = 0x1311f1bbc2e2b848a9441a17eb9f7898c1659a4cea89c6bbb04f8b129ab8fc5;
        pp.proof[123] = 0x2bd04628db213aa9f0dd2bf6e6f64efcd53e72e644760d934c5f61c0d399e9d4;
        pp.proof[124] = 0x2523590a2a6e7f75f7d056d688e3d0319ebd667bfd0aa66b9030900d0fffd318;
        pp.proof[125] = 0x70c83f433c76cb9630287c7e76920257140fd905dc7ffc043aece436db1b727;
        pp.proof[126] = 0x236b6ad4464f9fa8b87f2f2804c81d8b043deae47515e0a626691cd895a30024;
        pp.proof[127] = 0x15762864ff70a529920ec8094ecfe20ecc76a2282859f57f7a45fefb8bd22968;
        pp.proof[128] = 0x19fb23c5f13c4cd02073374c1d928546a4e1f62a5bea63ca1eaaf1c7366f8ee6;
        pp.proof[129] = 0x25023a7aa8a9e0a2a02b62237de31eacac54ce945919cecec9834e3d33ae2df5;
        pp.proof[130] = 0x1a6908d3c2508df8e0f627ca6e5aa9e8c028edcdc2bcc3d2653e822d4b7488ed;
        pp.proof[131] = 0x1b25e252ea10723a1a256f6438874fb1ac2e2f5e689f923622665699cab7dac0;
        pp.proof[132] = 0x425a7ed286d790bf44c96b4f592854c6d821f69c48cd17be40052c5278631de;
        pp.proof[133] = 0x795bb2dea59444945304cb30f0d2f180909e0a4f11851980ac051cc771bf82b;
        pp.proof[134] = 0x756e4f04be12f4bc4d17036d5500dbb1e1106f274227639d4992861299bc79e;
        pp.proof[135] = 0x1b40d3cb1e6744edc7570ecb3359d117843187267c6a2360bf7c593241c58790;
        pp.proof[136] = 0x25c82982e698cc6d5c4218147f8093bfd79a6fc28c8934ce0649f18d32c80c6f;
        pp.proof[137] = 0xc4b36e1c2b28f164c1ca7256c28bd8b244ed924ef751ee9a9db2b01f46e15c8;
        pp.proof[138] = 0x4a18c61a2c52248d9c6a4b8d196433cfe1f22a47138c8f960b82e5b67d91c19;
        pp.proof[139] = 0x27ba28ffba0a22680f688c07ee346226fadd17c0c63756bfae48012595eec6cf;
        pp.proof[140] = 0x86b682873b4824020b7d14cbc2d74dac1cc1df940bcbdee5a1fc65d4395dcfc;
        pp.proof[141] = 0x1148a55f374386563a81c43039ef903bc4c038af91e833a8036d5cadc8fc3f85;
        pp.proof[142] = 0x2858c31847e013411a19707498c7dfab5cd580ef383541ae31ff85e70a10fe1f;
        pp.proof[143] = 0x104ef6a6d4d0691b403ff9a7c29bffcec0e838cb6bdf3456b561bd437343ce66;
        pp.proof[144] = 0x97515e432c1b75357f4d209b300e5fe734c097ab5b84b054798336bed144786;

        pp.pair = new uint256[](4);
        pp.pair[0] = 8639398281568583195745543810727361683126111716677764693786436845797934326476;
        pp.pair[1] = 10008025483608756794486257442748651659515200453101863211203755091379161419123;
        pp.pair[2] = 19384061263476939563609921376769842609865162061918810795439396082482774507015;
        pp.pair[3] = 2994949873454822815381598779671562094664356126804871960294778420759678916516;

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
        ] = hex"0011608e4c1635552de6d1ce4ae41799e32c2d3baa13e05017029c8db503db4e6b0a4478c7f4040b98b10cd90a9e4a6c90f951dc8b65b73845471c49835b55415b";
        proof[
            1
        ] = hex"0002d35b1fb18e43d4cb3883a6a64cb882312db9ee0868c8e425919c2acc67b6fe2862cefc25230c03bcbd86bdb58b6f7662506e5d9af265239f2570df170089a9";
        proof[
            2
        ] = hex"002b189562f03c7653f51509a7840ba909c7bdbbab60778457f01c66649034f87f14c3c9034672c8e4adc74d672f746d96f24d703ab201f95b16c18e120caeb722";
        proof[
            3
        ] = hex"002bf1dc335fbd3c6252fa426d647d20ff86181a1a3ac1f4c53122b5a853436a592f737a81aa1525e0e68bd3e34dbc70f09386352308f59fc5c557ad6f0d389663";
        proof[
            4
        ] = hex"00188f4a9b605caa246413179613bffe6c4efcd717d4d68a2dbc531123c8d2a4531d12fb4a76dc568a4655d2e360e9bdcd9f86542d6f7d327e3140477c92aa2920";
        proof[
            5
        ] = hex"00244f31e0a770ed9dbdf583a0e0dbe036f5f9476c3cc761e18b2a475a96bcc10809387d2d4be643d4d8df653dcb8495eebc36eba91f993f78fe843a799a6d0fb9";
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
