docker-compose stop
docker-compose rm -v
docker volume rm backend_renderer-importer-cache
docker volume rm backend_renderer-importer-diff
docker volume rm backend_renderer-postgis-data
