import { LayerSpecification } from 'maplibre-gl';

const park_base: LayerSpecification = {
  id: 'park',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['all', ['==', 'type', 'park']],
  paint: {
    'fill-color': '#d8e8c8',
    'fill-opacity': 0.7,
    'fill-outline-color': 'rgba(95, 208, 100, 1)'
  }
};

const park_outline: LayerSpecification = {
  id: 'park_outline',
  type: 'line',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['all', ['==', 'type', 'park']],
  paint: {
    'line-dasharray': [1, 1.5],
    'line-color': 'rgba(228, 241, 215, 1)'
  }
};

const residential: LayerSpecification = {
  id: 'landuse_residential',
  type: 'fill',
  source: 'martin',
  maxzoom: 8,
  'source-layer': 'public.osm_landusages',
  filter: ['==', 'type', 'residential'],
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

const wood: LayerSpecification = {
  id: 'landcover_wood',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['all', ['in', 'type', 'wood', 'forest']],
  paint: {
    'fill-antialias': false,
    'fill-color': 'hsla(98, 61%, 72%, 0.7)',
    'fill-opacity': 0.4
  }
};

const grass: LayerSpecification = {
  id: 'landcover_grass',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['all', ['in', 'type', 'grass', 'garden', 'meadow', 'pitch', 'common', 'recreation_ground']],
  paint: {
    'fill-antialias': false,
    'fill-color': 'rgba(176, 213, 154, 1)',
    'fill-opacity': 0.3
  }
};

const cementry: LayerSpecification = {
  id: 'landuse_cemetery',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['==', 'type', 'cemetery'],
  paint: { 'fill-color': 'hsl(75, 37%, 81%)' }
};

const hospital: LayerSpecification = {
  id: 'landuse_hospital',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['==', 'type', 'hospital'],
  paint: { 'fill-color': '#fde' }
};

const school: LayerSpecification = {
  id: 'landuse_school',
  type: 'fill',
  source: 'martin',
  'source-layer': 'public.osm_landusages',
  filter: ['==', 'type', 'school'],
  paint: { 'fill-color': 'rgb(236,238,204)' }
};

export const landuse = {
  park_base,
  park_outline,
  residential,
  wood,
  grass,
  cementry,
  hospital,
  school
};
