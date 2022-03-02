import { LayerSpecification } from 'maplibre-gl';

const _3: LayerSpecification = {
  id: 'boundary_3',
  type: 'line',
  source: 'martin',
  'source-layer': 'public.osm_admin',
  minzoom: 8,
  filter: ['all', ['in', 'admin_level', 3, 4]],
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

const _2_z0_4: LayerSpecification = {
  id: 'boundary_2_z0-4',
  type: 'line',
  source: 'martin',
  'source-layer': 'public.osm_admin',
  maxzoom: 5,
  filter: ['all', ['==', 'admin_level', 2], ['!has', 'claimed_by']],
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

const _2_z5_: LayerSpecification = {
  id: 'boundary_2_z5-',
  type: 'line',
  source: 'martin',
  'source-layer': 'public.osm_admin',
  minzoom: 5,
  filter: ['all', ['==', 'admin_level', 2]],
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

export const boundaries = {
  _3,
  _2_z0_4,
  _2_z5_
};
