import { reactive } from 'vue';
import { SearchResult } from '../model/search';

interface MapStore {
  route?: any;
  searchResults: SearchResult[];
  selectedResult?: SearchResult;
  selectedIndex?: number;
}

export const mapStore = reactive<MapStore>({
  searchResults: []
});
