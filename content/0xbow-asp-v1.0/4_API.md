# REST / gRPC API

```admonish warning title='Work-in-Progress'

0xBow ASP API v1.0 is still in development and not production-ready.
Here is the list of features to be available at launch:

  - [ ] REST API
    - [ ] Endpoints for querying records
    - [ ] Endpoints for generating zk-proofs
    - [ ] Endpoints for service status & health checks.
  - [ ] gRPC API
    - [ ] Support for synchronous & asynchronous queries
    - [ ] Data streaming
    - [ ] Support for private channels via Waku [Waku](https://waku.org/)
  - [ ] Webhooks
    - [ ] Support for Event push-notifications & streams


⚠️ REST API Endpoints built for older revisions of Privacy Pool will be deprecated soon. ⚠️

```

---

## Overview:

<!-- toc -->

---

## 0xBow ASP REST API v1.0:

**Base URL**: `https://api.0xbow.io/api/{version}`

0xBow ASP Rest API v1.0 provides a set of API endpoints to for
querying records, generating `association-sets`, computing proofs and querying service status.

> ⚠️ These endpoints are not privacy-preserving as they are provided for convenience ⚠️

### 🟢 POST /api/v1/`{set}`

`Depreciated`

#### Description

> **Context:**
>
> Prior revisions of the ASP used binary classification to categorize records.
>
> Initial version maintained 2 large sets of records to reflect this classification:
>
> - Inclusion Set: Record Hashes of records that passed compliance checks.
> - Exclusion Set: Record Hashes of records that failed compliance checks.
>
> These sets were represented as a merkle-tree.
>
> Any inserts or removal of record hashes would result in onchain emission of the
> new merkle root.
>
> Later versions were optimised for onchain storage of sets
> to support onchain queries:
>
> - Rather than 1 large set, the sets were split into smaller sub-sets.
> - `mtID` is the unique Identifier were associated with each Subsets
>   - which is a Hash of the tuple (chainID, contract address, set type)

This API Endpoint generates a new `association-set` based on the provided `hashSet` and `hashFilter`.
It was tailored for `Privacy Pool` to support it's `proof-of-innocence` mechanism.

---

#### Path Parameters

| name  | type     | description                     |
| ----- | -------- | ------------------------------- |
| `set` | required | the target set to query against |

**Possible values for {setType}**

1. `inclusion`: query the inclusion set
2. `exclusion`: query the exclusion set

---

#### Query Parameters

| name          | type | data type | example       | description                                         |
| ------------- | ---- | --------- | ------------- | --------------------------------------------------- |
| `chain`       |      | string    | "sepolia"     | name of the chain where contract is deployed to     |
| `contract`    |      | string    | "0x8e3E..."   | privacy pool contract address                       |
| `mt_id`       |      | string    | "0x1e1294..." | unique identifier of a set                          |
| `hash_only`   |      | boolean   | false         | only return set of record hashes                    |
| `size_limit`  |      | integer   | 20            | limits the size of the returned set to `size_limit` |
| `pin_to_ipfs` |      | string    | false         | flag for pinning association set to ipfs            |
| `random`      |      | boolean   | true          | flag for randomising the record selection           |
| `needSort`    |      | boolean   | true          | flag for sorting set by record Index                |

---

#### Body Parameters

| name         | type | data type    | description          |
| ------------ | ---- | ------------ | -------------------- |
| `hashSet`    |      | string array | set of record hashes |
| `hashFilter` |      | string       | type of filtering    |

**Example for hashSet**:

```json
{
  "hashSet": [
    "113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
    "1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
    "18ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
    "2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e"
  ]
}
```

**Possible values for hashFilter:**:

1. `EXCEPT`: Exclude the records in `hashSet` from the response.

```json
{
  "hashFilter": ["EXCEPT"]
}
```

2. `INTERSECT`:

Return only set of records that are both members of the `hashSet` and `{set}`.

