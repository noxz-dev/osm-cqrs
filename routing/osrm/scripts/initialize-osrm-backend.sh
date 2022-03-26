#!/bin/bash
osrmFolder=/src/data/osrm

function init_backend() {
  mkdir -p ${osrmFolder}/"${1}"
  mkdir -p ${osrmFolder}/"${1}"-temp

  linkName=${osrmFolder}/"${1}"-temp/map-"${1}".pbf

  if [ ! -L "$linkName" ]; then
    ln -s /src/data/map/map.pbf "$linkName"
  else
    echo "Symlink already exists. Skipping..."
  fi

  echo "Initialized backend for ${1}"
}

if [[ ! -e /src/data/map ]]; then
  mkdir /src/data/map
fi

if [[ ! -e /src/data/statistics ]]; then
  mkdir /src/data/statistics
fi

if [ ! -L /src/data/map/map.pbf ]; then
  mv /src/map.pbf /src/data/map/map.pbf
fi

init_backend car
init_backend bicycle
init_backend foot
