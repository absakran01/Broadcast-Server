package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	ERR_INVALID_MESSAGE_ID = func(msgId string) error {
		return fmt.Errorf("invalid message ID format: %s", msgId)
	}
)


func GenerateMessageUID() string {
	return uuid.NewString()
}

func ExtractClientIpFromMsgId(msgId string) (string, error) {
	parts := strings.Split(msgId, "-")
	if len(parts) > 0 {
		return parts[0], nil
	}
	return "", ERR_INVALID_MESSAGE_ID(msgId)
}

func ExtractMsgIndxFromMsgId(msgId string) (int, error) {
	parts := strings.Split(msgId, "-")
	if len(parts) > 2 {
		msgIndxStr := parts[len(parts)-1]
		msgIndx, err := strconv.Atoi(msgIndxStr)
		if err != nil {
			return 0, err
		}
		return msgIndx, nil
	}
	return 0, ERR_INVALID_MESSAGE_ID(msgId)
}