```json
{
  "hashFilter": ["INTERSECT"]
}
```

3. `UNION`:

Return the union of the set of records from the response only for records that are members of the `{set}`.

```json
{
  "hashFilter": ["UNION"]
}
```

---

#### Example cURL

```bash

API="api.0xbow.io"
ENDPOINT="/api/v1/inclusion"
CHAIN="sepolia"
CONTRACT="0x8e3E4702B4ec7400ef15fba30B3e4bfdc72aBC3B"
HASH_ONLY="false"
SIZE_LIMIT="20"
PIN_TO_IPFS="false"

URI="${API}${ENDPOINT}?"
URI+="chain=${CHAIN}&"
URI+="contract=${CONTRACT}&"
URI+="hash_only=${HASH_ONLY}&"
URI+="size_limit=${SIZE_LIMIT}&"
URI+="pin_to_ipfs=${PIN_TO_IPFS}"

curl --location --request POST $URI \
--header "Content-Type: application/json" \
--data "{
    \"hashSet\": [],
    \"hashFilter\": \"\"
}"
```

#### Responses

| http code | content-type                      | response    |
| --------- | --------------------------------- | ----------- |
| `200`     | `application/json; charset=utf-8` | JSON Object |

#### Example Response

```json
{
    "uuid": "",
    "mtID": "1e1294aedb5c4bc78479c7cd09c163808d894bb37e61eadd73cdc8cedc85bf9f",
    "zero": "2fe54c60d3acabf3343a35b6eba15db4821b340f76e741e2249685ed4899af6c",
    "merkleRoot": "002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
    "hashSet": [
        "113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
        "1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
        "18ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
        "2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e",
        "10d6373c1464696f856fbfee98132e28166f0227a6e40ab501d5468ae73f1c22",
        "1658ef12bff2c2a6cd37f09e6f0686fba9514b8e17594752f898009f83cd6cfb",
        "2b06d56c6d1812babd87d3cd0127a8f4d92a56130bd57f11aacd51d8a4e634c3",
        "18e44125cbb1fe0d81d0c1694bde77ba35a2cb04dc1ee4d993809d919080da22",
        "03512b924c8c0d98a9ad40a1b9b934f83139adfc281fa120b755578a73457b63"
    ],
    "proofs": [
        {
            "record_hash": "113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
            "record_data": {
                "txHash": "0xaa2243999994946b104ecdcc41e8b392043d9478347fa11782ed6ae411021ae5",
                "outputCommitment1": "0x00b4a16ff4129dcdcd100bc1cad317980302f243d6ca184480455876d50eff5a",
                "outputCommitment2": "0x294f8fbc010ab687a719c5849420a49cec93bc831122684491f2527cd2011eeb",
                "nullifierHash1": "0x070cf43476880e27f1728a1f2446a57317a6892ef9af99a0bfa93f8a4792e341",
                "nullifierHash2": "0x0c3bbbce67df72abb37f6a0f603182c8984392f27bdda38a30452c985015562d",
                "recordIndex": 14,
                "publicAmount": "0000000000000000000000000000000000000000000000000de0b6b3a7640000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
                "leaf_index": 0,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    0,
                    0,
                    0,
                    0
                ],
                "path_positions": [
                    1,
                    1,
                    1,
                    1
                ],
                "path_elements": [
                    "0x1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
                    "0x2f7b5ca0810afc0422b315ebae2df141e67ed8487e864cb903e4590a7bd34403",
                    "0x12a24534a43a7a6f51a9beaa33b3676766f83ab6db907b590932a86f91ea0307",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
            "record_data": {
                "txHash": "0x9ac2b822f4147e4d915a846268cef946f988e67dd5da964049c30cf5bccb055c",
                "outputCommitment1": "0x0d45924f17aa19a6d5de9bf8c3ffcec906ae89f15989b065f93a44dc42fb7897",
                "outputCommitment2": "0x1e786708230d87cd775b2efd4f0822543a18b6c12d59eb2ab50bc8bd3b4d88aa",
                "nullifierHash1": "0x1e0e798d049cc18291c4212a90a0be8e795d1ae323c94f37613371eb1b3526e9",
                "nullifierHash2": "0x0b7b28d575321eef0b1b559ecc8826c5e0ccfa9b838b672784ca2371edfbc61c",
                "recordIndex": 22,
                "publicAmount": "00000000000000000000000000000000000000000000000000038d7ea4c68000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
                "leaf_index": 1,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    1,
                    0,
                    0,
                    0
                ],
                "path_positions": [
                    0,
                    1,
                    1,
                    1
                ],
                "path_elements": [
                    "0x113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
                    "0x2f7b5ca0810afc0422b315ebae2df141e67ed8487e864cb903e4590a7bd34403",
                    "0x12a24534a43a7a6f51a9beaa33b3676766f83ab6db907b590932a86f91ea0307",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "18ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
            "record_data": {
                "txHash": "0x8646ba48685dda1fd4b771448276f6b6812131baeb4ff8413999f00a59fc60e9",
                "outputCommitment1": "0x0a6fd9b9f65f4173feb5fb6745a2321700aac5c9039b39f8f98588e108756664",
                "outputCommitment2": "0x0453b7ff55700c143c99d44d6a2262fd89281c5eb2e0c7d057c5bdd4d8c8b00d",
                "nullifierHash1": "0x1adbd6e04911395701ed60358784b01a1188ccacf8e93e31bc15412893218ffa",
                "nullifierHash2": "0x2e7af471ba95a488bd6e8ad35939c6d292f8cee54ec349db14019c604c378cde",
                "recordIndex": 34,
                "publicAmount": "0000000000000000000000000000000000000000000000000de0b6b3a7640000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x18ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
                "leaf_index": 2,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    0,
                    1,
                    0,
                    0
                ],
                "path_positions": [
                    3,
                    0,
                    1,
                    1
                ],
                "path_elements": [
                    "0x2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e",
                    "0x0be6c215dddf4f423e478127405b6d33412378e10b191a6f093183dd45d7680b",
                    "0x12a24534a43a7a6f51a9beaa33b3676766f83ab6db907b590932a86f91ea0307",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e",
            "record_data": {
                "txHash": "0x8ac9b6bfc96bf159dca7c46328bd8121a7f70692c49cbe9c13df2292f1427c98",
                "outputCommitment1": "0x2d4e39c4f62f1c029f22e94a0b54a57a27e5d2919392f707bf31572b2df9576d",
                "outputCommitment2": "0x0b3833c986b86e6d4a16eae2bf14f5b9e34bc0911bd769f81e8604568eb939ee",
                "nullifierHash1": "0x178189eb340940570b5f3a74459262b1ff898057391fdd2ad97d012087baa14a",
                "nullifierHash2": "0x17fccb01a5ad7cfb5e4a2e7295dbf149de863b4874346974d1de15eff8f3dbde",
                "recordIndex": 36,
                "publicAmount": "0000000000000000000000000000000000000000000000000de0b6b3a7640000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e",
                "leaf_index": 3,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    1,
                    1,
                    0,
                    0
                ],
                "path_positions": [
                    2,
                    0,
                    1,
                    1
                ],
                "path_elements": [
                    "0x18ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
                    "0x0be6c215dddf4f423e478127405b6d33412378e10b191a6f093183dd45d7680b",
                    "0x12a24534a43a7a6f51a9beaa33b3676766f83ab6db907b590932a86f91ea0307",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "10d6373c1464696f856fbfee98132e28166f0227a6e40ab501d5468ae73f1c22",
            "record_data": {
                "txHash": "0xb6d72c74eb00a15bee2af1814cc1084b861d4538f8a4ef5650c248f3e418ca44",
                "outputCommitment1": "0x276c2fecf765137a15625f8696b25d1b39e402b8ed43950893962f72ca22c0fb",
                "outputCommitment2": "0x210d1818699a883053297c4ac920ab883c30472b821679255f2a6032e2a26316",
                "nullifierHash1": "0x015fa37c5f43504ba940b60361437b3be94830e6f4bee5a359381d2bf8e1e2bd",
                "nullifierHash2": "0x1f369fab9b7db862645106e6a3a07c7df3a5b2ce1b256aad5ee599410c14dff3",
                "recordIndex": 42,
                "publicAmount": "000000000000000000000000000000000000000000000000016345785d8a0000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x10d6373c1464696f856fbfee98132e28166f0227a6e40ab501d5468ae73f1c22",
                "leaf_index": 4,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    0,
                    0,
                    1,
                    0
                ],
                "path_positions": [
                    5,
                    3,
                    0,
                    1
                ],
                "path_elements": [
                    "0x1658ef12bff2c2a6cd37f09e6f0686fba9514b8e17594752f898009f83cd6cfb",
                    "0x2bb5136f5053629d470a7df2ea75ea49885714c98802c4b16fd42fd4359a2166",
                    "0x2357f678f06b3729cd17d0232af0cb5597aeb9695690b93f4ed613772712bb72",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "1658ef12bff2c2a6cd37f09e6f0686fba9514b8e17594752f898009f83cd6cfb",
            "record_data": {
                "txHash": "0xfff8cddf0a21328713a3d81c6b8c6b33bc80a45e21ee79a40720434bd25bf164",
                "outputCommitment1": "0x13137f57d077844c7f951d78120ba1f7925dfd30b5f1c8a20d34a2bb76ef18ce",
                "outputCommitment2": "0x1a3281e4d22ef165d86f67d956a922394ee5f587fb7be397d6a030e1d4f44c5b",
                "nullifierHash1": "0x20d7d4d025426a2f9d42fb626caf3a46217bb83abcb9b2278dc492e3a9badb0d",
                "nullifierHash2": "0x12841a879639231880052240d425e06092c948f9b7991ae6a501077908624404",
                "recordIndex": 44,
                "publicAmount": "00000000000000000000000000000000000000000000000000038d7ea4c68000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x1658ef12bff2c2a6cd37f09e6f0686fba9514b8e17594752f898009f83cd6cfb",
                "leaf_index": 5,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    1,
                    0,
                    1,
                    0
                ],
                "path_positions": [
                    4,
                    3,
                    0,
                    1
                ],
                "path_elements": [
                    "0x10d6373c1464696f856fbfee98132e28166f0227a6e40ab501d5468ae73f1c22",
                    "0x2bb5136f5053629d470a7df2ea75ea49885714c98802c4b16fd42fd4359a2166",
                    "0x2357f678f06b3729cd17d0232af0cb5597aeb9695690b93f4ed613772712bb72",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "2b06d56c6d1812babd87d3cd0127a8f4d92a56130bd57f11aacd51d8a4e634c3",
            "record_data": {
                "txHash": "0x09f5d7c15f730477e75089671c26c74edcf2d3c13c030ae7f000d20689feb920",
                "outputCommitment1": "0x0d85948b189c8f06ba0e0ddc10465e7b38346fff32a5e7653e1b127bcb61bad1",
                "outputCommitment2": "0x282eed96a8315d3fef4a8b44a481ea321164d419e155d0f5cc670b1cbd8d922c",
                "nullifierHash1": "0x1b561048777361b051d2bfc6ea865c7b6227a7c74e5dcd22a52e41932727d2f9",
                "nullifierHash2": "0x230520e4c6f2b20f7c60c385f08fc55bcd16f94b0ab24be01c8cb9add6675984",
                "recordIndex": 46,
                "publicAmount": "00000000000000000000000000000000000000000000000000038d7ea4c68000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x2b06d56c6d1812babd87d3cd0127a8f4d92a56130bd57f11aacd51d8a4e634c3",
                "leaf_index": 6,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    0,
                    1,
                    1,
                    0
                ],
                "path_positions": [
                    7,
                    2,
                    0,
                    1
                ],
                "path_elements": [
                    "0x18e44125cbb1fe0d81d0c1694bde77ba35a2cb04dc1ee4d993809d919080da22",
                    "0x2fc47895108f3de39eea196f7d2bcc12a7253943e9f359c9bea03a76fed0f03e",
                    "0x2357f678f06b3729cd17d0232af0cb5597aeb9695690b93f4ed613772712bb72",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "18e44125cbb1fe0d81d0c1694bde77ba35a2cb04dc1ee4d993809d919080da22",
            "record_data": {
                "txHash": "0x11258c721b6a7b3dd2dbd63423b0ee4a6d410f5161b0b35372b95c328a2f1d54",
                "outputCommitment1": "0x0d610b983dcbb8ee71abe7718c42e1c60153d58533dd7e3fbc4f5bf070e389eb",
                "outputCommitment2": "0x25722dfcb5efc4e537ffe27163a5e4c35642d96c8e925e4c8c9a1ed3a1ed2ece",
                "nullifierHash1": "0x15d9acfd0d5e541ada7b4132c2df7d2602dcda7708e086fa61b9f13b0c4b4054",
                "nullifierHash2": "0x0961572c8fe8b211d84bfe62479405dd87a0e67611ce9511bddb0b065025e0be",
                "recordIndex": 102,
                "publicAmount": "00000000000000000000000000000000000000000000000000b1a2bc2ec50000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x18e44125cbb1fe0d81d0c1694bde77ba35a2cb04dc1ee4d993809d919080da22",
                "leaf_index": 7,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    1,
                    1,
                    1,
                    0
                ],
                "path_positions": [
                    6,
                    2,
                    0,
                    1
                ],
                "path_elements": [
                    "0x2b06d56c6d1812babd87d3cd0127a8f4d92a56130bd57f11aacd51d8a4e634c3",
                    "0x2fc47895108f3de39eea196f7d2bcc12a7253943e9f359c9bea03a76fed0f03e",
                    "0x2357f678f06b3729cd17d0232af0cb5597aeb9695690b93f4ed613772712bb72",
                    "0x073f04e5838a95e2a635da7cbbf87b60cd974b5cad98b5638fd96e71cc5eb130"
                ]
            }
        },
        {
            "record_hash": "03512b924c8c0d98a9ad40a1b9b934f83139adfc281fa120b755578a73457b63",
            "record_data": {
                "txHash": "0xe5a319daa4ee50aa447c9a8ea0ac560d0d637ec4cac030e8919016d905f1071f",
                "outputCommitment1": "0x1240adadbb08ec7e69f0751b164b56e21521f78c1b0eb499c96d92caf47442b4",
                "outputCommitment2": "0x1da2712b9feb81f3da5ce2e360e4bb8d77346e6d33ab2f16ac3dc8fd5a318e0b",
                "nullifierHash1": "0x19e34548f6a584dab328f1c3fc2e9277a1653c751734a316d39d7e6f53175c99",
                "nullifierHash2": "0x05dc56e8cb7458085d6c221a208d8afb406a37c1680d92949a10167edcf2bb87",
                "recordIndex": 152,
                "publicAmount": "000000000000000000000000000000000000000000000000016345785d8a0000"
            },
            "merkle_proof": {
                "merkle_tree_max_depth": 4,
                "leaf": "0x03512b924c8c0d98a9ad40a1b9b934f83139adfc281fa120b755578a73457b63",
                "leaf_index": 8,
                "path_root": "0x002915b4928a5b34454158b06c50777f555f307b7fcace62f666e1586ee899b1",
                "path_indices": [
                    0,
                    0,
                    0,
                    1
                ],
                "path_positions": [
                    0,
                    0,
                    0,
                    0
                ],
                "path_elements": [
                    "0x2fe54c60d3acabf3343a35b6eba15db4821b340f76e741e2249685ed4899af6c",
                    "0x1a332ca2cd2436bdc6796e6e4244ebf6f7e359868b7252e55342f766e4088082",
                    "0x2fb19ac27499bdf9d7d3b387eff42b6d12bffbc6206e81d0ef0b0d6b24520ebd",
                    "0x2706ceb05a41606d32b2995e5586beecfb8dad66a6dc4f2e68f0b5a8e01ecf29"
                ]
            }
        }
    ],
    "ipfsHash": "",
    "txHash": "",
    "status": "SUCCESS",
    "timestamp": 1723181615
}
```

