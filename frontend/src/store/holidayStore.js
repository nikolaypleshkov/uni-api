import { defineStore } from "pinia";
import api from "@/services/apiService";

export const useHolidayStore = defineStore({
  id: "holidayStore",
  state: () => ({
    holidays: [],
    locations: [],
    reservations: [],
    loading: true,
    showEdit:false,
  }),
  actions: {
    // holidays
    async fetchHolidays() {
      try {
        this.loading = true;
        const response = await api.fetchHolidays();
        this.holidays = response;
      } catch (error) {
        console.error("Error fetching holidays:", error);
      } finally {
        this.loading = false;
      }
    },
    async fetchHolidaysForReservation(location, startDate, duration) {
      try {
        this.loading = true;
        const response = await api.fetchHolidaysForReservation(
          location,
          startDate,
          duration
        );
        return response;
      } catch (error) {
        console.error("Error fetching holidays:", error);
      } finally {
        this.loading = false;
      }
    },

    async fetchHolidayById(id) {
      try {
        const response = await api.fetchHolidayById(id);
        return response;
      } catch (error) {
        console.error("Error fetching holiday:", error);
      }
    },
    async createHoliday(newHoliday) {
      try {
        await api.createHoliday(newHoliday);
        this.fetchHolidays();
      } catch (error) {
        console.error("Error creating holiday:", error);
      }
    },

    async updateHoliday(updatedHoliday) {
      try {
        await api.updateHoliday(updatedHoliday);
        this.fetchHolidays();
      } catch (error) {
        console.error("Error updating holiday:", error);
      }
    },

    async deleteHoliday(holidayId) {
      try {
        await api.deleteHoliday(holidayId);
        this.fetchHolidays();
      } catch (error) {
        console.error("Error deleting holiday:", error);
      }
    },

    // locations
    async fetchLocations() {
      try {
        this.loading = true;
        const response = await api.fetchLocations();
        this.locations = response;
      } catch (error) {
        console.error("Error fetching locations:", error);
      } finally {
        this.loading = false;
      }
    },

    async fetchLocationById(id) {
      try {
        const response = await api.fetchLocationById(id);
        return response;
      } catch (error) {
        console.error("Error fetching location:", error);
      }
    },

    async createLocation(newLocation) {
      try {
        await api.createLocation(newLocation);
        this.fetchLocations();
      } catch (error) {
        console.error("Error creating holiday:", error);
      }
    },

    async updateLocation(updatedLocation) {
      try {
        await api.updateLocation(updatedLocation);
        this.fetchLocations();
      } catch (error) {
        console.error("Error updating holiday:", error);
      }
    },

    async deleteLocation(locationId) {
      try {
        await api.deleteLocation(locationId);
        this.fetchLocations();
      } catch (error) {
        console.error("Error deleting holiday:", error);
      }
    },

    // reservations
    async fetchReservations() {
      try {
        const response = await api.fetchReservations();
        this.reservations = response;
      } catch (error) {
        console.error("Error fetching reservations:", error);
      } finally {
        this.loading = false;
      }
    },

    async fetchReservationById(id) {
      try {
        const response = await api.fetchReservationById(id);
        return response;
      } catch (error) {
        console.error("Error fetching reservation:", error);
      }
    },

    findReservationByPhoneNumber(phoneNumber) {
      return this.reservations.find((r) => r.phoneNumber === phoneNumber);
    },

    async createReservation(newReservation) {
      try {
        await api.createReservation(newReservation);
        this.fetchReservations();
      } catch (error) {
        console.error("Error creating reservation:", error);
      }
    },

    async updateReservation(updatedReservation) {
      try {
        await api.updateReservation(updatedReservation);
        this.fetchReservations();
      } catch (error) {
        console.error("Error updating reservation:", error);
      }
    },

    async deleteReservation(reservationId) {
      try {
        await api.deleteReservation(reservationId);
        this.fetchReservations();
      } catch (error) {
        console.error("Error deleting reservation:", error);
      }
    },

    async fetchData() {
      await Promise.all([
        this.fetchHolidays(),
        this.fetchLocations(),
        this.fetchReservations(),
      ]);
    },
  },
});
