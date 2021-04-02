<template>
  <v-dialog
    v-model="showPrompt"
    :max-width="options.maxWidth"
    :style="{ zIndex: options.zIndex }"
    @keydown.esc="decline"
  >
    <v-card>
      <v-toolbar dark color="primary" dense flat>
        <v-toolbar-title>{{ title }}</v-toolbar-title>
      </v-toolbar>

      <v-form @submit.prevent class="mx-4 pb-2 pt-4">
        <!-- New password field -->
        <v-text-field
          v-model.trim="newPassword"
          label="Enter a new password"
          aria-autocomplete="off"
          outlined
          dense
          :type="displayNewPassword ? 'text' : 'password'"
          :append-icon="displayNewPassword ? 'mdi-eye' : 'mdi-eye-off'"
          @click:append="displayNewPassword = !displayNewPassword"
          :error-messages="newPasswordErrors"
          @input="$v.newPassword.$touch()"
          @blur="$v.newPassword.$touch()"
        >
        </v-text-field>

        <!-- Password confirmation field -->
        <v-text-field
          v-model.trim="confirmation"
          label="Confirm new password"
          aria-autocomplete="off"
          outlined
          dense
          :type="displayConfirmation ? 'text' : 'password'"
          :append-icon="displayConfirmation ? 'mdi-eye' : 'mdi-eye-off'"
          @click:append="displayConfirmation = !displayConfirmation"
          :error-messages="confirmationErrors"
          @input="$v.confirmation.$touch()"
          @blur="$v.confirmation.$touch()"
        >
        </v-text-field>
      </v-form>

      <v-card-actions class="pt-4">
        <v-spacer></v-spacer>
        <v-btn color="error" @click.native="decline">Cancel</v-btn>
        <v-btn color="primary" @click.native="onSubmit">Submit</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import {
  required,
  minLength,
  maxLength,
  sameAs,
} from 'vuelidate/lib/validators';

export default {
  name: 'ChangePasswordForm',
  data() {
    return {
      showPrompt: false,
      title: 'Change your password',
      options: {
        maxWidth: 400,
        zIndex: 200,
      },

      newPassword: null,
      displayNewPassword: false,

      confirmation: null,
      displayConfirmation: false,

      formEventName: 'changePasswordFormEvent',
    };
  },
  methods: {
    /**
     * Opens the change password form.
     */
    prompt() {
      this.showPrompt = true;
    },
    /**
     * Runs when the user submits the form. Checks if form data is valid
     * and closes the form if it is.
     */
    onSubmit() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      this.$emit(this.formEventName, {
        success: true,
        newPassword: this.newPassword,
        confirmation: this.confirmation,
      });

      this.resetForm();
      this.showPrompt = false;
    },
    /**
     * Runs when the user declines the form. Closes the form. Can be
     * triggered by keypress or mouse click.
     */
    decline() {
      this.$emit(this.formEventName, {
        success: false,
        newPassword: null,
        confirmation: null,
      });

      this.resetForm();
      this.showPrompt = false;
    },
    /**
     * Resets the change password form to an initial state.
     */
    resetForm() {
      this.$nextTick(() => this.$v.$reset());

      this.newPassword = null;
      this.displayNewPassword = false;
      this.confirmation = null;
      this.displayConfirmation = false;
    },
  },
  validations: {
    newPassword: {
      required,
      minLength: minLength(10),
      maxLength: maxLength(99),
    },
    confirmation: {
      required,
      sameAs: sameAs('newPassword'),
    },
  },
  computed: {
    newPasswordErrors() {
      const errors = [];
      if (!this.$v.newPassword.$dirty) return errors;
      if (!this.$v.newPassword.minLength) {
        errors.push('Password must be at least 10 characters.');
      }
      if (!this.$v.newPassword.maxLength) {
        errors.push('Password must not exceed 99 characters.');
      }
      if (!this.$v.newPassword.required) {
        errors.push('Password is required.');
      }
      return errors;
    },
    confirmationErrors() {
      const errors = [];
      if (!this.$v.confirmation.$dirty) return errors;
      if (!this.$v.confirmation.sameAs) {
        errors.push('Password confirmation must match.');
      }
      if (!this.$v.confirmation.required) {
        errors.push('Password confirmation is required.');
      }
      return errors;
    },
  },
};
</script>
