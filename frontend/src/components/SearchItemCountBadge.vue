<script setup lang="ts">
import { onDeactivated, onMounted, ref } from 'vue';
import { SearchService } from '../services/search';

const searchItemCount = ref<number>();
let searchInterval: any;

onMounted(() => {
  setSearchItemCount();
  pullSearchItemCount();
});

onDeactivated(() => {
  clearInterval(searchInterval);
});

async function setSearchItemCount() {
  const count = await SearchService.getSearchItemCount();

  searchItemCount.value = count;
}

function pullSearchItemCount() {
  searchInterval = setInterval(() => {
    setSearchItemCount();
  }, 30000);
}
</script>

<template>
  <div class="absolute bottom-4 left-4 z-50 bg-white p-2 rounded-lg shadow-xl">Search items: {{ searchItemCount }}</div>
</template>
