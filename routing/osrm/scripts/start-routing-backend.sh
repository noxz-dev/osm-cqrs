#!/bin/bash
osrmFolder=/src/data/osrm

function start_backend() {
  osrmFile=${osrmFolder}/${1}/map-${1}.osrm
  echo "$osrmFile"
  if [ -f "${osrmFile}" ]; then
    echo "Starting OSRM for ${1}"
    supervisorctl start osrm-"${1}"
  else
    echo "OSRM does not exist for ${1}. Skipping..."
  fi
}

start_backend car
start_backend foot
start_backend bicycle
