<script setup lang="ts">
import { ref, watch } from 'vue';
import { debounceRef } from '../debounceRef';
import { SearchResult as SearchResultModel } from '../model/search';
import { RoutingService } from '../services/routing';
import { SearchService } from '../services/search';
import { mapStore } from '../store/map';
import InputField from './form/InputField.vue';
import SearchResult from './SearchResult.vue';

const fromResult = ref<SearchResultModel>();
const fromResults = ref<SearchResultModel[]>();
const fromSearchString = debounceRef('', 300);
const toResult = ref<SearchResultModel>();
const toResults = ref<SearchResultModel[]>();

type LastChangedType = 'from' | 'to';

enum Profile {
  CAR = 'Car',
  BICYCLE = 'Bicycle',
  FOOT = 'Foot'
}

const lastChanged = ref<LastChangedType>();

const toSearchString = debounceRef('', 300);

const fromLoading = ref(false);

const toLoading = ref(false);

const selectedProfile = ref<Profile>(Profile.CAR);

watch(
  () => fromSearchString.value,
  () => {
    mapStore.routeError = false;
    fromResult.value = undefined;
    lastChanged.value = 'from';
    if (fromSearchString.value != '') {
      searchFromPosition();
    }
  }
);

watch(
  () => toSearchString.value,
  () => {
    mapStore.routeError = false;
    lastChanged.value = 'to';
    toResult.value = undefined;
    if (toSearchString.value != '') {
      searchToPosition();
    }
  }
);

async function searchFromPosition() {
  fromResults.value = [];
  if (toSearchString.value == '') {
    toResults.value = [];
  }
  if (fromSearchString.value != '') {
    fromLoading.value = true;

    fromResults.value = await SearchService.getPositionByName(fromSearchString.value);
    fromLoading.value = false;
  }
}

async function searchToPosition() {
  toResults.value = [];

  if (fromSearchString.value == '') {
    fromResults.value = [];
  }

  if (toSearchString.value != '') {
    toLoading.value = true;

    toResults.value = await SearchService.getPositionByName(toSearchString.value);
    toLoading.value = false;
  }
}

async function calculteRoute() {
  if (!fromResult.value || !toResult.value) return;

  const fromCoordinates = fromResult.value?.location;
  const toCoordinates = toResult.value?.location;
  mapStore.routeError = false;
  let route;
  if (selectedProfile.value === Profile.CAR) {
    try {
      route = await RoutingService.getCarRoute(fromCoordinates, toCoordinates);
    } catch (error: any) {
      if (error.response.status === 400) {
        mapStore.routeError = true;
      }
    }
  } else if (selectedProfile.value === Profile.BICYCLE) {
    try {
      route = await RoutingService.getBycicleRoute(fromCoordinates, toCoordinates);
    } catch (error: any) {
      if (error.response.status === 400) {
        mapStore.routeError = true;
      }
    }
  } else {
    try {
      route = await RoutingService.getFootRoute(fromCoordinates, toCoordinates);
    } catch (error: any) {
      if (error.response.status === 400) {
        mapStore.routeError = true;
      }
    }
  }

  mapStore.route = route;
}

function selectResult(eventPayload: { index: number; result: SearchResultModel }) {
  if (lastChanged.value === 'from') {
    fromResult.value = eventPayload.result;
    fromResults.value = [];
  }

  if (lastChanged.value === 'to') {
    toResult.value = eventPayload.result;
    toResults.value = [];
  }
}

function secondsToHoursMinuteSecondsString(totalSeconds: number) {
  const hours = totalSeconds / 3600;
  const minutes = Math.floor(totalSeconds / 60);
  const seconds = totalSeconds % 60;

  return `${hours >= 1 ? `${hours}h` : ''}${minutes >= 1 ? `${minutes}min` : ''} ${Number(seconds).toFixed(0)}s`;
}

function metersToKilometersMeters(totalMeters: number) {
  const kilometers = Math.floor(totalMeters / 1000);
  return `${kilometers >= 1.0 ? `${kilometers}km` : ''} ${Number(totalMeters % 1000).toFixed(0)}m`;
}
</script>

