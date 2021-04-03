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

      <v-card-text class="pa-4" v-show="!!message">{{ message }}</v-card-text>

      <v-card-actions class="pt-4">
        <v-spacer></v-spacer>
        <v-btn color="error" @click.native="decline">Cancel</v-btn>
        <v-btn color="primary" @click.native="accept">OK</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'ConfirmDialog',
  data() {
    return {
      showPrompt: false,
      resolve: null,
      reject: null,
      title: null,
      message: null,
      options: {
        maxWidth: 400,
        zIndex: 200,
      },
    };
  },
  methods: {
    /**
     * Opens the dialogue.
     * @param {String} title - The title of the dialog.
     * @param {String} message - The body message of the dialog.
     * @param {Object} options - Additional appearance customization.
     * @returns {Promise<Boolean>} - Resolves if the prompt was accepted.
     */
    prompt(title, message, options) {
      this.showPrompt = true;
      this.title = title;
      this.message = message;
      this.options = Object.assign(this.options, options);
      return new Promise((resolve, reject) => {
        this.resolve = resolve;
        this.reject = reject;
      });
    },
    /**
     * Runs when the user accepts the prompt. Closes the prompt.
     */
    accept() {
      this.resolve(true);
      this.showPrompt = false;
    },
    /**
     * Runs when the user declines the prompt. Closes the prompt. Can be
     * triggered by keypress or mouse click.
     */
    decline() {
      this.resolve(false);
      this.showPrompt = false;
    },
  },
};
</script>
