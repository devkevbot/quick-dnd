<!-- LoginForm.vue
This component is responsible for rendering the login fields that are
required for a player to authenticate.
--->

<template>
  <div class="login-form">
    <v-card height="100" class="my-2 align-center">
      <v-card-title>Campaign Milestones</v-card-title>
      <v-btn @click.prevent="sendLoginRequest"></v-btn>
    </v-card>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'LoginForm',
  data() {
    return {
      credentials: {
        username: '',
        password: '',
      },
    };
  },
  methods: {
    async sendLoginRequest() {
      await this.login({
        username: this.credentials.username,
        password: this.credentials.password,
      })
        .then(() => this.$router.push('/'))
        .catch((err) => {
          console.log(err.message);
        });
    },
    ...mapActions({
      login: 'authentication/loginPlayer',
    }),
  },
};
</script>
