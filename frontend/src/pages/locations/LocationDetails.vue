<template>
  <div class="container mt-5">
    <router-link :to="{ name: 'Locations' }" class="link-primary">
      &larr; Go Back
    </router-link>
    <h1 class="mt-4">Location Details</h1>
    <div v-if="locationDetails">
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">Street: {{ locationDetails.street }}</h5>
          <p class="card-text">
            <strong>Number:</strong> {{ locationDetails.number }}<br />
            <strong>City:</strong> {{ locationDetails.city }}<br />
            <strong>Country:</strong> {{ locationDetails.country }}
          </p>
        </div>
      </div>
    </div>
    <div v-if="!locationDetails">
      <p>Loading location details...</p>
    </div>
  </div>
</template>

<script setup>
import { useHolidayStore } from "@/store/holidayStore";
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

const holidayStore = useHolidayStore();
const route = useRoute();

const locationId = ref(null);
const locationDetails = ref(null);

onMounted(async () => {
  locationId.value = route.query.id;

  locationDetails.value = await holidayStore.fetchLocationById(
    locationId.value
  );
});
</script>
