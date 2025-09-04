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

FROM vm/ubuntu:18.04
RUN curl -sL https://deb.nodesource.com/setup_22 | bash
RUN apt install nodejs
COPY . .
RUN npm install
SECRET ENV CONFIGCAT_AUTH_KEY
RUN curl "https://api.configcat.com/v1/products/dcd53ddb-8104-4e48-8cc0-5df1088c6113/environments" \\
    -X POST \\
    -u $CONFIGCAT_AUTH_KEY \\
    -H "Content-Type: application/json" \\
    -d '{"name": "webapp.io-'$JOB_ID'"}'
RUN BACKGROUND REACT_CONFIGCAT_ENV="layerci-$JOB_ID" npm run start
EXPOSE WEBSITE localhost:8080

ENV CI_NAME=layerci \\
   CI_BUILD_NUMBER=$JOB_ID \\
   CI_BUILD_URL="https://webapp.io/qubuhub/commits?query=web4$JOB_ID" \\
   CI_BRANCH="$GIT_BRANCH" \\
   CI_PULL_REQUEST="$PULL_REQUEST_URL"

SECRET ENV COVERALLS_REPO_TOKEN

RUN (the test command)