<template>
  <div class="bg-gray-100 shadow-md border-x-2 border-b-2 border-slate-300 rounded-b-lg p-4 mb-2">
    <form class="flex flex-col gap-2" @submit.prevent>
      <div class="flex justify-center items-center gap-2 text-left">
        <span class="text-slate-600 font-semibold text-sm w-12">From: </span>

        <InputField class="w-full" v-model="fromSearchString" placeholder="Hochschule Hannover"></InputField>
      </div>
      <div class="flex justify-center items-center gap-2 text-right">
        <span class="text-slate-600 font-semibold text-sm w-12">To: </span>
        <InputField class="w-full" v-model="toSearchString" placeholder="Leibniz Universität"></InputField>
      </div>
      <div class="flex w-full justify-evenly my-2">
        <div
          :class="selectedProfile === profile && 'ring-2 ring-green-500 bg-green-200 '"
          class="px-2 py-0.5 rounded-3xl cursor-pointer"
          v-for="profile in Profile"
          @click="selectedProfile = profile"
        >
          {{ profile }}
        </div>
      </div>
      <button type="submit" @click="calculteRoute" class="bg-green-500 p-2 w-full rounded-lg text-white">
        Calculate Route
      </button>
    </form>
  </div>

  <div
    class="bg-gray-100 shadow-md rounded-lg p-4 flex flex-col items-center text-center"
    v-if="fromLoading || (fromResults?.length == 0 && fromSearchString != '' && !fromResult)"
  >
    <div v-if="fromLoading">
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
    <div v-if="!fromLoading">No results found</div>
  </div>
  <SearchResult class="ring-2 ring-green-500" v-if="fromResult" :index="0" :search-result="fromResult"></SearchResult>

  <div v-if="fromResults?.length !== 0" class="flex flex-col gap-2 max-h-[350px] overflow-y-auto p-1">
    <SearchResult
      class=""
      v-for="(item, index) in fromResults"
      :key="item.id"
      :index="index"
      :search-result="item"
      @result-selected="selectResult"
    />
  </div>
  <div v-if="toResult || fromResult" class="flex justify-center my-2">
    <svg
      width="64"
      viewBox="0 0 100 100"
      version="1.1"
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      xml:space="preserve"
      xmlns:serif="http://www.serif.com/"
      style="
        fill-rule: evenodd;
        clip-rule: evenodd;
        stroke-linecap: round;
        stroke-linejoin: round;
        stroke-miterlimit: 1.5;
      "
    >
      <g transform="matrix(1.20448,0,0,1.20448,-9.94379,-0.634241)">
        <circle
          cx="49.768"
          cy="12.164"
          r="9.065"
          style="fill: rgb(46, 204, 113); stroke: rgb(46, 204, 113); stroke-width: 1.25px"
        />
      </g>
      <g transform="matrix(0.48179,0,0,0.48179,26.0225,50.4211)">
        <circle
          cx="49.768"
          cy="12.164"
          r="9.065"
          style="fill: rgb(127, 140, 141); stroke: rgb(127, 140, 141); stroke-width: 2.08px"
        />
      </g>
      <g transform="matrix(0.48179,0,0,0.48179,26.0225,32.5645)">
        <circle
          cx="49.768"
          cy="12.164"
          r="9.065"
          style="fill: rgb(127, 140, 141); stroke: rgb(127, 140, 141); stroke-width: 2.08px"
        />
      </g>
      <g transform="matrix(1.34576,0,0,1.34576,-16.0071,-33.5049)">
        <g transform="matrix(0.135223,0,0,0.135223,36.0669,74.3077)">
          <path
            d="M156,78C156,132 96,174 96,174C96,174 36,132 36,78C36,45.085 63.085,18 96,18C128.915,18 156,45.085 156,78Z"
            style="
              fill: rgb(231, 76, 60);
              fill-rule: nonzero;
              stroke: rgb(231, 76, 60);
              stroke-width: 11.09px;
              stroke-miterlimit: 4;
            "
          />
        </g>
        <g transform="matrix(0.202692,0,0,0.202692,29.5899,69.0452)">
          <circle
            cx="96"
            cy="78"
            r="24"
            style="fill: white; stroke: rgb(231, 76, 60); stroke-width: 11.09px; stroke-miterlimit: 4"
          />
        </g>
      </g>
    </svg>
  </div>
  <SearchResult class="ring-2 ring-red-500" v-if="toResult" :index="0" :search-result="toResult"></SearchResult>
  <div
    class="bg-gray-100 shadow-md rounded-lg p-4 flex flex-col items-center text-center"
    v-if="toLoading || (toResults?.length == 0 && toSearchString != '' && !toResult)"
  >
    <div v-if="toLoading">
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
    <div v-if="!toLoading">No results found</div>
  </div>
  <div v-if="toResults?.length !== 0" class="max-h-[350px] overflow-y-auto p-1 flex flex-col gap-2">
    <SearchResult
      v-for="(item, index) in toResults"
      :key="item.id"
      :index="index"
      :search-result="item"
      @result-selected="selectResult"
    />
  </div>

  <div v-if="mapStore.route?.statistic || mapStore.routeError">
    <div
      v-if="mapStore.routeError || mapStore.route?.statistic.distance === 0 || mapStore.route?.statistic.duration === 0"
      class="flex gap-2 mt-6 p-2 shadow-xl border-2 border-red-800 rounded-lg justify-center items-center bg-red-50"
    >
      <div>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="30"
          height="30"
          class="stroke-red-800 fill-red-800"
          viewBox="0 0 256 256"
        >
          <circle cx="128" cy="128" r="96" fill="none" stroke-miterlimit="10" stroke-width="16"></circle>
          <circle cx="92" cy="108" r="12"></circle>
          <circle cx="164" cy="108" r="12"></circle>
          <circle cx="92" cy="108" r="12"></circle>
          <circle cx="164" cy="108" r="12"></circle>
          <path
            d="M169.6,176a48.1,48.1,0,0,0-83.2,0"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="16"
          ></path>
        </svg>
      </div>
      <div class="text-red-800 font-semibold">Oh no! There is no route available</div>
    </div>
    <div
      v-if="
        !mapStore.routeError &&
        mapStore.route?.statistic &&
        mapStore.route?.statistic.distance > 0 &&
        mapStore.route?.statistic.duration > 0
      "
      class="flex gap-4 justify-evenly bg-white rounded-lg mt-6 p-2 shadow-xl border-2 border-slate-300"
    >
      <div class="flex items-center justify-center gap-1">
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" class="stroke-slate-500" viewBox="0 0 256 256">
            <circle cx="128" cy="128" r="96" fill="none" stroke-miterlimit="10" stroke-width="16"></circle>
            <polyline
              points="128 72 128 128 184 128"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></polyline>
          </svg>
        </div>
        <div>{{ secondsToHoursMinuteSecondsString(mapStore.route.statistic.duration) }}</div>
      </div>

      <div class="flex items-center justify-center gap-1">
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" class="stroke-slate-500" viewBox="0 0 256 256">
            <rect width="256" height="256" fill="none"></rect>
            <rect
              x="26.2"
              y="82.7"
              width="203.6"
              height="90.51"
              rx="8"
              transform="translate(-53 128) rotate(-45)"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></rect>
            <line
              x1="132"
              y1="60"
              x2="164"
              y2="92"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></line>
            <line
              x1="96"
              y1="96"
              x2="128"
              y2="128"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></line>
            <line
              x1="60"
              y1="132"
              x2="92"
              y2="164"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></line>
          </svg>
        </div>
        <div>{{ metersToKilometersMeters(mapStore.route.statistic.distance) }}</div>
      </div>
      <div class="flex items-center justify-center gap-1">
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" class="stroke-slate-500" viewBox="0 0 256 256">
            <circle
              cx="40"
              cy="200"
              r="24"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></circle>
            <circle
              cx="96"
              cy="96"
              r="24"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></circle>
            <circle
              cx="160"
              cy="160"
              r="24"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></circle>
            <circle
              cx="216"
              cy="56"
              r="24"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></circle>
            <line
              x1="84.6"
              y1="117.1"
              x2="51.4"
              y2="178.9"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></line>
            <line
              x1="143"
              y1="143"
              x2="113"
              y2="113"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></line>
            <line
              x1="204.6"
              y1="77.1"
              x2="171.4"
              y2="138.9"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="16"
            ></line>
          </svg>
        </div>
        <div>{{ mapStore.route.statistic.steps.length }} Steps</div>
      </div>
    </div>
  </div>
</template>
