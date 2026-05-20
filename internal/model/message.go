package model

type Message struct {
	// msgID = clientIP-randUID-msgIndex
	ID	  string `json:"id"`
	Content  string `json:"text"`
}