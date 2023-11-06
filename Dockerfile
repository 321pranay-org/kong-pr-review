FROM golang:1.21-alpine3.17

WORKDIR /tmp

COPY ./* ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/kong-pr-review main.go

RUN chmod +x /bin/kong-pr-review

RUN rm -rf /tmp

WORKDIR /go

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
