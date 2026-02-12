# -------------------------
# Stage 1: Build
# -------------------------
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tg_bot main.go

# -------------------------
# Stage 2: Minimal runtime
# -------------------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/tg_bot .
CMD ["./tg_bot"]
