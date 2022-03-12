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
  // paint: {
  //   'fill-extrusion-color': 'hsl(35, 8%, 85%)',
  //   'fill-extrusion-height': {
  //     property: 'building_levels',
  //     type: 'identity'
  //   },
  //   'fill-extrusion-base': {
  //     property: 'building_min_levels',
  //     type: 'identity'
  //   },
  //   'fill-extrusion-opacity': 0.8
  // },

  paint: {
    'fill-extrusion-color': 'hsl(35, 8%, 85%)',

    // use an 'interpolate' expression to add a smooth transition effect to the
    // buildings as the user zooms in
    'fill-extrusion-height': ['interpolate', ['linear'], ['zoom'], 15, 0, 15.05, ['get', 'height']],
    'fill-extrusion-base': ['interpolate', ['linear'], ['zoom'], 15, 0, 15.05, ['get', 'min_height']],
    'fill-extrusion-opacity': 0.6
  }

  // paint: {
  //   'fill-extrusion-color': '#aaa',

  //   // use an 'interpolate' expression to add a smooth transition effect to the
  //   // buildings as the user zooms in
  //   'fill-extrusion-height': 12,
  //   'fill-extrusion-base': 0,
  //   'fill-extrusion-opacity': 0.6
  // }
};

export const building = {
  area,
  threeD
};
