docker-compose stop
docker-compose rm -vf
docker volume rm renderer_renderer-importer-cache
docker volume rm renderer_renderer-importer-diff
docker volume rm renderer_renderer-postgis-data
