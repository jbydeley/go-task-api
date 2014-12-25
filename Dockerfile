# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
# ADD . /Go/src/github.com/gorilla/mux
# ADD . /Go/src/gopkg.in/unrolled/render.v1
ADD . /Go/src/github.com/jbydeley/go-task-api
ADD . /Go/src/github.com/jbydeley/go-task-lib
ENV GOPATH /Go

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
# RUN go get -u github.com/gorilla/mux
# RUN go get -u gopkg.in/unrolled/render.v1
RUN go get -u github.com/gorilla/mux
RUN go get -u gopkg.in/unrolled/render.v1
RUN go build github.com/jbydeley/go-task-lib
RUN go install github.com/jbydeley/go-task-api

# Run the outyet command by default when the container starts.
ENTRYPOINT /Go/bin/go-task-api

# Document that the service listens on port 8080.
EXPOSE 8080
