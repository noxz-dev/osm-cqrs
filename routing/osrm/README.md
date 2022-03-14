# Routing

Service for routing spatial data. It can consume osmChange events for updating the local map. Changes are not instantly
reflected in the routing database.

## Infrastructure

[OSRM](https://github.com/Project-OSRM/osrm-backend) is currently used for routing. Therefore, a full rebuild of the
routing database is needed to reflect updates.

The rebuild of the routing database is automated via a cron job which runs every 5 minutes.

To store the current state of the map and consume osmChange events [Osmosis](https://github.com/osmosis-labs/osmosis) is
used. A local PBF file of the map is stored on the file system and Osmosis used to apply Changesets to the file.

To automate the map update and routing database as well as server restart a number of shell scripts are used. These can
be found under `./scripts`

## Getting Started

To start the routing service run

```shell
docker-compose up --build
```

The routing api is available under multiple ports. Each port reflect a different routing profile:

| Address        | Profile |
|----------------|---------|
| localhost:5000 | car     |
| localhost:5001 | bicycle |
| localhost:5002 | foot    |

For further usage of the api, please refer to
the [OSRM documentation](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md)

