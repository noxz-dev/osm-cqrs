##
## Build
##

FROM golang:1.17-buster AS build

WORKDIR /app

RUN apt-get update && apt-get install pkg-config

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /changeset-watcher

##
## Deploy
##

FROM debian:stable-slim

WORKDIR /

COPY --from=build /changeset-watcher /changeset-watcher
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


ENTRYPOINT ["/changeset-watcher"]