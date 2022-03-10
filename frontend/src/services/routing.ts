import polyline from '@mapbox/polyline';
import axios from 'axios';
import { Coordinate } from '../model/coordinate';
export class RoutingService {
  private static _carUrl = 'http://localhost:5000/route/v1/';
  private static _bicycleUrl = 'http://localhost:5001/route/v1/';
  private static _footUrl = 'http://localhost:5002/route/v1/';

  public static async getCarRoute(from: Coordinate, to: Coordinate) {
    // http://localhost:5000/route/v1/driving/9.674266,52.354086;9.663847,52.359593?steps=true
    const response = await axios.get(
      this._carUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&alternatives=true`
    );
    // const response = await axios.get(this._apiUrl + `driving/9.674266,52.354086;9.663847,52.359593`);
    const route = response.data!.routes[0].geometry;

    const geoJson = polyline.toGeoJSON(route);
    return geoJson;
  }

  public static async getBycicleRoute(from: Coordinate, to: Coordinate) {
    const response = await axios.get(
      this._bicycleUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&alternatives=true`
    );
    const route = response.data!.routes[0].geometry;

    const geoJson = polyline.toGeoJSON(route);
    return geoJson;
  }

  public static async getFootRoute(from: Coordinate, to: Coordinate) {
    const response = await axios.get(
      this._footUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&alternatives=true`
    );
    const route = response.data!.routes[0].geometry;

    const geoJson = polyline.toGeoJSON(route);
    return geoJson;
  }
}
