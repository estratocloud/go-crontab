ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}-alpine

RUN go install golang.org/x/tools/cmd/goimports@v0.24

WORKDIR /app

ENTRYPOINT ["sh"]
