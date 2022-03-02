<script setup lang="ts">
import { Map } from 'maplibre-gl';
import { onMounted } from 'vue';
import { aeroway } from '../mapStyles/aeroways';
import { boundaries } from '../mapStyles/boundaries';
import { bridges } from '../mapStyles/bridges';
import { building } from '../mapStyles/building';
import { landuse } from '../mapStyles/landuse';
import { roads } from '../mapStyles/roads';
import { tunnel } from '../mapStyles/tunnels';
import { water } from '../mapStyles/water';
onMounted(() => {
  const map = new Map({
    container: 'map',
    center: [9.7255, 52.36643],
    zoom: 15,
    style: {
      version: 8,
      sources: {
        martin: {
          type: 'vector',
          url: 'http://localhost:3000/public.osm_waterways,public.osm_landusages,public.osm_admin,public.osm_buildings,public.osm_roads,public.osm_aeroways,public.osm_waterways,public.osm_waterareas.json'
        }
      },
      layers: [
        {
          id: 'background',
          type: 'background',
          paint: { 'background-color': 'rgb(239,239,239)' }
        },
        landuse.park_base,
        landuse.park_outline,
        landuse.residential,
        landuse.wood,
        landuse.grass,
        landuse.cementry,
        landuse.hospital,
        landuse.school,

        water.river,
        water.other,
        water.areas,

        aeroway.fill,
        aeroway.runway,
        aeroway.taxiway,

        tunnel.motorway_link_casing,
        tunnel.service_track_casing,
        tunnel.link_casing,
        tunnel.street_casing,
        tunnel.seconday_tetiary_casing,
        tunnel.motorway_casing,
        tunnel.path_pedestrian,
        tunnel.motorway_link,
        tunnel.service_track,
        tunnel.link,
        tunnel.residental,
        tunnel.secondary_tertiary,
        tunnel.trunk_primary,
        tunnel.motorway,
        tunnel.major_rail,
        tunnel.major_rail_hatching,

        roads.residential_base,
        roads.residential_casting,
        roads.motorway_link_casing,
        roads.service_track_casing,
        roads.link_casing,
        roads.secondary_tertiary_casing,
        roads.trunk_primary_casing,
        roads.motorway_casing,
        roads.path_pedestrian,
        roads.motorway_link,
        roads.service_track,
        roads.link,
        roads.seconday_tetiary,
        roads.trunk_primary,
        roads.motorway,
        roads.major_rail,
        roads.major_rail_hatching,

        bridges.motorway_link_casing,
        bridges.service_track_casing,
        bridges.street_casing,
        bridges.path_pedestrian_casing,
        bridges.secondary_tertiary_casing,
        bridges.trunk_primary_casing,
        bridges.motorway_casing,
        bridges.motorway_link,
        bridges.service_track,
        bridges.street,
        bridges.secondary_tertiary,
        bridges.trunk_primary,
        bridges.motorway,
        bridges.major_rail,
        bridges.major_rail_hatching,

        building.area,

        boundaries._3,
        boundaries._2_z0_4,
        boundaries._2_z5_
      ]
    }
  });
});
</script>

<template>
  <div class="map-wrap">
    <div id="map"></div>
  </div>
</template>

<style scoped>
@import 'maplibre-gl/dist/maplibre-gl.css';

.map-wrap {
  position: relative;
  width: 100%;
  height: 100vh;
}

#map {
  position: absolute;
  width: 100%;
  height: 100%;
}
</style>
