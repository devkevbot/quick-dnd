<!-- AddCharacterDialog.vue
This component is the dialog that pops up when a user wishes to
add a new character to their campaign.
-->

<template>
  <div class="add-character-dialog">
    <v-dialog
      v-model="showPrompt"
      :max-width="400"
      :style="{ zIndex: 200 }"
      @keydown.esc="decline"
    >
      <v-card>
        <v-toolbar dark color="primary" dense flat>
          <v-toolbar-title>Add New Character to Campaign</v-toolbar-title>
        </v-toolbar>
        <v-card-text class="pa-4"
          >Enter character ID to add to this campaign:</v-card-text
        >

        <v-text-field
          class="ml-4 mr-4"
          v-model="id"
          :counter="50"
          label="ID"
          required
          :error-messages="idErrors"
          @input="$v.id.$touch()"
          @blur="$v.id.$touch()"
        ></v-text-field>

        <v-card-actions class="pt-4">
          <v-spacer></v-spacer>
          <v-btn color="error" @click.native="decline">Cancel</v-btn>
          <v-btn color="primary" @click.native="accept">OK</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import {
  required,
  numeric,
  maxLength,
} from 'vuelidate/lib/validators';

export default {
  name: 'AddCharacterDialog',
  data() {
    return {
      showPrompt: false,
      resolve: null,
      reject: null,
      id: null,
    };
  },
  validations: {
    id: {
      required,
      numeric,
      maxLength: maxLength(50),
    },
  },
  methods: {
    prompt() {
      this.showPrompt = true;
      return new Promise((resolve, reject) => {
        this.resolve = resolve;
        this.reject = reject;
      });
    },
    /**
     * Runs when the user accepts the prompt. Closes the prompt.
     */
    accept() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      this.resolve(true);
      this.showPrompt = false;
      this.$v.$reset();
    },
    /**
     * Runs when the user declines the prompt. Closes the prompt. Can be
     * triggered by keypress or mouse click.
     */
    decline() {
      this.resolve(false);
      this.showPrompt = false;
      this.$v.$reset();
    },
    /**
     * Clears ID from the text box
     */
    clearID() {
      this.id = null;
    },
  },
  computed: {
    idErrors() {
      const errors = [];
      if (!this.$v.id.$dirty) return errors;
      if (!this.$v.id.required) {
        errors.push('ID is required.');
      }
      if (!this.$v.id.numeric) {
        errors.push('ID must be numeric.');
      }
      if (!this.$v.id.maxLength) {
        errors.push('Max length exceeded.');
      }
      return errors;
    },
    getID() {
      return this.id;
    },
  },
};
</script>
