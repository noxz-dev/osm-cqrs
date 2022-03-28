import polyline from '@mapbox/polyline';
import axios from 'axios';
import { Coordinate } from '../model/coordinate';
import { RoutingResult } from '../model/route';

export class RoutingService {
  private static _carUrl = 'http://localhost:5000/route/v1/';
  private static _bicycleUrl = 'http://localhost:5001/route/v1/';
  private static _footUrl = 'http://localhost:5002/route/v1/';

  // private static _carUrl = 'https://osm.noxz.dev/routing/car/route/v1/';
  // private static _bicycleUrl = 'https://osm.noxz.dev/routing/bicycle/route/v1/';
  // private static _footUrl = 'https://osm.noxz.dev/routing/foot/route/v1/';

  public static async getCarRoute(from: Coordinate, to: Coordinate): Promise<RoutingResult> {
    const response = await axios.get(
      this._carUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&steps=true`
    );

    console.log(response);

    // const response = await axios.get(
    //   this._carUrl +
    //     `driving/9.674198279999999,52.354088280000006;9.76564744090909,52.36859876363637?overview=full&steps=true`
    // );

    return RoutingService.generateResult(response);
  }

  public static async getBycicleRoute(from: Coordinate, to: Coordinate) {
    const response = await axios.get(
      this._bicycleUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&steps=true`
    );
    return RoutingService.generateResult(response);
  }

  public static async getFootRoute(from: Coordinate, to: Coordinate) {
    const response = await axios.get(
      this._footUrl + `driving/${from.Lng},${from.Lat};${to.Lng},${to.Lat}?overview=full&steps=true`
    );
    return RoutingService.generateResult(response);
  }

  private static generateResult(response: any) {
    const route = response.data!.routes[0];
    const statistic = route.legs[0];
    const geoJson = polyline.toGeoJSON(route.geometry);
    return {
      geometry: geoJson,
      statistic: statistic
    };
  }
}
