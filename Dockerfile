# Use the official Go image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files and download dependencies

COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Install buf
RUN curl -sSL \
    "https://github.com/bufbuild/buf/releases/latest/download/buf-Linux-x86_64" \
    -o /usr/local/bin/buf && \
    chmod +x /usr/local/bin/buf

# Verify installation
RUN buf --version

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Command to run the application
CMD ["go", "run", "main.go"]