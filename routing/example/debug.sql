SELECT * FROM osm_ways WHERE osm_ways.tags -> 'name' = 'Ostafrikastraße';

SELECT COUNT(*) FROM osm_ways;

SELECT *, ST_Transform(geometry, 4326) FROM network WHERE name = 'Ostafrikastraße';
SELECT *, ST_Transform(geometry, 4326) FROM network WHERE name = 'Im Born';

SELECT COUNT(*) FROM network;

SELECT ST_Transform(geometry, 4326) FROM network;

SELECT * FROM imposm2pgr."osm_ways_junctions";

SELECT ST_Transform(geometry, 4326) FROM network WHERE target_vertex_id=35637;


SELECT * FROM network where source_vertex_id in (
	SELECT ST_Transform(point, 4326) FROM "imposm2pgr".osm_ways_junctions ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.674266 52.354086)'::geometry LIMIT 20
);



	SELECT ST_Transform(point, 4326) FROM "imposm2pgr".osm_ways_junctions ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.674266 52.354086)'::geometry LIMIT 20

SELECT "imposm2pgr".initialize_osm_ways_junctions(); 
SELECT "imposm2pgr".initialize_network();