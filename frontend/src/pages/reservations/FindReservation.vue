<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-4">
        <div class="card pe-4" style="height: 300px">
          <div class="card-body">
            <h1>Find A Reservation</h1>
            <div class="mt-4">
              <label for="searchPhoneNumber" class="mb-2 form-label"
                >Search by Phone Number:</label
              >
              <input
                type="text"
                class="form-control mb-5"
                id="searchPhoneNumber"
                v-model="searchPhoneNumber"
              />
            </div>
            <button class="btn btn-primary" @click="findReservation">
              Search
            </button>
          </div>
        </div>
      </div>
      <div class="col-md-8">
        <div
          v-if="foundReservation"
          class="card pe-4"
          style="height: 400px; overflow: scroll"
        >
          <div class="card-body">
            <h1 class="mb-4">Found Reservation</h1>
            <div class="card p-4">
              <div class="row">
                <div class="col-md-10">
                  <h4>{{ foundReservation.contactName }}</h4>
                  <h5>Phone Number: {{ foundReservation.phoneNumber }}</h5>

                  <p>Holiday: {{ foundReservation.holiday.title }}</p>
                  <p>Start Date: {{ foundReservation.holiday.startDate }}</p>
                  <p>
                    Location: {{ foundReservation.holiday.location.city }},
                    {{ foundReservation.holiday.location.country }}
                  </p>
                </div>
                <div class="col-md-2">
                  <button
                    class="btn btn-primary m-1"
                    @click="editFoundReservation"
                  >
                    Edit
                  </button>
                  <button
                    class="btn btn-danger"
                    @click="deleteFoundReservation"
                  >
                    Delete
                  </button>
                </div>
              </div>
            </div>
            <!-- edit -->
            <div v-if="showEditForm" class="mt-3">
              <h6>Edit Reservation</h6>
              <div class="mb-3">
                <label for="editContactName" class="form-label"
                  >Contact Name:</label
                >
                <input
                  type="text"
                  class="form-control"
                  id="editContactName"
                  v-model="editForm.contactName"
                />
              </div>
              <div class="mb-3">
                <label for="editPhoneNumber" class="form-label"
                  >Phone Number:</label
                >
                <input
                  type="text"
                  class="form-control"
                  id="editPhoneNumber"
                  v-model="editForm.phoneNumber"
                />
              </div>
              <button
                class="btn btn-success mt-4"
                @click="updateFoundReservation"
              >
                Update Reservation
              </button>
              <button
                @click="cancelEdit"
                class="btn btn-secondary mt-4 btn-cancel"
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useHolidayStore } from "@/store/holidayStore";
import { ref, onMounted } from "vue";

const holidayStore = useHolidayStore();

const searchPhoneNumber = ref("");
const foundReservation = ref(null);
const showEditForm = ref(false);

const editForm = ref({
  id: null,
  contactName: "",
  phoneNumber: "",
});

const findReservation = () => {
  const reservation = holidayStore.findReservationByPhoneNumber(
    searchPhoneNumber.value
  );

  foundReservation.value = reservation ? reservation : null;
};

const editFoundReservation = () => {
  editForm.value.contactName = foundReservation.value.contactName;
  editForm.value.phoneNumber = foundReservation.value.phoneNumber;
  showEditForm.value = true;
};

const cancelEdit = () => {
  showEditForm.value = false;
};

const updateFoundReservation = async () => {
  try {
    const updatedReservation = {
      id: foundReservation.value.id,
      contactName: editForm.value.contactName,
      phoneNumber: editForm.value.phoneNumber,
      holiday: foundReservation.value.holiday.id,
    };

    await holidayStore.updateReservation(updatedReservation);
    showEditForm.value = false;

    const updated = await holidayStore.fetchReservationById(
      foundReservation.value.id
    );
    if (updated) {
      foundReservation.value = updated;
    }
  } catch (error) {
    console.error("Error updating reservation:", error);
  }
};

const deleteFoundReservation = async () => {
  try {
    await holidayStore.deleteReservation(foundReservation.value.id);
    foundReservation.value = null;
  } catch (error) {
    console.error("Error deleting reservation:", error);
  }
};

onMounted(() => {
  holidayStore.fetchData();
});
</script>