---

### 🟡 POST /api/v1/records/filter

> `Endpoint is a Work-In-Progress`

#### Description

> **Context:**
>
> Latest revision of the ASP implements the categorization process per ASP specification v1.0,
> where Records are classified with multiple categories mapped to 252 bits (`category bitmap`).
>
> Recrod to Category Bitmap mapping is stored in the onchain Registry.

This endpoint provides a way to filter a set of record hashes

---

#### Body Parameters

| name       | type     | data type    | description                             |
| ---------- | -------- | ------------ | --------------------------------------- |
| `scope`    | required | string       | unique identifier for protocol          |
| `subSet`   |          | string array | set of record hashes                    |
| `filter`   |          | string       | hex-encoded bitmap filter               |
| `type`     |          | enum         | defines the type of predicate to apply  |
| `complete` |          | boolean      | flag for including complete record data |

**Example Body**

```json
{
  "scope": "0xd234x67851b11a21",
  "subset": [
    "0x113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
    "0x1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
    "0x8ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
    "0x2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e"
  ],
  "filter": "0x1234567891011121",
  "type": 1,
  "complete": true
}
```

**Example for hashSet**:

```json
{
  "hashSet": [
    "113143e9dae0aa58d13b26dec085606d28fafe70582ec52fd5bbc08ae8d5b5c9",
    "1aa21d201f72b61e0e59bdd7a0ef62dced57e4e80fa180ff113a58dc3aeb8ea9",
    "18ba306635d7838c1378a9243c22487f906ec929a5a8d5c30f172a9bc5824d64",
    "2dca7e37ec7e31d0e56b456e6ed435ced4c506b6dada186f6a14907ecc50a37e"
  ]
}
```

