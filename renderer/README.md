# Renderer

Service for serving vector tiles in PBF format. Changes made to the map via osmChange events do not require a rerender
of the vector tiles. Updates are instantly reflected in the next request of the vector tile.

## Getting Started

To start the renderer service run

````shell
docker-compose up --build
````

The Tile server should now be accessible under `localhost:8080`.

The importer only works if a Nats-Server is running and accessible in the docker network `osm_cqrs_network` under the
ip `nats`.

To delete all containers and volumes run the `./docker-cleanup.sh` shell script.

## Infrastructure

The renderer internally uses a PostGIS database to store the current state of the map. To make updates with the events
received from Nats [Imposm3](https://github.com/omniscale/imposm3) is used. To serve the vector tiles for map rendering
the [Tegola](https://github.com/go-spatial/tegola) library used. Which data is included in each vector tile depending on
the requested zoom layer is defined in the specific Toml file which can be found under `./tegola/config.toml`

## Related research and commands

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

### Imposm3

Insert initial osm file with setup for continuous osm changeset update

````shell
imposm import \
    -config config.json \ 
    -read path_to_osm.osm.pbf \
    -write -diff
````

Config file sums als configuration arguments together (
see [Config file](https://imposm.org/docs/imposm3/latest/tutorial.html#config-file)). The`mapping.json` file contains
the database schema
(see [Data Mapping](https://imposm.org/docs/imposm3/latest/mapping.html)).

````json5
// imposm.json
{
  "cachedir": "./cache",
  "connection": "postgis://postgres:password@localhost/postgres",
  "mapping": "./mapping.json",
  "diffdir": "./diff"
}
````

To make the database available for continuous updates the tables must be made ready for production. They will be added
to the public schema and the import schema will be removed.

````shell
imposm import -config config.json -deployproduction
````

Automatic updates can then be started with:

````shell
imposm run -config config.json
````

To import just a one or multiple specific change file run:

````shell
imposm diff -config config.json changes-1.osc.gz changes-2.osc.gz
````