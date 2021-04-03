<template>
  <div class="registration-form">
    <v-card width="500" class="mx-auto my-10">
      <v-card-title class="primary justify-center white--text">
        QuickDND
      </v-card-title>

      <v-card-text class="mt-2 subtitle-1 text-center text--primary">
        Register a new account
      </v-card-text>

      <v-form @submit.prevent class="mx-4 pb-2">
        <!-- Name field -->
        <v-text-field
          v-model.trim="name"
          label="Enter full name"
          aria-autocomplete="off"
          outlined
          dense
          type="text"
          required
          :error-messages="nameErrors"
          @input="$v.name.$touch()"
          @blur="$v.name.$touch()"
        >
        </v-text-field>

        <!-- Username field -->
        <v-text-field
          v-model.trim="username"
          label="Enter username"
          aria-autocomplete="off"
          outlined
          dense
          type="text"
          required
          :error-messages="usernameErrors"
          @input="$v.username.$touch()"
          @blur="$v.username.$touch()"
        >
        </v-text-field>

        <!-- Password field -->
        <v-text-field
          v-model.trim="password"
          label="Enter password"
          aria-autocomplete="off"
          outlined
          dense
          :type="displayPassword ? 'text' : 'password'"
          :append-icon="displayPassword ? 'mdi-eye' : 'mdi-eye-off'"
          @click:append="displayPassword = !displayPassword"
          :error-messages="passwordErrors"
          @input="$v.password.$touch()"
          @blur="$v.password.$touch()"
        >
        </v-text-field>

        <!-- Password confirmation field -->
        <v-text-field
          v-model.trim="passwordConf"
          label="Confirm password"
          aria-autocomplete="off"
          outlined
          dense
          :type="displaypasswordConf ? 'text' : 'password'"
          :append-icon="displaypasswordConf ? 'mdi-eye' : 'mdi-eye-off'"
          @click:append="displaypasswordConf = !displaypasswordConf"
          :error-messages="passwordConfErrors"
          @input="$v.passwordConf.$touch()"
          @blur="$v.passwordConf.$touch()"
        >
        </v-text-field>

        <v-btn
          class="mb-4"
          color="primary"
          block
          type="submit"
          @click.prevent="onSubmit"
          :loading="submitBtn.isLoading"
        >
          {{ submitBtn.text }}
          <v-icon right>{{ submitBtn.icon }}</v-icon>
        </v-btn>

        <v-card-text class="mb-n2 subtitle-1 text-center text--primary">
          Have an account already?
        </v-card-text>

        <v-card-actions>
          <v-spacer />

          <v-btn class="mb-2" color="primary" outlined to="/login">
            <v-icon left>mdi-arrow-left</v-icon>
            Back to login
          </v-btn>

          <v-spacer />
        </v-card-actions>

        <v-alert :type="alert.type" :value="alert.show">
          {{ alert.text }}
        </v-alert>
      </v-form>
    </v-card>
  </div>
</template>

<script>
import {
  required,
  minLength,
  maxLength,
  alphaNum,
  sameAs,
  helpers,
} from 'vuelidate/lib/validators';

export default {
  name: 'RegistrationForm',
  data() {
    return {
      name: '',
      username: '',
      password: '',
      passwordConf: '',

      displayPassword: false,
      displaypasswordConf: false,

      submitBtn: {
        isLoading: false,
        text: 'Complete Registration',
        icon: 'mdi-account',
      },

      alert: {
        type: 'success',
        text: '',
        show: false,
      },
    };
  },
  validations: {
    name: {
      required,
      maxLength: maxLength(50),
      /* This regex obviously does not cover all cases, but is
      sufficient for our project. */
      isNameValid: helpers.regex('nameValidation', /^[a-zA-Z '.-]*$/),
    },
    username: {
      required,
      minLength: minLength(3),
      maxLength: maxLength(25),
      alphaNum,
    },
    password: {
      required,
      minLength: minLength(10),
      maxLength: maxLength(99),
    },
    passwordConf: {
      required,
      sameAs: sameAs('password'),
    },
  },
  computed: {
    nameErrors() {
      const errors = [];
      if (!this.$v.name.$dirty) return errors;
      if (!this.$v.name.maxLength) {
        errors.push('Name must not exceed 50 characters.');
      }
      if (!this.$v.name.isNameValid) {
        errors.push('Name contains disallowed characters.');
      }
      if (!this.$v.name.required) {
        errors.push('Name is required.');
      }
      return errors;
    },
    usernameErrors() {
      const errors = [];
      if (!this.$v.username.$dirty) return errors;
      if (!this.$v.username.minLength) {
        errors.push('Username must be at least 3 characters.');
      }
      if (!this.$v.username.maxLength) {
        errors.push('Username must not exceed 25 characters.');
      }
      if (!this.$v.username.alphaNum) {
        errors.push('Username must contain only letters and numbers 0-9.');
      }
      if (!this.$v.username.required) {
        errors.push('Username is required.');
      }
      return errors;
    },
    passwordErrors() {
      const errors = [];
      if (!this.$v.password.$dirty) return errors;
      if (!this.$v.password.minLength) {
        errors.push('Password must be at least 10 characters.');
      }
      if (!this.$v.password.maxLength) {
        errors.push('Password must not exceed 99 characters.');
      }
      if (!this.$v.password.required) {
        errors.push('Password is required.');
      }
      return errors;
    },
    passwordConfErrors() {
      const errors = [];
      if (!this.$v.passwordConf.$dirty) return errors;
      if (!this.$v.passwordConf.sameAs) {
        errors.push('Password confirmation must match.');
      }
      if (!this.$v.passwordConf.required) {
        errors.push('Password confirmation is required.');
      }
      return errors;
    },
  },
  methods: {
    async onSubmit() {
      this.alert.show = false;

      this.$v.$touch();
      if (this.$v.$invalid) return;

      this.changeLoadingState(true);
      await this.sendRegistrationRequest();
      this.changeLoadingState(false);
    },
    changeLoadingState(isLoading) {
      this.submitBtn.isLoading = !!isLoading;
    },
    async sendRegistrationRequest() {
      const requestURI = '/register';
      const data = {
        username: this.username,
        password: this.password,
        name: this.name,
      };
      const method = 'POST';

      await this.$http({ url: requestURI, data, method })
        .then(() => {
          this.onSuccess();
          setTimeout(() => {
            this.$router.push('/login');
          }, 5000);
        })
        .catch(() => {
          this.onFailure();
        });
    },
    onSuccess() {
      this.$v.$reset();

      this.name = '';
      this.username = '';
      this.password = '';
      this.passwordConf = '';

      this.submitBtn.text = 'Successfully registered';
      this.submitBtn.icon = 'mdi-account-check';

      this.alert.type = 'success';
      this.alert.show = true;
      this.alert.text = `Registration complete,
        taking you back to the login page in 5 seconds.`;
    },
    onFailure() {
      this.alert.type = 'error';
      this.alert.show = true;
      this.alert.text = `Registration failed!
        Please try again.`;
    },
  },
};
</script>
