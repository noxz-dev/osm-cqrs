<script setup lang="ts">
import { LngLatBounds, LngLatLike, Map, Marker } from 'maplibre-gl';
import { onMounted, ref, watch } from 'vue';
import { layers } from '../mapStyles/styles';
import { Coordinate } from '../model/Coordinate';
import { Route } from '../model/route';
import { mapStore } from '../store/map';

const map = ref<Map>();

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
let locationMarker: Marker = new Marker();

let resultMarkerList: Marker[] = [];

watch(
  () => mapStore.route,
  () => {
    resetMarkers();
    if (mapStore.route) {
      updateRoute();
    }
  }
);

watch(
  () => mapStore.selectedResult,
  () => {
    resetMarkers();

    if (mapStore.selectedResult) {
      zoomToLocation();
    }
  }
);

watch(
  () => mapStore.searchResults,
  () => {
    resetMarkers();

    if (mapStore.searchResults) {
      fitBoundsAndPlaceMarkers();
    }
  }
);

onMounted(async () => {
  map.value = new Map({
    container: 'map',
    center: [9.724391386801031, 52.353106612757465],
    zoom: 17,
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
        'line-color': '#9B59B6',
        'line-width': 5,
        'line-opacity': 0.7
      }
    });
  });
});

function updateRoute() {
  if (mapStore.route) {
    // const coordinates = mapRouteCoordinates(mapStore.route);
    geojsonRoute.geometry = mapStore.route;
    map.value?.getSource('route')?.setData(geojsonRoute);

    const coordinates = geojsonRoute.geometry.coordinates;
    map.value?.setLayoutProperty('route-line', 'visibility', 'visible');
    startMarker = new Marker({ color: 'rgb(46, 204, 113)' }).setLngLat(coordinates[0]).addTo(map.value!);
    endMarker = new Marker({ color: 'rgb(231, 76, 60)' })
      .setLngLat(coordinates[coordinates.length - 1])
      .addTo(map.value!);

    let bounds = coordinates.reduce(function (bounds, coord) {
      return bounds.extend(coord);
    }, new LngLatBounds(coordinates[0], coordinates[0]));

    map.value?.fitBounds(bounds, {
      padding: 200
    });
  }
}

function zoomToLocation() {
  const coordinate: [number, number] = [mapStore.selectedResult!.location.Lng, mapStore.selectedResult!.location.Lat];
  map.value?.flyTo({
    center: coordinate,
    zoom: 17
  });

  if (mapStore.selectedIndex != null && resultMarkerList.length > mapStore.selectedIndex) {
    resultMarkerList.forEach((m, index) => {
      const tempCoords = m._lngLat;
      m.remove();
      resultMarkerList[index] = new Marker({ color: '#3FB1CE' }).setLngLat(tempCoords).addTo(map.value!);
    });
    resultMarkerList[mapStore.selectedIndex].remove();
    resultMarkerList[mapStore.selectedIndex] = new Marker({ color: 'rgb(231, 76, 60)' })
      .setLngLat(coordinate)
      .addTo(map.value!);
  }
}

function mapRouteCoordinates(route: Route): LngLatLike[] {
  return route.routePoints.map((coordinate: Coordinate) => [coordinate.Lng, coordinate.Lat]);
}

function fitBoundsAndPlaceMarkers() {
  if (mapStore.searchResults.length === 0) return;

  const coordinates = mapStore.searchResults.map((result) => [
    result.location.Lng,
    result.location.Lat
  ]) as LngLatLike[];

  resultMarkerList.forEach((marker) => marker.remove());
  resultMarkerList = [];

  for (const coord of coordinates) {
    resultMarkerList.push(new Marker().setLngLat(coord).addTo(map.value!));
  }

  let bounds = coordinates.reduce((bounds, coord) => {
    return bounds.extend(coord);
  }, new LngLatBounds(coordinates[0], coordinates[0]));

  map.value?.fitBounds(bounds, {
    padding: 400
  });
}

function resetMarkers() {
  resultMarkerList.forEach((m) => {
    m.remove();
  });
  startMarker?.remove();
  endMarker?.remove();
  if (map.value?.getLayer('route-line')) {
    map.value?.setLayoutProperty('route-line', 'visibility', 'none');
  }
}
</script>

<template>
  <div class="absolute w-full h-screen">
    <div id="map" class="w-full h-full"></div>
  </div>
</template>

<style>
@import 'maplibre-gl/dist/maplibre-gl.css';
</style>
