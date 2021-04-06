<template>
  <v-form>
    <v-container>
      <p class="headline primary--text">Select character stats</p>
      <!-- Display a slider for each character stat. -->
      <v-row v-for="stat in stats" :key="stat.name" class="pb-4">
        <!-- The stat name is in a separate column so that all the
        slider bar lengths will be equal. -->
        <v-col cols="12" md="1">
          <p class="primary--text subtitle-1">{{ stat.name }}</p>
        </v-col>

        <v-col cols="12" md="6">
          <v-slider
            v-model="stat.value"
            :min="stat.min"
            :max="stat.max"
            thumb-label="always"
            thumb-size="45"
          >
            <!-- Resets the stat value to the initial value when clicked. -->
            <template v-slot:append>
              <v-btn text color="primary" @click="stat.value = stat.initValue">
                Reset
              </v-btn>
            </template>
          </v-slider>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
export default {
  name: 'Stats',
  data() {
    return {
      /* Character stats which are selected by the player. */
      stats: [
        {
          name: 'Speed',
          initValue: 640,
          value: 640,
          min: 0,
          max: 1280,
        },
        {
          name: 'Strength',
          initValue: 15,
          value: 15,
          min: 0,
          max: 30,
        },
        {
          name: 'Dexterity',
          initValue: 15,
          value: 15,
          min: 0,
          max: 30,
        },
        {
          name: 'Intelligence',
          initValue: 15,
          value: 15,
          min: 0,
          max: 30,
        },
        {
          name: 'Wisdom',
          initValue: 15,
          value: 15,
          min: 0,
          max: 30,
        },
        {
          name: 'Charisma',
          initValue: 15,
          value: 15,
          min: 0,
          max: 30,
        },
        {
          name: 'Constitution',
          initValue: 15,
          value: 15,
          min: 0,
          max: 30,
        },
        {
          name: 'Max HP',
          initValue: 220,
          value: 220,
          min: 1,
          max: 440,
        },
        {
          name: 'Ability Points',
          initValue: 10,
          value: 10,
          min: 0,
          max: 180,
        },
        {
          /* TODO: Calculate level based on XP points. */
          name: 'XP Points',
          initValue: 0,
          value: 0,
          min: 0,
          max: 355000,
        },
      ],
    };
  },
};
</script>
