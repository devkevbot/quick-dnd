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
          @change="fetchCharacterItems()"
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
          :items="itemNames"
          v-model="selectedItemName"
        >
          <!-- Button used to refresh the item list. -->
          <template v-slot:append-outer>
            <v-btn
              color="primary"
              fab
              small
              class="mt-n2"
              @click="fetchCharacterItems()"
            >
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </template>
        </v-select>
      </v-col>

      <!-- Item selection dropdown -->
    </v-row>

    <!-- Item Name -->
    <v-card v-if="characters.length && items.length" class="pl-2">
      <v-card-title class="display-1 primary--text">
        {{ selectedItem.item_name }}
      </v-card-title>

      <!-- Delete button -->
      <v-card-actions>
        <v-btn class="error" @click="displayItemDeletionPrompt">
          Delete Item
          <v-icon class="ml-2">mdi-delete-forever</v-icon>
        </v-btn>

        <ConfirmDialog ref="confirmDelete" />
      </v-card-actions>

      <!-- This container displays all item information -->
      <v-container>
        <p class="headline primary--text">Item Information</p>
        <v-row>
          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Type"
              :value="selectedItem.type"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Rarity"
              :value="selectedItem.rarity"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Weight"
              :value="selectedItem.weight"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Gold Value"
              :value="selectedItem.gold_value"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Quantity"
              :value="selectedItem.quantity"
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
              :value="selectedItem.description"
            ></v-textarea>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-container>
</template>

<script>
import { mapActions } from 'vuex';
import ConfirmDialog from '@/components/ConfirmDialog.vue';

export default {
  name: 'ItemDetails',
  components: {
    ConfirmDialog,
  },
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
    /* Only fetch items if the user has characters and they were
    successfully loaded. */
    if (this.characters.length === 0) return;
    await this.fetchCharacterItems(this.characters[0].id);
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
     * @returns {Array<String>} - The names of the user's items (if
     * any).
     */
    itemNames() {
      if (this.items === null) return [];
      return this.items.map((x) => x.item_name);
    },
    /**
     * @returns {String} - The character currently selected by the user.
     */
    selectedCharacter() {
      if (!this.selectedCharName) return {};

      const target = this.selectedCharName;
      return this.characters.filter((x) => x.name === target)[0];
    },
    /**
     * @returns {String} - The item currently selected by the user.
     */
    selectedItem() {
      if (!this.selectedItemName) return {};

      const target = this.selectedItemName;
      return this.items.filter((x) => x.item_name === target)[0];
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
    /**
     * Fetches items for the currently-selected character.
     */
    async fetchCharacterItems() {
      const requestURI = `auth/character/${this.selectedCharacter.id}/item`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          const fetchedItems = resp.data.data.items;
          if (fetchedItems) {
            this.items = fetchedItems;
            this.selectedItemName = fetchedItems[0].item_name;
          } else {
            this.items = [];
            this.selectedItemName = null;
          }
        })
        .catch((err) => {
          let message = 'Could not fetch items. Please try again.';
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
    /**
     * Displays a prompt to the user, asking them to confirm their
     * item deletion.
     */
    async displayItemDeletionPrompt() {
      const title = 'Confirm deletion';
      const message = 'Delete your item? This action is irreversible.';
      if (await this.$refs.confirmDelete.prompt(title, message)) {
        await this.requestItemDeletion();
      }
    },
    /**
     * Sends an HTTP request to delete the user's item for a specific character.
     */
    async requestItemDeletion() {
      const integerID = parseInt(this.selectedCharacter.id, 10);

      const requestURI = `auth/character/${integerID}/item/${this.selectedItemName}`;
      const method = 'DELETE';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then(() => {
          this.display({
            message: 'Item was successfully deleted!',
            color: 'success',
            timeout: 6000,
          });
        })
        .catch((err) => {
          let message = 'Something went wrong. Please try again.';
          if (err.response) {
            message = `Error: ${err.response.data.message}. Please try again.`;
          }
          this.display({
            message,
            color: 'error',
            timeout: 6000,
          });
        })
        .finally(() => {
          this.fetchCharacterItems();
        });
    },
  },
};
</script>
