# 2.4 Security Model

```admonish important
 A comprehensive security analysis and formal proofs are beyond the scope of this overview and are addressed in subsequent sections.
```

```admonish abstract  title='Principles'
The ASP security model is based on the following principles:

1. **Immutability**: The Public Registry is append-only and immutable, ensuring the integrity of stored records.

2. **Privacy**: Zero-knowledge proofs enable compliance verification without revealing sensitive transaction details.

3. **Decentralization**: The system operates across multiple protocols and doesn't rely on a single point of trust.

4. **Access Control**: Strict write policies prevent unauthorized modifications to the Public Registry.

5. **Cryptographic Integrity**: All records and proofs are cryptographically secured.
```

**The security of the ASP system relies on the following assumptions:**

1. The underlying blockchain protocol's security (e.g., Ethereum's consensus mechanism).
2. The cryptographic security of the hash functions used (e.g., Keccak-256, Poseidon).
3. The soundness and zero-knowledge properties of the ZKP system employed.

**Key security considerations include:**

1. **Sybil Resistance**: The system must be resistant to Sybil attacks on the Watcher.

2. **Front-Running Protection**: Measures to prevent front-running of compliance proofs in the on-chain components.

3. **Privacy Leakage**: Careful design & implementation of Interfaces & communication channels to prevent
  inadvertent privacy leaks through query patterns or request origin traces.

4. **Upgrade Security**: Secure processes for updating classification rules and compliance policies.
