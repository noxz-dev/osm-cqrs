import { LayerSpecification } from 'maplibre-gl';

const _1_2: LayerSpecification = {
  id: 'boundary_1_2',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'boundaries_1_2',
  minzoom: 5,
  layout: { 'line-cap': 'round', 'line-join': 'round' },
  paint: {
    'line-color': 'hsl(248, 1%, 41%)',
    'line-opacity': {
      base: 1,
      stops: [
        [0, 0.4],
        [4, 1]
      ]
    },
    'line-width': {
      base: 1,
      stops: [
        [3, 1],
        [5, 1.2],
        [12, 3]
      ]
    }
  }
};

const _3_4: LayerSpecification = {
  id: 'boundary_3_4',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'boundaries_3_4',
  minzoom: 8,
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#9e9cab',
    'line-dasharray': [5, 1],
    'line-width': {
      base: 1,
      stops: [
        [4, 0.4],
        [5, 1],
        [12, 1.8]
      ]
    }
  }
};

const _5_6: LayerSpecification = {
  id: 'boundary_5_6',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'boundaries_5_6',
  minzoom: 11,
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#BBB',
    'line-dasharray': [5, 1],
    'line-width': {
      base: 1,
      stops: [
        [4, 0.4],
        [5, 1],
        [12, 1.8]
      ]
    }
  }
};

const _7_8: LayerSpecification = {
  id: 'boundary_7_8',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'boundaries_7_8',
  minzoom: 11,
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#BBB',
    'line-dasharray': [5, 1],
    'line-width': {
      base: 1,
      stops: [
        [4, 0.4],
        [5, 1],
        [12, 1.8]
      ]
    }
  }
};

const _9_10: LayerSpecification = {
  id: 'boundary_9_10',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'boundaries_9_10',
  minzoom: 12,
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#BBB',
    'line-dasharray': [5, 1],
    'line-width': {
      base: 1,
      stops: [
        [4, 0.4],
        [5, 1],
        [12, 1.8]
      ]
    }
  }
};

export const boundaries = {
  _1_2,
  _3_4,
  _5_6,
  _7_8,
  _9_10
};
