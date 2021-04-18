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
        <v-row class="ml-2">
          <v-btn
            color="primary"
            @click="requestCharacterUpdate"
          >
            Update Character
          </v-btn>
          <v-btn class="error ml-4" @click="displayCharacterDeletionPrompt">
            Delete character
            <v-icon>mdi-delete-forever</v-icon>
          </v-btn>
        </v-row>

        <ConfirmDialog ref="confirmDelete" />
      </v-card-actions>

      <!-- This container displays all character information -->
      <v-container>
        <p class="headline primary--text">General information</p>
        <v-row>
          <v-col cols="12" lg="2">
            <v-select
              :items="sexOptions"
              v-model="selectedCharacter.sex"
              label="Sex"
              required
            ></v-select>
          </v-col>

          <v-col cols="12" lg="2">
            <v-select
              :items="raceOptions"
              v-model="selectedCharacter.race"
              label="Race"
              required
            ></v-select>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              v-model="selectedCharacter.weight"
              label="Weight"
              suffix="kg"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-text-field
              v-model="selectedCharacter.height"
              label="Height"
              suffix="cm"
            ></v-text-field>
          </v-col>

          <v-col cols="12" lg="2">
            <v-select
              :items="alignmentOptions"
              v-model="selectedCharacter.alignment"
              label="Alignment"
              required
            ></v-select>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" lg="10">
            <v-textarea
              no-resize
              v-model="selectedCharacter.background"
              label="Background"
              outlined
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
              v-model="selectedCharacter[`${stat.dataName}`]"
              :label="stat.name"
            >
            </v-text-field>
          </v-col>
        </v-row>

        <p class="headline primary--text">Modifiers</p>
        <v-row>
          <v-col
            cols="12"
            md="6"
            lg="1"
            v-for="stat in statModifiers"
            :key="stat.name"
            class="pb-4"
          >
            <v-text-field readonly :label="stat.name" :value="stat.value">
            </v-text-field>
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
  name: 'CharacterInfo',
  components: {
    ConfirmDialog,
  },
  data() {
    return {
      characters: [],
      selectedCharName: null,

      sexes: ['Male', 'Female', 'Other'],
      races: ['Dragonborn', 'Dwarf', 'Elf', 'Gnome',
        'Half-Elf', 'Halfling', 'Half-Orc', 'Human', 'Tiefling'],
      alignments: ['Lawful Good', 'Neutral Good', 'Chaotic Good', 'Lawful Neutral',
        'True Neutral', 'Chaotic Neutral', 'Lawful Evil', 'Neutral Evil', 'Chaotic Evil'],

      /* These stats are used to dynamically fetch the stat date from a
      given character object. */
      stats: [
        {
          name: 'Charisma',
          dataName: 'charisma',
        },
        {
          name: 'Constitution',
          dataName: 'constitution',
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
          name: 'Strength',
          dataName: 'strength',
        },
        {
          name: 'Wisdom',
          dataName: 'wisdom',
        },
        {
          name: 'Speed',
          dataName: 'speed',
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
    statModifiers() {
      return [
        {
          name: 'Charisma',
          value: Math.floor((this.selectedCharacter.charisma - 10) / 2),
        },
        {
          name: 'Constitution',
          value: Math.floor((this.selectedCharacter.constitution - 10) / 2),
        },
        {
          name: 'Dexterity',
          value: Math.floor((this.selectedCharacter.dexterity - 10) / 2),
        },
        {
          name: 'Intelligence',
          value: Math.floor((this.selectedCharacter.intelligence - 10) / 2),
        },
        {
          name: 'Strength',
          value: Math.floor((this.selectedCharacter.strength - 10) / 2),
        },
        {
          name: 'Wisdom',
          value: Math.floor((this.selectedCharacter.wisdom - 10) / 2),
        },
      ];
    },
    sexOptions() {
      return this.sexes;
    },
    raceOptions() {
      return this.races;
    },
    alignmentOptions() {
      return this.alignments;
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
    async requestCharacterUpdate() {
      const integerID = parseInt(this.selectedCharacter.id, 10);
      const requestURI = `auth/character/${integerID}`;
      const method = 'PUT';

      await this.$http({
        url: requestURI,
        data: {
          id: integerID,
          name: this.selectedCharacter.name,
          weight: this.selectedCharacter.weight,
          height: this.selectedCharacter.height,
          alignment: this.selectedCharacter.alignment,
          sex: this.selectedCharacter.sex,
          background: this.selectedCharacter.background,
          race: this.selectedCharacter.race,
          speed: this.selectedCharacter.speed,
          strength: this.selectedCharacter.strength,
          dexterity: this.selectedCharacter.dexterity,
          intelligence: this.selectedCharacter.intelligence,
          wisdom: this.selectedCharacter.wisdom,
          charisma: this.selectedCharacter.charisma,
          constitution: this.selectedCharacter.constitution,
          hp_max: this.selectedCharacter.hp_max,
          ability_points: this.selectedCharacter.ability_points,
          xp_points: this.selectedCharacter.xp_points,
          class: this.selectedCharacter.class,
          class_attribute: this.selectedCharacter.class_attribute,
          player_username: this.selectedCharacter.player_username,
        },
        method,
      })
        .then(() => {
          this.display({
            message: 'Character was updated successfully!',
            color: 'success',
            timeout: 6000,
          });
        })
        .catch((err) => {
          let message = 'Could not update character. Please try again.';
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
