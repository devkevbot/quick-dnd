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
        <!-- New campaign location field -->
        <v-text-field
          v-model.trim="location"
          label="Update the campaign location"
          outlined
          counter="50"
          :error-messages="locationErrors"
          @input="$v.location.$touch()"
          @blur="$v.location.$touch()"
        >
        </v-text-field>

        <!-- New campaign state field -->
        <v-textarea
          v-model.trim="state"
          label="Update the campaign state"
          outlined
          counter="1024"
          :error-messages="stateErrors"
          @input="$v.state.$touch()"
          @blur="$v.state.$touch()"
        >
        </v-textarea>
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
import { required, maxLength } from 'vuelidate/lib/validators';

export default {
  name: 'ModifyCampaignForm',
  data() {
    return {
      showPrompt: false,
      title: 'Update the campaign',
      options: {
        maxWidth: 400,
        zIndex: 200,
      },

      state: '',
      location: '',
      formEventName: 'complete',
    };
  },
  methods: {
    /**
     * Opens form.
     */
    prompt(state, location) {
      this.state = state;
      this.location = location;
      this.showPrompt = true;
    },
    /**
     * Runs when the user submits the form. Checks if form data is valid
     * and closes the form if it is.
     */
    onSubmit() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      const event = { state: this.state, location: this.location };
      this.$emit(this.formEventName, event);

      this.resetForm();
      this.showPrompt = false;
    },
    /**
     * Runs when the user declines the form. Closes the form. Can be
     * triggered by keypress or mouse click.
     */
    decline() {
      this.resetForm();
      this.showPrompt = false;
    },
    /**
     * Resets the form to an initial state.
     */
    resetForm() {
      this.$nextTick(() => this.$v.$reset());
      this.state = null;
      this.location = null;
    },
  },
  validations: {
    state: {
      required,
      maxLength: maxLength(1024),
    },
    location: {
      required,
      maxLength: maxLength(50),
    },
  },
  computed: {
    stateErrors() {
      const errors = [];
      if (!this.$v.state.$dirty) return errors;
      if (!this.$v.state.maxLength) {
        errors.push('State must not exceed 1024 characters.');
      }
      if (!this.$v.state.required) {
        errors.push('State is required.');
      }
      return errors;
    },
    locationErrors() {
      const errors = [];
      if (!this.$v.location.$dirty) return errors;
      if (!this.$v.location.maxLength) {
        errors.push('Location must not exceed 50 characters.');
      }
      if (!this.$v.location.required) {
        errors.push('Location is required.');
      }
      return errors;
    },
  },
};
</script>
