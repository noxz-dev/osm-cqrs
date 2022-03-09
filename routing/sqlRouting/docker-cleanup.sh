docker-compose stop
docker-compose rm -f
docker volume rm routing_routing-postgis-data
docker volume rm routing_routing-importer-cache
docker volume rm routing_routing-importer-diff
