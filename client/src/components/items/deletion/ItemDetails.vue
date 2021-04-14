<!-- ItemDetails.vue -->
<!-- Displays all details for a specific character's items.
The player is then able to delete the item. -->

<template>
  <v-container class="item-info">
    <!-- Character selection dropdown -->
    <v-row>
      <v-col cols="4">
        <p class="headline primary--text">Select your character</p>
        <v-select
          solo
          prepend-inner-icon="mdi-account-search"
          :items="characterNames"
          v-model="selectedCharName"
        >
          <!-- Button used to refresh the character list. -->
          <template v-slot:append-outer>
            <v-btn
              color="primary"
              fab
              small
              class="mt-n2"
              @click="fetchUserCharacters"
            >
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </template>
        </v-select>
      </v-col>

      <!-- Item selection dropdown -->
      <v-col cols="4" v-if="selectedCharName">
        <p class="headline primary--text">Select your item</p>
        <v-select
          solo
          no-data-text="This character has no items."
          prepend-inner-icon="mdi-account-search"
          :items="['Character Items Here']"
        >
          <!-- Button used to refresh the item list. -->
          <template v-slot:append-outer>
            <v-btn color="primary" fab small class="mt-n2">
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </template>
        </v-select>
      </v-col>

      <!-- Item selection dropdown -->
    </v-row>

    <!-- Item Name -->
    <v-card class="pl-2">
      <v-card-title class="display-1 primary--text">
        Item Name Here
      </v-card-title>

      <!-- Delete button -->
      <v-btn class="error"> Delete Item </v-btn>

      <!-- This container displays all item information -->
      <v-container>
        <p class="headline primary--text">Item Information</p>
        <v-row>
          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Type"
              :value="'Item Type Here'"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Rarity"
              :value="'Item Rarity Here'"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Weight"
              :value="'Item Weight Here'"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Gold Value"
              :value="'Item Gold Value Here'"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Quantity"
              :value="'Item Quantity Here'"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" lg="12">
            <v-textarea
              readonly
              no-resize
              label="Description"
              outlined
              :value="'Item Description Here'"
            ></v-textarea>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-container>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'ItemDetails',
  data() {
    return {
      characters: [],
      selectedCharName: null,

      items: [],
      selectedItemName: null,
    };
  },
  /**
   * Fetches all of the user's character and item data when the component loads.
   */
  async created() {
    await this.fetchUserCharacters();
  },
  computed: {
    /**
     * @returns {Array<String>} - The names of the user's characters (if
     * any).
     */
    characterNames() {
      if (this.characters === null) return [];
      return this.characters.map((x) => x.name);
    },
    /**
     * @returns {String} - The character currently selected by the user.
     */
    selectedCharacter() {
      if (!this.selectedCharName) return {};

      const target = this.selectedCharName;
      return this.characters.filter((x) => x.name === target)[0];
    },
  },
  methods: {
    ...mapActions({
      display: 'notifications/display',
    }),
    /**
     * Sends an HTTP request to fetch the user's characters.
     */
    async fetchUserCharacters() {
      const requestURI = 'auth/character/me';
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          const fetchedCharacters = resp.data.data.characters;
          if (fetchedCharacters) {
            this.characters = fetchedCharacters;
            this.selectedCharName = fetchedCharacters[0].name;
          } else {
            this.characters = [];
            this.selectedCharName = null;
          }
        })
        .catch((err) => {
          let message = 'Could not fetch characters. Please try again.';
          if (err.response) {
            message = `Error: ${err.response.data.message}. Please try again.`;
          }
          this.display({
            message,
            color: 'error',
            timeout: 10000,
          });
        });
    },
  },
};
</script>
