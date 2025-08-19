# Build stage
FROM golang:1.20 AS builder
WORKDIR /app

# Copy Go source
COPY . .

# Initialize module if missing, tidy dependencies
RUN go mod init web4app || true
RUN go mod tidy

# Build binary
RUN go build -o web4app main.go

# Run stage
FROM alpine:latest
WORKDIR /root/

# Copy binary + static assets
COPY --from=builder /app/web4app .
COPY static ./static

EXPOSE 8080
ENV PORT=8080
CMD ["./web4app"]
