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
        <Background />
        <v-btn class="primary" @click="incrementStepNumber">Next Step</v-btn>
      </v-stepper-content>

      <!-- Step 2 -->
      <v-stepper-content step="2">
        <Stats />
        <v-btn text @click="decrementStepNumber">Back</v-btn>
        <v-btn class="primary" @click="incrementStepNumber">Next Step</v-btn>
      </v-stepper-content>

      <!-- Step 3 -->
      <v-stepper-content step="3">
        <Class />
        <v-btn text @click="decrementStepNumber">Back</v-btn>
        <v-btn class="primary">Create</v-btn>
      </v-stepper-content>
    </v-stepper-items>
  </v-stepper>
</template>

<script>
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
    };
  },
  methods: {
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
     * Executes when the player has confirmed their intent to create a
     * character. Used as a click event on the "create" button at the
     * end of the `<v-stepper>` steps.
     */
    onSubmit() {
      /* TODO: fill this function. */
    },
    /**
     * Send a request to the backend server's character creation API endpoint.
     */
    sendCharacterCreationRequest() {
      /* TODO: fill this function. */
    },
  },
};
</script>
