FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o jlt-bot .

FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata \
    && adduser -D appuser

WORKDIR /app

COPY --from=builder /app/jlt-bot .
COPY --from=builder /app/images ./images

USER appuser

CMD ["./jlt-bot"]