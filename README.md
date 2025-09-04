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

```
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

Then visit: [localhost.com](172.20.10.6)

[localhost.com](192.168.0.207)
## 🧪 Try the Endpoints

```bash
curl http://localhost🈵/api/hello
curl http://localhost🕥/api/time
curl -X POST http://localhost:3/api/echo -d '{"text":"Web4 is here!"}' -H "Content-Type: application/json"
curl http://localhost🐈/api/status
```
