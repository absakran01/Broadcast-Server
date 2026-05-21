package model

import (
	"broadcast-server/internal/util"
	"fmt"
)

type Message struct {
	// msgID = clientID-randUID-msgIndex
	ID      string `json:"id"`
	Content []byte `json:"text"`
}

func NewMessage(clientID string, msgIndex int, content []byte) *Message {
	randUID := util.GenerateUID()
	id := fmt.Sprintf("%s-%s-%d", clientID, randUID, msgIndex)
	return &Message{
		ID:      id,
		Content: content,
	}
}