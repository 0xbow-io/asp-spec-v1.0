package plugins

type PluginCl interface {
	Connect(string) error
	Process(string, [32]byte) interface{}
}

// Common interface for all plugins
type _Plugin interface {
	// Ping the plugin to check if it is available
	Ping() error
	// Connect to the plugin
	Connect() error
	// Disconnect from the plugin
	Disconnect() error
	// Get the plugin Description
	// This is used to get the plugin ID & version
	Describe() []string
	// IO: Send bytes to Plugin for processing
	// The plugin will process the data and return bytes
	IO([]byte) []byte
}

type _PluginCL struct {
	Sockets map[string]_Plugin
	DataBus chan []byte
}
