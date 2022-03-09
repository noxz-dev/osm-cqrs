SELECT 'SRID=4326;POINT(9.73322 52.37052)'::geometry, 
ST_Transform(point, 4326), ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.73322 52.37052)'::geometry 
AS dist FROM "imposm2pgr".osm_ways_junctions ORDER BY dist LIMIT 3;


SELECT 'SRID=4326;POINT(9.73322 52.37052)'::geometry, 
ST_Transform(needed.point, 4326), ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.73322 52.37052)'::geometry 
AS dist FROM 
(
	SELECT * FROM "imposm2pgr".osm_ways_junctions as wj
	JOIN (
		SELECT source_vertex_id FROM network WHERE name LIKE 'Ostafrikastraße'
	) as net
	ON
	net.source_vertex_id = wj.id
) as needed
ORDER BY dist LIMIT 3;

SELECT 'SRID=4326;POINT(9.73322 52.37052)'::geometry, 
ST_Transform(needed.point, 4326), ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.73322 52.37052)'::geometry 
AS dist FROM "imposm2pgr".osm_ways_junctions
(
	SELECT * FROM "imposm2pgr".osm_ways_junctions as wj
	JOIN (
		SELECT source_vertex_id FROM network WHERE name LIKE 'Ostafrikastraße'
	) as net
	ON
	net.source_vertex_id = wj.id
) as needed
ORDER BY dist LIMIT 3;


SELECT 'SRID=4326;POINT(9.663847 52.359593)'::geometry, 
ST_Transform(needed.point, 4326), ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.663847 52.359593)'::geometry 
AS dist FROM 
(
	SELECT * FROM "imposm2pgr".osm_ways_junctions as wj
	JOIN (
		SELECT target_vertex_id FROM network WHERE name LIKE 'Im Born'
	) as net
	ON
	net.target_vertex_id = wj.id
) as needed
ORDER BY dist LIMIT 1;


SELECT 'SRID=4326;POINT(9.663847 52.359593)'::geometry, 
ST_Transform(point, 4326), ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.663847 52.359593)'::geometry 
AS dist FROM "imposm2pgr".osm_ways_junctions ORDER BY dist LIMIT 1;

SELECT  FROM "imposm2pgr".osm_ways_junctions as wj
JOIN (
 SELECT source_vertex_id FROM network WHERE name LIKE 'Ostafrikastraße'
) as net
ON
net.source_vertex_id = wj.id

SELECT * FROM "imposm2pgr".osm_ways_junctions as wj
JOIN (
 SELECT target_vertex_id FROM network WHERE name LIKE 'Im Born'
) as net
ON
net.target_vertex_id = wj.id

SELECT target_vertex_id FROM network WHERE name LIKE 'Im Born';

-- Queries needed for routing
SELECT
	ST_TRANSFORM(ST_Collect(network.geometry), 4326)  AS geometry, sum (path.cost) AS total_cost
FROM
	pgr_dijkstra('SELECT id, source_vertex_id as source, target_vertex_id as target, cost FROM network', 
				 (
				 	SELECT id FROM "imposm2pgr".osm_ways_junctions 
				  	ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.674266 52.354086)'::geometry 
				  	LIMIT 1
				 )
				 , 
				 (
				  	SELECT id FROM "imposm2pgr".osm_ways_junctions 
				  	ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.663847 52.359593)'::geometry 
				  	LIMIT 1
				 ), 
				 true
				)
	AS PATH
JOIN network ON network.id = path.edge;

SELECT id FROM "imposm2pgr".osm_ways_junctions 
ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.674266 52.354086)'::geometry 
LIMIT 1;

SELECT id 
FROM "imposm2pgr".osm_ways_junctions 
ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.663847 52.359593)'::geometry 
LIMIT 1;


