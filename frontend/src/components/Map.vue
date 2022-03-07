<script setup lang="ts">
import { LngLatBounds, LngLatLike, Map, Marker } from 'maplibre-gl';
import { onMounted, PropType, ref, watch } from 'vue';
import { layers } from '../mapStyles/styles';
import { Coordinate, Route } from '../model/route';

const map = ref<Map>();

const props = defineProps({
  route: Object as PropType<Route>
});

const geojsonRoute = {
  type: 'Feature',
  properties: {},
  geometry: {
    type: 'LineString',
    coordinates: [[0, 0]]
  }
};

let startMarker: Marker = new Marker();
let endMarker: Marker = new Marker();

watch(
  () => props.route,
  () => {
    if (props.route) {
      updateRoute();
    }
  }
);

onMounted(async () => {
  map.value = new Map({
    container: 'map',
    center: [9.7255, 52.36643],
    zoom: 15,
    style: {
      version: 8,
      sources: {
        osm_cqrs: {
          type: 'vector',
          url: 'http://localhost:8080/capabilities/osm_cqrs.json'
        }
      },
      sprite: 'https://go-spatial.github.io/carto-assets/spritesets/osm_tegola_spritesheet',
      glyphs: 'https://go-spatial.github.io/carto-assets/fonts/{fontstack}/{range}.pbf',
      layers: layers
    }
  });

  map.value.on('load', () => {
    map.value?.addSource('route', {
      type: 'geojson',
      data: geojsonRoute
    });

    map.value?.addLayer({
      id: 'route-line',
      type: 'line',
      source: 'route',
      layout: {
        'line-cap': 'round',
        'line-join': 'round'
      },
      paint: {
        'line-color': '#3498DB',
        'line-width': 5,
        'line-opacity': 0.7
      }
    });
  });
});

function updateRoute() {
  if (props.route) {
    const coordinates = mapRouteCoordinates(props.route);
    geojsonRoute.geometry.coordinates = coordinates as number[][];
    map.value?.getSource('route')?.setData(geojsonRoute);

    startMarker = new Marker().setLngLat(coordinates[0]).addTo(map.value!);
    endMarker = new Marker().setLngLat(coordinates[coordinates.length - 1]).addTo(map.value!);

    let bounds = coordinates.reduce(function (bounds, coord) {
      return bounds.extend(coord);
    }, new LngLatBounds(coordinates[0], coordinates[0]));

    map.value?.fitBounds(bounds, {
      padding: 30
    });
  }
}

function mapRouteCoordinates(route: Route): LngLatLike[] {
  return route.routePoints.map((coordinate: Coordinate) => [coordinate.Lng, coordinate.Lat]);
}
</script>

<template>
  <div class="relative w-full h-screen">
    <div id="map" class="absolute w-full h-full"></div>
  </div>
</template>

<style scoped>
@import 'maplibre-gl/dist/maplibre-gl.css';
</style>
