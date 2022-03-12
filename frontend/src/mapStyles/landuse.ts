import { LayerSpecification } from 'maplibre-gl';

const park_base: LayerSpecification = {
  id: 'park',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_park',
  paint: {
    'fill-color': '#d8e8c8',
    'fill-opacity': 0.7,
    'fill-outline-color': 'rgba(95, 208, 100, 1)'
  }
};

const park_outline: LayerSpecification = {
  id: 'park_outline',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'landuse_park',
  paint: {
    'line-dasharray': [1, 1.5],
    'line-color': 'rgba(228, 241, 215, 1)'
  }
};

const residential: LayerSpecification = {
  id: 'landuse_residential',
  type: 'fill',
  source: 'osm_cqrs',
  maxzoom: 8,
  'source-layer': 'landuse_residential',
  paint: {
    'fill-color': {
      base: 1,
      stops: [
        [9, 'hsla(0, 3%, 85%, 0.84)'],
        [12, 'hsla(35, 57%, 88%, 0.49)']
      ]
    }
  }
};

const wood_z_2_7: LayerSpecification = {
  id: 'landcover_wood_z_2_7',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_wood_z_2_7',
  paint: {
    'fill-antialias': false,
    'fill-color': 'hsla(98, 61%, 72%, 0.7)',
    'fill-opacity': 0.4
  },
  minzoom: 2,
  maxzoom: 7
};

const wood_z_7_9: LayerSpecification = {
  id: 'landcover_wood_z_7_9',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_wood_z_7_9',
  paint: {
    'fill-antialias': false,
    'fill-color': 'hsla(98, 61%, 72%, 0.7)',
    'fill-opacity': 0.4
  },
  minzoom: 7,
  maxzoom: 9
};

const wood_z_9_13: LayerSpecification = {
  id: 'landcover_wood_z_9_13',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_wood_z_9_13',
  paint: {
    'fill-antialias': false,
    'fill-color': 'hsla(98, 61%, 72%, 0.7)',
    'fill-opacity': 0.4
  },
  minzoom: 9,
  maxzoom: 13
};

const wood_z_13_22: LayerSpecification = {
  id: 'landcover_wood_z_13_22',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_wood',
  paint: {
    'fill-antialias': false,
    'fill-color': 'hsla(98, 61%, 72%, 0.7)',
    'fill-opacity': 0.4
  },
  minzoom: 13,
  maxzoom: 22
};

const grass: LayerSpecification = {
  id: 'landcover_grass',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_grass',
  paint: {
    'fill-antialias': false,
    'fill-color': 'rgba(176, 213, 154, 1)',
    'fill-opacity': 0.3
  }
};

const cementry: LayerSpecification = {
  id: 'landuse_cemetery',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_cemetery',
  paint: { 'fill-color': 'hsl(75, 37%, 81%)' }
};

const hospital: LayerSpecification = {
  id: 'landuse_hospital',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_hospital',
  paint: { 'fill-color': '#fde' }
};

const school: LayerSpecification = {
  id: 'landuse_school',
  type: 'fill',
  source: 'osm_cqrs',
  'source-layer': 'landuse_school',
  paint: { 'fill-color': 'rgb(236,238,204)' }
};

export const landuse = {
  park_base,
  park_outline,
  residential,
  wood_z_2_7,
  wood_z_7_9,
  wood_z_9_13,
  wood_z_13_22,
  grass,
  cementry,
  hospital,
  school
};
