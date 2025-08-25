# API Documentation

## GET `/api/hello`
Returns a welcome message.
```json
{ "message": "Hello from Go 🚀" }
```

## GET `/api/time`
Returns the current UTC server time.
```json
{ "time": "2025-07-18T20:20:00Z" }
```

## POST `/api/echo`
Echoes back the text sent in the request.
### Request:
```json
{ "text": "Hello" }
```
### Response:
```json
{ "echo": "Hello" }
```

## GET `/api/status`
Returns system status.
```json
{ "status": "ok", "uptime": "10s" }
```
