# Web4App Fullstack

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

Then visit: [web4app](index.html)

[web4app.io](index)
## 🧪 Try the Endpoints

```bash
curl http://localhost/api/hello
curl http://localhost/api/time
curl -X POST http://localhost/api/echo -d '{"text":"Web4 is here!"}' -H "Content-Type: application/json"
curl http://localhost.com/api/status
```
