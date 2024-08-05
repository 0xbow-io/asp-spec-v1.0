# 3.3 Record Generation

For each valid state transition, a Record $R$ is generated with the following structure:

$$R = (Scope, e, h, h')$$

Where:

- $Scope$ is the unique identifier of the protocol instance
- $e$ is the reference to the state transition event (i.e. the block number, transaction hash, log index)
- $h$ is the pre-state hash
- $h'$ is the post-state hash
- $h \not = h'$

The Record Generator performs the following steps:

1. Compute the $Scope$ using the protocol's $Scope$ function if not allready provided.
2. Compute $h$ and $h'$ using the protocol's state hash function $H$.
4. Assemble the Record $R$.
5. Compute the Record hash $r = Hash(R)$.

$R$ along with its hash $r$ is passed to the Category Engine for further processing.
