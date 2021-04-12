<template>
  <v-form @submit.prevent>
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
            :error-messages="nameErrors"
            @input="$v.character.name.$touch()"
            @blur="$v.character.name.$touch()"
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
            :error-messages="weightErrors"
            @input="$v.character.weight.$touch()"
            @blur="$v.character.weight.$touch()"
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
            :error-messages="heightErrors"
            @input="$v.character.height.$touch()"
            @blur="$v.character.height.$touch()"
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
            label="Enter background (optional)"
            placeholder="A dark and mysterious past."
            outlined
            no-resize
            counter="250"
            :error-messages="backgroundErrors"
            @input="$v.character.background.$touch()"
            @blur="$v.character.background.$touch()"
          >
          </v-textarea>
        </v-col>
      </v-row>

      <v-btn class="primary" @click="onComplete">Next Step</v-btn>
    </v-container>
  </v-form>
</template>

<script>
import {
  required,
  minLength,
  maxLength,
  integer,
  minValue,
  maxValue,
  helpers,
} from 'vuelidate/lib/validators';

export default {
  name: 'Background',
  data() {
    return {
      character: {
        name: null,
        sex: 'Male',
        race: 'Human',
        weight: 75,
        height: 185,
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
  validations: {
    character: {
      name: {
        required,
        minLength: minLength(1),
        maxLength: maxLength(50),
        /* This regex obviously does not cover all cases, but is
      sufficient for our project. */
        isNameValid: helpers.regex('nameValidation', /^[a-zA-Z, '.-]*$/),
      },
      weight: {
        required,
        integer,
        minValue: minValue(1),
        maxValue: maxValue(1000),
      },
      height: {
        required,
        integer,
        minValue: minValue(1),
        maxValue: maxValue(1000),
      },
      background: {
        maxLength: maxLength(250),
      },
    },
  },
  computed: {
    nameErrors() {
      const errors = [];
      if (!this.$v.character.name.$dirty) return errors;
      if (!this.$v.character.name.minLength) {
        errors.push('Name must be at least 1 character.');
      }
      if (!this.$v.character.name.maxLength) {
        errors.push('Name must not exceed 50 characters.');
      }
      if (!this.$v.character.name.isNameValid) {
        errors.push('Name contains disallowed characters.');
      }
      if (!this.$v.character.name.required) {
        errors.push('Name is required.');
      }
      return errors;
    },
    weightErrors() {
      const errors = [];
      if (!this.$v.character.weight.$dirty) return errors;
      if (!this.$v.character.weight.integer) {
        errors.push('Weight must be a number.');
      }
      if (!this.$v.character.weight.minValue) {
        errors.push('Weight must be greater than 0 kg.');
      }
      if (!this.$v.character.weight.maxValue) {
        errors.push('Weight must not exceed 1000 kg.');
      }
      if (!this.$v.character.weight.required) {
        errors.push('Weight is required.');
      }
      return errors;
    },
    heightErrors() {
      const errors = [];
      if (!this.$v.character.height.$dirty) return errors;
      if (!this.$v.character.height.integer) {
        errors.push('Height must be a number.');
      }
      if (!this.$v.character.height.minValue) {
        errors.push('Height must be greater than 0 cm.');
      }
      if (!this.$v.character.height.maxValue) {
        errors.push('Height must not exceed 1000 cm.');
      }
      if (!this.$v.character.height.required) {
        errors.push('Height is required.');
      }
      return errors;
    },
    backgroundErrors() {
      const errors = [];
      if (!this.$v.character.background.$dirty) return errors;
      if (!this.$v.character.background.maxLength) {
        errors.push('Background must not exceed 250 characters.');
      }
      return errors;
    },
  },
  methods: {
    /**
     * Check if the user has completed the step. If so, emits the
     * up-to-date background data entered by the user. Should be used as
     * a click handler for the "next step" button.
     */
    onComplete() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      this.$emit('complete', this.character);
    },
  },
};
</script>
