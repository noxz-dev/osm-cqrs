import axios from 'axios';
import { Coordinate, Route } from '../model/route';

export class RoutingService {
  private static _apiUrl = 'http://localhost:3003';

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
}
