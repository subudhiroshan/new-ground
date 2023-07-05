FROM golang:1.19-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o /main .

EXPOSE 8090
ENTRYPOINT ["/main"]
