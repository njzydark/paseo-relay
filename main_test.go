package main

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebSocketMessageLimitMatchesHostedRelay(t *testing.T) {
	const hostedRelayWebSocketMessageLimit = 32 * 1024 * 1024

	if maxMessageBytes != hostedRelayWebSocketMessageLimit {
		t.Fatalf("maxMessageBytes = %d, want %d", maxMessageBytes, hostedRelayWebSocketMessageLimit)
	}
}

func TestPendingBufferCanHoldOneMaxSizeMessage(t *testing.T) {
	if maxPendingBytes != maxMessageBytes+1 {
		t.Fatalf("maxPendingBytes = %d, want maxMessageBytes + 1 (%d)", maxPendingBytes, maxMessageBytes+1)
	}

	buf := newFrameBuffer(1)
	payload := make([]byte, maxMessageBytes)
	buf.push(websocket.BinaryMessage, payload)

	frames := buf.flush()
	if len(frames) != 1 {
		t.Fatalf("flushed %d frames, want 1", len(frames))
	}
	if got := len(frames[0]); got != maxMessageBytes+1 {
		t.Fatalf("frame size = %d, want %d", got, maxMessageBytes+1)
	}
	if got := frames[0][0]; got != byte(websocket.BinaryMessage) {
		t.Fatalf("message type prefix = %d, want %d", got, websocket.BinaryMessage)
	}
}
