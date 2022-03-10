<script setup lang="ts">
import { shallowRef } from 'vue';
import { mapStore } from '../store/map';
import Routing from './Routing.vue';
import Search from './Search.vue';

const components = [
  {
    name: 'Search',
    component: Search
  },
  {
    name: 'Routing',
    component: Routing
  }
];

const currentComponent = shallowRef<{ name: string; component: any }>(components[0]);

function isCurrentComponent(component: string): boolean {
  return currentComponent.value.name === component;
}

function changeComponent(component: any) {
  currentComponent.value = component;
  mapStore.route = undefined;
  mapStore.searchResults = [];
  mapStore.selectedIndex = undefined;
  mapStore.selectedResult = undefined;
}
</script>

<template>
  <div class="absolute top-4 left-4 z-50 w-96">
    <div
      class="flex justify-center p-2 gap-4 shadow-xl border-x-2 border-t-2 border-slate-300 bg-gray-100 rounded-t-lg"
    >
      <div
        v-for="component in components"
        class="cursor-pointer"
        :class="isCurrentComponent(component.name) && 'text-green-500'"
        @click="changeComponent(component)"
      >
        <span>{{ component.name }}</span>
      </div>

      <!-- <div :class="isCurrentComponent('Search') && 'text-green-500'">Search</div>
      <div :class="isCurrentComponent('Routing') && 'text-green-500'">Routing</div> -->
    </div>
    <component :is="currentComponent.component"></component>
  </div>
</template>
