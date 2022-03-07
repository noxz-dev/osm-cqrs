<script setup lang="ts">
import { RoutingService } from '../services/routing';
import InputField from './form/InputField.vue';

const emit = defineEmits(['route-changed']);

async function searchRoute() {
  const route = await RoutingService.getDijkstraRoute(
    {
      Lat: 52.354086,
      Lng: 9.674266
    },
    {
      Lat: 52.359593,
      Lng: 9.663847
    }
  );

  // const route = await RoutingService.getDijkstraRoute(
  //   {
  //     lat: 52.355171729958165,
  //     lng: 9.725275628389035
  //   },
  //   {
  //     lat: 52.354086,
  //     lng: 9.674266
  //   }
  // );

  emit('route-changed', route);
}
</script>

<template>
  <div class="absolute bg-slate-600 shadow-md top-4 left-4 z-50 rounded-lg p-4">
    <form>
      <InputField></InputField>
    </form>
    <button @click="searchRoute">Search</button>
  </div>
</template>

<style scoped>
#searchContainer {
  background-color: white;
  top: 1rem;
  left: 1rem;
  position: absolute;
  z-index: 999;
}
</style>
