# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.9.2

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace

ADD . /go/src/dojo-microservices/participantes

# Setting up working directory
WORKDIR /go/src/dojo-microservices/participantes

# Build the showtimes command inside the container.
RUN go install dojo-microservices/participantes

# Run the showtimes microservices when the container starts.
ENTRYPOINT /go/bin/participantes

# Service listens on port 8080.
EXPOSE 8080

