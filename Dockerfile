FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o /app/push-watch .

FROM alpine:3.14

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/push-watch /usr/local/bin/push-watch

ENTRYPOINT ["/usr/local/bin/push-watch"]
