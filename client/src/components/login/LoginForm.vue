<!-- LoginForm.vue
This component is responsible for rendering the login fields that are
required for a player to authenticate.
--->

<template>
  <div class="login-form">
    <v-card width="500" class="mx-auto my-10">
      <v-card-title class="primary justify-center white--text">
        QuickDND
      </v-card-title>

      <v-card-text class="mt-2 subtitle-1 text-center text--primary">
        Log in to your account
      </v-card-text>

      <v-form @submit.prevent class="mx-4 pb-2">
        <!-- Username field -->
        <v-text-field
          v-model.trim="credentials.username"
          label="Enter username"
          aria-autocomplete="off"
          outlined
          dense
          type="text"
        >
        </v-text-field>

        <!-- Password field -->
        <v-text-field
          v-model.trim="credentials.password"
          label="Enter password"
          aria-autocomplete="off"
          outlined
          dense
          :type="displayPassword ? 'text' : 'password'"
          :append-icon="displayPassword ? 'mdi-eye' : 'mdi-eye-off'"
          @click:append="displayPassword = !displayPassword"
        >
        </v-text-field>

        <v-btn
          class="mb-4"
          @click.prevent="sendLoginRequest"
          color="primary"
          block
          type="submit"
        >
          Log in
        </v-btn>

        <v-card-text class="mb-n2 subtitle-1 text-center text--primary">
          Need an account?
        </v-card-text>

        <v-card-actions>
          <v-spacer />

          <v-btn class="mb-2" outlined to="/register">
            <v-icon left>mdi-account</v-icon>
            Register
          </v-btn>

          <v-spacer />
        </v-card-actions>

        <FailedLoginAlert :errorText="errorAlertText" :show="showErrorAlert">
        </FailedLoginAlert>
      </v-form>
    </v-card>
  </div>
</template>

<script>
import { mapActions } from 'vuex';
import FailedLoginAlert from './FailedLoginAlert.vue';

export default {
  name: 'LoginForm',
  components: {
    FailedLoginAlert,
  },
  data() {
    return {
      credentials: {
        username: '',
        password: '',
      },

      displayPassword: false,

      errorAlertText: '',
      showErrorAlert: false,
    };
  },
  methods: {
    async sendLoginRequest() {
      await this.login({
        username: this.credentials.username,
        password: this.credentials.password,
      })
        .then(() => {
          this.showErrorAlert = false;
          this.$router.push('/');
        })
        .catch(() => {
          this.errorAlertText = 'Login unsuccessful. Try again.';
          this.showErrorAlert = true;
        });
    },
    ...mapActions({
      login: 'authentication/loginPlayer',
    }),
  },
};
</script>
