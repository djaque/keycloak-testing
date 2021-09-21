<template>
  <div id="app">
    <Nav />
    <div class="auth-wrapper">
      <div class="container">
        <div class="row">
          <div class="col"></div>
          <div class="col">
            <router-view />
          </div>
          <div class="col"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

import Nav from "./components/Nav.vue";

export default {
  name: "App",
  components: {
    Nav,
  },
  async created() {
    const response = await axios.get("userinfo");
    this.$store.dispatch("user", response.data);
  },
};
</script>

<style>
body {
  background-color: lightblue;
}
.auth-wrapper {
  padding-top: 60px;
}
.auth-inner {
  background-color: #fff;
  border-radius: 25px;
  border: 2px solid #73ad21;
  padding: 20px;
  height: 150px;
}
</style>
