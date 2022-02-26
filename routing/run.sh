#!/bin/bash
sleep 5 # Workaround to wait for postgis

export PGCON="${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT:-5432}/${POSTGRES_DB}"

echo $PGCON

# Import PBF
echo "Start PBF import"
psql "postgresql://${PGCON}" -c "DROP SCHEMA IF EXISTS imposm2pgr CASCADE; DROP SCHEMA IF EXISTS import CASCADE;"
/src/imposm/imposm import -config /src/imposm/config.json -read /src/imposm/base.pbf -write -diff
/src/imposm/imposm import -config /src/imposm/config.json -deployproduction

# Import SQL
echo "Load ImpOsm2pgRouting into database"
echo "-- 00_init.sql"
psql "postgresql://${PGCON}" </src/pgRouting/00_init.sql
echo "-- 01_vertices.sql"
psql "postgresql://${PGCON}" </src/pgRouting/01_vertices.sql
echo "-- 02_edge.sql"
psql "postgresql://${PGCON}" </src/pgRouting/02_edge.sql
echo "-- 03_update.sql"
psql "postgresql://${PGCON}" </src/pgRouting/03_update.sql

# Import SQL
echo "Load custom SQL"
psql "postgresql://${PGCON}" </src/pgRouting/10_network.sql

/src/imposm/imposm run -config /src/imposm/config.json
