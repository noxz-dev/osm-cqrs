# Tile renderer

## Useful command

Prepare OsmChange for inserting into PostGIS

`osmosis --read-xml-change 126.osc --simplify-change --write-xml-change out_2.xml`

Insert created file into PostGIS

````shell
osm2pgsql --append \ 
    -r xml -s -C 300 -G --hstore --number-processes 24 \
    --style openstreetmap-carto.style \ 
    --tag-transform-script openstreetmap-carto.lua \ 
    -d gis -H localhost -U renderer -W out_2.xml
````
