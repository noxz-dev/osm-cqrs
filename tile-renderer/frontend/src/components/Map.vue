<script setup lang="ts">
import { onMounted } from 'vue';
import { Map } from 'maplibre-gl';

// style: 'https://demotiles.maplibre.org/style.json', // stylesheet location,

onMounted(() => {
  const map = new Map({
    container: 'map',
    style: {
      version: 8,
      name: '',
      sources: {
        streets: {
          type: 'vector',
          url: 'http://localhost:3000/public.osm_roads.json'
        },
        landusages: {
          type: 'vector',
          url: 'http://localhost:3000/public.osm_landusages.json'
        }
      },
      layers: [
        {
          id: 'landuse',
          type: 'fill',
          source: 'landusages',
          'source-layer': 'public.osm_landusages',
          minzoom: 5,

          paint: {
            'fill-color': [
              'interpolate',
              ['linear'],
              ['zoom'],
              15,
              [
                'match',
                ['get', 'class'],
                'park',
                'hsl(99, 57%, 75%)',
                'airport',
                'hsl(230, 15%, 91%)',
                'cemetery',
                'hsl(81, 28%, 81%)',
                'glacier',
                'hsl(196, 60%, 85%)',
                'hospital',
                'hsl(340, 37%, 87%)',
                'pitch',
                'hsl(99, 58%, 70%)',
                'sand',
                'hsl(56, 47%, 87%)',
                'school',
                'hsl(50, 48%, 81%)',
                'hsl(35, 16%, 85%)'
              ],
              16,
              [
                'match',
                ['get', 'class'],
                'park',
                'hsl(99, 57%, 75%)',
                'airport',
                'hsl(230, 29%, 89%)',
                'cemetery',
                'hsl(81, 28%, 81%)',
                'glacier',
                'hsl(196, 60%, 85%)',
                'hospital',
                'hsl(340, 63%, 89%)',
                'pitch',
                'hsl(99, 58%, 70%)',
                'sand',
                'hsl(56, 47%, 87%)',
                'school',
                'hsl(50, 48%, 81%)',
                'hsl(35, 16%, 85%)'
              ]
            ],
            'fill-opacity': [
              'interpolate',
              ['linear'],
              ['zoom'],
              5,
              0,
              6,
              ['match', ['get', 'class'], 'glacier', 0.5, 1]
            ]
          }
        },
        {
          id: 'streets',
          source: 'streets',
          'source-layer': 'public.osm_roads',
          type: 'line',
          paint: {
            'line-color': '#000'
          }
        }
      ]
    },
    center: [0, 0],
    zoom: 0
  });

  console.log(map);
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
