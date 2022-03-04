import { LayerSpecification } from 'maplibre-gl';

const park: LayerSpecification = {
  minzoom: 14,
  layout: {
    'text-field': '{name}',
    'text-size': 11
  },
  maxzoom: 24,
  type: 'symbol',
  source: 'osm_cqrs',
  id: 'landuse_areaslabels_park',
  paint: {
    'text-color': 'rgba(122, 143, 61, 1)',
    'text-halo-color': 'rgba(228, 235, 209, 1)',
    'text-halo-width': 1
  },
  'source-layer': 'amenities'
};

export const amenities = {
  park
};
