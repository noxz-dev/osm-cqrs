#!/bin/bash
/src/scripts/initialize-osrm-backend.sh
/usr/bin/supervisord
/src/scripts/start-routing-backend.sh
