#!/bin/bash
sleep 10 # Workaround to wait for postgis

imposm import -config /src/imposm/config.json -read /src/imposm/base.pbf -write -diff
imposm import -config /src/imposm/config.json -deployproduction

/src/tile-renderer
