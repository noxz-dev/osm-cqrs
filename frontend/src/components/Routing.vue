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
  if (!fromSearchString.value || !toSearchString.value) return;

  const fromResults = await SearchService.getPositionByName(fromSearchString.value);
  const toResults = await SearchService.getPositionByName(toSearchString.value);

  let route;
  if (selectedProfile.value === Profile.CAR) {
    route = await RoutingService.getCarRoute(fromResults[0].location, toResults[0].location);
  } else if (selectedProfile.value === Profile.BICYCLE) {
    route = await RoutingService.getBycicleRoute(fromResults[0].location, toResults[0].location);
  } else {
    route = await RoutingService.getFootRoute(fromResults[0].location, toResults[0].location);
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
</template>
