<template>
  <form @submit.prevent="handleSubmit">
    <h3>Ingresar</h3>
    <div class="form-group">
      <label for="email">Email</label>
      <input
        type="email"
        name="email"
        id="email"
        placeholder="Email"
        class="form-control"
        v-model="email"
      />
    </div>
    <div class="form-group">
      <label for="password">Contraseña</label>
      <input
        type="password"
        name="password"
        id="password"
        placeholder="Contraseña"
        class="form-control"
        v-model="password"
      />
    </div>
    <button class="btn btn-primary btn-block">Ingresar</button>
  </form>
</template>
<script>
import axios from "axios";
export default {
  name: "Login",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    async handleSubmit() {
      localStorage.removeItem("token");
      localStorage.removeItem("refresh_token");
      const response = await axios.post("login", {
        email: this.email,
        password: this.password,
      });

      localStorage.setItem("token", response.data.access_token);
      localStorage.setItem("refresh_token", response.data.refresh_token);

      const userinfo = await axios.get("userinfo", {
        headers: {
          Authorization: `Bearer ${response.data.access_token}`,
        },
      });
      this.$store.dispatch("user", userinfo.data);

      this.$router.push("/");
    },
  },
};
</script>
