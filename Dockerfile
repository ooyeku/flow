# Start from the latest golang base image
FROM golang:1.21

# Add Maintainer Info
LABEL maintainer="Ola Yeku <ooyeku@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Install flow cli
RUN go install .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]
CMD ["chat"]

# For CLI build:
# docker build -t flow-cli-build .
# docker run -it --name flow-cli flow-cli-build

# For server build:
# docker build -t flow-server-build .
# docker run --name flow-server -p 8080:8080 flow-server-build