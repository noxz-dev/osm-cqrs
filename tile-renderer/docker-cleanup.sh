docker-compose stop
docker-compose rm -v
docker volume rm tile-renderer_renderer-importer-cache
docker volume rm tile-renderer_renderer-importer-diff
docker volume rm tile-renderer_renderer-postgis-data
