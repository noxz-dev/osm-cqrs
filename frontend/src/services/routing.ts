import polyline from '@mapbox/polyline';
import axios from 'axios';
import { Coordinate } from '../model/coordinate';
import { Route } from '../model/route';
export class RoutingService {
  private static _apiUrl = 'http://localhost:5000/route/v1/';

  public static async getDijkstraRoute(from: Coordinate, to: Coordinate): Promise<Route> {
    const response = await axios.get(this._apiUrl + '/routeDijkstra', {
      params: {
        fromLat: from.Lat,
        fromLng: from.Lng,
        toLat: to.Lat,
        toLng: to.Lng
      }
    });
    return {
      routePoints: response.data!
    };
  }

  public static async getRoute(from: Coordinate, to: Coordinate) {
    // http://localhost:5000/route/v1/driving/9.674266,52.354086;9.663847,52.359593?steps=true
    const response = await axios.get(
      this._apiUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&alternatives=true`
    );
    // const response = await axios.get(this._apiUrl + `driving/9.674266,52.354086;9.663847,52.359593`);
    const route = response.data!.routes[0].geometry;

    const geoJson = polyline.toGeoJSON(route);
    return geoJson;
  }
}
