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
