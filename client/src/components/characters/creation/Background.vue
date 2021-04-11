<template>
  <v-form>
    <v-container>
      <!-- Character name, sex, race. -->
      <p class="headline primary--text">Identity</p>
      <v-row>
        <v-col cols="12" md="2">
          <v-text-field
            v-model.trim="character.name"
            label="Enter name"
            outlined
            required
            placeholder="Trogdor, Destroyer of Worlds"
          >
          </v-text-field>
        </v-col>

        <v-col cols="12" md="2">
          <v-select
            v-model="character.sex"
            :items="characterSexOptions"
            label="Choose sex"
            outlined
            required
          >
          </v-select>
        </v-col>

        <v-col cols="12" md="2">
          <v-select
            v-model.trim="character.race"
            :items="characterRaceOptions"
            label="Choose race"
            outlined
            required
          >
          </v-select>
        </v-col>
      </v-row>

      <!-- Character weight and height. -->
      <p class="headline primary--text">Body Shape</p>
      <v-row>
        <v-col cols="12" md="3">
          <v-text-field
            v-model.trim="character.weight"
            label="Enter weight"
            suffix="kg"
            outlined
            required
          >
          </v-text-field>
        </v-col>

        <v-col cols="12" md="3">
          <v-text-field
            v-model.trim="character.height"
            label="Enter height"
            suffix="cm"
            outlined
            required
          >
          </v-text-field>
        </v-col>
      </v-row>

      <p class="headline primary--text">Alignment</p>
      <v-row>
        <v-col cols="12" md="6">
          <v-select
            v-model.trim="character.alignment"
            :items="characterAlignmentOptions"
            label="Choose alignment"
            outlined
            required
          >
          </v-select>
        </v-col>
      </v-row>

      <!-- Character background. -->
      <p class="headline primary--text">Background</p>
      <v-row>
        <v-col cols="12" md="6">
          <v-textarea
            v-model.trim="character.background"
            label="Enter background"
            placeholder="A dark and mysterious past."
            outlined
            clearable
            no-resize
            required
            counter="250"
          >
          </v-textarea>
        </v-col>
      </v-row>

      <v-btn class="primary" @click="onComplete">Next Step</v-btn>
    </v-container>
  </v-form>
</template>

<script>
export default {
  name: 'Background',
  data() {
    return {
      character: {
        name: null,
        sex: 'Male',
        race: 'Human',
        weight: 180,
        height: 6.2,
        alignment: 'True Neutral',
        background: 'A dark and mysterious past.',
      },

      characterSexOptions: ['Male', 'Female', 'Other'],

      characterRaceOptions: [
        'Dragonborn',
        'Dwarf',
        'Elf',
        'Gnome',
        'Half-Elf',
        'Halfling',
        'Half-Orc',
        'Human',
        'Tiefling',
      ],

      characterAlignmentOptions: [
        'Lawful Good',
        'Neutral Good',
        'Chaotic Good',
        'Lawful Neutral',
        'True Neutral',
        'Chaotic Neutral',
        'Lawful Evil',
        'Neutral Evil',
        'Chaotic Evil',
      ],
    };
  },
  methods: {
    /**
     * Emits the up-to-date background data entered by the user.
     * Should be used a click handler for the "next step" button.
     */
    onComplete() {
      this.$emit('complete', this.character);
    },
  },
};
</script>
