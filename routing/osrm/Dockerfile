FROM golang:buster as routing_updater_build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY .. .

RUN go build -o /routing-updater

FROM debian:stable-slim as osmium_build

# Adapted from: https://github.com/stefda/docker-osmium-tool

ENV LIBOSMIUM_VERSION 2.18.0
ENV OSMIUM_TOOL_VERSION 1.14.0

RUN apt-get update
RUN apt-get install -y \
    wget g++ cmake cmake-curses-gui make libprotozero-dev libboost-dev libboost-program-options-dev libbz2-dev zlib1g-dev \
    liblz4-dev libexpat1-dev cmake pandoc

RUN mkdir /var/install
WORKDIR /var/install

RUN wget https://github.com/osmcode/libosmium/archive/refs/tags/v${LIBOSMIUM_VERSION}.tar.gz && \
    tar xzvf v${LIBOSMIUM_VERSION}.tar.gz && \
    rm v${LIBOSMIUM_VERSION}.tar.gz && \
    mv libosmium-${LIBOSMIUM_VERSION} libosmium

RUN cd libosmium && \
    mkdir build && cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_EXAMPLES=OFF -DBUILD_TESTING=OFF -DINSTALL_PROTOZERO=ON .. && \
    make

RUN wget https://github.com/osmcode/osmium-tool/archive/refs/tags/v${OSMIUM_TOOL_VERSION}.tar.gz && \
    tar xzvf v${OSMIUM_TOOL_VERSION}.tar.gz && \
    rm v${OSMIUM_TOOL_VERSION}.tar.gz && \
    mv osmium-tool-${OSMIUM_TOOL_VERSION} osmium-tool

RUN cd osmium-tool && \
    mkdir build && cd build && \
    cmake -DOSMIUM_INCLUDE_DIR=/var/install/libosmium/include/ .. && \
    make

FROM debian:stable-slim as osrm_build

ENV OSRM_VERSION 5.26.0



RUN apt-get update
RUN apt-get install -y \
    build-essential wget cmake pkg-config \
    libbz2-dev libstxxl-dev libstxxl1v5 libxml2-dev \
    libzip-dev libboost-all-dev lua5.2 liblua5.2-dev libtbb-dev


RUN wget https://github.com/Project-OSRM/osrm-backend/archive/refs/tags/v${OSRM_VERSION}.tar.gz && \
    tar -xzf v${OSRM_VERSION}.tar.gz && \
    rm v${OSRM_VERSION}.tar.gz && \
    mv osrm-backend-${OSRM_VERSION} osrm-backend

RUN cd osrm-backend && \
    mkdir -p build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release && \
    cmake --build . && \
    cmake --build . --target install && \
    cd ../profiles && \
    cp -r * /opt && \
    strip /usr/local/bin/* && \
    rm -rf /src /usr/local/lib/libosrm*


FROM debian:stable-slim as runstage

RUN mkdir -p /src/osrm  && mkdir -p /opt

RUN apt-get update && \
    apt-get install -y --no-install-recommends libboost-program-options1.74.0 libboost-regex1.74.0 \
        libboost-date-time1.74.0 libboost-chrono1.74.0 libboost-filesystem1.74.0 \
        libboost-iostreams1.74.0 libboost-system1.74.0 libboost-thread1.74.0 \
        expat liblua5.2-0 libtbb2 supervisor &&\
    rm -rf /var/lib/apt/lists/*

COPY --from=osrm_build /usr/local /usr/local
COPY --from=osrm_build /opt /src/osrm
COPY --from=routing_updater_build /routing-updater /usr/bin/routing-updater
COPY --from=osmium_build /var/install/osmium-tool/build/src/osmium /usr/bin/osmium

COPY base.pbf /src/data/data/map/map.pbf
COPY data/config/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

ADD scripts/ /src/scripts/
RUN /src/scripts/initialize-osrm-backend.sh

CMD ["/src/scripts/run.sh"]
