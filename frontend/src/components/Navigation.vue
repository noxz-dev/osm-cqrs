<script setup lang="ts">
import { ref, shallowRef } from 'vue';
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

const showNavigation = ref(true);

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

function toggleNavigation() {
  showNavigation.value = !showNavigation.value;
}
</script>

<template>
  <div class="absolute top-4 left-4 z-50" :class="showNavigation ? 'w-96' : 'w-12 h-12'">
    <div
      class="flex justify-center p-2 gap-4 shadow-xl border-x-2 border-t-2 border-slate-300 bg-gray-100 rounded-t-lg h-full"
      :class="!showNavigation && 'rounded-b-lg border-b-2'"
    >
      <div class="absolute left-2 top-2">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="32"
          height="32"
          class="stroke-slate-500 cursor-pointer"
          viewBox="0 0 256 256"
          @click="toggleNavigation"
        >
          <line
            x1="40"
            y1="128"
            x2="216"
            y2="128"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="16"
          ></line>
          <line
            x1="40"
            y1="64"
            x2="216"
            y2="64"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="16"
          ></line>
          <line
            x1="40"
            y1="192"
            x2="216"
            y2="192"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="16"
          ></line>
        </svg>
      </div>
      <div
        v-show="showNavigation"
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
    <div v-show="showNavigation">
      <component :is="currentComponent.component"></component>
    </div>
  </div>
</template>
