const buildings: LayerSpecification = {
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