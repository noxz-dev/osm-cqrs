# change-set watcher

Changeset Watcher watches the changesets from Open Street Maps to create events from newly added data each minute.
This module is the core of the pipeline, as it is responsible for supplying the read projections with data

## setup

- install go
- run the docker compose to start nats

- run: "go get" inside the changeset-watcher folder

```vim
- go run main.go
```

### run the demo subscriber:

- cd subscriber-demo

```vim
- go run sub.go
```

## PBF importer

```vim
go run main.go --import <filepath>

EXAMPLE:

go run main.go --import ./data/Hannover.pbf
```
