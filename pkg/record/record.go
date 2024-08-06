package record

// 32-byte Scope
// 32-byte Tx Hash
// 32-bits Log Index
// 32-byte pre-state hash
// 32-byte post-state hash
type RecordT [130]byte

type Record interface {
	Scope() [32]byte
	TxHash() [32]byte
	LogIndex() uint
	PreStateHash() [32]byte
	PostStateHash() [32]byte
}

var _ Record = (*RecordT)(nil)

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

func (r RecordT) PreStateHash() (preStateHash [32]byte) {
	copy(preStateHash[:], r[65:97])
	return
}

func (r RecordT) PostStateHash() (postStateHash [32]byte) {
	copy(postStateHash[:], r[97:129])
	return
}
