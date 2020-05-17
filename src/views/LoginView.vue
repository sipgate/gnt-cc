<template>
  <div class="login">
    <section class="logo">
      <img class="brand-logo" src="../assets/ganeti_logo.svg" />
    </section>
    <section class="login-form">
      <form @submit.prevent="login">
        <b-field>
          <b-input
            icon="user"
            placeholder="Username"
            required
            type="text"
            v-model="credentials.username"
          />
        </b-field>
        <b-field>
          <b-input
            password-reveal
            icon="lock"
            placeholder="Password"
            required
            type="password"
            v-model="credentials.password"
          />
        </b-field>
        <div class="login-error">
          <p class="error">{{ error }}</p>
        </div>
        <b-button :loading="loading" native-type="submit" type="is-primary">Login</b-button>
      </form>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ClusterStatsDigits from "@/components/ClusterStatsDigits.vue";
import Component from "vue-class-component";
import Api from "@/store/api";
import PageNames from "@/data/enum/PageNames";
import { REDIRECT_TO_QUERY_KEY } from "@/router";

@Component({
  name: "LoginView",
  components: {
    ClusterStatsDigits
  }
})
export default class LoginView extends Vue {
  credentials = {
    username: "",
    password: ""
  };

  loading = false;
  error = "";

  async login() {
    this.error = "";
    this.loading = true;
    const response = await Api.login(this.credentials);
    this.loading = false;

    if (response.code === 401) {
      this.error = "Wrong username or password.";
      return;
    }

    if (response.code === 500) {
      this.error = "Internal server error.";
      return;
    }

    if (response.code !== 200) {
      this.error = "Unknown error.";
      return;
    }

    await this.$store.dispatch("saveToken", response.token);
    if (typeof this.$route.query[REDIRECT_TO_QUERY_KEY] !== "undefined") {
      await this.$router.push(this.$route.query[REDIRECT_TO_QUERY_KEY] as string);
    } else {
      await this.$router.push({ name: PageNames.Statistics });
    }
  }
}
</script>

<style scoped lang="scss">
.login {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 4rem;

  .logo {
    margin: 2rem 0;

    .brand-logo {
      width: 160px;
      height: auto;
    }
  }

  .error {
    height: 4rem;
    display: flex;
    justify-content: center;
    align-items: center;
    color: red;
  }
}
</style>
