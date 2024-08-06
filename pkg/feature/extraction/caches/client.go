package caches

type CacheCl interface {
	Connect(string) error
	Read(string, []byte) ([32]byte, error)
	Write(string, [32]byte) error
}
