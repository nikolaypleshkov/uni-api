<template>
  <form @submit.prevent="updateHoliday(holiday)" class="m-3">
    <div class="row g-3">
      <div class="col-md-4">
        <label for="title" class="form-label">Title</label>
        <input
          v-model="editFormData.title"
          type="text"
          class="form-control"
          id="title"
          placeholder="Title"
        />
      </div>
      <div class="col-md-2">
        <label for="startDate" class="form-label">Start Date</label>
        <input
          v-model="editFormData.startDate"
          type="date"
          class="form-control"
          id="startDate"
        />
      </div>

      <div class="col-md-3">
        <label for="duration" class="form-label">Duration (days)</label>
        <input
          v-model="editFormData.duration"
          type="number"
          class="form-control"
          id="duration"
          placeholder="Duration"
        />
      </div>
      <div class="col-md-3">
        <label for="price" class="form-label">Price</label>
        <input
          v-model="editFormData.price"
          type="text"
          class="form-control"
          id="price"
          placeholder="Price"
        />
      </div>

      <div class="col-md-3">
        <label for="freeSlots" class="form-label">Free Slots</label>
        <input
          v-model="editFormData.freeSlots"
          type="number"
          class="form-control"
          id="freeSlots"
          placeholder="Free Slots"
        />
      </div>

      <div class="col-md-3">
        <label for="location" class="form-label">Location</label>
        <select
          v-model="editFormData.location"
          class="form-select"
          id="location"
        >
          <option value="" disabled>Select a location</option>
          <option
            v-for="location in holidayStore.locations"
            :key="location.id"
            :value="location.id"
          >
            {{ location.street }} {{ location.number }}, {{ location.city }},
            {{ location.country }}
          </option>
        </select>
      </div>
    </div>
    <button type="submit" class="btn btn-success mt-4">Update Holiday</button>
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
  holiday: {
    type: Object,
    required: true,
  },
});

const editFormData = ref({
  title: props.holiday.title,
  startDate: props.holiday.startDate,
  duration: props.holiday.duration,
  price: props.holiday.price,
  freeSlots: props.holiday.freeSlots,
  location: props.holiday.location.id,
});

const editedHoliday = ref(props.holiday);

const updateHoliday = async () => {
  try {
    const updateDTO = {
      id: props.holiday.id,
      title: editFormData.value.title,
      startDate: editFormData.value.startDate,
      duration: editFormData.value.duration,
      price: Number(editFormData.value.price),
      freeSlots: editFormData.value.freeSlots,
      location: editFormData.value.location,
    };

    await holidayStore.updateHoliday(updateDTO);
    holidayStore.showEdit = false;
    editedHoliday.value = null;
  } catch (error) {
    console.log(error);
  }
};

const cancelEdit = () => {
  holidayStore.showEdit = false;
  editedHoliday.value = null;
};
</script>
