// Component contains the primary characteristics for a character

<template>
  <div class="primaryStats">
    <v-btn>
      <h3>Save Changes</h3>
    </v-btn>

    <!-- This row contains the character's: Level, XP, and Alignment -->
    <v-container class="characterSelect">
      <v-row>
        <v-col cols="12" lg="4">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ experiencePointsRelated[0].attribute }}
            </v-card-title>
            <v-card-actions>
              <v-text-field
                v-model="experiencePointsRelated[0].value"
                label="Unloaded"
                solo
                disabled
              >
              </v-text-field>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col cols="12" lg="4">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ experiencePointsRelated[1].attribute }}
            </v-card-title>
            <v-card-actions>
              <v-text-field
                v-model="experiencePointsRelated[1].value"
                label="XP Earned"
                filled
              >
              </v-text-field>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col cols="12" lg="4">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ statSelect[4].attribute }}
            </v-card-title>
            <v-card-actions>
              <v-text-field
                v-model="statSelect[4].value"
                label="Unloaded"
                solo
                disabled
              >
              </v-text-field>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>

      <!-- This row contains the character's: Abilities, Items, and Spells -->
      <v-row>
        <v-col cols="12" lg="4">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ statSelect[1].attribute }}
            </v-card-title>
            <v-card-actions>
              <v-select
                :items="characterAbilities"
                filled
                label="Abilities Learned"
              >
              </v-select>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col cols="12" lg="4">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ statSelect[2].attribute }}
            </v-card-title>
            <v-card-actions>
              <v-select :items="characterItems" filled label="Items Owned">
              </v-select>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col cols="12" lg="4">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ statSelect[3].attribute }}
            </v-card-title>
            <v-card-actions>
              <v-select :items="characterSpells" filled label="Spells Known">
              </v-select>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>

    <!-- This row contains the character's: Race, Class, Weight, Height, Sex, and Armor Class -->
    <v-container class="characterTraits">
      <v-row wrap>
        <v-col
          cols="12"
          lg="2"
          v-for="row in characterTraits"
          :key="row.attribute"
        >
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{ row.attribute }}
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
  data() {
    return {
      characterAbilities: ['Unloaded Ability 1', 'Unloaded Ability 2'],
      characterItems: ['Unloaded Item 1', 'Unloaded Item 2'],
      characterSpells: ['Unloaded Spell 1', 'Unloaded Spell 2'],
    };
  },
  computed: {
    characterTraits() {
      return [
        { attribute: 'Race', value: this.characterData.race },
        { attribute: 'Class', value: this.characterData.class },
        { attribute: 'Weight', value: this.characterData.weight, max: 1000 },
        { attribute: 'Height', value: this.characterData.height, max: 1000 },
        { attribute: 'Sex', value: this.characterData.sex },
        { attribute: 'Armor Class', value: 'DERIVED' },
      ];
    },
    statSelect() {
      return [
        { attribute: 'Character Name', value: this.characterData.name },
        { attribute: 'Abilities', value: 'Unloaded' },
        { attribute: 'Items', value: 'Unloaded' },
        { attribute: 'Spells', value: 'Unloaded' },
        { attribute: 'Alignment', value: this.characterData.alignment },
      ];
    },
    experiencePointsRelated() {
      return [
        { attribute: 'Level', value: this.level },
        { attribute: 'XP', value: this.characterData.xp_points },
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
