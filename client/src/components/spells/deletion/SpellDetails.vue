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
      <v-col cols="4">
        <p v-if="selectedCharName" class="headline primary--text">
          Select your spell
        </p>
        <v-select
          v-if="selectedCharName"
          solo
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
              @click="fetchCharacterSpells(61)"
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
        {{ characterSpells[selectedSpellIndex].spell_name }}
      </v-card-title>
      <v-card-subtitle>
        Character ID: {{ selectedCharacter.id }}
      </v-card-subtitle>

      <!-- Delete button -->
      <v-card-actions>
        <v-btn class="error" @click="displayCharacterDeletionPrompt">
          Delete {{ characterSpells[selectedSpellIndex].spell_name }}
          <v-icon class="ml-2">mdi-delete-forever</v-icon>
        </v-btn>
      </v-card-actions>

      <!-- This container displays all spell information -->
      <v-container>
        <p class="headline primary--text">Spell Information</p>
        <v-row>
          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Level"
              :value="characterSpells[selectedSpellIndex].level"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="School"
              :value="characterSpells[selectedSpellIndex].school"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Concentration"
              :value="characterSpells[selectedSpellIndex].concentration"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Casting Time"
              :value="characterSpells[selectedSpellIndex].casting_time"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Range"
              :value="characterSpells[selectedSpellIndex].range"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Duration"
              :value="characterSpells[selectedSpellIndex].duration"
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
              :value="characterSpells[selectedSpellIndex].description"
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
  name: 'SpellDetails',
  data() {
    return {
      characters: [],
      spells: [],
      selectedCharName: null,
      selectedSpellName: null,
      selectedSpellIndex: 0,
      characterSpells: [
        {
          character_id: '',
        },
        {
          spell_name: '',
        },
        {
          level: '',
        },
        {
          school: '',
        },
        {
          concentration: '',
        },
        {
          description: '',
        },
        {
          casting_time: '',
        },
        {
          range: '',
        },
        {
          duration: '',
        },
      ],
    };
  },
  /**
   * Fetches all of the user's character and spell data when the component loads.
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
      return this.characterSpells.map((x) => x.spell_name);
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
          if (this.spells) {
            const targetSpell = this.spells[0];
            this.selectedSpellName = targetSpell.spell_name;
          }
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
  },
};
</script>
