<template>
  <v-stepper v-model="stepNumber">
    <!-- Display the names and numbers of each step. -->
    <v-stepper-header>
      <v-stepper-step :complete="stepNumber > 1" step="1">
        General Information
      </v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step :complete="stepNumber > 2" step="2">
        Stats
      </v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step :complete="stepNumber > 3" step="3">
        Class
      </v-stepper-step>
    </v-stepper-header>

    <!-- The actual content displayed in each respective step. -->
    <v-stepper-items>
      <!-- Step 1 -->
      <v-stepper-content step="1">
        <Background
          @complete="
            updateCharacterBackgroundData($event);
            incrementStepNumber();
          "
        >
        </Background>
      </v-stepper-content>

      <!-- Step 2 -->
      <v-stepper-content step="2">
        <Stats
          @cancel="decrementStepNumber"
          @complete="
            updateCharacterStatData($event);
            incrementStepNumber();
          "
        >
        </Stats>
      </v-stepper-content>

      <!-- Step 3 -->
      <v-stepper-content step="3">
        <Class
          @cancel="decrementStepNumber"
          @complete="
            updateCharacterClassData($event);
            sendCharacterCreationRequest();
          "
        >
        </Class>
      </v-stepper-content>
    </v-stepper-items>
  </v-stepper>
</template>

<script>
import { mapActions } from 'vuex';
import Background from './Background.vue';
import Stats from './Stats.vue';
import Class from './Class.vue';

export default {
  name: 'CreationStepper',
  components: {
    Background,
    Stats,
    Class,
  },
  data() {
    return {
      /* Tracks which step the user is at in the character creation
      process. */
      stepNumber: 1,
      character: {},
    };
  },
  methods: {
    ...mapActions({
      display: 'notifications/display',
    }),
    /**
     * Increments the current step number by 1. Used as a click event to
     * continue to the next step of a`<v-stepper>`.
     */
    incrementStepNumber() {
      this.stepNumber += 1;
    },
    /**
     * Decrements the current step number by 1. Used as a click event to
     * continue to the previous step of a `<v-stepper>`.
     */
    decrementStepNumber() {
      this.stepNumber -= 1;
    },
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
     *  classAttribute: Object
     *  }} character - The event data emitted after the
     *  class step has been completed.
     */
    updateCharacterClassData({ className, classAttribute }) {
      this.character.class = className;
      this.character.class_attribute = classAttribute;
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
  },
};
</script>
