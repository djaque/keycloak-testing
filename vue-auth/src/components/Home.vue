<template>
  <div class="auth-inner p-5">
    <h1 v-if="user">Hola, {{ user.name }}</h1>
    <h1 v-if="!user">¡Bienvenido!</h1>
  </div>
</template>
<script>
import { mapGetters } from "vuex";
export default {
  name: "Home",
  computed: {
    ...mapGetters(["user"]),
  },
};
</script>
