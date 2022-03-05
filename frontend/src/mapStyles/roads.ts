import { LayerSpecification } from 'maplibre-gl';

const residential_casting: LayerSpecification = {
  id: 'road_minor_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_residential',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
  paint: {
    'line-color': '#cfcdca',
    'line-opacity': {
      stops: [
        [12, 0],
        [12.5, 1]
      ]
    },
    'line-width': {
      base: 1.2,
      stops: [
        [12, 0.5],
        [13, 1],
        [14, 4],
        [20, 20]
      ]
    }
  }
};

const residential_base: LayerSpecification = {
  id: 'road_minor',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_residential',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
  paint: {
    'line-color': '#fff',
    'line-width': {
      base: 1.2,
      stops: [
        [13.5, 0],
        [14, 2.5],
        [20, 18]
      ]
    }
  }
};

const motorway_link_casing: LayerSpecification = {
  id: 'road_motorway_link_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_motorways_link',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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

const link_casing: LayerSpecification = {
  id: 'road_link_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_link',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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
  id: 'road_service_track_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_service_track',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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

const secondary_tertiary_casing: LayerSpecification = {
  id: 'road_secondary_tertiary_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_secondary_tertiary',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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
  id: 'road_trunk_primary_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_trunk_primary',
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
  id: 'road_motorway_casing',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_motorways',
  minzoom: 5,
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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

const path_pedestrian: LayerSpecification = {
  id: 'road_path_pedestrian',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_pedestrian',
  minzoom: 14,
  layout: { 'line-join': 'round' },
  paint: {
    'line-color': 'hsl(0, 0%, 100%)',
    'line-dasharray': [1, 0.7],
    'line-width': {
      base: 1.2,
      stops: [
        [14, 1],
        [20, 10]
      ]
    }
  }
};

const motorway_link: LayerSpecification = {
  id: 'road_motorway_link',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_motorways_link',
  minzoom: 12,
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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
  id: 'road_service_track',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_service_track',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
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
  id: 'road_link',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_link',
  minzoom: 13,
  layout: { 'line-cap': 'round', 'line-join': 'round' },
  paint: {
    'line-color': '#fea',
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

const seconday_tetiary: LayerSpecification = {
  id: 'road_secondary_tertiary',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_secondary_tertiary',
  layout: { 'line-cap': 'round', 'line-join': 'round' },
  paint: {
    'line-color': '#fea',
    'line-width': {
      base: 1.2,
      stops: [
        [6.5, 0],
        [8, 0.5],
        [20, 13]
      ]
    }
  }
};

const trunk_primary: LayerSpecification = {
  id: 'road_trunk_primary',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_trunk_primary',
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
  id: 'road_motorway',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'roads_motorways',
  minzoom: 5,
  layout: { 'line-cap': 'round', 'line-join': 'round' },
  paint: {
    'line-color': {
      base: 1,
      stops: [
        [5, 'hsl(26, 87%, 62%)'],
        [6, '#fc8']
      ]
    },
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
  id: 'road_major_rail',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'major_rail',
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
  id: 'road_major_rail_hatching',
  type: 'line',
  source: 'osm_cqrs',
  'source-layer': 'major_rail',
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

export const roads = {
  residential_casting,
  residential_base,
  motorway_link_casing,
  service_track_casing,
  link_casing,
  secondary_tertiary_casing,
  trunk_primary_casing,
  motorway_casing,
  path_pedestrian,
  motorway_link,
  service_track,
  link,
  seconday_tetiary,
  trunk_primary,
  motorway,
  major_rail,
  major_rail_hatching
};

// {
//   id: 'road_service_track_casing',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['in', 'type', 'service', 'track']],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#cfcdca',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [13, 0],
//         [15, 1],
//         [16, 4],
//         [20, 11]
//       ]
//     }
//   }
// },
// {

// },
// {
//   id: 'road_path_pedestrian',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   minzoom: 14,
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['in', 'type', 'path', 'pedestrian']],
//   layout: { 'line-join': 'round' },
//   paint: {
//     'line-color': 'hsl(0, 0%, 100%)',
//     'line-dasharray': [1, 0.7],
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [14, 1],
//         [20, 10]
//       ]
//     }
//   }
// },
// {
//   id: 'road_minor_casing',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['in', 'type', 'residential']],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#cfcdca',
//     'line-opacity': {
//       stops: [
//         [12, 0],
//         [12.5, 1]
//       ]
//     },
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [12, 0.5],
//         [13, 1],
//         [14, 4],
//         [20, 20]
//       ]
//     }
//   }
// },
// {
//   id: 'road_minor',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['in', 'type', 'residential']],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#fff',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [13.5, 0],
//         [14, 2.5],
//         [20, 18]
//       ]
//     }
//   }
// },
// // {
// //   id: 'road_link_casing',
// //   type: 'line',
// //   source: 'osm_cqrs',
// //   'source-layer': 'public.osm_roads',
// //   minzoom: 13,
// //   filter: [
// //     'all',
// //     ['!in', 'brunnel', 'bridge', 'tunnel'],
// //     ['!in', 'class', 'pedestrian', 'path', 'track', 'service', 'motorway'],
// //     ['==', 'ramp', 1]
// //   ],
// //   layout: { 'line-cap': 'round', 'line-join': 'round' },
// //   paint: {
// //     'line-color': '#e9ac77',
// //     'line-width': {
// //       base: 1.2,
// //       stops: [
// //         [12, 1],
// //         [13, 3],
// //         [14, 4],
// //         [20, 15]
// //       ]
// //     }
// //   }
// // },
// {
//   id: 'road_trunk_primary_casing',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['!in', 'brunnel', 'bridge', 'tunnel'], ['in', 'type', 'primary', 'trunk']],
//   layout: { 'line-join': 'round' },
//   paint: {
//     'line-color': '#e9ac77',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [5, 0.4],
//         [6, 0.7],
//         [7, 1.75],
//         [20, 22]
//       ]
//     }
//   }
// },
// {
//   id: 'road_trunk_primary',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['!in', 'brunnel', 'bridge', 'tunnel'], ['in', 'type', 'primary', 'trunk']],
//   layout: { 'line-join': 'round' },
//   paint: {
//     'line-color': '#fea',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [5, 0],
//         [7, 1],
//         [20, 18]
//       ]
//     }
//   }
// },
// {
//   id: 'road_motorway_casing',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   minzoom: 5,
//   filter: ['all', ['!in', 'brunnel', 'bridge', 'tunnel'], ['==', 'type', 'motorway'], ['!=', 'ramp', 1]],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#e9ac77',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [5, 0.4],
//         [6, 0.7],
//         [7, 1.75],
//         [20, 22]
//       ]
//     }
//   }
// },
// {
//   id: 'road_motorway',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   minzoom: 5,
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['==', 'type', 'motorway'], ['!=', 'ramp', 1]],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': {
//       base: 1,
//       stops: [
//         [5, 'hsl(26, 87%, 62%)'],
//         [6, '#fc8']
//       ]
//     },
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [5, 0],
//         [7, 1],
//         [20, 18]
//       ]
//     }
//   }
// },
// {
//   id: 'road_secondary_tertiary_casing',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: [
//     'all',
//     ['!=', 'bridge', 1],
//     ['!=', 'tunnel', 1],
//     ['in', 'type', 'secondary', 'tertiary'],
//     ['!=', 'ramp', 1]
//   ],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#e9ac77',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [8, 1.5],
//         [20, 17]
//       ]
//     }
//   }
// },
// {
//   id: 'road_secondary_tertiary',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['in', 'type', 'secondary', 'tertiary']],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#fea',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [6.5, 0],
//         [8, 0.5],
//         [20, 13]
//       ]
//     }
//   }
// },
// {
//   id: 'road_motorway_link',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   minzoom: 12,
//   filter: ['all', ['!=', 'bridge', 1], ['!=', 'tunnel', 1], ['==', 'type', 'motorway_link']],
//   layout: { 'line-cap': 'round', 'line-join': 'round' },
//   paint: {
//     'line-color': '#fc8',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [12.5, 0],
//         [13, 1.5],
//         [14, 2.5],
//         [20, 11.5]
//       ]
//     }
//   }
// },

// {
//   id: 'bridge_motorway_link_casing',
//   type: 'line',
//   source: 'osm_cqrs',
//   'source-layer': 'public.osm_roads',
//   filter: ['all', ['==', 'class', 'motorway_link'], ['==', 'bridge', 1]],
//   layout: { 'line-join': 'round' },
//   paint: {
//     'line-color': '#e9ac77',
//     'line-width': {
//       base: 1.2,
//       stops: [
//         [12, 1],
//         [13, 3],
//         [14, 4],
//         [20, 15]
//       ]
//     }
//   }
// },
