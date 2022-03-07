<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import { ref } from 'vue';
import Map from './components/Map.vue';
import Search from './components/Search.vue';
import { Route } from './model/route';

const route = ref<Route>();

function updateRoute(newRoute: Route) {
  route.value = newRoute;
}
</script>

<template>
  <Search @route-changed="updateRoute"></Search>
  <Map :route="route"></Map>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  margin: 0;
}
</style>
