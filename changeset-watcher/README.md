# change-set watcher

## setup

- install go
- run the docker compose to start nats

- run: "go get" inside the changeset-watcher folder

```bash
- go run main.go
```

### run the demo subscriber:

- cd subscriber-demo

```bash
- go run sub.go
```

## PBF importer

```bash
go run main.go --import <filepath>

EXAMPLE:

go run main.go --import ./data/Hannover.pbf
```
