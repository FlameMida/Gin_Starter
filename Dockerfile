FROM golang:alpine

WORKDIR /go/src/gin_starter
COPY . .

RUN go generate && go env && go build -o server .

FROM alpine:latest

WORKDIR /go/src/gin-starter

COPY --from=0 /go/src/gin_starter ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
