# change-set watcher

Changeset Watcher (CSW) watches the changesets from Open Street Maps to create events from newly added data each minute.
This module is the core of the pipeline, as it is responsible for supplying the read projections with data
## project structure
the project contains of some packages. Non-self explaining packages are explained in the following section.
### Package `config`
This package contains constants which affect the behavior of the CSW. It contains three files bundling similar constants:

- `config.go`: general constants
- `statistics.go`: headers for statistics
- `urls.go`: urls accessed by the CSW
### Package `statistics`
This package implements statistic features and provides them to the CSW. 

### Package `types`
This package defines structs and methods. All structs are defined in `types.go`. All methods and factories are defined 
in the files:

- `action.go`: modifications, deletions and creations are defined as _Actions_. This file contains methods to handle, manage or modify these actions
- `change.go`: these file contains all methods for the `OsmChangeNormalized` and `OsmChange` structs.
- `filters.go`: contains factories for `NodeFilter` and `WayFilter`
- `node.go`: methods and helper functions for `node`struct
- `way.go`: methods and helper functions for `way`struct


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
