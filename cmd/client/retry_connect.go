package client

import (
	"log"
	"strings"
)

func RetryConnect(addr string) {
	log.Printf("Attempting to reconnect to %s...", addr)
	hostUrl := GetRemoteAddrParts(addr)
	CreateSocketConnection(hostUrl, "8080")
}

func GetRemoteAddrParts(addr string) string {
	// Remove "ws://" prefix if present
	if len(addr) > 5 && addr[:5] == "ws://" {
		addr = addr[5:]
	}

	// Remove port if present
	if colonIndex := strings.LastIndex(addr, ":"); colonIndex != -1 {
		addr = addr[:colonIndex]
	}

	return addr
}
