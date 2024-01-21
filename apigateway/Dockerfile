FROM golang:alpine3.19

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]