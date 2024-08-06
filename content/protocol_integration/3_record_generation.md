# 3.3 Record Generation

A Record ($R$) is a data structure that captures the state transition of a protocol instance and
can be represented as a tuple:

$$R = (Scope, e, h, h')$$

Where:

- $Scope$ is the unique identifier of the protocol instance
- $e$ is the reference to the state transition event (i.e. the block number, transaction hash, log index)
- $h$ is the pre-state hash
- $h'$ is the post-state hash
- $h \not = h'$

Records exsists as serialized binary-object.
The construction of the Recrod object is performed by the Record Generator, a component of the ASP system.

The Record Generator performs the following steps:

1. Compute the $Scope$ using the protocol's $Scope$ function if not allready provided.
2. Compute $h$ and $h'$ using the protocol's state hash function $H$.
3. Assemble the Record $R$ object.
4. Compute the Record hash $r = Hash(R)$.

Example go-implementation of $R$:

```go

package record

// 32-byte Scope
// 32-byte Tx Hash
// uint Log Index
// 32-byte pre-state hash
// 32-byte post-state hash
type RecordT [130]byte

type Record interface {
	Hash() [32]byte
	Scope() [32]byte
	TxHash() [32]byte
	LogIdx() uint
	PreState() [32]byte
	PostState() [32]byte
}

var _ Record = (*RecordT)(nil)

func (r RecordT) Hash() [32]byte {
	return HashAlgorithm(r)
}

func (r RecordT) Scope() (scope [32]byte) {
	copy(scope[:], r[:32])
	return
}

func (r RecordT) TxHash() (txHash [32]byte) {
	copy(txHash[:], r[32:64])
	return
}

func (r RecordT) LogIndex() (logIndex uint) {
	return uint(r[64])
}

func (r RecordT) PreState() (preStateHash [32]byte) {
	copy(preStateHash[:], r[65:97])
	return
}

func (r RecordT) PostState() (postStateHash [32]byte) {
	copy(postStateHash[:], r[97:129])
	return
}

```
