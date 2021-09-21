<template>
  <nav class="navbar navbar-expand navbar-light fixed-top bg-light">
    <div class="container">
      <router-link to="/" class="navbar-brand">Inicio</router-link>
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav ml-auto" v-if="!user">
          <li class="nav-item">
            <router-link to="/login" class="nav-link">Ingresar</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/register" class="nav-link">Registro</router-link>
          </li>
        </ul>
        <ul class="navbar-nav ml-auto" v-if="user">
          <li class="nav-item">
            <a href="javascript:void(0)" @click="handleClick" class="nav-link"
              >Salir</a
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "Nav",
  methods: {
    handleClick() {
      localStorage.removeItem("token");
      localStorage.removeItem("refresh_token");
      this.$store.dispatch("user", null);
      this.$router.push("/");
    },
  },
  computed: {
    ...mapGetters(["user"]),
  },
};
</script>

<style></style>
