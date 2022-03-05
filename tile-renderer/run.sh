#!/bin/bash
export PGCON="postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT:-5432}/${POSTGRES_DB}"

until psql "${PGCON}" -c '\q'; do
  echo >&2 "Postgres is unavailable - sleeping"
  sleep 1
done

TABLE=osm_buildings
SQL_EXISTS=$(printf '\dt "%s"' "$TABLE")

if [[ $(psql "$PGCON" -c "$SQL_EXISTS") ]]; then
  echo "Imposm and postgis already set up. Skipping..."
else
  echo "Initializing imposm an postgis..."
  imposm import -config /src/imposm/config.json -read /src/imposm/niedersachsen.pbf -write -diff
  imposm import -config /src/imposm/config.json -deployproduction
fi

/src/tile-renderer
