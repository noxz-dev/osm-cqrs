import { Coordinate } from './coordnite';

export interface Tag {
  K: string;
  V: string;
}

export interface SearchResult {
  id: string;
  name: string;
  location: Coordinate;
  tags: Tag[];
}

export interface RawSearchResult {
  _index: string;
  _id: string;
  _score: number;
  _source: RawSource;
}

interface SearchCoordinate {
  lat: number;
  lon: number;
}
export interface RawSource {
  name: string;
  location: SearchCoordinate;
  tags: Tag[];
}

export function getSearchResult(result: RawSearchResult[]): SearchResult[] {
  return result.map((el: RawSearchResult) => {
    return {
      id: el._id,
      name: el._source.name,
      location: {
        Lat: el._source.location.lat,
        Lng: el._source.location.lon
      },
      tags: el._source.tags
    };
  });
}
