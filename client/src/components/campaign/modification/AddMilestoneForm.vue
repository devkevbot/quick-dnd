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
        <v-textarea
          v-model.trim="milestone"
          label="Enter a milestone"
          no-resize
          outlined
          counter="100"
          :error-messages="milestoneErrors"
          @input="$v.milestone.$touch()"
          @blur="$v.milestone.$touch()"
        >
        </v-textarea>
      </v-form>

      <v-card-actions class="pt-4">
        <v-spacer></v-spacer>
        <v-btn color="error" @click.native="decline">Cancel</v-btn>
        <v-btn color="primary" @click.native="onSubmit">Add</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';

export default {
  name: 'AddMilestoneForm',
  data() {
    return {
      showPrompt: false,
      title: 'Add a campaign milestone',
      options: {
        maxWidth: 400,
        zIndex: 200,
      },

      milestone: '',
      formEventName: 'complete',
    };
  },
  methods: {
    /**
     * Opens form.
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

      const event = { milestone: this.milestone };
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
      this.milestone = null;
    },
  },
  validations: {
    milestone: {
      required,
      minLength: minLength(1),
      maxLength: maxLength(100),
    },
  },
  computed: {
    milestoneErrors() {
      const errors = [];
      if (!this.$v.milestone.$dirty) return errors;
      if (!this.$v.milestone.minLength) {
        errors.push('Milestone must be at least 1 character.');
      }
      if (!this.$v.milestone.maxLength) {
        errors.push('Milestone must not exceed 100 characters.');
      }
      if (!this.$v.milestone.required) {
        errors.push('Milestone is required.');
      }
      return errors;
    },
  },
};
</script>
