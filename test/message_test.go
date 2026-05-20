package test

import (
	"broadcast-server/internal/model"
	"broadcast-server/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	tests := []struct {
		clientIp string
		msgIndx  int
		content  []byte
	}{
		{"127.0.0.1", 1, []byte("Hello, World!")},
		{"192.168.1.100", 2, []byte("Test Message")},
		{"localhost", 3, []byte("Another Message")},
	}
	for _, tt := range tests {
		msg := model.NewMessage(tt.clientIp, tt.msgIndx, tt.content)

		clientIp, err := util.ExtractClientIpFromMsgId(msg.ID)
		assert.NoError(t, err)
		assert.Equal(t, tt.clientIp, clientIp)

		msgIndx, err := util.ExtractMsgIndxFromMsgId(msg.ID)
		assert.NoError(t, err)
		assert.Equal(t, tt.msgIndx, msgIndx)
	}
}
