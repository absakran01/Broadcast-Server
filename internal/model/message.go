package model

import (
	"broadcast-server/internal/util"
	"fmt"
)

type Message struct {
	// msgID = clientIP-randUID-msgIndex
	ID      string `json:"id"`
	Content []byte `json:"text"`
}

func NewMessage(clientIp string, msgIndex int, content []byte) *Message {
	randUID := util.GenerateMessageUID()
	id := fmt.Sprintf("%s-%s-%d", clientIp, randUID, msgIndex)
	return &Message{
		ID:      id,
		Content: content,
	}
}