# WebSocket GO

## Introduction

WebSocket is a protocol that allows interactive communicational session between user browser and a server

Go has tools like goroutines and channels that process real-time data safely which make it suitable for using WebSocket

## WebSocket server in Go
- In this example we'll use `gorilla/websocket` 
- `websocket.Upgrader` updates HTTP connection to WebSocket protocol
- `websocket.Conn` is a connection between client and server. Provides interfaces

## How it works

1. `websocket.Upgrader` returns `*websocket.Conn` and now our server can start websocket session
2. `Conn.ReadMessage` & `Conn.WriteMessage` used to write and read messages.
`ReadMessage` blocks calling stream until a message is received. `WriteMessage` is used for sending messages to client
3. `NextWriter` & `NextReader` allows low-level access to read and write streams. returns `Writer` & `Reader`

- `SetReadDeadline` sets time limit waiting for message, if no message received within this time seconds, connection will be closed

we can periodically send ping to clients

```go
ticker := time.NewTicker(pingPeriod)
defer ticker.Stop()

for {
    select {
    case <-ticker.C:
        conn.SetWriteDeadline(time.Now().Add(writeWait))
        if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
            return // или обработать ошибку
        }
    }
}
```

as well we can pong

```go
conn.SetPongHandler(func(string) error {
    conn.SetReadDeadline(time.Now().Add(pongWait));
    return nil 
      }
    )
```

## Security
- WebSocket uses `WSS` analog to `HTTPS` that encrypts data between client and server
- Also Server should have limits for connections, message size and other params to protect from DDoS attacks

## Scaling 
- is reachable by using goroutines to handle each connection or balancing between goroutines
- async reading and writing of messages can speed up the server

## Conclusion
- Websocket in go has a lot of advantages and is a good choice for building real-time applications