sudo docker volume create niedersachsen-osm-data
sudo docker volume create niedersachsen-rendered-tiles
sudo docker run \
        -e UPDATES=enabled \
        -e THREADS=24 \
        -v /home/james/Projects/osm-tile-server/niedersachsen-latest.osm.pbf:/data.osm.pbf \
        -v /home/james/Projects/osm-tile-server/niedersachsen.poly:/data.poly \
        -v niedersachsen-osm-data:/var/lib/postgresql/12/main \
        -v niedersachsen-rendered-tiles:/var/lib/mod_tiles \
        overv/openstreetmap-tile-server \
        import
