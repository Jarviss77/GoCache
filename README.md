# GoCache

GoCache is a simple in-memory key-value store written in Go. It is thread-safe and supports expiration of keys. It uses AVL tree to store the keys and values. It can handle multiple requests concurrently from multiple clients and is optimized for fast read access. 

## Features
- Multiple clients support
- Thread-safe
- AVL tree for key-value storage
- Fast read access (O(log n) using AVL tree)
- Concurrent requests handling
- RESP protocol support (https://redis.io/topics/protocol)
- In-memory storage
- Supports expiration of keys

## Installation

```bash
git clone https://github.com/Jarviss77/GoCache.git
```

## Supported Commands

```bash
PING
```
Returns PONG if the server is running.

```bash
ECHO message
```
Returns the message.

```bash
SET key value
```
Set the key -> value. First the key is hashed and then stored in the AVL tree.

```bash
GET key
```
Get the value of the key. Searches the AVL tree for the key and returns the value.

## Usage

Run the server using the following command:

```bash
go run server.go commands.go database.go hasher.go parser.go
```

The server will start running on port 6379. You can connect to the server using telnet to test the working of the server.

```bash
telnet localhost 6379
```

## Testing Commands

Go to the client directory and open the client.go file and check the lines of code. You can add the commands you want to test in the client.go file. Then run the client.go file using the following command:

```go
    _, err = conn.Write([]byte("*2\r\n$4\r\nECHO\r\n$4\r\nHello\r\n"))   <-- This
```
This is a RESP encoded command to test the ECHO command. You can read about the RESP protocol [here](https://redis.io/topics/protocol)

What this line outputs in the client side:

```bash
Hello
```

To run the client.go file, use the following command:

```bash
cd client
go run client.go
```

For example, to test the ECHO command, you can use the following command:

```go
    _, err = conn.Write([]byte("*2\r\n$4\r\nECHO\r\n$4\r\nHello\r\n"))
```
To test the PING command, you can use the following command:

```go
    _, err = conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
```

To test the SET command, you can use the following command:

```go
    _, err = conn.Write([]byte("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"))
```

To test the GET command, you can use the following command:

```go
    _, err = conn.Write([]byte("*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"))
```
