package server

func Shutdown() {
	// Stop accepting new connections
	// Close existing connections gracefully
	// Perform any necessary cleanup
	cache.Close()
}