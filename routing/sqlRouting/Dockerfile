## Build ##
FROM golang:buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .. .

RUN go build -o /routing



## Deploy ##
FROM debian:stable-slim

RUN apt-get update && apt-get install -y wget && apt-get install -y postgresql-client

COPY --from=build /routing /src/routing


WORKDIR /src/imposm
RUN wget https://github.com/omniscale/imposm3/releases/download/v0.11.1/imposm-0.11.1-linux-x86-64.tar.gz -O imposm3.tar.gz \
&& tar --strip-components=1 -xvzf imposm3.tar.gz imposm-0.11.1-linux-x86-64/imposm imposm-0.11.1-linux-x86-64/lib \
&& rm -rf imposm3.tar.gz


ADD data/Hannover.pbf .
ADD data/niedersachsen.pbf .
ADD base.pbf .
ADD imposm .

WORKDIR /src/pgRouting
ADD pgRouting .

ENV PATH="/src/imposm:${PATH}"

WORKDIR /src
#COPY test-data/create.xml .
#COPY test-data/update.xml .
COPY run.sh .
RUN mkdir -p /src/out-files
ENTRYPOINT ["/src/run.sh"]