// Component contains the primary characteristics for a character

<template>
<div class="primaryStats">

  <v-btn>
    <h3> Save Changes</h3>
  </v-btn>

<v-container class="characterSelect">
      <v-row>
          <v-col cols="12" lg="4">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[1].attribute}}
              </v-card-title>
              <v-card-actions>
                <v-text-field v-model="statUniqueElements[1].value" label="TEMP" solo disabled>
                </v-text-field>
              </v-card-actions>
            </v-card>
          </v-col>
          <v-col cols="12" lg="4">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[3].attribute}}
              </v-card-title>
              <v-card-actions>
                <v-text-field v-model="statUniqueElements[3].value" label="5000" filled>
                </v-text-field>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" lg="4">
            <v-card flat class="text-xs-center" elevation="8" outlined tile >
              <v-card-title>
                {{statSelect[3].attribute}}
              </v-card-title>
              <v-card-actions>
                <v-select :items="items" filled label="Filled style">
                </v-select>
              </v-card-actions>
            </v-card>
          </v-col>
          <v-col cols="12" lg="4">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statSelect[4].attribute}}
              </v-card-title>
              <v-card-actions>
                <v-select :items="items" filled label="Filled style">
                </v-select>
              </v-card-actions>
            </v-card>
          </v-col>
          <v-col cols="12" lg="4">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statSelect[5].attribute}}
              </v-card-title>
              <v-card-actions>
                <v-select :items="items" filled label="Filled style">
                </v-select>
              </v-card-actions>
            </v-card>
          </v-col>
      </v-row>
  </v-container>

    <v-container class="characterTraits">
    <v-row wrap>
        <v-col cols="12" lg="2" v-for="row in characterTraits" :key="row.attribute">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{row.attribute}}
            </v-card-title>
            <v-card-actions>
              <v-text-field v-model="row.value" label="Temp" solo disabled>
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
    characterTraits() {
      return [
        { attribute: 'Race', value: this.characterData.race },
        { attribute: 'Class', value: this.characterData.class },
        { attribute: 'Weight', value: this.characterData.weight, max: 1000 },
        { attribute: 'Height', value: this.characterData.height, max: 1000 },
        { attribute: 'Sex', value: this.characterData.sex },
        { attribute: 'Armour Class', value: 'DERIVED' },
      ];
    },
    statSelect() {
      return [
        { attribute: 'Campaign', value: 'TEMP' },
        { attribute: 'State', value: 'TEMP' },
        { attribute: 'Character Name', value: this.characterData.name },
        { attribute: 'Abilities', value: 'TEMP' },
        { attribute: 'Items', value: 'TEMP' },
        { attribute: 'Spells', value: 'TEMP' },
        { attribute: 'Alignment', value: this.characterData.alignment },
      ];
    },
    statUniqueElements() {
      return [
        { attribute: 'Location', value: 'TEMP' },
        { attribute: 'Level', value: this.level },
        { attribute: 'Character Background', value: this.characterData.background },
        { attribute: 'XP', value: this.characterData.xp_points },
      ];
    },
  },
};
</script>
