# OSM-CQRS

The project OSM-CQRS implements an exemplary architecture how the Command Query Segregation Pattern (CQRS) can be
applied to spatial data. Within the scope of this project, three read projections were implemented which, together with
a frontend, enable the implementation of a simple map application. The project uses changesets provided by Open Street
Maps which are transferred to the respective projections using our event
processor ([changeset watcher](./changeset-watcher/)).

## Folder Structure

### Module: [Changeset-Watcher](./changeset-watcher//)

- The changeset watcher is the core of the pipeline, it takes care that new changes which are published as changesets on
  planet.osm.org are sent to the read projektions after respective pre-processing steps.
- includes its own filter format to adapt changesets to the needs of the applications

### Module: [Search](./search/)

- Search service to be able to find houses an amenties from the frontend

- build using elastic search and nodejs

### Module: [Routing](./routing/)

- We need to get from one point to another. This module uses [OSRM](http://project-osrm.org/) with our own event
  processing to build a routing table and provide an api to use in the frontend

### Module: [Renderer](./renderer/)

- The rendering module manages a Vector Tile Server and additional event processing pipeline to keep its data up to date
  in order to render it in the frontend in the shortest possible time.

### Module: [Frontend](./frontend/)

- Map application built with Vuejs which integrates all services to provide a sample map application from all processed
  data

### [Docs](./docs/)

- The Docs folder contains diagrams general notes and research on the individual technologies used.

### [Watcher-Config](./watcher-config/)

- This Folder contains filter configurations for the change-set watcher

### [Metric-Analysis](./metric-analysis)

- We tested our implementation on a server for one week. This folder contains the analyses and results.

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

## License

MIT License © 2022 [Finn Beer](https://github.com/noxz-dev), [Jannes Neemann](https://github.com/JamesNeumann)
, [Darius Alter](https://github.com/herralter) 