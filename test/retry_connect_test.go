package test

import (
	"broadcast-server/cmd/client"
	"testing"
)

func TestGetRemoteAddrParts(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Standard WebSocket URL",
			input:    "ws://127.0.0.1:8080",
			expected: "127.0.0.1",
		},
		{
			name:     "Localhost WebSocket URL",
			input:    "ws://localhost:8080",
			expected: "localhost",
		},
		{
			name:     "Different IP address",
			input:    "ws://192.168.1.100:8080",
			expected: "192.168.1.100",
		},
		{
			name:     "Domain name",
			input:    "ws://example.com:8080",
			expected: "example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.GetRemoteAddrParts(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
