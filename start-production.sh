docker-compose \
  -f docker-compose.production.yml \
  -f ./search/docker-compose.production.yml \
  -f ./tile-renderer/backend/docker-compose.production.yml \
  -f ./routing/docker-compose.production.yml \
  up --build
