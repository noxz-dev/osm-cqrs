##
## Build
##

FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /changeset-watcher

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /changeset-watcher /changeset-watcher

USER nonroot:nonroot

ENTRYPOINT ["/changeset-watcher"]