# Start from a lightweight Golang base image
FROM golang:1.18-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
# CGO_ENABLED=0, creates a statically linked binary which basically means that the binary has everything it needs to run, so that a base alpine image can be used and 
# there is no need to install any libraries promoting (ease of deployment and security)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api ./cmd/server

# Start a new stage from scratch for smaller image size
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/go-api .

# Command to run the executable
CMD ["./go-api"]
