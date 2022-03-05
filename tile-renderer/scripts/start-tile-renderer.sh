sudo docker run \
        -p 8081:80 \
        -p 5433:5432 \
        -e THREADS=24 \
        -e UPDATES=enabled \
        -v niedersachsen-osm-data:/var/lib/postgresql/12/main \
        -v niedersachsen-rendered-tiles:/var/lib/mod_tile \
        -e ALLOW_CORS=enabled \
        -d overv/openstreetmap-tile-server \
        run 
