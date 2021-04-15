<template>
  <v-form>
    <v-container>
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
      </v-row>

      <v-card class="pa-4" v-if="selectedCharacter">
        <!-- Item name, type, and rarity. -->
        <p class="headline primary--text">General information</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="item.name"
              label="Enter item name"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model="item.type"
              :items="typeOptions"
              label="Choose item type"
              outlined
              required
            >
            </v-select>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model.trim="item.rarity"
              :items="rarityOptions"
              label="Choose rarity"
              outlined
              required
            >
            </v-select>
          </v-col>
        </v-row>

        <!-- Item level, weight, gold value, and quantity. -->
        <p class="headline primary--text">Numeric details</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="item.weight"
              label="Enter item weight"
              suffix="lbs"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="item.goldValue"
              label="Enter item's gold value"
              suffix="gold"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="item.quantity"
              label="Enter quantity of item(s)"
              outlined
              required
            >
            </v-text-field>
          </v-col>
        </v-row>

        <!-- Item description. -->
        <p class="headline primary--text">Description</p>
        <v-row>
          <v-col cols="12" md="6">
            <v-textarea
              v-model.trim="item.description"
              label="Enter description"
              placeholder=""
              outlined
              clearable
              no-resize
              required
              counter="250"
            >
            </v-textarea>
          </v-col>
        </v-row>

        <v-btn class="primary" @click="onComplete">Create</v-btn>
      </v-card>
    </v-container>
  </v-form>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'CreateItem',
  data() {
    return {
      characters: [],
      selectedCharName: null,

      item: {
        name: '',
        type: 'Armor',
        rarity: 'Common',
        weight: 15,
        goldValue: 10,
        quantity: 1,
        description: 'The most humble of items.',
      },

      typeOptions: [
        'Armor',
        'Potion',
        'Ring',
        'Rod',
        'Scroll',
        'Staff',
        'Wand',
        'Weapon',
        'Wondrous Item',
      ],

      rarityOptions: [
        'Common',
        'Uncommon',
        'Rare',
        'Very Rare',
        'Legendary',
        'Artifact',
      ],
    };
  },
  /**
   * Fetches all of the user's character data when the component loads.
   * Does not handle fetch errors.
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
      return this.characters.map((x) => x.name);
    },
    /**
     * @returns {String} - The character currently selected by the user.
     */
    selectedCharacter() {
      if (this.selectedCharName === null) return '';

      const target = this.selectedCharName;
      return this.characters.filter((x) => x.name === target)[0];
    },
  },
  methods: {
    ...mapActions({
      display: 'notifications/display',
    }),
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
          if (fetchedCharacters === null) return;

          this.characters = fetchedCharacters;
          const targetChar = this.characters[0];
          this.selectedCharName = targetChar.name;
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
     * Click handler for the 'create' button for the item form.
     */
    async onComplete() {
      /* TODO: validate inputs. */
      await this.sendItemCreationRequest();
      /* TODO: cleanup creation form. */
    },
    /**
     * @returns {Object} Properly-formatted item data that can be sent
     * the body of an HTTP request for item creation.
     */
    prepareDataForRequest() {
      return {
        item_name: this.item.name,
        type: this.item.type,
        rarity: this.item.rarity,
        weight: parseInt(this.item.weight, 10),
        gold_value: parseInt(this.item.goldValue, 10),
        quantity: parseInt(this.item.quantity, 10),
        description: this.item.description,
      };
    },
    /**
     * Sends an HTTP request to the backend's item creation API
     * endpoint.
     */
    async sendItemCreationRequest() {
      const requestURI = `auth/character/${this.selectedCharacter.id}/item`;
      const method = 'POST';

      await this.$http({
        url: requestURI,
        data: this.prepareDataForRequest(),
        method,
      })
        .then(() => {
          this.display({
            message: 'Item was successfully created!',
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
        });
    },
  },
};
</script>
