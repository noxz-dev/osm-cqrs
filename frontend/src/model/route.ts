export interface Coordinate {
  Lat: number;
  Lng: number;
}

export interface Route {
  routePoints: Coordinate[];
}
