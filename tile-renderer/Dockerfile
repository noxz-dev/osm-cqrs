## Build ##
FROM golang:buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /tile-renderer



## Deploy ##
FROM debian:stable-slim

RUN apt-get update && apt-get install -y wget

COPY --from=build /tile-renderer /src/tile-renderer

WORKDIR /src/imposm
RUN wget https://github.com/omniscale/imposm3/releases/download/v0.11.1/imposm-0.11.1-linux-x86-64.tar.gz -O imposm3.tar.gz \
&& tar --strip-components=1 -xvzf imposm3.tar.gz imposm-0.11.1-linux-x86-64/imposm imposm-0.11.1-linux-x86-64/lib \
&& rm -rf imposm3.tar.gz


COPY base.pbf ./base.pbf
#COPY data/niedersachsen.pbf ./niedersachsen.pbf
COPY imposm/config.json ./config.json
COPY imposm/mapping.json ./mapping.json
ENV PATH="/src/imposm:${PATH}"

WORKDIR /src

COPY run.sh .
ENTRYPOINT ["/src/run.sh"]

# TODO: ADD working container for accepting events and publishing them to the database \

