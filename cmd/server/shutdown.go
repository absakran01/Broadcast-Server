package server

func Shutdown() {
	cache.Close()
}