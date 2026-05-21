package model

import (
	"broadcast-server/internal/util"
	"fmt"
)

type Message struct {
	// msgID = clientID:msgUID-msgIndex
	//clientID = client.RemoteAddr() + ":" + instanceUID
	ID      string `json:"id"`
	Content []byte `json:"text"`
}

func NewMessage(clientID string, msgIndex int, content []byte) *Message {
	msgUID := util.GenerateUID()
	id := fmt.Sprintf("%s:%s-%d", clientID, msgUID, msgIndex)
	return &Message{
		ID:      id,
		Content: content,
	}
}