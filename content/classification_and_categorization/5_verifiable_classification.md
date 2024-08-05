## 5.4 Verifiable Classification

> To ensure trustless operation of the ASP system and mitigate the risk of a malicious ASP,
> we introduce the concept of `Verifiable Classification`

### Key components of Verifiable Classification:

1. **Zero-Knowledge Proof Generation**:

   The ASP generates a ZKP proving:

   - Correct application of classification rules to extracted features
   - Relevance of extracted features to the specific record
   - Correct utility of feature extractor
   - Verification of any data sourced externally that were used in feature extraction (i.e proof of payload origin)

   ```admonish info Title="ZKP Implementation"
   **The specific implementation of the ZKP system is WIP**.

   Here are the current thoughts on the implementation:
   - Encoding of all possible Feature Types and Names and associated feature-extractors Values a
    with a 254bit-map (Feature Space), i.e:
      - 4 bits to encode Feature Types, i.e.:
        - CATEGORICAL: 0001
        - BOOLEAN: 0010
        - INT: 0011
        - FLOAT: 0100
        - STRING: 0101
        - BYTES: 0110
        - TIMESTAMP: 0111
      - 8 bits to encode identifiers of feature-extractors, i.e.:
        - AML: 00000001
        - SOURCE_OF_ASSET: 00000010
      - 64 bits to encode Feature Name constants:
        - ILLICT_EXPOSURE = 0x0000000000000001
      - 180 bits to encode unsigned Feature Values.
   - Category Space (all possible categories) are also encoded with a bitmap.
   - Classification rules can be expressed as a circuit of logical gates mapping Feature Space to Category Space.
      - Correctness of logical operations proven with Zk (i.e. Circom circuits).
      - Modular design allows for easy addition of new rules.
      - Zk Circuit operate on a private input of encoded Feature.
   - Feature Documents can be represented as a Merkle Tree (Feature Tree) using the Feature Space encoding.
      - As suggested in section [4.2 Feature Types and Formats](feature_extraction/2_feature_types_and_formats.md)
      - Feature Extractors can generate a signature over the Merkle Root of the Feature Tree.
    - Classificaation rules can be expressed as logical expressions in a zk-circuit (i.e. Circom circuit)
      - Assume at least 1 circuit per classifier.
      - Receive

    - Classifiers can generate:
      - Merkle-proofs of features against the Feature Tree used in classificatio (Feature Proof).
      - Merkle-proof of the category against the Category Tree (Category Proof).
    - Categorizers
   ```

2. **On-chain Attestations**:

Other ASPs or external parties can attest to the validity or correctness of an ASP record categorization
on-chain through attestation channels such as [EAS](https://attest.org/).

3. **Proof Verification**:

External parties can verify the ZKP without accessing private features or extractor code.

```

```
