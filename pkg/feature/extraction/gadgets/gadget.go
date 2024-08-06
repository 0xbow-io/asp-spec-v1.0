package gadget

type GadgetCl interface {
	Connect(string) error
	process(string, [32]byte) interface{}
}
