<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-8">
        <div class="card pe-4 mb-5">
          <div class="card-body">
            <h1>Holidays</h1>
            <div v-if="holidayStore.loading" class="mb-3">Loading...</div>
            <div v-if="!holidayStore.loading">
              <div
                v-for="holiday in holidayStore.holidays"
                :key="holiday.id"
                class="card mb-3"
              >
                <div class="card-body">
                  <div class="row">
                    <div class="col-md-4">
                      <h5 class="card-title mb-2">
                        <strong>{{ holiday.title }}</strong>
                      </h5>
                      <p class="card-text mb-2">
                        <strong>Start Date:</strong> {{ holiday.startDate }}
                      </p>
                      <p class="card-text mb-2">
                        <strong>Duration:</strong> {{ holiday.duration }}
                      </p>
                      <p class="card-text mb-2">
                        <strong>Price:</strong> ${{ holiday.price }}
                      </p>
                      <p class="card-text mb-2">
                        <strong>Free Slots:</strong> {{ holiday.freeSlots }}
                      </p>
                    </div>
                    <div class="col-md-4">
                      <p class="card-text mb-0">
                        <span class="fw-bold">Location:</span><br />
                        {{ holiday.location.street }},
                        {{ holiday.location.number }},<br />
                        {{ holiday.location.city }},
                        {{ holiday.location.country }}
                      </p>
                    </div>
                    <div
                      class="col-md-4 d-flex flex-column justify-content-end align-items-end mr-3"
                    >
                      <button
                        @click="showEditForm(holiday)"
                        class="btn btn-primary mb-2"
                      >
                        Edit
                      </button>
                      <button class="btn btn-secondary mb-2">
                        <router-link
                          :to="{
                            name: 'HolidayDetails',
                            query: { id: holiday.id },
                          }"
                          class="text-white text-decoration-none"
                        >
                          Details
                        </router-link>
                      </button>
                      <button
                        @click="deleteHoliday(holiday.id)"
                        class="btn btn-danger mb-2"
                      >
                        Delete
                      </button>
                    </div>
                  </div>
                </div>
                <!-- edit -->
                <EditHolidayForm
                  v-if="holidayStore.showEdit && editedHoliday === holiday.id"
                  :holiday="holiday"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- create location  -->
      <div class="col-md-4">
        <div class="card pe-4" style="height: 700px">
          <div class="card-body">
            <h1>Create A New Holiday</h1>
            <form @submit.prevent="createHoliday">
              <div class="mb-3">
                <label for="location" class="form-label">Location</label>
                <select
                  v-model="formData.location"
                  class="form-select"
                  id="location"
                >
                  <option value="" disabled>Select a location</option>
                  <option
                    v-for="location in holidayStore.locations"
                    :key="location.id"
                    :value="location.id"
                  >
                    {{ location.street }} {{ location.number }},
                    {{ location.city }},
                    {{ location.country }}
                  </option>
                </select>
              </div>
              <div class="mb-3">
                <label for="title" class="form-label">Title</label>
                <input
                  v-model="formData.title"
                  type="text"
                  class="form-control"
                  id="title"
                  placeholder="Title"
                />
              </div>
              <div class="mb-3">
                <label for="startDate" class="form-label">Start Date</label>
                <input
                  v-model="formData.startDate"
                  type="date"
                  class="form-control"
                  id="startDate"
                />
              </div>
              <div class="mb-3">
                <label for="duration" class="form-label">Duration (days)</label>
                <input
                  v-model="formData.duration"
                  type="number"
                  class="form-control"
                  id="duration"
                  placeholder="Duration"
                />
              </div>
              <div class="mb-3">
                <label for="price" class="form-label">Price</label>
                <input
                  v-model="formData.price"
                  type="text"
                  class="form-control"
                  id="price"
                  placeholder="Price"
                />
              </div>
              <div class="mb-3">
                <label for="freeSlots" class="form-label">Free Slots</label>
                <input
                  v-model="formData.freeSlots"
                  type="number"
                  class="form-control"
                  id="freeSlots"
                  placeholder="Free Slots"
                />
              </div>
              <button type="submit" class="btn btn-primary">
                Create Holiday
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useHolidayStore } from "@/store/holidayStore";
import EditHolidayForm from "./EditHolidayForm.vue";

const holidayStore = useHolidayStore();

const formData = ref({
  title: "",
  startDate: "",
  duration: "",
  price: "",
  freeSlots: "",
  location: "",
});

const editedHoliday = ref(null);

const showEditForm = (holiday) => {
  editedHoliday.value = holiday.id;
  holidayStore.showEdit = true;
};

const createHoliday = async () => {
  try {
    console.log(formData.value);
    await holidayStore.createHoliday(formData.value);
  } catch (error) {
    console.error("Error creating holiday:", error);
  }
};

const deleteHoliday = async (id) => {
  try {
    const confirmed = window.confirm(
      "Are you sure you want to delete this holiday?"
    );
    if (confirmed) {
      await holidayStore.deleteHoliday(id);
    } else {
      console.log("Deletion canceled");
    }
  } catch (error) {
    console.error("Error deleting holiday:", error);
  }
};

onMounted(() => {
  holidayStore.fetchData();
});
</script>
