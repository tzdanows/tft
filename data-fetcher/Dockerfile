FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o data-fetcher .

# Linux distribution alpine
FROM alpine:latest

# Necessary certification for go to make https requests
RUN apk --no-cache add ca-certificates #

WORKDIR /root/
COPY --from=builder /app/data-fetcher .

# Execute data-fetcher binary
CMD ["./data-fetcher"]