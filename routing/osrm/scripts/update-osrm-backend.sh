#!/bin/bash
osrmFolder=/src/data/osrm

function delete_temp() {
  rm -rf "${osrmFolder}"/"${1}"-temp/map-"${1}".osrm*
}

function run_osrm_backend_generation() {
  osrm-extract "${osrmFolder}"/"${1}"-temp/map-"${1}".pbf -p /src/osrm/"${1}".lua
  if [ $? -eq 0 ]; then
      echo "Routing data extraction finished for ${1}"
    else
      echo "Routing data extraction failed for ${1}"
      return
  fi
  osrmFile="${osrmFolder}"/"${1}"-temp/map-"${1}".osrm
  osrm-partition "$osrmFile"
  osrm-customize "$osrmFile"

  if [ $? -eq 0 ]; then
    echo "Routing data generation finished for ${1}"
  else
    echo "Routing data generation failed for ${1}"
  fi
}

function restart_osrm_backend() {
  supervisorctl stop osrm-"${1}"
  rm -rf ${osrmFolder:?}/"${1}"/*
  pattern="map-${1}.osrm*"
  find ${osrmFolder}/"${1}"-temp/ -type f -name "$pattern" -exec mv -i {} ${osrmFolder}/"${1}"/ \;
  supervisorctl start osrm-"${1}"
}

delete_temp car
run_osrm_backend_generation car
restart_osrm_backend car

delete_temp bicycle
run_osrm_backend_generation bicycle
restart_osrm_backend bicycle

delete_temp foot
run_osrm_backend_generation foot
restart_osrm_backend foot
