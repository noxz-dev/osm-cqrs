-- Teilweg von Im Born
SELECT ST_TRANSFORM(geometry, 4326), * FROM network WHERE id = 134916;

-- Source Node Im Born
SELECT ST_TRANSFORM(point, 4326), * FROM imposm2pgr.osm_ways_junctions WHERE id = 2651249;

SELECT * FROM network WHERE name = 'Ostafrikastraße';

-- Teilweg der Ostafrikastraße
SELECT ST_TRANSFORM(geometry, 4326), * FROM network WHERE id = 136970;

-- Source Node auf der Ostafrikastraße
SELECT ST_TRANSFORM(point, 4326), * FROM imposm2pgr.osm_ways_junctions WHERE id = 809603;

-- Berechne Pfad zwischen Osftafrikastraße und Im Born
SELECT
	ST_TRANSFORM(ST_Collect(network.geometry), 4326)  AS geometry, sum (path.cost) AS total_cost
FROM
	pgr_dijkstra('SELECT id, source_vertex_id as source, target_vertex_id as target, cost FROM network', 809603, 2651249, true)
	AS PATH
JOIN network ON network.id = path.edge;
