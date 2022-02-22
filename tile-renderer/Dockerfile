FROM golang:buster

RUN apt-get update && apt-get install -y wget

RUN wget https://github.com/omniscale/imposm3/releases/download/v0.11.1/imposm-0.11.1-linux-x86-64.tar.gz -O imposm3.tar.gz \
&& tar -xvzf imposm3.tar.gz

# TODO: ADD working container for accepting events and publishing them to the database \

