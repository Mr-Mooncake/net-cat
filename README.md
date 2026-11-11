# Net-Cat (TCPChat)

A group chat server in Go that works like NetCat (`nc`). Clients connect over TCP, pick a name, and chat with everyone in the room.

## Features

- Up to 10 clients at the same time
- Messages show a timestamp and the sender's name: `[YYYY-MM-DD HH:MM:SS][name]:message`
- New clients see the chat history when they join
- Everyone is told when a client joins or leaves
- Empty names and empty messages are rejected
- Uses goroutines and mutexes to keep shared data safe

## Usage

```bash
go run .          # listens on port 8989
go run . 2525     # listens on port 2525
```

Connect from another terminal:

```bash
nc localhost 8989
```

## Testing

```bash
go test ./...
```

## Authors

- Mohammed Abu Hasan
- Husain Hani Tammam
- Baqer Alasmawi
