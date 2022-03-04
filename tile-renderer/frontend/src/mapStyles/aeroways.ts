import { LayerSpecification } from 'maplibre-gl';

const fill: LayerSpecification = {
  id: 'aeroway_fill',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'aeroway_fill',
  minzoom: 11,
  paint: { 'fill-color': 'rgba(229, 228, 224, 1)', 'fill-opacity': 0.7 }
};

const runway: LayerSpecification = {
  id: 'aeroway_runway',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'aeroway_runway',
  minzoom: 11,
  paint: {
    'line-color': '#f0ede9',
    'line-width': {
      base: 1.2,
      stops: [
        [11, 3],
        [20, 16]
      ]
    }
  }
};

const taxiway: LayerSpecification = {
  id: 'aeroway_taxiway',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'aeroway_taxiway',
  minzoom: 11,
  paint: {
    'line-color': '#f0ede9',
    'line-width': {
      base: 1.2,
      stops: [
        [11, 0.5],
        [20, 6]
      ]
    }
  }
};

export const aeroway = {
  fill,
  runway,
  taxiway
};
