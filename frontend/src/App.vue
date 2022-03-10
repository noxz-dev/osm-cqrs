<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import Map from './components/Map.vue';
import Navigation from './components/Navigation.vue';
</script>

<template>
  <Navigation></Navigation>
  <Map></Map>
</template>

<style>
#app {
  font-family: 'Inter', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  margin: 0;
  scrollbar-width: none;
}

* {
  scrollbar-width: none;
}
</style>
