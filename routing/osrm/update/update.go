package update

import (
	"bytes"
	"github.com/withmandala/go-log"
	"noxz.dev/routing/osrm/config"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var logger = log.New(os.Stderr)

var isMapUpdating = false
var isRoutingUpdating = false

// RunLocalMapUpdate updates the local state of the OSM map with the current changeset
func RunLocalMapUpdate() {
	isMapUpdating = true
	osmium := "osmium"
	osmiumCommand := "apply-changes"
	currentMap := config.MapDir + "map.pbf"
	changeSet := config.DataDir + "change.osc.gz"
	outputFileOption := "-o"
	outputFile := config.MapDir + "map2.pbf"
	//overwrite := "--overwrite"

	cmd := exec.Command(osmium, osmiumCommand, currentMap, changeSet, outputFileOption, outputFile)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Info("Updating local map")

	config.LogIfFailing(config.MapUpdateStat.StartTimer(config.MapUpdateDuration))

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while updating the local map: %s", stderr.String())
		return
	}

	logger.Info(out.String())

	err = os.Remove(config.MapDir + "map.pbf")
	if err != nil {
		logger.Infof("Error while deleting map: %s", err.Error())
		return
	}

	err = os.Rename(config.MapDir+"map2.pbf", config.MapDir+"map.pbf")
	if err != nil {
		logger.Infof("Error while renaming map: %s", err.Error())
		return
	}

	elapsed := time.Since(start)
	logger.Infof("Updated map in %s", elapsed)
	config.LogIfFailing(config.MapUpdateStat.StopTimerAndSetDuration(config.MapUpdateDuration))

	if config.CollectStatistics {
		fi, err := os.Stat(config.MapDir + "map.pbf")
		if err != nil {
			config.Logger.Errorf("Error while loading file stats %s: ", err.Error())
		}
		config.LogIfFailing(config.MapUpdateStat.SetValue(config.MapSize, strconv.FormatInt(fi.Size(), 10)))
	}

	isMapUpdating = false
}

// RunRoutingServerUpdate updates the routing servers for all profiles
func RunRoutingServerUpdate() {
	for isRoutingUpdating {
		logger.Info("Routing database is already updating. Skipping...")
		return
	}

	for isMapUpdating {
		logger.Info("The map is currently updating. Waiting...")
		time.Sleep(1)
	}

	isRoutingUpdating = true

	cmd := exec.Command("/src/scripts/update-osrm-backend.sh")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Info("Updating routing database")

	config.LogIfFailing(config.RoutingUpdateStat.BeginnColum())

	config.LogIfFailing(config.RoutingUpdateStat.StartTimer(config.RoutingServerUpdateDuration))

	err := cmd.Run()

	config.LogIfFailing(config.RoutingUpdateStat.StopTimerAndSetDuration(config.RoutingServerUpdateDuration))

	if err != nil {
		logger.Errorf("Error occurred while updating the routing database: %s", stderr.String())
		return
	}

	logger.Info(out.String())
	elapsed := time.Since(start)
	logger.Infof("Updated routing database in %s", elapsed)

	config.LogIfFailing(config.RoutingUpdateStat.EndColum())

	isRoutingUpdating = false
}
