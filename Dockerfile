FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tg-bot ./cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/tg-bot .


RUN adduser -D -g '' appuser
USER appuser

CMD ["./tg-bot"]