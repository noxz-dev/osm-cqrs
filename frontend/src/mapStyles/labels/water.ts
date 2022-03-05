import { LayerSpecification } from 'maplibre-gl';

const areas_z15: LayerSpecification = {
  minzoom: 15,
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
  id: 'water_areaslabels_z15',
  paint: {
    'text-color': 'rgba(68, 136, 136, 1)',
    'text-halo-width': 1,
    'text-halo-color': 'rgba(178, 220, 220, 1)'
  },
  'source-layer': 'water_areas'
};

export const water_labels = {
  areas_z15
};
