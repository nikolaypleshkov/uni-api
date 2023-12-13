export const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/pages/holidays/Holidays.vue"),
  },
  {
    path: "/locations",
    name: "Locations",
    component: () => import("@/pages/locations/Locations.vue"),
  },
  {
    path: "/location-details",
    name: "LocationDetails",
    component: () => import("@/pages/locations/LocationDetails.vue"),
  },
  {
    path: "/holidays",
    name: "Holidays",
    component: () => import("@/pages/holidays/Holidays.vue"),
  },
  {
    path: "/holiday-details",
    name: "HolidayDetails",
    component: () => import("@/pages/holidays/HolidayDetails.vue"),
  },
  {
    path: "/reservations",
    name: "Reservations",
    component: () => import("@/pages/reservations/Reservations.vue"),
  },
  {
    path: "/find-reservation",
    name: "FindReservation",
    component: () => import("@/pages/reservations/FindReservation.vue"),
  },
];
