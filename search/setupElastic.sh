# curl -XDELETE "http://localhost:9200/osm"

echo "setup elasticsearch schema"
curl -XPUT "http://elasticsearch:9200/osm" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "search_analyzer": "whitespace"
      },
      "location": {
        "type": "geo_point"
      }
    }
  }
}'