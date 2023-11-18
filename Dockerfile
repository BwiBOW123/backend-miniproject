# Use the official Golang image as a parent image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the Go Mod and Sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# Assuming main.go is your entry point and it is located in the cmd/main directory
RUN go build -o backend.exe cmd/main.go

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./backend.exe"]
