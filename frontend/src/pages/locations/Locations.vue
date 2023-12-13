<template>
  <div class="container mt-5">
    <div class="row">
      <!-- List of Locations -->
      <div class="col-md-8">
        <div class="card pe-4" >
          <div class="card-body">
            <h1>Locations</h1>
            <div v-if="holidayStore.loading" class="mb-3">Loading...</div>
            <div v-if="!holidayStore.loading">
              <div
                v-for="location in holidayStore.locations"
                :key="location.id"
                class="card mb-3"
              >
                <div
                  class="card-body d-flex justify-content-between align-items-center"
                >
                  <p class="card-text">
                    "{{ location.street }}" №{{ location.number }}<br />
                    {{ location.city }}, {{ location.country }}
                  </p>
                  <div>
                    <button
                      @click="showEditForm(location)"
                      class="btn btn-primary me-2"
                    >
                      Edit
                    </button>
                    <router-link
                      :to="{
                        name: 'LocationDetails',
                        query: { id: location.id },
                      }"
                      class="btn btn-secondary me-2"
                    >
                      Details
                    </router-link>
                    <button
                      @click="deleteLocation(location.id)"
                      class="btn btn-danger me-2"
                    >
                      Delete
                    </button>
                  </div>
                </div>
                <!-- Edit Form -->
                <EditLocationForm
                  v-if="holidayStore.showEdit && editedLocation === location.id"
                  :location="location"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- create -->
      <div class="col-md-4">
        <div class="card pe-4" style="height: 550px">
          <div class="card-body">
            <h1>Create A New Location</h1>
            <form @submit.prevent="createLocation" class="mt-4">
              <div class="form-group">
                <label for="street" class="form-label">Street</label>
                <input
                  v-model="formData.street"
                  type="text"
                  class="form-control"
                  id="street"
                  placeholder="Street"
                  required
                />
              </div>
              <div class="form-group">
                <label for="number" class="form-label">Number</label>
                <input
                  v-model="formData.number"
                  type="text"
                  class="form-control"
                  id="number"
                  placeholder="Number"
                  required
                />
              </div>
              <div class="form-group">
                <label for="city" class="form-label">City</label>
                <input
                  v-model="formData.city"
                  type="text"
                  class="form-control"
                  id="city"
                  placeholder="City"
                  required
                />
              </div>
              <div class="form-group">
                <label for="country" class="form-label">Country</label>
                <input
                  v-model="formData.country"
                  type="text"
                  class="form-control"
                  id="country"
                  placeholder="Country"
                  required
                />
              </div>
              <button type="submit" class="btn btn-primary mt-3">
                Create Location
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
import EditLocationForm from "./EditLocationForm.vue";

const holidayStore = useHolidayStore();

const formData = ref({
  street: "",
  number: "",
  city: "",
  country: "",
});

const editedLocation = ref(null);

const showEditForm = (location) => {
  editedLocation.value = location.id;
  holidayStore.showEdit = true;
};

const createLocation = async () => {
  try {
    await holidayStore.createLocation(formData.value);
    holidayStore.fetchData();
  } catch (error) {
    console.error("Error creating location:", error);
  }
};

const deleteLocation = async (id) => {
  try {
    await holidayStore.deleteLocation(id);
    holidayStore.fetchData();
  } catch (error) {
    console.error("Error deleting location:", error);
  }
};

onMounted(() => {
  holidayStore.fetchData();
});
</script>
