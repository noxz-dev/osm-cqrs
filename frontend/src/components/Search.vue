<script setup lang="ts">
import { ref, watch } from 'vue';
import { debounceRef } from '../debounceRef';
import { SearchResult as SearchResultModel } from '../model/search';
import { SearchService } from '../services/search';
import { mapStore } from '../store/map';
import InputField from './form/InputField.vue';
import SearchResult from './SearchResult.vue';
const emit = defineEmits(['route-changed', 'result-selected']);

const searchString = debounceRef('', 300);

const searchResult = ref<SearchResultModel[]>([]);

const loading = ref(false);

watch(
  () => searchString.value,
  () => {
    if (searchString.value != '') {
      searchPosition();
    }
  }
);

async function searchPosition() {
  mapStore.searchResults = [];
  searchResult.value = [];
  if (searchString.value != '') {
    loading.value = true;

    searchResult.value = await SearchService.getPositionByName(searchString.value);
    mapStore.selectedIndex = -1;
    if (searchResult.value.length > 0) {
      mapStore.searchResults = searchResult.value;
    }
    loading.value = false;
  }
}

function selectResult(eventPayload: { index: number; result: SearchResultModel }) {
  mapStore.selectedResult = eventPayload.result;
  mapStore.selectedIndex = eventPayload.index;
}
</script>

<template>
  <div>
    <div class="bg-gray-100 shadow-md rounded-b-lg p-4 mb-2 border-b-2 border-x-2 border-slate-300">
      <form class="flex items-center justify-center gap-4" @submit.prevent>
        <InputField v-model="searchString" placeholder="Hochschule Hannover"></InputField>
        <button type="submit" @click="searchPosition" class="bg-green-500 p-2 w-full rounded-lg text-white">
          Search
        </button>
      </form>
    </div>
    <div
      class="bg-gray-100 shadow-md rounded-lg p-4 mb-2 flex flex-col items-center text-center"
      v-if="loading || (searchResult.length == 0 && searchString != '')"
    >
      <div v-if="loading">
        <svg
          class="animate-spin h-[24px] w-[24px] text-slate-600"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
      </div>
      <div v-if="!loading">No results found</div>
    </div>
    <div class="max-h-[700px] overflow-y-auto p-1">
      <SearchResult
        class="mb-2"
        v-for="(item, index) in searchResult"
        :key="item.id"
        :index="index"
        :search-result="item"
        @result-selected="selectResult"
      />
    </div>
  </div>
</template>

<style>
::-webkit-scrollbar {
  display: none;
}
</style>
