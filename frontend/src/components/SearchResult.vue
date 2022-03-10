<script setup lang="ts">
import { PropType } from 'vue';
import { SearchResult } from '../model/search';
import { mapStore } from '../store/map';

const emit = defineEmits(['result-selected']);

const props = defineProps({
  searchResult: {
    required: true,
    type: Object as PropType<SearchResult>
  },
  index: {
    required: true,
    type: Number
  }
});

const flatTags = new Map();

const tags = props.searchResult.tags.map((t) => {
  flatTags.set(t.K, t.V);
});
</script>

<template>
  <div
    :class="mapStore.selectedIndex == index && 'bg-gray-200 ring-2 ring-slate-500'"
    class="bg-gray-100 shadow-md rounded-lg p-4 flex flex-col items-start text-left cursor-pointer hover:bg-gray-200 transition-color"
    @click="emit('result-selected', { index: props.index, result: props.searchResult })"
  >
    <span class="text-slate-500 text-sm">{{ props.searchResult.id }}</span>
    <div class="mt-2 mb-1">
      <span class="font-bold">{{ props.searchResult.name }}</span>
    </div>
    <span class="text-sm text-slate-700">
      {{ flatTags.get('addr:street') }} {{ flatTags.get('addr:housenumber') }} {{ flatTags.get('addr:postcode') }}
    </span>
  </div>
</template>
