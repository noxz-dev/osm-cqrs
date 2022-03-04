import { LayerSpecification } from 'maplibre-gl';

const river: LayerSpecification = {
  id: 'waterway_river',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'water_river',
  layout: { 'line-cap': 'round' },
  paint: {
    'line-color': '#a0c8f0',
    'line-width': {
      base: 1.2,
      stops: [
        [11, 0],
        [20, 6]
      ]
    }
  }
};

const other: LayerSpecification = {
  id: 'waterway_other',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'water_other',
  layout: { 'line-cap': 'round' },
  paint: {
    'line-color': '#a0c8f0',
    'line-width': {
      base: 1.3,
      stops: [
        [8, 0],
        [13, 0.5],
        [20, 6]
      ]
    }
  }
};

const areas: LayerSpecification = {
  id: 'water',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'water_areas',
  paint: { 'fill-color': 'rgb(158,189,255)' }
};

export const water = {
  river,
  other,
  areas
};
