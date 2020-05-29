<template>
  <div class="login">
    <section class="logo">
      <img class="brand-logo" src="../assets/ganeti_logo.svg" />
    </section>
    <section class="login-form">
      <Form @submit="login" :fields="formFields" :values="credentials" />
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import Api from "@/store/api";
import PageNames from "@/data/enum/PageNames";
import { REDIRECT_TO_QUERY_KEY } from "@/router";
import Input from "@/components/form/Input.vue";
import Form from "@/components/form/Form.vue";

@Component({
  name: "LoginView",
  components: { Input, Form }
})
export default class LoginView extends Vue {
  credentials = {
    username: "",
    password: "",
    remember: false
  };

  formFields = [
    {
      name: "username",
      label: "Username",
      required: true
    },
    {
      name: "password",
      label: "Password",
      type: "password",
      required: true
    },
    {
      name: "remember",
      label: "Remember on this computer",
      type: "checkbox"
    }
  ];

  loading = false;
  error = "";

  async login() {
    this.error = "";
    this.loading = true;
    const { username, password } = this.credentials;
    const response = await Api.login({ username, password });
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
