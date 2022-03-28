import { LayerSpecification } from 'maplibre-gl';
import { aeroway } from './aeroways';
import { boundaries } from './boundaries';
import { bridges } from './bridges';
import { building } from './building';
import { amenities } from './labels/landuse';
import { road_labels } from './labels/roads';
import { water_labels } from './labels/water';
import { landuse } from './landuse';
import { roads } from './roads';
import { tunnel } from './tunnels';
import { water } from './water';

export const layers: LayerSpecification[] = [
  {
    id: 'background',
    type: 'background',
    paint: { 'background-color': 'rgb(239,239,239)' }
  },
  landuse.park_base,
  landuse.park_outline,
  landuse.residential,
  landuse.wood_z_2_7,
  landuse.wood_z_7_9,
  landuse.wood_z_9_13,
  landuse.wood_z_13_22,
  landuse.grass,
  landuse.cementry,
  landuse.hospital,
  landuse.school,

  water.river,
  water.other,
  water.areas,

  aeroway.fill,
  aeroway.runway,
  aeroway.taxiway,

  tunnel.motorway_link_casing,
  tunnel.service_track_casing,
  tunnel.link_casing,
  tunnel.seconday_tetiary_casing,
  tunnel.primary_casing,
  tunnel.motorway_casing,
  tunnel.path_pedestrian,
  tunnel.motorway_link,
  tunnel.service_track,
  tunnel.link,
  tunnel.residental,
  tunnel.secondary_tertiary,
  tunnel.trunk_primary,
  tunnel.motorway,
  tunnel.major_rail,
  tunnel.major_rail_hatching,

  roads.service_track_casing,
  roads.service_track,
  roads.residential_casting,
  roads.residential_base,
  roads.motorway_link_casing,

  roads.link_casing,
  roads.secondary_tertiary_casing,
  roads.trunk_primary_casing,
  roads.motorway_casing,
  roads.path_pedestrian,
  roads.motorway_link,

  roads.link,
  roads.seconday_tetiary,
  roads.trunk_primary,
  roads.motorway,
  roads.major_rail,
  roads.major_rail_hatching,

  bridges.motorway_link_casing,
  bridges.service_track_casing,
  bridges.path_pedestrian_casing,
  bridges.secondary_tertiary_casing,
  bridges.trunk_primary_casing,
  bridges.motorway_casing,
  bridges.motorway_link,
  bridges.service_track,
  bridges.secondary_tertiary,
  bridges.trunk_primary,
  bridges.motorway,
  bridges.major_rail,
  bridges.major_rail_hatching,

  building.area,
  building.threeD,

  boundaries._1_2,
  boundaries._3_4,
  boundaries._5_6,
  boundaries._7_8,
  boundaries._9_10,

  road_labels.z14_motorway,
  road_labels.z14_trunk,
  road_labels.z14_residential,
  road_labels.z14_roads_secondary_tertiary,
  road_labels.z14_service_track,
  road_labels.z14_roads_trunk_primary,
  road_labels.z14_roads_pedestrian,
  water_labels.areas_z13,
  water_labels.ways_z13,
  water_labels.others_z13,
  amenities.park
];
