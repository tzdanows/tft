FROM golang:latest

WORKDIR /app/data-fetcher

COPY . ..

RUN go build -o main .

CMD ["./main"]