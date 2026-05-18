package client

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	headers = http.Header{
		"User-Agent": []string{"RedisRelayClient/1.0"},
	}
	input string
)

func Connect(host string, port string) {
	addr := "ws://" + host + ":" + port + "/ws"



		dialer := websocket.Dialer{
		HandshakeTimeout:  http.DefaultClient.Timeout,
		} 
		remoteConn, _, err := dialer.Dial(addr, headers)
		if err != nil {
			log.Printf("failed to connect to %s: %v", addr, err)
			return
		}
		defer remoteConn.Close()

		go write(remoteConn)

		recieve(remoteConn)

	
}
