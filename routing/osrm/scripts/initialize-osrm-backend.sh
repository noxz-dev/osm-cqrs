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

init_backend car
init_backend bicycle
init_backend foot
