<template>
  <header>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container">
        <router-link to="/" class="navbar-brand">Travel Agency</router-link>
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
          aria-controls="navbarNav"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav ms-auto">
            <li class="nav-item">
              <router-link to="/holidays" class="nav-link">Holidays</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/locations" class="nav-link">Locations</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/reservations" class="nav-link">Make a reservation</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/find-reservation" class="nav-link">Find a reservation</router-link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>
</template>

<script>
import { RouterLink } from "vue-router";
</script>

<style scoped>
.navbar-brand,
.navbar-nav .nav-link {
  color: #fff;
}

</style>
