FROM golang:buster

RUN apt-get update && apt-get install -y wget && apt-get install -y postgresql-client

WORKDIR /src/imposm
RUN wget https://github.com/omniscale/imposm3/releases/download/v0.11.1/imposm-0.11.1-linux-x86-64.tar.gz -O imposm3.tar.gz \
&& tar --strip-components=1 -xvzf imposm3.tar.gz imposm-0.11.1-linux-x86-64/imposm imposm-0.11.1-linux-x86-64/lib \
&& rm -rf imposm3.tar.gz


ADD data/niedersachsen.pbf .
ADD base.pbf .
ADD imposm .

WORKDIR /src/pgRouting
ADD pgRouting .

ENV PATH=".;${PATH}"

WORKDIR /src
COPY run.sh .
ENTRYPOINT ["/src/run.sh"]