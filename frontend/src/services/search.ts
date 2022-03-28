import axios from 'axios';
import { getSearchResult, RawSearchResult, SearchResult } from '../model/search';

export class SearchService {
  // private static _apiUrl = 'https://osm.noxz.dev/search';
  private static _apiUrl = 'http://localhost:12000';

  public static async getPositionByName(name: string): Promise<SearchResult[]> {
    const response = await axios.get<RawSearchResult[]>(this._apiUrl + '/searchByName', {
      params: {
        name
      }
    });
    return getSearchResult(response.data!);
  }

  public static async getSearchItemCount(): Promise<number> {
    const response = await axios.get<{ count: number }>(this._apiUrl + '/count');
    return response!.data.count;
  }
}
