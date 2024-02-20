FROM golang:latest AS builder
WORKDIR /app

# Since the context is 'data-fetcher', the paths should not be prefixed with './data-fetcher/'
COPY go.mod .
COPY go.sum .
RUN go mod tidy

# Copy the rest of the application's source code from the host to the container's workspace.
COPY main.go .

# Build Application inside the container (Using `-o df` to specify the output binary name as df)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o df .

# Start from alpine:latest with the CA certificates installed for HTTPS requests.
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

# Set the working directory to /root/ (the default directory when you login to the alpine image).
WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/df .

# Run the df binary.S
CMD ["./df"]
