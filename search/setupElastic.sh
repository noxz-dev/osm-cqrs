curl -XDELETE "http://localhost:9200/osm"

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