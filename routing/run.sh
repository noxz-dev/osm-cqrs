#!/bin/bash
export PGCON="postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT:-5432}/${POSTGRES_DB}"

until psql "${PGCON}" -c '\q'; do
  echo >&2 "Postgres is unavailable - sleeping"
  sleep 1
done

TABLE=network
SQL_EXISTS=$(printf '\dt "%s"' "$TABLE")

if [[ $(psql "$PGCON" -c "$SQL_EXISTS") ]]; then
  echo "Imposm, postgis and routing already set up. Skipping..."
else
  echo "Initializing imposm, postig and routing..."
  # Import PBF
  echo "Start PBF import"
  psql "${PGCON}" -c "DROP SCHEMA IF EXISTS imposm2pgr CASCADE; DROP SCHEMA IF EXISTS import CASCADE;"
  /src/imposm/imposm import -config /src/imposm/config.json -read /src/imposm/base.pbf -write -diff
  /src/imposm/imposm import -config /src/imposm/config.json -deployproduction

  # Import SQL
  echo "Load ImpOsm2pgRouting into database"
  echo "-- 00_init.sql"
  psql "${PGCON}" </src/pgRouting/00_init.sql
  echo "-- 01_vertices.sql"
  psql "${PGCON}" </src/pgRouting/01_vertices.sql
  echo "-- 02_edge.sql"
  psql "${PGCON}" </src/pgRouting/02_edge.sql
  echo "-- 03_update.sql"
  psql "${PGCON}" </src/pgRouting/03_update.sql

  # Import SQL
  echo "Load custom SQL"
  psql "${PGCON}" </src/pgRouting/10_network.sql
fi

/src/routing
