# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory to the GOPATH
WORKDIR $GOPATH/src

# Copy only the go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Install any dependencies required by the project
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN go build -o /app/main ./

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container
CMD ["/app/main"]
