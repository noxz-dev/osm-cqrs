import { LayerSpecification } from 'maplibre-gl';

const motorway_link_casing: LayerSpecification = {
  id: 'bridge_motorway_link_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_motorways_link',
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

const service_track_casing: LayerSpecification = {
  id: 'bridge_service_track_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_service_track',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#cfcdca',
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

const path_pedestrian_casing: LayerSpecification = {
  id: 'bridge_path_pedestrian_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_pedestrian',
  paint: {
    'line-color': 'hsl(35, 6%, 80%)',
    'line-dasharray': [1, 0],
    'line-width': {
      base: 1.2,
      stops: [
        [14, 1.5],
        [20, 18]
      ]
    }
  }
};

const secondary_tertiary_casing: LayerSpecification = {
  id: 'bridge_secondary_tertiary_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_secondary_tertiary',
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

const trunk_primary_casing: LayerSpecification = {
  id: 'bridge_trunk_primary_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_trunk_primary',
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
  id: 'bridge_motorway_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_motorways',
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

const motorway_link: LayerSpecification = {
  id: 'bridge_motorway_link',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_motorways_link',
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
  id: 'bridge_service_track',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_service_track',
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

const secondary_tertiary: LayerSpecification = {
  id: 'bridge_secondary_tertiary',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_secondary_tertiary',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fea',
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
  id: 'bridge_trunk_primary',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_trunk_primary',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fea',
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
  id: 'bridge_motorway',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_motorways',
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': '#fc8',
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
  id: 'bridge_major_rail',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_rail',
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
  id: 'bridge_major_rail_hatching',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'bridge_rail',
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

export const bridges = {
  motorway_link_casing,
  service_track_casing,
  path_pedestrian_casing,
  secondary_tertiary_casing,
  trunk_primary_casing,
  motorway_casing,
  motorway_link,
  service_track,
  secondary_tertiary,
  trunk_primary,
  motorway,
  major_rail,
  major_rail_hatching
};
