## Setup ElasticSearch

```bash
docker-compose up --build

docker-compose up --build --attach search-backend
```

## setup the schema

```bash
curl -XPUT "http://localhost:9200/osm" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "name": {
        "type": "text"
      },
      "location": {
        "type": "geo_point"
      }
    }
  }
}'
```

## test the setup:

- https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started.html

### find everything

```bash
curl -XGET "http://localhost:9200/_search" -H 'Content-Type: application/json' -d'
{
  "query": {
    "match_all": {}
  }
}'
```

### delete the index

```bash
curl -XDELETE "http://localhost:9200/osm"
```
