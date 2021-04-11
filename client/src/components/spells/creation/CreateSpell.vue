<template>
  <v-form>
    <v-container>
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
      <v-card v-if="selectedCharName" class="pl-2">
        <!-- Spell name, shcool, and concentration. -->
        <p class="headline primary--text">Spell Information</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="character.name"
              label="Enter name"
              outlined
              required
              placeholder="Fireball"
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model="character.sex"
              :items="characterSexOptions"
              label="Choose School"
              outlined
              required
            >
            </v-select>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model.trim="character.race"
              :items="characterRaceOptions"
              label="Choose Concentration"
              outlined
              required
            >
            </v-select>
          </v-col>
        </v-row>

        <!-- Character weight and height. -->
        <p class="headline primary--text">Values</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="character.weight"
              label="Enter cast time"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="character.height"
              label="Enter duration"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="character.height"
              label="Enter Range"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>
        </v-row>

        <!-- Character background. -->
        <p class="headline primary--text">Description</p>
        <v-row>
          <v-col cols="12" md="6">
            <v-textarea
              v-model.trim="character.background"
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
      character: {},
      spell: {},
      selectedCharName: null,
    };
  },
  methods: {
    ...mapActions({
      display: 'notifications/display',
    }),
    /**
     * Updates the current character background data. Triggered when the
     * 'Background' step of the character creation process has been completed.
     * @param {{
     *  name: String
     *  sex: String
     *  race: String
     *  weight: String
     *  height: String
     *  alignment: String
     *  background: String
     *  }} character - The event data emitted after the
     *  background step has been completed.
     */
    updateCharacterBackgroundData({
      name,
      sex,
      race,
      weight,
      height,
      alignment,
      background,
    }) {
      this.character.name = name;
      this.character.sex = sex;
      this.character.race = race;
      this.character.weight = parseInt(weight, 10);
      this.character.height = parseInt(height, 10);
      this.character.alignment = alignment;
      this.character.background = background;
    },
    /**
     * Updates the current character stats data. Triggered when the
     * 'Stats' step of the character creation process has been completed.
     * @param {{
     *  speed: Number
     *  strength: Number
     *  dexterity: Number
     *  intelligence: Number
     *  wisdom: Number
     *  charisma: Number
     *  constitution: Number
     *  hpMax: Number
     *  abilityPoints: Number
     *  xpPoints: Number
     *  }} character - The event data emitted after the
     *  stats step has been completed.
     */
    updateCharacterStatData({
      speed,
      strength,
      dexterity,
      intelligence,
      wisdom,
      charisma,
      constitution,
      hpMax,
      abilityPoints,
      xpPoints,
    }) {
      this.character.speed = speed;
      this.character.strength = strength;
      this.character.dexterity = dexterity;
      this.character.intelligence = intelligence;
      this.character.wisdom = wisdom;
      this.character.charisma = charisma;
      this.character.constitution = constitution;
      this.character.hp_max = hpMax;
      this.character.ability_points = abilityPoints;
      this.character.xp_points = xpPoints;
    },
    /**
     * Updates the current character class data. Triggered when the
     * 'Class' step of the character creation process has been completed.
     * @param {{
     *  className: String
     *  classAttributes: Object
     *  }} character - The event data emitted after the
     *  class step has been completed.
     */
    updateCharacterClassData({ className, classAttributes }) {
      this.character.class = className;
      this.character.class_attributes = classAttributes;
    },
    /**
     * Send a request to the backend server's character creation API endpoint.
     */
    async sendCharacterCreationRequest() {
      const requestURI = 'auth/character';
      const method = 'POST';

      await this.$http({
        url: requestURI,
        data: this.character,
        method,
      })
        .then(() => {
          this.display({
            message: 'Character was successfully created!',
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
  },
};
</script>
