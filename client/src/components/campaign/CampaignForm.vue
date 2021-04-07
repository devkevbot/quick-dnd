<!-- CampaignForm.vue
This component is the view that allows the user to create a new
campaign. Assumes the user is authenticated.
-->

<template>
  <div class="campaign-form">
    <v-form>
      <v-container class="mb-1">
        <v-row class="ml-n3 mb-1">
          <p class="display-1 primary--text">Create a Campaign</p>
        </v-row>

        <v-row class="ml-n2 mt-2 mb-n1">
          <p class="subtitle-2 primary--text">Info</p>
        </v-row>

        <v-row class="mb-3">
          <v-col cols="12" md="6">
            <v-text-field
              :counter="50"
              label="Campaign Name"
              required
            ></v-text-field>
          </v-col>

          <v-col cols="12" md="6">
            <v-text-field
              :counter="50"
              label="Location"
              required
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-textarea filled label="Current State"></v-textarea>
        </v-row>

        <v-row class="ml-n2 mb-1">
          <p class="subtitle-2 primary--text">Campaign Characters</p>
        </v-row>

        <AddCharacterTable ref="addCharacterTable" />

        <v-row class="ml-0 mt-3">
          <v-btn @click="showAddCharacterDialog()">Add Character</v-btn>
        </v-row>

        <v-row class="ml-0 mt-10">
          <v-btn color="primary">Create</v-btn>
        </v-row>

        <AddCharacterDialog ref="addCharacterDialog" />
      </v-container>
    </v-form>
  </div>
</template>

<script>
import AddCharacterDialog from './AddCharacterDialog.vue';
import AddCharacterTable from './AddCharacterTable.vue';

export default {
  name: 'CampaignForm',
  components: {
    AddCharacterDialog,
    AddCharacterTable,
  },
  methods: {
    async showAddCharacterDialog() {
      if (await this.$refs.addCharacterDialog.prompt()) {
        await this.$refs.addCharacterTable.addItem('null', 'null', 'null'); // Test values for now
      }
    },
  },
};
</script>
