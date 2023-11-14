# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory to the GOPATH
WORKDIR $GOPATH/src

# Copy the entire project to the working directory
COPY . .

# Install any dependencies required by the project
RUN go get -d -v ./...

# Build the application
RUN go build -o /app/main ./

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container
CMD ["/app/main"]
