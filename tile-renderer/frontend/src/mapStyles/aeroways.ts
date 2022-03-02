import { LayerSpecification } from 'maplibre-gl';

const fill: LayerSpecification = {
  id: 'aeroway_fill',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  minzoom: 11,
  filter: ['==', 'type', 'runway'],
  paint: { 'fill-color': 'rgba(229, 228, 224, 1)', 'fill-opacity': 0.7 }
};

const runway: LayerSpecification = {
  id: 'aeroway_runway',
  type: 'line',
  source: 'martin',
  'source-layer': 'public.osm_aeroways',
  minzoom: 11,
  filter: ['all', ['==', 'type', 'runway']],
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
  source: 'martin',
  'source-layer': 'public.osm_aeroways',
  minzoom: 11,
  filter: ['all', ['==', 'type', 'taxiway']],
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
