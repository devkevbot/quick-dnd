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
        <!-- Spell name, school, and concentration. -->
        <p class="headline primary--text">General information</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="spell.name"
              label="Enter spell name"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model="spell.school"
              :items="schoolOptions"
              label="Choose school"
              outlined
              required
            >
            </v-select>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model.trim="spell.concentration"
              :items="concentrationOptions"
              label="Choose concentration"
              outlined
              required
            >
            </v-select>
          </v-col>
        </v-row>

        <!-- Spell level, casting time, duration, and range. -->
        <p class="headline primary--text">Numeric details</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="spell.level"
              label="Enter level number"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="spell.casting_time"
              label="Enter casting time"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="spell.duration"
              label="Enter duration"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="spell.range"
              label="Enter range"
              suffix="ft"
              outlined
              required
            >
            </v-text-field>
          </v-col>
        </v-row>

        <!-- Spell description. -->
        <p class="headline primary--text">Description</p>
        <v-row>
          <v-col cols="12" md="6">
            <v-textarea
              v-model.trim="spell.description"
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
  name: 'CreateSpell',
  data() {
    return {
      characters: [],
      selectedCharName: null,

      spell: {
        name: '',
        school: 'Abjuration',
        concentration: true,
        level: 0,
        casting_time: 5,
        duration: 1,
        range: 10,
        description: 'Blasts a foe into oblivion.',
      },

      schoolOptions: [
        'Abjuration',
        'Conjuration',
        'Divination',
        'Enchantment',
        'Evocation',
        'Illusion',
        'Necromancy',
        'Transmutation',
      ],

      concentrationOptions: [true, false],
    };
  },
  /**
   * Fetches all of the user's character data when the component loads.
   * Does not handle fetch errors.
   * Not sure how to handle spell loading.
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
     * Click handler for the 'create' button for the spell form.
     */
    async onComplete() {
      /* TODO: validate inputs. */
      await this.sendSpellCreationRequest();
      /* TODO: cleanup creation form. */
    },
    /**
     * @returns {Object} Properly-formatted spell data that can be sent
     * the body of an HTTP request for spell creation.
     */
    prepareDataForRequest() {
      return {
        spell_name: this.spell.name,
        level: parseInt(this.spell.level, 10),
        school: this.spell.school,
        concentration: this.spell.concentration,
        casting_time: parseInt(this.spell.casting_time, 10),
        duration: parseInt(this.spell.duration, 10),
        range: parseInt(this.spell.range, 10),
        description: this.spell.description,
      };
    },
    /**
     * Sends an HTTP request to the backend's spell creation API
     * endpoint.
     */
    async sendSpellCreationRequest() {
      const requestURI = `auth/character/${this.selectedCharacter.id}/spell`;
      const method = 'POST';

      await this.$http({
        url: requestURI,
        data: this.prepareDataForRequest(),
        method,
      })
        .then(() => {
          this.display({
            message: 'Spell was successfully created!',
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
