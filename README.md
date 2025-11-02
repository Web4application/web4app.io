## DISCORD GO WEBSOCKET CLIENT FLOW DIAGRAM

```xlsl


+-------------------+                             +----------------------+
|   Go Discord      |                             |   Discord Gateway    |
|   Client (Session)|                             |                      |
+-------------------+                             +----------------------+
          |                                                 |
          |  Open()                                        |
          |------------------------------------------------>|
          |    Op 10 Hello packet (with heartbeat interval) |
          |<------------------------------------------------|
          |   Start heartbeat goroutine                      |
          |   Start listen goroutine                         |
          |                                                 |
          |  Send Op 2 Identify (new session) or Op 6 Resume|
          |------------------------------------------------>|
          |                                                 |
          |           READY / RESUMED event                |
          |<------------------------------------------------|
          | Update internal session state                   |
          |                                                 |
          |--Heartbeat Loop--                               |
          |   Every interval send Op 1 heartbeat           |
          |------------------------------------------------>|
          |   Receive Op 11 heartbeat ACK                  |
          |<------------------------------------------------|
          |                                                 |
          |---Event Handling Loop---                         |
          |  Listen for messages from Gateway              |
          |<------------------------------------------------|
          |   onEvent() decodes messages                   |
          |   Dispatch to event handlers (message, member,|
          |   voice, presence, etc.)                       |
          |                                                 |
          |---Voice Connection---                            |
          | ChannelVoiceJoin() sends Op 4 Join             |
          |------------------------------------------------>|
          |  Receive VoiceServerUpdate                     |
          |<------------------------------------------------|
          |  Establish UDP / WS connection for voice       |
          |                                                 |
          |---Error / Disconnect---                          |
          | Detect connection lost or invalid heartbeat    |
          | Close(), then reconnect() if ShouldReconnectOnError |
          |------------------------------------------------>|
          | (loops back to Open())                          |
          
```

# Web4App.io Fullstack

🚀 A simple full-stack starter kit using Go for the backend and vanilla HTML for the frontend.

## 🌐 Features

- Serve static frontend from `./static`
- Multiple backend API endpoints:
  - `/api/hello`
  - `/api/time`
  - `/api/echo`
  - `/api/status`
- HTML fetches API response and displays it dynamically

## 📁 Project Structure

```bash
web4app_fullstack/
├── main.go              # Go backend with multiple endpoints
├── static/
│   └── index.html       # Frontend
└── docs/
    └── API.md           # API Documentation
```

## ▶️ Getting Started

```bash
go run main.go
```

Then visit: [web4app](index.html)

[web4app.io](index)
## 🧪 Try the Endpoints

```bash
curl http://localhost/api/hello
curl http://localhost/api/time
curl -X POST http://localhost/api/echo -d '{"text":"Web4 is here!"}' -H "Content-Type: application/json"
curl http://localhost.com/api/status
```
