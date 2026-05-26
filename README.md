![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
# Broadcast Server

A lightweight real-time broadcast server built with Go.

The server allows multiple clients to connect and exchange messages in real time while providing reliable message delivery through reconnection handling and missed-message recovery.

Designed for environments where clients can temporarily disconnect and reconnect without losing messages.


# Features

- Real-time message broadcasting
- Multiple concurrent clients  
- Automatic reconnection support  
- Missed message recovery  
- Server-generated message indexing  
- Message deduplication using unique IDs  
- Configurable server settings
- Lightweight architecture 

# How It Works

The server maintains:

- Active client connections
- Global message index
- Message storage
- Client synchronization state

When a client reconnects:

1. Client sends its last known message index
2. Server compares the index with its current index
3. Missing messages are identified
4. Missed messages are replayed
5. Client resumes normal communication

# Installation

## Prerequisites

Install:

- Go 1.24+
- Git

Verify:

```bash
go version
```

Clone repository:

```bash
git clone https://github.com/yourusername/broadcast-server.git

cd broadcast-server
```

Install dependencies:

```bash
go mod download
```

Build:

```bash
go build -o broadcast-server
```

# Quick Start

## Start server

```bash
./broadcast-server start
```

## Connect client

Open another terminal:

```bash
./broadcast-server connect
```

Open additional terminals:

```bash
./broadcast-server connect
```

## Send messages

Input:

```text
hello everyone
```

Output:

```text
Client 1:

Client 0:hello everyone

Client 2:

Client 0:hello everyone
```


# Connection Flow

```text
Client
   │
   │ Connect
   │
Server
   │
   │ Send current global index
   │
Client
   │
   │ Send local index
   │
Server
   │
   │ Compare indexes
   │
   │ Replay cached missing messages
   │
Connected
```


# Message Flow

```text
Client
    │
    │ Message
    │
Server
    │
    │ Assign message index
    │
    │ Store message
    │
    │ Broadcast
    │
Clients
```

# License
This project is licensed under the MIT License - see the LICENSE file for details.

# Project Reference

This project was inspired by [roadmap.sh](https://roadmap.sh/projects/broadcast-server).


---


Built by abdulmajeed alsakran
