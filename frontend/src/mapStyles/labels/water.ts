import { LayerSpecification } from 'maplibre-gl';

const areas_z13: LayerSpecification = {
  minzoom: 13,
  layout: {
    'text-field': '{name}',
    'text-font': ['Open Sans Italic'],
    'text-padding': 2,
    'text-allow-overlap': false,
    'text-size': {
      stops: [
        [15, 14],
        [20, 24]
      ]
    }
  },
  maxzoom: 24,
  type: 'symbol',
  source: 'osm_cqrs',
  id: 'water_areaslabels_z13',
  paint: {
    'text-color': 'rgba(68, 136, 136, 1)',
    'text-halo-width': 1,
    'text-halo-color': 'rgba(178, 220, 220, 1)'
  },
  'source-layer': 'water_areas'
};

const ways_z13: LayerSpecification = {
  minzoom: 13,
  layout: {
    'text-field': '{name}',
    'text-font': ['Open Sans Italic'],
    'text-padding': 2,
    'text-allow-overlap': false,
    'symbol-placement': 'line',
    'text-rotation-alignment': 'auto',
    'text-pitch-alignment': 'auto',
    'text-size': {
      stops: [
        [15, 14],
        [20, 24]
      ]
    }
  },
  maxzoom: 24,
  type: 'symbol',
  source: 'osm_cqrs',
  id: 'water_wayslabels_z13',
  paint: {
    'text-color': 'rgba(68, 136, 136, 1)',
    'text-halo-width': 1,
    'text-halo-color': 'rgba(178, 220, 220, 1)'
  },
  'source-layer': 'water_river'
};

const others_z13: LayerSpecification = {
  minzoom: 13,
  layout: {
    'text-field': '{name}',
    'text-font': ['Open Sans Italic'],
    'text-padding': 2,
    'text-allow-overlap': false,
    'symbol-placement': 'line',
    'text-rotation-alignment': 'auto',
    'text-pitch-alignment': 'auto',
    'text-size': {
      stops: [
        [15, 14],
        [20, 24]
      ]
    }
  },
  maxzoom: 24,
  type: 'symbol',
  source: 'osm_cqrs',
  id: 'water_ways_others_labels_z13',
  paint: {
    'text-color': 'rgba(68, 136, 136, 1)',
    'text-halo-width': 1,
    'text-halo-color': 'rgba(178, 220, 220, 1)'
  },
  'source-layer': 'water_other'
};

export const water_labels = {
  areas_z13,
  ways_z13,
  others_z13
};
