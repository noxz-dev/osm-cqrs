# osm-cqrs

## watch dev setup

```bash
docker-compose up
```

## search dev setup

```bash
docker-compose -f ./search/docker-compose.yml up
```

## full infrastructure production setup

```bash
docker-compose -f docker-compose.production.yml -f ./search/docker-compose.production.yml up --build
```
