// Component for stats and computed modifiers for a character

<template>
  <div class="modifiers">
    <v-btn>
      <h3>Save Changes</h3>
    </v-btn>

    <!-- Main Character Stats -->
    <v-container class="characterSliders">
      <v-row wrap>
        <v-col cols="12" lg="3" v-for="row in statSlider" :key="row.attribute">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ row.attribute }}
            </v-card-title>
            <v-card-actions>
              <v-slider v-model="row.value" :max="row.max" class="align-center">
                <template v-slot:append>
                  <v-text-field
                    v-model="row.value"
                    class="mt-0 pt-0"
                    type="text"
                  >
                  </v-text-field>
                </template>
              </v-slider>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>

    <!-- Derived Character Stats  -->
    <v-container class="characterModifiers">
      <v-row wrap>
        <v-col
          cols="12"
          md="6"
          lg="2"
          v-for="row in statModifiers"
          :key="row.modifier"
        >
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ row.modifier }}
            </v-card-title>
            <v-card-actions>
              <v-text-field v-model="row.value" label="Unloaded" solo disabled>
              </v-text-field>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  props: {
    characterData: {
      type: Array,
      required: false,
    },
  },
  computed: {
    statSlider() {
      return [
        { attribute: 'Strength', value: this.characterData.strength, max: 30 },
        {
          attribute: 'Constitution',
          value: this.characterData.constitution,
          max: 30,
        },
        { attribute: 'Wisdom', value: this.characterData.wisdom, max: 30 },
        { attribute: 'Hit Points', value: this.characterData.hp_max, max: 440 },
        {
          attribute: 'Dexterity',
          value: this.characterData.dexterity,
          max: 30,
        },
        {
          attribute: 'Ability Points',
          value: this.characterData.ability_points,
          max: 180,
        },
        { attribute: 'Charisma', value: this.characterData.charisma, max: 30 },
        {
          attribute: 'Intellect',
          value: this.characterData.intelligence,
          max: 30,
        },
      ];
    },
    statModifiers() {
      return [
        { modifier: 'Initiative', value: 'Unloaded' },
        { modifier: 'Number of Dies', value: 'Unloaded' },
        {
          modifier: 'Strength',
          value: Math.floor((this.characterData.strength - 10) / 2),
        },
        {
          modifier: 'Constitution',
          value: Math.floor((this.characterData.constitution - 10) / 2),
        },
        {
          modifier: 'Wisdom',
          value: Math.floor((this.characterData.wisdom - 10) / 2),
        },
        { modifier: 'Hit Points', value: 'Unloaded', max: 440 },
        {
          modifier: 'Dexterity',
          value: Math.floor((this.characterData.dexterity - 10) / 2),
        },
        { modifier: 'Ability Points', value: 'Unloaded' },
        {
          modifier: 'Charisma',
          value: Math.floor((this.characterData.charisma - 10) / 2),
        },
        {
          modifier: 'Intellect',
          value: Math.floor((this.characterData.intelligence - 10) / 2),
        },
        { modifier: 'Proficency', value: Math.ceil(this.level / 4) + 1 },
        { modifier: 'Speed', value: this.characterData.speed, max: 1280 },
      ];
    },
    level() {
      let tempLevel = 1;
      if (this.characterData.xp_points < 300) {
        tempLevel = 1;
      } else if (this.characterData.xp_points < 900) {
        tempLevel = 2;
      } else if (this.characterData.xp_points < 2700) {
        tempLevel = 3;
      } else if (this.characterData.xp_points < 6500) {
        tempLevel = 4;
      } else if (this.characterData.xp_points < 14000) {
        tempLevel = 5;
      } else if (this.characterData.xp_points < 23000) {
        tempLevel = 6;
      } else if (this.characterData.xp_points < 34000) {
        tempLevel = 7;
      } else if (this.characterData.xp_points < 48000) {
        tempLevel = 8;
      } else if (this.characterData.xp_points < 64000) {
        tempLevel = 9;
      } else if (this.characterData.xp_points < 85000) {
        tempLevel = 10;
      } else if (this.characterData.xp_points < 100000) {
        tempLevel = 11;
      } else if (this.characterData.xp_points < 120000) {
        tempLevel = 12;
      } else if (this.characterData.xp_points < 140000) {
        tempLevel = 13;
      } else if (this.characterData.xp_points < 165000) {
        tempLevel = 14;
      } else if (this.characterData.xp_points < 195000) {
        tempLevel = 15;
      } else if (this.characterData.xp_points < 225000) {
        tempLevel = 16;
      } else if (this.characterData.xp_points < 265000) {
        tempLevel = 17;
      } else if (this.characterData.xp_points < 305000) {
        tempLevel = 18;
      } else if (this.characterData.xp_points < 355000) {
        tempLevel = 19;
      } else {
        tempLevel = 20;
      }
      return tempLevel;
    },
  },
};
</script>
