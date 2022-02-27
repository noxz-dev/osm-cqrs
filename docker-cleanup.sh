docker-compose rm -f -s -v
prefix=osm-cqrs

docker container ls -a -q --filter="name=${prefix}" | while read -r line; do
  docker container stop "$line"
  docker container rm -v "$line"
done

docker volume ls -q | grep $prefix | while read -r line; do
  docker volume rm "$line"
done
