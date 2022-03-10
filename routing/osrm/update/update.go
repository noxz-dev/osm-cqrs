package update

import (
	"bytes"
	"noxz.dev/routing/osrm/config"
	"os"
	"os/exec"
	"time"

	"github.com/withmandala/go-log"
)

var logger = log.New(os.Stderr)

func RunLocalMapUpdate() {
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

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while updating the local map: %s", stderr.String())
		return
	}

	logger.Info(out.String())
	elapsed := time.Since(start)

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
	logger.Infof("Updated map in %s", elapsed)
}

func RunRoutingServersUpdate() {
	RunCarRoutingServerUpdate()
}

func RunCarRoutingServerUpdate() {
	filename := "map.osrm"
	RunOsrmExtract(&filename)
	RunOsrmPartition(&filename)
	RunOsrmCustomize(&filename)
	RestartOsrmServer(&filename)
}

func RunOsrmExtract(filename *string) {
	cmd := exec.Command("osrm-extract", config.OsrmDir+*filename)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Infof("Pre-process %s for OSRM", *filename)

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while running osrm-extract: %s", stderr.String())
		return
	}

	logger.Info(out.String())

	elapsed := time.Since(start)
	logger.Infof("Pre-processed done in %s", elapsed)
}

func RunOsrmPartition(filename *string) {
	cmd := exec.Command("osrm-partition", config.OsrmDir+*filename)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Infof("Partition %s for OSRM", *filename)

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while running osrm-partition: %s", stderr.String())
		return
	}

	logger.Info(out.String())

	elapsed := time.Since(start)
	logger.Infof("Partition done in %s", elapsed)
}

func RunOsrmCustomize(filename *string) {
	cmd := exec.Command("osrm-customize", config.OsrmDir+*filename)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Infof("Customize %s for OSRM", *filename)

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while running osrm-customize: %s", stderr.String())
		return
	}

	logger.Info(out.String())

	elapsed := time.Since(start)
	logger.Infof("Partition done in %s", elapsed)
}

func RestartOsrmServer(filename *string) {
	StopOsrmServer()

	cmd := exec.Command("osrm-routed", "--algorithm", "mld", config.OsrmDir+*filename)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Infof("Restarting OSRM server for %s ", *filename)

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while running osrm-routed: %s", stderr.String())
		return
	}

	logger.Info(out.String())

	elapsed := time.Since(start)
	logger.Infof("Partition done in %s", elapsed)
}

func StopOsrmServer() {
	cmd := exec.Command("kill", `ps -e | grep osrm | egrep -v grep | awk '{print $1}'`)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	start := time.Now()

	logger.Infof("Stopping OSRM server ")

	err := cmd.Run()

	if err != nil {
		logger.Errorf("Error occurred while stopping osrm-server: %s", stderr.String())
		return
	}

	logger.Info(out.String())

	elapsed := time.Since(start)
	logger.Infof("Stopping done in %s", elapsed)
}
