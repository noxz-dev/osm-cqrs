SELECT * FROM osm_ways WHERE osm_ways.tags -> 'name' = 'Ostafrikastraße';

SELECT COUNT(*) FROM osm_ways;

SELECT *, ST_Transform(geometry, 4326) FROM network WHERE name = 'Ostafrikastraße';
SELECT *, ST_Transform(geometry, 4326) FROM network WHERE name = 'Im Born';

SELECT COUNT(*) FROM network;

SELECT ST_Transform(geometry, 4326) FROM network;
SELECT ST_Transform(geometry, 4326) FROM osm_ways;

SELECT * FROM imposm2pgr."osm_ways_junctions";

SELECT ST_Transform(geometry, 4326) FROM network WHERE target_vertex_id=35637;


SELECT * FROM network where source_vertex_id in (
	SELECT ST_Transform(point, 4326) FROM "imposm2pgr".osm_ways_junctions ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.674266 52.354086)'::geometry LIMIT 20
);



	SELECT ST_Transform(point, 4326) FROM "imposm2pgr".osm_ways_junctions ORDER BY ST_Transform(point, 4326) <-> 'SRID=4326;POINT(9.674266 52.354086)'::geometry LIMIT 20

SELECT "imposm2pgr".initialize_osm_ways_junctions(); 
SELECT "imposm2pgr".initialize_network();

SELECT * FROM "imposm2pgr".updates;

SELECT ST_Transform(point, 4326) FROM "imposm2pgr".osm_ways_junctions;

SHOW config_file;

SELECT * FROM public.osm_ways

INSERT INTO "public"."osm_ways" ("osm_id", "geometry", "tags") VALUES (1234, '0102000020110F0000070000004EA55DBB6C6F304144F22A8F772F5A413C0B62BB136F30411AD13103772F5A41BAE2B289DB6E3041A81DB6AB762F5A41262BFCDDA96E3041ED4A8F5D762F5A412A0F2B5A786E30410B79680F762F5A41AB4E179F266E30412FDAEE8D752F5A4186F9EF26CC6D304129C2F501752F5A41', '"name"=>"Ostafrikastraße", "highway"=>"residential", "maxspeed"=>"30", "sidewalk"=>"both"')

SELECT * FROM "imposm2pgr"."osm_ways_diff";

SELECT * FROM TEMP
