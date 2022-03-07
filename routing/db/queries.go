package db

import (
	"context"
	"fmt"
)

func CountRowsInNetwork() (int64, error) {
	var rowNum int64

	conn, err := GetConnection().Begin(context.Background())

	if err != nil {
		return -1, err
	}

	err = conn.QueryRow(context.Background(), "select count(*) from network").Scan(&rowNum)

	if err != nil {
		return -1, err
	}

	return rowNum, nil
}

type Point struct {
	Lng float64
	Lat float64
}

type Route = []Point

func GetRouteDijkstra(fromLat float64, fromLng float64, toLat float64, toLng float64) (Route, error) {

	route := make(Route, 0)

	conn, err := GetConnection().Begin(context.Background())

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(context.Background(), fmt.Sprintf(`SELECT ST_X(g), ST_Y(g) FROM (SELECT (ST_DumpPoints(ST_TRANSFORM(ST_Collect(network.geometry), 4326))).geom as g FROM pgr_dijkstra('SELECT id, source_vertex_id as source, target_vertex_id as target, cost FROM network', (SELECT id FROM "imposm2pgr".osm_ways_junctions ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(%f %f)'::geometry LIMIT 1), (SELECT id FROM "imposm2pgr".osm_ways_junctions ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(%f %f)'::geometry LIMIT 1), true) AS PATH JOIN network ON network.id = path.edge) as route`, fromLng, fromLat, toLng, toLat))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var lat float64
		var lng float64

		err := rows.Scan(&lng, &lat)
		if err != nil {
			return nil, err
		}
		//route = append(route, Point{Lat: lat, Lng: lng})
		route = append(route, Point{lng, lat})
	}

	return route, nil

}
