// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Hashing } from "../../libraries/Hashing.sol";
import { Types } from "../../libraries/Types.sol";
import { RLPWriter } from "../../libraries/rlp/RLPWriter.sol";
import { Colosseum } from "../../L1/Colosseum.sol";

library ColosseumTestData {
    uint256 internal constant INVALID_BLOCK_NUMBER = 21;
    bytes32 internal constant PREV_OUTPUT_ROOT =
        0x50db02c3a09408b9a4d244bc49c23b5af54270edd55e94bf2c9ba568a4e7bd39;
    bytes32 internal constant TARGET_OUTPUT_ROOT =
        0x0103a1da2dd7354e185a2cdfffad16491e2836dbcff0c92177c4ca218eb5a823;

    function outputRootProof()
        internal
        pure
        returns (Types.OutputRootProof memory, Types.OutputRootProof memory)
    {
        Types.OutputRootProof memory src = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x0b1fd0c9beef9cb4ac073bc7ac702f237060926b59ad9b58c34225c3c2883042,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0xc5b85aa0fcef96933ee0e2629c1cd30dd7d14083c03f08b1ab95f1e9f3757b46,
            nextBlockHash: 0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b
        });

        Types.OutputRootProof memory dst = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: 0x2ff9f8349c59d025d078de07d9f7175c3422109c9a11a4075d383ca62a685716,
            messagePasserStorageRoot: 0x24f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127,
            blockHash: 0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b,
            nextBlockHash: 0x7b8aa6641024ecc150b38d9710d40aa6d77f800d2fd1788660914f190840af20
        });

        return (src, dst);
    }

    function publicInput() internal pure returns (Types.PublicInput memory) {
        bytes32[] memory txHashes = new bytes32[](1);
        txHashes[0] = 0x2c9d326676e7b90ef6e9bf330592b667ac081d01ef57ca4f14281e7afaab2dd3;

        return
            Types.PublicInput({
                blockHash: 0xd3555c9d09ce9863ee8fbe80ce7006ffa0750115b215f0c5c7d8f658bfd8420b,
                parentHash: 0xc5b85aa0fcef96933ee0e2629c1cd30dd7d14083c03f08b1ab95f1e9f3757b46,
                timestamp: 0x660e9c50,
                number: 0x15,
                gasLimit: 0x1c9c380,
                baseFee: 0x1,
                transactionsRoot: 0x91c1194e5d35a6f40e1aa4913664dacbcfbb36e8246e9b60ef1c00184f6ff05b,
                stateRoot: 0x2ff9f8349c59d025d078de07d9f7175c3422109c9a11a4075d383ca62a685716,
                withdrawalsRoot: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421,
                txHashes: txHashes,
                blobGasUsed: 0x0,
                excessBlobGas: 0x0,
                parentBeaconRoot: 0xf4db98f24ef4e2ce756ee5c926671ee2d68fba50c72ae29e25f45a1a8ec5966f
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
                    hex"d6016faa5d7979ca763d5cbc6a228df98194cd3ecbd74e0b22833d55bd9028bf"
                ),
                logsBloom: RLPWriter.writeBytes(
                    hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
                ),
                difficulty: RLPWriter.writeUint(0),
                gasUsed: RLPWriter.writeUint(0xdb30),
                extraData: RLPWriter.writeBytes(hex""),
                mixHash: RLPWriter.writeBytes(
                    hex"363cf43f812d151c2e9d64de797403c7c73c973ec31b5de66db208de4a5c9d1b"
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
        pp.proof[0] = 0xc465b6d797615f4ed6837b783b9f94f4efa3a67a631e83dfeed2f9c56517b1a;
        pp.proof[1] = 0x272556a23e1b872b9197437467b244f31b4af935aaa6cbc87e7db8026fbbc06f;
        pp.proof[2] = 0x65f7f0badc319d8ea8796d16229e16387a37bef209104bca09e1be0b795af4b;
        pp.proof[3] = 0x1ef6eaa311722d38106ba3a9d31f36aadd4a72d78ea7e7aaf58a12b55f091554;
        pp.proof[4] = 0x234f179cba9fdda4eb2a6d472022658642f7d95ac7484d286f689da489da5c0;
        pp.proof[5] = 0x178ea5ea382ff64448b76f7f471232915d41e506f04138848fb5cdf56ba3b54;
        pp.proof[6] = 0x68e206dc4f91957685c395e8d782b65c5a9bfa4dc0065c58faac185f3e883b7;
        pp.proof[7] = 0x87da57de86cc118cfd5ec886ffdc0b2976d17983f87aa67c401667c8cf68093;
        pp.proof[8] = 0x1db7dfc21d3cc45c7cadc869e348820aecaaee5a56e4480a5f85bd0b3c252344;
        pp.proof[9] = 0xfe83ad7ee58d5144013a6919f5be73ca99a43f485498c54acd3bff8f6db8698;
        pp.proof[10] = 0x293f0b038cacdbde19c99fd9dcfc1f4bf14f85e62e26893eb5e4f78a4ee4fdc9;
        pp.proof[11] = 0x2acfa730f297b02a568ff51b2ea9890beddbd3c03a07ba1a99c247b63e9ce7c0;
        pp.proof[12] = 0x10ff58e3ff7c926eeb20bd2ead6389d3a02d661d79ec0ace5bd429fbe84242b7;
        pp.proof[13] = 0xb862bb9981b28f333602c4f442baac3a8b86d7f37b9f11b465b5b3026808c91;
        pp.proof[14] = 0x2954bf783f7f311eaaec0d03ab047daaa32626cf2edb159e0edaa28fcc438818;
        pp.proof[15] = 0x2331fd94b75c69f7bc0e1af61f691c91ccc1d63358475bd56ad716747d70824f;
        pp.proof[16] = 0x1a842c311419ddd6edb94029258a5fa57a8f25cd1edce673291ca4d0aebeae61;
        pp.proof[17] = 0x58b76cbf77fdb056b880df4d877b3b76b5750d92949455da0d72c5072f10b88;
        pp.proof[18] = 0xbaa8c26132a2b3db28ac92c99316df7f8a604028658ee292129d3db9db1e917;
        pp.proof[19] = 0xebfe27829fa223b95479fdc3900cb455a782ed9fd3b4ddb52662cf28d740eb7;
        pp.proof[20] = 0xf560e26ea271f9abd2029bcedd260e9519410324f43fe55cc4eb6acf970e98a;
        pp.proof[21] = 0x19324a06c0b913a49bd1ab2e71dbfc9fc625583e71ec3059c1be0cf88a819ca6;
        pp.proof[22] = 0x14c0af298216fc68f3d7f2ed86095a32b9cedc4ea238ea692b9b9e16cdea0b0b;
        pp.proof[23] = 0x26afcb563a1172282207b5354d025881ebec5dd0312c3e8695f41ee03df3f5e8;
        pp.proof[24] = 0xd83a9328c74fe2ccfa4e024f97f1141413d0365b10fff5992cdb2e2413860eb;
        pp.proof[25] = 0x87dea11223eb0590f823c3ed59b59c8d58c98e2ee08744e27e28d92bae0af8e;
        pp.proof[26] = 0x1bf1bbe76b6944899b6f50a43de91ff2ab5100d417ccd45c2d22e469d13e611b;
        pp.proof[27] = 0x1fbcb630e5f339ec57d6712dfde7d6d8e48db584b5145b4e8818cefa638e168f;
        pp.proof[28] = 0x998845b727bb858e4ad30f4dd74894d93620982fa6e2616fc04edc31ded0c2;
        pp.proof[29] = 0x239bd89afac66ef8a91350026ad81099beb161a97205fb24181eb5bfc80405a0;
        pp.proof[30] = 0x8530008500b253bdc4c97b72cc03184f37b4bb8fac18408d4efec3951548bdc;
        pp.proof[31] = 0x2712e1db9b893cf074456318e1eb5d32ccd823e0e90932af865d80892e9d147c;
        pp.proof[32] = 0x18b0fd7ea5064b1dd2827b5ba99b68d1a622002c46b6ff1e22edb22cc296643;
        pp.proof[33] = 0x26b4c4b2e4f1c1f5ff0d10d9b68b60710ed456841554c45391be0a91c0b66f1;
        pp.proof[34] = 0x44afd9dd5359d4d21c3ad873abb62894e2a5d2dde84c30615782da018b4ccdd;
        pp.proof[35] = 0x4678286b497f1b46bc004b4389e542ffbc433800ccc6379d5b9cb1ec3acaa99;
        pp.proof[36] = 0x28bc0a63edefb0dc39594da11239ab9b339486d17a53ad94cbbec07752bbece6;
        pp.proof[37] = 0x1b1d63140103821c911425663a034f86bf7c1e7781220190dd2b9722e1c5ae93;
        pp.proof[38] = 0x2abb1d8960e0ab2044c4f1f6d2e0b05286a1452442cc41b48d3c9993d3a8fd13;
        pp.proof[39] = 0xb86cd3670472bdae3fd7825605fcf2796d62949e8ae0f645c83597af62d6d12;
        pp.proof[40] = 0x123ce94be75385c1b78d9c905b4068852a5a9d1500e575dec7afcd2a97a8c5db;
        pp.proof[41] = 0x294399b080cacd277ff77bde0a2342cc55257b85842c08dec0e15eb704f93756;
        pp.proof[42] = 0x297f54db0810b5e1e96d665190be1bcdd2c25a0a40310ea1cd2f1d45e67072e3;
        pp.proof[43] = 0x7eb61dcf33ac68d51beaceb8661b5289d9e96f8cbaded0c45642802e92172dc;
        pp.proof[44] = 0x4a44e4d83e6a7ac321615b3e3d93475c584cdebf9bdaac504b6e9d49331dca3;
        pp.proof[45] = 0xdb485c5f33c6f445ac168474a06b4cfabd3e5a4a58593680797a8055bb68762;
        pp.proof[46] = 0x8634b7448790129e14797798e144ce0136ee1b33357d6281820e5c6b79daac5;
        pp.proof[47] = 0x2df3be3209b171b59f3fbd2cf18f5ca7b678ea42875a799c3c426013caff513;
        pp.proof[48] = 0x1ea6261c9b9a06f0f94d74a335a3d8956547bd01efcc8eacf414d55dd8414105;
        pp.proof[49] = 0x25c8d7e96b0301ddf4f4fb3938f7dcb393e3d8adb5afb83e04f2dc21b39ce3b2;
        pp.proof[50] = 0x14b990bb21594e92106bf4617a130e63543fb82e7dd067f64403f12bb3c86065;
        pp.proof[51] = 0x147fc2bfc89a88ca91fff31ee844dc6d3f818a58595a3a8d5e6c8942296399b2;
        pp.proof[52] = 0x28093e0333c1d1c5905b03422f1b0cf7dd0007e835891a04cb119a0305b0ea74;
        pp.proof[53] = 0xe44b1ee87767f57c00bee1377600286cbd0339d88da7d1622e208888f761eba;
        pp.proof[54] = 0x116c6d3568bb233b159a56b9ea6424200cca3524e25f83d06cbc48ecfaba9b4a;
        pp.proof[55] = 0x16abe8b507386bccd0db17f74cb8c5d042594b8b6e05bf841bc0a3c2fe66d925;
        pp.proof[56] = 0x01;
        pp.proof[57] = 0x02;
        pp.proof[58] = 0x2d84e5e6e7d0f6fca025544ff3aa761d77b7170ae0d4ed6d46010597a935ca4;
        pp.proof[59] = 0xe9fa30dc923f4b59dba81e1a3ef0d7e1a4f0b48c1c89cfa5f207433c25358f3;
        pp.proof[60] = 0x24b5b214d6f051169613fea140cba16f1ab11dd28976f296a35dff5e432b224;
        pp.proof[61] = 0x15b939da5a6ff9502f891b2f98d0159f3fa6a312244931882446ce1fcb772d1f;
        pp.proof[62] = 0x160e3248d006c84b9d6c98557b02118cbee05fd851a92892ad3492491ecad705;
        pp.proof[63] = 0x2dc78ef4a77ce1f90ae22a46779e711889d5a204d112db880c471e6e1d9c4ac5;
        pp.proof[64] = 0x2e411ed3f13f57063ec51459bc7be380b9528e785e9be1c943814e0f8f4bb201;
        pp.proof[65] = 0x2e2c64f5925364824f9fe0640e1f0fdaee57287f788f1de7065a7d8c95961a36;
        pp.proof[66] = 0x2c80245edaf387862bf0c1945b6aaeaa16c624af54062f97dbb3c9690cc9c981;
        pp.proof[67] = 0xe68d3910ff72ffe74dc61a5ce9816224493dfafac39d59bd5e602eba88d3b00;
        pp.proof[68] = 0x170315391948998bafa32200024f9adb5bb2e3c212d9f0512b164126fad47c91;
        pp.proof[69] = 0x2833894b32f24e6ed2f99b628b0ad2220bfe15cc5c580d0e1a3c2902bc2ab0c2;
        pp.proof[70] = 0x164cbeb31b9e86bdfdc16ecb2e2c5fbecaaf17a8d07396e74e4699f15cb8b653;
        pp.proof[71] = 0x2d9af87312f7591404fff0ef076eb505c37ffc21c14a9aa234980a402143bbd5;
        pp.proof[72] = 0x2a1297599f6757208560e57c824d28862cc0483e5d6403919b41884844f7cc43;
        pp.proof[73] = 0x2a62c30424435df62d212869b6174baa9357f2d34457659398b19114893e5089;
        pp.proof[74] = 0x105ac1283b72f8b0d5fa3b790823ea65b407da63482f27bc4a70c8f06a03724e;
        pp.proof[75] = 0x2c98cb20cfca5fcc1bbac38ca67c06fe987e67d2626ccb1e5d04eb9756504cfe;
        pp.proof[76] = 0x871930b404eb88a8b967962699228b3b851aca258593b698be306078c0f2112;
        pp.proof[77] = 0x2729a77b2580c015d71251914c7fba0a8b6e7e717985ecece4c8b74bbec17770;
        pp.proof[78] = 0x1aa43cd2b55efc37a10812d303b156daa52d6a4951eefe9c21b96af1dbbdc1d7;
        pp.proof[79] = 0x489c214058fa1ee1e06550e72ca4efc35142a3d1b9359bde92d1ddff21c4200;
        pp.proof[80] = 0xee866ca57953895fb45c7c09b1af1f4269fc79ad6b77f765e57ff9253363972;
        pp.proof[81] = 0x2bbefe87c925e04eeb22b7243d57338d81efe66f99d305e69afed7c2a227952c;
        pp.proof[82] = 0x23cd195635e1b05a9fc6940ca1418afb77104d2c17b08f80655c1f6f9d8ccb40;
        pp.proof[83] = 0x8fa9431b2edf32baa738a3d238970cfa586c3729639704de891f92e60e4f033;
        pp.proof[84] = 0x10c1e095a9af2b2cfd508fb39462164e92b6244097c256116df25d4de50cabe5;
        pp.proof[85] = 0x2824a17dcda4264819371da14b0a828509a13bb23409c61e03f7a380273b2c48;
        pp.proof[86] = 0xaba8dc330671e63fa29b6d52f0816f500300a5028c5d457123a4dc1dfe66334;
        pp.proof[87] = 0x1b7c467230a06a10f21f2356cd3b2b594966f9255a5960ae66db8401f1c5bce4;
        pp.proof[88] = 0x2393ddbd37d3dc687fa7d6e9dd007fbbef9ecd358758008ffdb2f4510e0593b;
        pp.proof[89] = 0x2824a17dcda4264819371da14b0a828509a13bb23409c61e03f7a380273b2c48;
        pp.proof[90] = 0x01;
        pp.proof[91] = 0x58a29f5ea419af2ec4700e3363abc629fb24fff8277cca0afffb8ded8a5ecc3;
        pp.proof[92] = 0xeb66c3ad4a1983077a74b495b19102eb89a52877c14396d785c5d1a8dcc9b28;
        pp.proof[93] = 0x2ff64515b3a9dda2816ea2a23ab0b4a4b11da391bd9d6e8fbae3adb638573f49;
        pp.proof[94] = 0x3041c8c8c79b2552b9f47a0021edc87fe92e9a3a625e8d7872672c0edb8fb4f5;
        pp.proof[95] = 0x1d5a658fa87f92af09bd17d7d30da2788b630a25c989e85afe6f418ae2675801;
        pp.proof[96] = 0x2aa4572188cb3b9798c2d8c30129bc51de2f316d7f6e17770086c6d45cf5ba9f;
        pp.proof[97] = 0x1c00a86a871d99b008400936ecdbb74bebc48c650c06db19177bd22fac15ebe4;
        pp.proof[98] = 0x244f12dbaf072876d86a13537bc8db1f78e052ad500ae12b777f00ff5f5cb226;
        pp.proof[99] = 0x2b7de4f091740334d21c771d8b9485b7df1cfefa533848be6744c7dd92207af8;
        pp.proof[100] = 0x21b63cec3ad6de8c42186c9075153d0b5f6c723138f283efcfe968354b5705e7;
        pp.proof[101] = 0xd54e9b628d5950b0c897196cdbf615342b3419d2306aba3eb5128bd1ca8163e;
        pp.proof[102] = 0x1450f875378f7ace8cb74df94dfe34dae0b2785f45644d8a5f640ddfca30bb5f;
        pp.proof[103] = 0x214d3d9e0aed1054942380bc0739e5b179f65c962f87f28997f9c75cce4c82b6;
        pp.proof[104] = 0x21edada2e293696734adceb685df0d2393a0df06e62a9b4c51f44ca35fbb0777;
        pp.proof[105] = 0x239f374c3ac8d315f95199e94274ccff9f3504b467c9c7c41a8b3bb01e48c793;
        pp.proof[106] = 0xac10444ab8587014ed068972c3c2627573b9983416e7bdb38e2a2f849e80c47;
        pp.proof[107] = 0x159a32eb5278c1aeba6b1731b35ae5575c28bd7d84e528e64c1c52d66583afc7;
        pp.proof[108] = 0x17b9ae519384adda9b832b40ed1cf090db2e54a1d19152a5fb7983de4468d61a;
        pp.proof[109] = 0x6036318b6e4d03eb5567cf831778ccaa8561f5ed581f1f88aa5cd6a71a20ae8;
        pp.proof[110] = 0x24e2c9118171d1b74cf9c16e118337e05791e834013aa73f36e41d096942220a;
        pp.proof[111] = 0x1e9bce565f043165d71592b85c5b54f956bdb3cb7ae7b67eeb59027281256f30;
        pp.proof[112] = 0x1266e1b3c13f0fa0f339aa4094a4d10dcc1da1b8f1c83302fb2963e8718f7249;
        pp.proof[113] = 0xfae30b97194b280bf7c20b95cf1cdb4e10ac38a19746b787aa8baaa2a57f9f6;
        pp.proof[114] = 0x4867310a19065c490390c5ea909f17adb32c8d1692794ad1cb52cd9785eae55;
        pp.proof[115] = 0x26f9b926ff7e07c541edf80f42c77a6a0b2abeb64f2a1a20056794479fe6e25c;
        pp.proof[116] = 0xa94e19a3ad0b53bd9eee4c2225a815e984a91402d9f87c6e786af92aff5df81;
        pp.proof[117] = 0x4047890b029b64b2c9c41cd1b1c793bb1f2315aae04315a59d87ff3b081572f;
        pp.proof[118] = 0xb2a796ced1acfe2875e9c155bcc94bf4a1e33bb4d8ee6a909713d75439af5fc;
        pp.proof[119] = 0x1a341ad1d377f265670f5d986d6661870756020c1bc0cf6ce4083122789dbef9;
        pp.proof[120] = 0x52820af63e31976facd81c2e7d3aa6aa2296b32706a56339ee252e340d11e4a;
        pp.proof[121] = 0x10b5aec02eb43c2aef717bfe940c61fe915dae00365637dd7e8736e7cd663b2;
        pp.proof[122] = 0x138d18a208783cd235a58a707abba2413b23105513eb6344b9980b9681f2480e;
        pp.proof[123] = 0x71bc38ffbc24794997bd042c921ffe5eca71110768abf5bba58647ff1877e71;
        pp.proof[124] = 0x1231d82e18b39a015a95c1c98e249967baec8f615afa373c236aa3f0f13236f4;
        pp.proof[125] = 0x2decd4cfd999f63f4e5ffc976f7e7aeec79bf88656e99d7d7408b4ee9bac02aa;
        pp.proof[126] = 0x31e245cd85ad08fa24b433eb10a54c302c8801ab00ee60c0c763667b2711d83;
        pp.proof[127] = 0x1dbaf4582b2685deaf2ab6a41bf0fce056a1e5c366ed08ff04afd6d20fe57bdb;
        pp.proof[128] = 0x789e791da888633b73fb6bb166f8b8702acf71d8c19974a805dfef9e2cebb8b;
        pp.proof[129] = 0x170f73bcf8486f3f1c1b18afb0cb0dee032de60b6ce922325d4ca4619861f2c0;
        pp.proof[130] = 0x2b2e754b2f5f3396b4cdfa0e356b2f4727a38356d4ab3a5030ed7821a7118911;
        pp.proof[131] = 0x12e496af7b1da9189899b6260f7b1fd1cf1639c93236b5937940ab0a63470ac6;
        pp.proof[132] = 0x542882d1f1aa48c397e432574046fd8f71239600e1369060f9ed4f1e405d62c;
        pp.proof[133] = 0xbb8fc37ba043574be43d1a792971851fb3b7b090b8fbc187b0120532a9f924;
        pp.proof[134] = 0x2bb00f4f8aae0cc1bf505e326694973b7e974e622faa795ef81f3b54d87fb2c;
        pp.proof[135] = 0x14e410c1010a122adf12d14f73d31e114c1af547d9c4c0a4ed2daca27dac5f5;
        pp.proof[136] = 0xdc692285f01b04f477a9f3ed2215e68e3b6366052ab70e9d4358b163c7d8c3d;
        pp.proof[137] = 0x2bd0c997b6f2d2d47af893087ec47ef32e5463e95ad0e07611536d9faac3fd5;
        pp.proof[138] = 0x47b192e36f9d15b1859da50290ea7c174d60398019add0f1237532557f7d299;
        pp.proof[139] = 0x668da094b2b01d315f0084d0af50c32a845e839a3d2451877cf0d1f4937b903;
        pp.proof[140] = 0xa2651f4e0fc5cba9c5a2c999955a1267315d9ffb020dca10c893e8f126ece6a;
        pp.proof[141] = 0x26f7cca832ea44aec2d08c55508be2ba4c0d847ddf6e7c5b966d3c6e26e84b4b;
        pp.proof[142] = 0xaf5dc17a95a4650e1d81e9d6279ba0725f2ff5641efb6a391751869580ffea9;
        pp.proof[143] = 0x2bdb8a6dc3b9831144da43f6d91e2b9c856d8d0e5699227aa02028c968f85f6d;
        pp.proof[144] = 0x20f67644bf008c664445b1651084e76b53129a2162e4d8e62b05e8493bc9799c;

        pp.pair = new uint256[](4);
        pp.pair[0] = 10255283654756376057447115068073595544895286591765668024488430230775279850409;
        pp.pair[1] = 2259542749915571388087660230495832020457703829347077426244731916748197175933;
        pp.pair[2] = 10943686387138024022331955772248598513456581885161670804733170571820524450237;
        pp.pair[3] = 14070295494573564542741915732770469685462860797646649187818885174019520818087;

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
        ] = hex"002c93216458be4d48afd625db7b1f59df392099a3fbb908f1a665e963a5352e93093f4c7d37c8ed44d34d346f08a39995c238ebcd98fddcc5541d55093271c159";
        proof[
            1
        ] = hex"002a543be0e6a209964727ae31f1a57521dedb2a549fb59e5a190b42c5becbb2a82cf819ee96fa5bf76e6ce7712caf6673df7fc3af269cfe193bed443eeb89527a";
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
        ] = hex"012de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c304040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc3524f53397bd92b66fda812b6e1191a00b60fc8e304033518006cbeedcab7f2127202de4ca10cb48fa7ae483633127295fecab2f03da9355f4ca12ca0c820096f9c3";
        proof[
            12
        ] = hex"5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449";

        return (account, proof);
    }
}
