# Use the official Golang image
FROM golang:1.20

# Set the working directory
WORKDIR /PortPatrol

# Copy the Go program into the container
COPY . .

# Build the Go program
RUN go build -o server_ip .

# Create an output directory
RUN mkdir /output

# Run the Go program
CMD ["./server_ip"]
