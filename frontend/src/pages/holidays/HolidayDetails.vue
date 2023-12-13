<template>
  <div class="container mt-5">
    <router-link :to="{ name: 'Holidays' }" class="link-primary">
      &larr; Go Back
    </router-link>
    <h1 class="mt-4">Holiday Details</h1>
    <div v-if="holiday">
      <div class="card" style="width: 40%">
        <div class="card-body">
          <h5 class="card-title">{{ holiday.title }}</h5>
          <p class="card-text">
            <strong>Start Date:</strong> {{ holiday.startDate }}<br />
            <strong>Duration:</strong> {{ holiday.duration }} days<br />
            <strong>Price:</strong> ${{ holiday.price }}<br />
            <strong>Location:</strong> {{ holiday.location.city }},
            {{ holiday.location.country }} <br />
            <strong>Address:</strong> {{ holiday.location.street }}, №{{
              holiday.location.number
            }}
          </p>
        </div>
      </div>
    </div>
    <div v-if="!holiday">
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

const holidayId = ref(null);
const holiday = ref(null);

onMounted(async () => {
  holidayId.value = route.query.id;

  holiday.value = await holidayStore.fetchHolidayById(holidayId.value);
});
</script>
