# Use Golang base image to build the app
FROM golang:1.23.5 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests (go.mod and go.sum)
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are not changed
RUN go mod tidy

# Copy the source code from the cmd folder
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

# Build the Go app
RUN go mod tidy 
RUN go build -o /app/main ./cmd 

# Start a new stage from scratch (smaller image)
FROM ubuntu:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the built Go binary from the previous stage
COPY --from=builder /app/main .

# Expose the port your app will run on (adjust as needed)
EXPOSE 9999

# Command to run the executable
CMD ["./main"]

