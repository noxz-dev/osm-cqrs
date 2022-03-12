import { LayerSpecification } from 'maplibre-gl';

const amenity: LayerSpecification = {
  minzoom: 14,
  layout: {
    'icon-image': '{type}-12',
    visibility: 'visible',
    'text-field': '{name}',
    'text-size': 8,
    'text-anchor': 'top',
    'text-offset': [0, 1]
  },
  maxzoom: 16,
  filter: [
    'all',
    [
      'in',
      'type',
      'fire_station',
      'bank',
      'border_control',
      'embassy',
      'government',
      'hospital',
      'police',
      'school',
      'taxi',
      'townhall',
      'university'
    ]
  ],
  type: 'symbol',
  source: 'osm',
  id: 'points_of_interest_fromareasz14',
  paint: {
    'text-color': 'rgba(108, 132, 137, 1)',
    'text-halo-color': 'rgba(255, 255, 255, 1)',
    'text-halo-width': 0.5,
    'text-halo-blur': 1
  },
  'source-layer': 'amenities'
};
