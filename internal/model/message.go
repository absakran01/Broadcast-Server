package model

import (
	"fmt"
	"strings"
)
// ParseUIDAndContent splits a message in the format UID|content and returns UID and content.
func ParseUIDAndContent(msg []byte) (string, []byte, error) {
	parts := strings.SplitN(string(msg), "|", 2)
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("invalid message format: missing UID or content")
	}
	return parts[0], []byte(parts[1]), nil
}

type Message struct {
	// msgID = clientID:msgUID-msgIndex
	//clientID = client.RemoteAddr() + ":" + instanceUID
	ID      string `json:"id"`
	Content []byte `json:"text"`
}

func NewMessage(clientID string, msgUID string, msgIndex int, content []byte) *Message {
       id := fmt.Sprintf("%s:%s-%d", clientID, msgUID, msgIndex)
       return &Message{
	       ID:      id,
	       Content: content,
       }
}