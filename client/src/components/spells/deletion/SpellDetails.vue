<!-- SpellDetails.vue -->
<!-- Displays all details for a specific character's spells.
The player is then able to delete the spell. -->

<template>
  <v-container class="spell-info">
    <!-- Character selection dropdown -->
    <v-row>
      <v-col cols="4">
        <p class="headline primary--text">Select your character</p>
        <v-select
          solo
          prepend-inner-icon="mdi-account-search"
          :items="characterNames"
          v-model="selectedCharName"
          @change="fetchCharacterSpells()"
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

      <!-- Spell selection dropdown -->
      <v-col cols="4" v-if="selectedCharName">
        <p class="headline primary--text">Select your spell</p>
        <v-select
          solo
          no-data-text="This character has no spells."
          prepend-inner-icon="mdi-account-search"
          :items="spellNames"
          v-model="selectedSpellName"
        >
          <!-- Button used to refresh the spell list. -->
          <template v-slot:append-outer>
            <v-btn
              color="primary"
              fab
              small
              class="mt-n2"
              @click="fetchCharacterSpells()"
            >
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </template>
        </v-select>
      </v-col>
    </v-row>

    <!-- Spell Name and Character ID -->
    <v-card v-if="characters.length && spells.length" class="pl-2">
      <v-card-title class="display-1 primary--text">
        {{ selectedSpell.spell_name }}
      </v-card-title>

      <!-- Delete button -->
      <v-card-actions>
        <v-btn class="error" @click="displaySpellDeletionPrompt">
          Delete Spell
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
              :value="selectedSpell.level"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="School"
              :value="selectedSpell.school"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Concentration"
              :value="selectedSpell.concentration"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Casting Time"
              :value="selectedSpell.casting_time"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Range"
              :value="selectedSpell.range"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Duration"
              :value="selectedSpell.duration"
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
              :value="selectedSpell.description"
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
   * Fetches all of the user's character and spell data when the component loads.
   */
  async created() {
    await this.fetchUserCharacters();
    /* Only fetch spells if the user has characters and they were
    successfully loaded. */
    if (this.characters.length === 0) return;
    await this.fetchCharacterSpells(this.characters[0].id);
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
     * @returns {Array<String>} - The names of the user's spells (if
     * any).
     */
    spellNames() {
      if (this.spells === null) return [];
      return this.spells.map((x) => x.spell_name);
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
     * @returns {String} - The spell currently selected by the user.
     */
    selectedSpell() {
      if (!this.selectedSpellName) return {};

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
     * Fetches spells for the currently-selected character.
     */
    async fetchCharacterSpells() {
      const requestURI = `auth/character/${this.selectedCharacter.id}/spell`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          const fetchedSpells = resp.data.data.spells;
          if (fetchedSpells) {
            this.spells = fetchedSpells;
            this.selectedSpellName = fetchedSpells[0].spell_name;
          } else {
            this.spells = [];
            this.selectedSpellName = null;
          }
        })
        .catch((err) => {
          let message = 'Could not fetch spells. Please try again.';
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
     * spell deletion.
     */
    async displaySpellDeletionPrompt() {
      const title = 'Confirm deletion';
      const message = 'Delete your spell? This action is irreversible.';
      if (await this.$refs.confirmDelete.prompt(title, message)) {
        await this.requestSpellDeletion();
      }
    },
    /**
     * Sends an HTTP request to delete the user's spells for a specific character.
     */
    async requestSpellDeletion() {
      const integerID = parseInt(this.selectedCharacter.id, 10);

      const requestURI = `auth/character/${integerID}/spell/${this.selectedSpellName}`;
      const method = 'DELETE';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then(() => {
          this.display({
            message: 'Spell was successfully deleted!',
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
          this.fetchCharacterSpells();
        });
    },
  },
};
</script>
