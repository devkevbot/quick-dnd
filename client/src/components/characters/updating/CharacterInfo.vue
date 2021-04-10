<!-- CharacterInfo.vue -->
<!-- Displays general, stat, and other information about user
characters. Enables other actions such as deleting characters. -->

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

    <!-- Character name and ID -->
    <v-card v-if="selectedCharName" class="pl-2">
      <v-card-title class="display-1 primary--text">
        {{ selectedCharacter.name }}
      </v-card-title>
      <v-card-subtitle>
        Character ID: {{ selectedCharacter.id }}
      </v-card-subtitle>

      <!-- Delete button -->
      <v-card-actions>
        <v-btn class="error" @click="displayCharacterDeletionPrompt">
          Delete character
          <v-icon class="ml-2">mdi-delete-forever</v-icon>
        </v-btn>

        <ConfirmDialog ref="confirmDelete" />
      </v-card-actions>

      <!-- This container displays all character information -->
      <v-container>
        <p class="headline primary--text">General information</p>
        <v-row>
          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Sex"
              :value="selectedCharacter.sex"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Race"
              :value="selectedCharacter.race"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Weight"
              suffix="lbs"
              :value="selectedCharacter.weight"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Height"
              suffix="ft"
              :value="selectedCharacter.height"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              readonly
              label="Alignment"
              :value="selectedCharacter.alignment"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" lg="10">
            <v-textarea
              readonly
              no-resize
              label="Background"
              outlined
              :value="selectedCharacter.background"
            ></v-textarea>
          </v-col>
        </v-row>

        <p class="headline primary--text">Stats</p>
        <v-row>
          <v-col
            cols="12"
            md="6"
            lg="1"
            v-for="stat in stats"
            :key="stat.name"
            class="pb-4"
          >
            <v-text-field
              readonly
              :label="stat.name"
              :value="selectedCharacter[`${stat.dataName}`]"
            >
            </v-text-field>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-container>
</template>

<script>
import ConfirmDialog from '@/components/ConfirmDialog.vue';

export default {
  name: 'CharacterInfo',
  components: {
    ConfirmDialog,
  },
  data() {
    return {
      characters: [],
      selectedCharName: null,

      /* These stats are used to dynamically fetch the stat date from a
      given character object. */
      stats: [
        {
          name: 'Speed',
          dataName: 'speed',
        },
        {
          name: 'Strength',
          dataName: 'strength',
        },
        {
          name: 'Dexterity',
          dataName: 'dexterity',
        },
        {
          name: 'Intelligence',
          dataName: 'intelligence',
        },
        {
          name: 'Wisdom',
          dataName: 'wisdom',
        },
        {
          name: 'Charisma',
          dataName: 'charisma',
        },
        {
          name: 'Constitution',
          dataName: 'constitution',
        },
        {
          name: 'Max HP',
          dataName: 'hp_max',
        },
        {
          name: 'Ability Points',
          dataName: 'ability_points',
        },
        {
          name: 'XP Points',
          dataName: 'xp_points',
        },
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
        .then(() => this.fetchUserCharacters())
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
  },
};
</script>
