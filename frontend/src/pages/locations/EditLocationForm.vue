<template>
  <form @submit.prevent="updateLocation" class="m-3">
    <div class="row g-3">
      <div class="col-md-4">
        <label for="editStreet" class="form-label">Street</label>
        <input
          v-model="editFormData.street"
          type="text"
          class="form-control"
          id="editStreet"
          placeholder="Street"
        />
      </div>
      <div class="col-md-2">
        <label for="editNumber" class="form-label">Number</label>
        <input
          v-model="editFormData.number"
          type="text"
          class="form-control"
          id="editNumber"
          placeholder="Number"
        />
      </div>
      <div class="col-md-3">
        <label for="editCity" class="form-label">City</label>
        <input
          v-model="editFormData.city"
          type="text"
          class="form-control"
          id="editCity"
          placeholder="City"
        />
      </div>
      <div class="col-md-3">
        <label for="editCountry" class="form-label">Country</label>
        <input
          v-model="editFormData.country"
          type="text"
          class="form-control"
          id="editCountry"
          placeholder="Country"
        />
      </div>
    </div>
    <button type="submit" class="btn btn-success mt-4">Update Location</button>
    <button @click="cancelEdit" class="btn btn-secondary mt-4 btn-cancel">
      Cancel
    </button>
  </form>
</template>

<script setup>
import { ref } from "vue";
import { useHolidayStore } from "@/store/holidayStore";

const holidayStore = useHolidayStore();

const props = defineProps({
  location: {
    type: Object,
    required: true,
  },
});

const editFormData = ref({
  street: props.location.street,
  number: props.location.number,
  city: props.location.city,
  country: props.location.country,
});

const editedLocation = ref(props.location);

const updateLocation = async () => {
  try {
    const updateDTO = {
      id: props.location.id,
      street: editFormData.value.street,
      number: editFormData.value.number,
      city: editFormData.value.city,
      country: editFormData.value.country,
    };
    await holidayStore.updateLocation(updateDTO);
    holidayStore.fetchData();
    holidayStore.showEdit = false;
    editedLocation.value = null;
  } catch (error) {
    console.error("Error updating location:", error);
  }
};

const cancelEdit = () => {
  editedLocation.value = null;
  holidayStore.showEdit = false;
};
</script>