**Possible values for hashFilter:**:

1. `EXCEPT`: Exclude the records in `hashSet` from the response.

```json
{
  "hashFilter": ["EXCEPT"]
}
```

2. `INTERSECT`:

Return only set of records that are both members of the `hashSet` and `{set}`.

```json
{
  "hashFilter": ["INTERSECT"]
}
```

3. `UNION`:

Return the union of the set of records from the response only for records that are members of the `{set}`.

```json
{
  "hashFilter": ["UNION"]
}
```

#### Responses

| http code | content-type                      | response    |
| --------- | --------------------------------- | ----------- |
| `200`     | `application/json; charset=utf-8` | JSON Object |

#### Example cURL

```bash

API="api.0xbow.io"
ENDPOINT="/api/v1/inclusion"
CHAIN="sepolia"
CONTRACT="0x8e3E4702B4ec7400ef15fba30B3e4bfdc72aBC3B"
HASH_ONLY="false"
SIZE_LIMIT="20"
PIN_TO_IPFS="false"

URI="${API}${ENDPOINT}?"
URI+="chain=${CHAIN}&"
URI+="contract=${CONTRACT}&"
URI+="hash_only=${HASH_ONLY}&"
URI+="size_limit=${SIZE_LIMIT}&"
URI+="pin_to_ipfs=${PIN_TO_IPFS}"

curl --location --request POST $URI \
--header "Content-Type: application/json" \
--data "{
    \"hashSet\": [],
    \"hashFilter\": \"\"
}"
```
