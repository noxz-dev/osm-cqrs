import { LayerSpecification } from 'maplibre-gl';

const motorway_link_casing: LayerSpecification = {
  id: 'tunnel_motorway_link_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_motorways_link',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#e9ac77',
    'line-dasharray': [0.5, 0.25],
    'line-width': {
      base: 1.2,
      stops: [
        [12, 1],
        [13, 3],
        [14, 4],
        [20, 15]
      ]
    }
  }
};

const service_track_casing: LayerSpecification = {
  id: 'tunnel_service_track_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_service_track',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#cfcdca',
    'line-dasharray': [0.5, 0.25],
    'line-width': {
      base: 1.2,
      stops: [
        [15, 1],
        [16, 4],
        [20, 11]
      ]
    }
  }
};

const link_casing: LayerSpecification = {
  id: 'tunnel_link_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_link',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#e9ac77',
    'line-width': {
      base: 1.2,
      stops: [
        [12, 1],
        [13, 3],
        [14, 4],
        [20, 15]
      ]
    }
  }
};

const seconday_tetiary_casing: LayerSpecification = {
  id: 'tunnel_secondary_tertiary_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_secondary_tertiary',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#e9ac77',
    'line-width': {
      base: 1.2,
      stops: [
        [8, 1.5],
        [20, 17]
      ]
    }
  }
};

const primary_casing: LayerSpecification = {
  id: 'tunnel_trunk_primary_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_trunk_primary',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#e9ac77',
    'line-width': {
      base: 1.2,
      stops: [
        [5, 0.4],
        [6, 0.7],
        [7, 1.75],
        [20, 22]
      ]
    }
  }
};

const motorway_casing: LayerSpecification = {
  id: 'tunnel_motorway_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_motorways',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#e9ac77',
    'line-dasharray': [0.5, 0.25],
    'line-width': {
      base: 1.2,
      stops: [
        [5, 0.4],
        [6, 0.7],
        [7, 1.75],
        [20, 22]
      ]
    }
  }
};

const path_pedestrian: LayerSpecification = {
  id: 'tunnel_path_pedestrian',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_pedestrian',
  paint: {
    'line-color': 'hsl(0, 0%, 100%)',
    'line-dasharray': [1, 0.75],
    'line-width': {
      base: 1.2,
      stops: [
        [14, 0.5],
        [20, 10]
      ]
    }
  }
};

const motorway_link: LayerSpecification = {
  id: 'tunnel_motorway_link',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_motorways_link',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fc8',
    'line-width': {
      base: 1.2,
      stops: [
        [12.5, 0],
        [13, 1.5],
        [14, 2.5],
        [20, 11.5]
      ]
    }
  }
};

const service_track: LayerSpecification = {
  id: 'tunnel_service_track',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_service_track',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fff',
    'line-width': {
      base: 1.2,
      stops: [
        [15.5, 0],
        [16, 2],
        [20, 7.5]
      ]
    }
  }
};

const link: LayerSpecification = {
  id: 'tunnel_link',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_link',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fff4c6',
    'line-width': {
      base: 1.2,
      stops: [
        [12.5, 0],
        [13, 1.5],
        [14, 2.5],
        [20, 11.5]
      ]
    }
  }
};

const residental: LayerSpecification = {
  id: 'tunnel_minor',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_residential',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fff',
    'line-width': {
      base: 1.2,
      stops: [
        [13.5, 0],
        [14, 2.5],
        [20, 11.5]
      ]
    }
  }
};

const secondary_tertiary: LayerSpecification = {
  id: 'tunnel_secondary_tertiary',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_secondary_tertiary',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fff4c6',
    'line-width': {
      base: 1.2,
      stops: [
        [6.5, 0],
        [7, 0.5],
        [20, 10]
      ]
    }
  }
};

const trunk_primary: LayerSpecification = {
  id: 'tunnel_trunk_primary',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_trunk_primary',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fff4c6',
    'line-width': {
      base: 1.2,
      stops: [
        [5, 0],
        [7, 1],
        [20, 18]
      ]
    }
  }
};

const motorway: LayerSpecification = {
  id: 'tunnel_motorway',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_motorways',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#ffdaa6',
    'line-width': {
      base: 1.2,
      stops: [
        [5, 0],
        [7, 1],
        [20, 18]
      ]
    }
  }
};

const major_rail: LayerSpecification = {
  id: 'tunnel_major_rail',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_rail',
  paint: {
    'line-color': '#bbb',
    'line-width': {
      base: 1.4,
      stops: [
        [14, 0.4],
        [15, 0.75],
        [20, 2]
      ]
    }
  }
};

const major_rail_hatching: LayerSpecification = {
  id: 'tunnel_major_rail_hatching',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'tunnel_rail',
  paint: {
    'line-color': '#bbb',
    'line-dasharray': [0.2, 8],
    'line-width': {
      base: 1.4,
      stops: [
        [14.5, 0],
        [15, 3],
        [20, 8]
      ]
    }
  }
};

export const tunnel = {
  motorway_link_casing,
  service_track_casing,
  link_casing,
  seconday_tetiary_casing,
  primary_casing,
  motorway_casing,
  path_pedestrian,
  motorway_link,
  service_track,
  link,
  residental,
  secondary_tertiary,
  trunk_primary,
  motorway,
  major_rail,
  major_rail_hatching
};
