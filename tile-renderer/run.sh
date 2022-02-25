#!/bin/bash
sleep 10 # Workaround to wait for postgis

./imposm import -config config.json -read base.pbf -write -diff
./imposm import -config config.json -deployproduction
