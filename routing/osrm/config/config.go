package config

import (
	"github.com/withmandala/go-log"
	"noxz.dev/routing/osrm/statistics"
	"os"
)

const (
	MapDir  = "/src/data/map/"
	OsrmDir = "/src/data/osrm/"
	DataDir = "/src/data/"
)

const (
	MapUpdateDuration           = "duration of map update (in ms)"
	MapSize                     = "size of map (in bytes)"
	WriteDuration               = "duration of writing zip file to disk (in ms)"
	SizeOfIncomingEvent         = "size of incoming event (in bytes)"
	RoutingServerUpdateDuration = "duration of routing servers update (in ms)"
	CollectStatistics           = true
)

var MapUpdateStat = statistics.NewStatistic(
	"/src/data/statistics/map.statistics.csv",
	MapUpdateDuration,
	MapSize,
	WriteDuration,
	SizeOfIncomingEvent,
)

var RoutingUpdateStat = statistics.NewStatistic(
	"/src/data/statistics/routing-server.statistics.csv",
	RoutingServerUpdateDuration,
)

var Logger = log.New(os.Stderr)

func LogIfFailing(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
