import { LayerSpecification } from 'maplibre-gl';

const area: LayerSpecification = {
  id: 'building',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'buildings',
  minzoom: 13,
  paint: {
    'fill-color': 'hsl(35, 8%, 85%)',
    'fill-outline-color': {
      default: 'hsla(35, 6%, 79%, 0.32)',
      property: 'fill-outline-color',
      type: 'exponential',
      stops: [
        [13, 'hsla(35, 6%, 79%, 0.32)'],
        [14, 'hsl(35, 6%, 79%)']
      ]
    }
  }
};

const threeD: LayerSpecification = {
  id: 'building-3d',
  type: 'fill-extrusion',
  source: 'osm_cqrs',
  'source-layer': 'buildings',
  minzoom: 14,
  paint: {
    'fill-extrusion-color': 'hsl(35, 8%, 85%)',
    'fill-extrusion-height': {
      property: 'render_height',
      type: 'identity'
    },
    'fill-extrusion-base': {
      property: 'render_min_height',
      type: 'identity'
    },
    'fill-extrusion-opacity': 0.8
  }
};

export const building = {
  area,
  threeD
};
