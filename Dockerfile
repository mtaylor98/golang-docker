# golang image where workspace (GOPATH) configured at /go
FROM golang:alpine

# Copy the local package files to container's workspace
# Need to look at best practice here for adding code to container
ADD . /go/src/github.com/mtaylor98/golang-docker

# Build the golang-docker command inside container
RUN go install github.com/mtaylor98/golang-docker

# Run the golang-docker command when the container starts
ENTRYPOINT /go/bin/golang-docker

# http server listens on port 8080
EXPOSE 8080
