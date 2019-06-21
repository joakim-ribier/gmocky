# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Joakim Ribier <joakim.ribier@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/joakim-ribier/gmocky

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN GO111MODULE=on go get -d -v ./...

# Install the package
RUN GO111MODULE=on go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8585

ENV GMOCKY_PORT 8585

# Run the executable
CMD ["gmocky", "start"]
