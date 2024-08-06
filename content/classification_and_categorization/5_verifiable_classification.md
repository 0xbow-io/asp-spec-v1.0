## 5.4 Verifiable Classification

> To ensure trustless operation of the ASP system and mitigate the risk of a malicious ASP,
> we introduce the concept of `Verifiable Classification`

### Key components of Verifiable Classification:

```admonish warning Title="Work in Progress"
**The specific implementation of the ZKP system is WIP**
```

1. **Zero-Knowledge Proof Generation**:

The classification process is to be implemented with zero-knoweldge DSL (i.e. Circom). This allows
the ASP to generate a computation proof which verifies that the classification was done correctly
without revealing the actual features or the feature extractor code.

Current thoughts on the approach:

- Auto-Generate Circom Templates based on the category-feature schema.
  - At least 1 circuit per category.
- Utilise folding schemes (i.e. NOVA) with libraries such as [Sonobe](https://github.com/privacy-scaling-explorations/sonobe)
  - The bitmap is the shared state between all steps
  - Each step is 1 classification that sets the appropriate bit in the bitmap
  - 1 public output at the final step which is the final bitmap.

2. **On-chain Attestations**:

Other ASPs or external parties can attest to the validity or correctness of an ASP record categorization
on-chain through attestation channels such as [EAS](https://attest.org/).

3. **Proof Verification**:

External parties can verify the ZKP without accessing private features or extractor code.
