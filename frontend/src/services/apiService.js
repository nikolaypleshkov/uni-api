import axios from "axios";

const BASE_URL = "http://localhost:5173/travel-agency";

const instance = axios.create({
  baseURL: BASE_URL,
});

const api = {
  // holidays
  async fetchHolidays() {
    try {
      const response = await instance.get("/holidays");
      return response.data;
    } catch (error) {
      console.error("Error fetching journeys:", error);
      throw error;
    }
  },

  async fetchHolidaysForReservation(location, startDate, duration) {
    try {
      const params = {};

      if (location) {
        params.location = location;
      }
      if (startDate) {
        params.startDate = startDate;
      }
      if (duration) {
        params.duration = duration;
      }

      const response = await instance.get("/holidays", {
        params: params,
      });

      return response.data;
    } catch (error) {
      console.error("Error fetching holidays:", error);
      throw error;
    }
  },

  async fetchHolidayById(id) {
    try {
      const response = await instance.get(`/holidays/${id}`);
      return response.data;
    } catch (error) {
      console.error("Error fetching journey:", error);
      throw error;
    }
  },

  async createHoliday(holiday) {
    try {
      const response = await instance.post("/holidays", holiday);
      return response.data;
    } catch (error) {
      console.error("Error creating journey:", error);
      throw error;
    }
  },

  async updateHoliday(holiday) {
    try {
      await instance.put(`/holidays`, holiday);
    } catch (error) {
      console.error("Error updating journey:", error);
      throw error;
    }
  },

  async deleteHoliday(id) {
    try {
      await instance.delete(`/holidays/${id}`);
    } catch (error) {
      console.error("Error deleting journey:", error);
      throw error;
    }
  },

  // locations
  async fetchLocations() {
    try {
      const response = await instance.get("/locations");
      return response.data;
    } catch (error) {
      console.error("Error fetching locations:", error);
      throw error;
    }
  },

  async fetchLocationById(id) {
    try {
      const response = await instance.get(`/locations/${id}`);
      return response.data;
    } catch (error) {
      console.error("Error fetching location:", error);
      throw error;
    }
  },

  async createLocation(locationData) {
    try {
      await instance.post("/locations", locationData);
    } catch (error) {
      console.error("Error creating location:", error);
      throw error;
    }
  },

  async updateLocation(location) {
    try {
      await instance.put(`/locations`, location);
    } catch (error) {
      console.error("Error updating location:", error);
      throw error;
    }
  },

  async deleteLocation(id) {
    try {
      await instance.delete(`/locations/${id}`);
    } catch (error) {
      console.error("Error deleting location:", error);
      throw error;
    }
  },

  // reservations
  async fetchReservations() {
    try {
      const response = await instance.get("/reservations");
      return response.data;
    } catch (error) {
      console.error("Error fetching reservations:", error);
      throw error;
    }
  },

  async fetchReservationById(id) {
    try {
      const response = await instance.get(`/reservations/${id}`);
      return response.data;
    } catch (error) {
      console.error("Error fetching reservation:", error);
      throw error;
    }
  },

  async createReservation(reservation) {
    try {
      await instance.post("/reservations", reservation);
    } catch (error) {
      console.error("Error creating reservation:", error);
      throw error;
    }
  },

  async updateReservation(reservation) {
    try {
      await instance.put(`/reservations`, reservation);
    } catch (error) {
      console.error("Error updating reservation:", error);
      throw error;
    }
  },

  async deleteReservation(id) {
    try {
      await instance.delete(`/reservations/${id}`);
    } catch (error) {
      console.error("Error deleting reservation:", error);
      throw error;
    }
  },
};

export default api;
