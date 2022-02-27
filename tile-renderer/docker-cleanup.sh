docker-compose stop
docker-compose rm -v
docker volume rm tile-renderer_importer-cache
docker volume rm tile-renderer_importer-diff
docker volume rm tile-renderer_postgis-data
