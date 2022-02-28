docker-compose \
  -f docker-compose.production.yml \
  -f ./search/docker-compose.production.yml \
  -f ./tile-renderer/docker-compose.production.yml \
  -f ./routing/docker-compose.production.yml \
  up --build
