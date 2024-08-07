# 3.1 Protocol Requirements

For a protocol to integrate with the ASP system, it must satisfy the following requirements:

1. **State Machine Representation**

   The protocol must be representable as a state machine with the following tuple:

   $$Protocol: StateMachine = (Scope, S, I, O, T, H, V, \Delta S)$$

   Where:

   - $Scope$: Hash function to compute unique identifier of the protocol instance
     - Implementation of a protocol in a specific chain (i.e. Contract deployed in ethereum) is an instance of a protocol
     - The Scope function computes the unique identifier of the protocol instance
     - I.e. Keccak256 Hash of (address, chainID, contractCode)
   - $S$: State space
     - The agreed upon state space of the protocol
   - $I$: Input dictionary
     - The agreed upon inputs that can trigger a state transition in the protocol
   - $O$: Output dictionary
     - The agreed upon outputs that are returned by the state transition function of the protocol
   - $T$: Transition function, $T(s, i) \rightarrow s'$
     - s = the agreed upon prestate satisfying S,
     - s' = the post state satisfying S,
     - i = the state transition inputs satisfying I
   - $H$: State hash function, $H(s, e) \rightarrow h$
     - $s$ = the state of the protocol
     - $e$ = reference to the state-transition event which can be a singular value or a tuple of values (i.e. block number, transaction number, log number)
     - $h$ = the hash of the state and must be a unique identifier of the state at $e$
   - $V$: Verification function, $V(s', s, e) \rightarrow \{0, 1\}$
     - A function that can verify that $s'$ is the pre-image of $s$ at $e$ satisfying the transition function $T$
   - $\Delta S$: State comparator function, $\Delta S(s, s') \rightarrow \delta$
     - A function that can compare two states and return the difference between them.
     - $\delta \not = 0$ indicates a state transition

2. **Deterministic State Transitions**: The protocol must have deterministic state transitions to ensure consistent record generation.

3. **Observable State**: The protocol must expose sufficient information to reconstruct its state at any given epoch.

4. **Unique Identifiers**: Each protocol instance must have a unique identifier computable by the $Scope$ function.

5. **Event Emission**: The protocol should emit events for all state-changing operations to facilitate efficient monitoring.
