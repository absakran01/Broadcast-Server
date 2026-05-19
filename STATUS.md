# RedisRelay - Project Status Report

**Date:** May 18, 2026  
**Status:** In Development

## What We've Built

A WebSocket server that can run in two modes:

### 1. Server Mode (Broadcast)
- Listens for WebSocket connections
- Any message sent by one client gets broadcast to all connected clients
- Simple chat/pub-sub functionality

### 2. Client Mode (Relay)
- Connects to a remote WebSocket server
- Forwards messages from that remote server to local clients
- Acts as a proxy/relay

## How to Run

**Start a broadcast server:**
```bash
go run main.go -mode server -port 8080
```

**Start a relay (connects to another server):**
```bash
go run main.go -mode client -port 7070 -remote-host localhost -remote-port 8080
```

## Testing

Created Postman test files:
- `postman_test_requests.json` - WebSocket connection setup
- `websocket_message_examples.json` - 14 sample messages to test with
- `POSTMAN_TESTING.md` - Step-by-step testing guide

## What Still Needs Work

1. **Thread Safety** - Client list isn't protected from concurrent access
2. **Two-Way Relay** - Client mode only reads from remote, doesn't send back
3. **Redis Integration** - Not implemented yet
4. **Client Cleanup** - Disconnected clients aren't removed from list
5. **Error Handling** - Basic implementation only

## File Organization

```
main.go                          - App entry point, routing
cmd/server/                      - Server mode code
cmd/client/                      - Client relay code
postman_test_requests.json       - Postman collection
websocket_message_examples.json  - Test messages
POSTMAN_TESTING.md              - Testing instructions
```

## Next Session

Priority fixes:
- Add mutex to protect client list
- Implement proper client removal on disconnect
- Add two-way communication in client mode

---

*Simple, working foundation is in place. Ready for improvements.*
