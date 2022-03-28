import { Coordinate } from './coordinate';

export interface Route {
  routePoints: Coordinate[];
}

export interface RoutingResult {
  geometry: any;
  statistic: RouteStatistic;
}
export interface RouteStatistic {
  steps: any;
  weight: number;
  duration: number;
  distance: number;
}
