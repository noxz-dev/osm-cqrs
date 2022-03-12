import { LayerSpecification } from 'maplibre-gl';

const z14_motorway: LayerSpecification = {
  id: 'roadlabels_z14_motorways',
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

const z14_trunk: LayerSpecification = {
  id: 'roadlabels_z14_trunk_primary',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_trunk_primary',
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

const z14_residential: LayerSpecification = {
  id: 'roadlabels_z14_residential',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_residential',
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

const z14_service_track: LayerSpecification = {
  id: 'roadlabels_z14_service_track',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_service_track',
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

const z14_roads_secondary_tertiary: LayerSpecification = {
  id: 'roadlabels_z14_roads_secondary_tertiary',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_secondary_tertiary',
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

const z14_roads_trunk_primary: LayerSpecification = {
  id: 'roadlabels_z14_roads_trunk_primary',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_trunk_primary',
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

const z14_roads_pedestrian: LayerSpecification = {
  id: 'roadlabels_z14_roads_pedestrian',
  type: 'symbol',
  source: 'osm_cqrs',
  'source-layer': 'roads_pedestrian',
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
  z14_motorway,
  z14_trunk,
  z14_residential,
  z14_service_track,
  z14_roads_secondary_tertiary,
  z14_roads_trunk_primary,
  z14_roads_pedestrian
};
