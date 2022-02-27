# osm-cqrs

## watch dev setup

```bash
docker-compose up --build
```

## search dev setup

```bash
docker-compose -f ./search/docker-compose.yml up --build
```

## full infrastructure production setup

```bash
docker-compose -f docker-compose.production.yml -f ./search/docker-compose.production.yml up --build
```
