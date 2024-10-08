# 2.3 Data Flow

1. **State Transition Detection**:

   - Watcher observes registered protocols for protocol interactions.
   - State changes is detected by the State Transition Detector.

2. **Record Generation and Classification**:

   - Record Generator creates a cryptographic record $R$ of the state transition.
   - Feature Extractor processes $R$ to extract relevant features.
   - Classifier categorises $R$ based on extracted features.
   - Categorizer creates a 256-bit category bitmap $B$ for $R$.

3. **On-Chain Storage**:

   - The category bitmap and the associated record hash $r$ are stored on-chain in the Public Registry.

4. **Querying**:

   - External systems can query the Public Registry to retrieve subsets of
     record hashes based on scope, category or feature criteria.

5. **Privacy-Preserving Compliance Verification**:

   - An entity requests a compliance proof for a set of records: $\Nu=\Set{r_1, ..., r_n}$.
   - ZKP Generator computes a proof $\pi$ proving $\forall r \in \Nu: P(r)$ where predicate $P$ is the
     given compliance policy expressed as a catgory bitmap.
