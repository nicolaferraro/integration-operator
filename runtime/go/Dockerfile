FROM golang:1.10

# Please: build this Dockerfile from the project root directory

# Env
ENV KAMEL_HOME=/go/src/github.com/nicolaferraro/integration-operator

# Srouce
WORKDIR $KAMEL_HOME
COPY . .

# Build
RUN go build ./runtime/go/cmd/camel-go-runtime

# Run
CMD ./camel-go-runtime