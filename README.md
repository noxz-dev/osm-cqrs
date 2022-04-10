# OSM-CQRS

## Module: Changeset-Watcher

- Module to observe the changesets from planet.osm.org

## Module: Search

- Search service to be able to find houses an amenties from the frontend

- build using elastic search and nodejs


## Module: Routing

- We need to get from one point to another. This module uses [OSRM](http://project-osrm.org/) with our own event processing to build a routing table and provide an api to use in the frontend


## Module: Renderer


## Frontend
- Map application built with Vuejs which integrates all services to provide a sample map application from all processed data

## Full Pipeline Setup

```vim
docker-compose -f docker-compose.full.yml up --build
```

## Start nats

```vim
docker-compose up
```

## Start nats and changeset watcher

```vim
docker-compose -f docker-compose.production.yml up --build
```