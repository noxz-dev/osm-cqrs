import { reactive } from 'vue';
import { RoutingResult } from '../model/route';
import { SearchResult } from '../model/search';
interface MapStore {
  route?: RoutingResult;
  searchResults: SearchResult[];
  selectedResult?: SearchResult;
  selectedIndex?: number;
  routeError: boolean;
}

export const mapStore = reactive<MapStore>({
  searchResults: [],
  routeError: false
});
