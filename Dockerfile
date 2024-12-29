# Use the official Go image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files and download dependencies

COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Command to run the application
CMD ["go", "run", "main.go"]