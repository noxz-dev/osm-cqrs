import { LayerSpecification } from 'maplibre-gl';

const z14: LayerSpecification = {
  id: 'roadlabels_z14',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_motorways',
  layout: {
    'text-size': {
      stops: [
        [13, 10],
        [20, 18]
      ]
    },
    'text-allow-overlap': false,
    'symbol-avoid-edges': false,
    'symbol-spacing': 250,
    'text-font': ['Open Sans Regular'],
    'symbol-placement': 'line',
    'text-padding': 2,
    'text-rotation-alignment': 'auto',
    'text-pitch-alignment': 'auto',
    'text-field': '{name}'
  },
  paint: {
    'text-color': 'rgba(82, 82, 82, 1)',
    'text-halo-width': 1,
    'text-halo-color': 'rgba(255, 255, 255, 0.8)'
  }
};

export const road_labels = {
  z14
};
