<!-- SpellDetails.vue -->
<!-- Displays all details for a specific character's spells.
The player is then able to delete the spell. -->

<template>
  <v-container class="character-info">
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

    <!-- Spell selection dropdown -->
    <v-row>
      <v-col v-if="selectedCharName" cols="4">
        <p class="headline primary--text">Select your spell</p>
        <v-select
          v-if="selectedCharName"
          solo
          prepend-inner-icon="mdi-account-search"
          :items="characterNames"
          v-model="selectedSpellName"
        >
          <!-- Button used to refresh the spell list. -->
          <template v-slot:append-outer>
            <v-btn
              color="primary"
              fab
              small
              class="mt-n2"
              @click="fetchCharacterSpells(selectedCharacter.id)"
            >
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </template>
        </v-select>
      </v-col>
    </v-row>

    <!-- Spell Name and Character ID -->
    <v-card v-if="selectedCharName" class="pl-2">
      <v-card-title class="display-1 primary--text">
        {{ characterSpells[0].spell_name }}
      </v-card-title>
      <v-card-subtitle>
        Character ID: {{ selectedCharacter.id }}
      </v-card-subtitle>

      <!-- Delete button -->
      <v-card-actions>
        <v-btn class="error" @click="displayCharacterDeletionPrompt">
          Delete {{ spellNames[0] }}
          <v-icon class="ml-2">mdi-delete-forever</v-icon>
        </v-btn>

        <ConfirmDialog ref="confirmDelete" />
      </v-card-actions>

      <!-- This container displays all spell information -->
      <v-container>
        <p class="headline primary--text">Spell Information</p>
        <v-row>
          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Level"
              :value="characterSpells[0].level"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="School"
              :value="characterSpells[0].school"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Concentration"
              :value="characterSpells[0].concentration"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Casting Time"
              :value="characterSpells[0].casting_time"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Range"
              :value="characterSpells[0].range"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Duration"
              :value="characterSpells[0].duration"
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
              :value="characterSpells[0].description"
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
  name: 'SpellDetails',
  components: {
    ConfirmDialog,
  },
  data() {
    return {
      characters: [],
      selectedCharName: null,
      spells: [],
      selectedSpellName: null,
    };
  },
  /**
   * Fetches all of the user's character data when the component loads.
   * Does not handle fetch errors.
   * Not sure how to handle spell loading.
   */
  async created() {
    await this.fetchUserCharacters();
    await this.fetchCharacterSpells(61);
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
     * @returns {Array<String>} - The names of the user's spell (if
     * any).
     */
    spellNames() {
      return this.data.map((x) => x.spell_name);
    },
    /**
     * @returns {String} - The character currently selected by the user.
     */
    selectedCharacter() {
      if (this.selectedCharName === null) return '';

      const target = this.selectedCharName;
      return this.characters.filter((x) => x.name === target)[0];
    },
    /**
     * @returns {String} - The spell currently selected by the user.
     */
    selectedSpell() {
      if (this.selectedSpellName === null) return '';

      const target = this.selectedSpellName;
      return this.spells.filter((x) => x.spell_name === target)[0];
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
          this.characters = resp.data.data.characters;
          if (this.characters) {
            const targetChar = this.characters[0];
            this.selectedCharName = targetChar.name;
          }
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * @param {String} charID - The character ID which can help to identify the spells
     * to fetch.
     */
    async fetchCharacterSpells(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/spell`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.characterSpells = resp.data.data.spells;
          this.data = this.characterSpells;
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * Displays a prompt to the user, asking them to confirm their
     * character deletion.
     */
    async displayCharacterDeletionPrompt() {
      const title = 'Confirm deletion';
      const message = 'Delete your character? This action is irreversible.';
      if (await this.$refs.confirmDelete.prompt(title, message)) {
        await this.requestCharacterDeletion();
      }
    },
    /**
     * Sends an HTTP request to delete the user's character.
     */
    async requestCharacterDeletion() {
      const integerID = parseInt(this.selectedCharacter.id, 10);

      const requestURI = `auth/character/${integerID}`;
      const method = 'DELETE';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then(() => {
          this.display({
            message: 'Character was successfully deleted!',
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
          this.fetchUserCharacters();
        });
    },
  },
};
</script>
