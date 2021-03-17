<!-- Characters.vue
This component hold the temporary UI layout for character selection.
There are currently 4 types of cards: Sliders, Selects, TextBoxes, and TextBoxesDisabled.
Currently a work in progress.
-->

<template>
<div class="characters">

  <v-btn>
    Save Changes
  </v-btn>

  <v-container class="characterSliders">
      <v-row wrap>
          <v-col cols="12" md="6" lg="2" v-for="row in statSlider" :key="row.attribute">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{row.attribute}}
              </v-card-title>
              <v-card-subtitle>
                <!-- Reserved for subtitle -->
              </v-card-subtitle>
              <v-card-text>
                <!--Reserved for Text -->
              </v-card-text>
                <v-card-actions>
                  <v-slider v-model="statSliderValue" :max="100" class="align-center">
                    <template v-slot:append>
                      <v-text-field v-model="statSliderValue" class="mt-0 pt-0" type="number">
                      </v-text-field>
                    </template>
                  </v-slider>
                </v-card-actions>
              </v-card>
          </v-col>
      </v-row>
  </v-container>

  <v-container class="characterSelects">
  <v-row wrap>
      <v-col cols="12" md="6" lg="2" v-for="row in statSelect" :key="row.attribute">
        <v-card flat class="text-xs-center" elevation="8" outlined tile>
          <v-card-title>
            {{row.attribute}}
          </v-card-title>
          <v-card-subtitle>
            <!-- Reserved for subtitle -->
          </v-card-subtitle>
          <v-card-text>
            <!--Reserved for Text -->
          </v-card-text>
            <v-card-actions>
              <v-select :items="items" filled label="Filled style">
              </v-select>
            </v-card-actions>
          </v-card>
      </v-col>
  </v-row>
</v-container>

<v-container class="characterTextBox">
  <v-row wrap>
      <v-col cols="12" md="6" lg="2" v-for="row in statTextBox" :key="row.attribute">
        <v-card flat class="text-xs-center" elevation="8" outlined tile>
          <v-card-title>
            {{row.attribute}}
          </v-card-title>
          <v-card-subtitle>
            <!-- Reserved for subtitle -->
          </v-card-subtitle>
          <v-card-text>
            <!--Reserved for Text -->
          </v-card-text>
            <v-card-actions>
                <v-text-field label="Temp" solo >
                </v-text-field>
            </v-card-actions>
          </v-card>
      </v-col>
  </v-row>
</v-container>

  <v-container class="characterTextBoxDisabled">
    <v-row wrap>
        <v-col cols="12" md="6" lg="2" v-for="row in statTextBoxDisabled" :key="row.modifier">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{row.modifier}}
            </v-card-title>
            <v-card-subtitle>
              <!-- Reserved for subtitle -->
            </v-card-subtitle>
            <v-card-text>
              <!--Reserved for Text -->
            </v-card-text>
              <v-card-actions>
                  <v-text-field label="Temp" solo disabled>
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
  data() {
    return {
      statSlider: [
        { attribute: 'Strength' },
        { attribute: 'Contribution' },
        { attribute: 'Wisdom' },
        { attribute: 'Hit Points' },
        { attribute: 'Dexterity' },
        { attribute: 'Ability_Points' },
        { attribute: 'Charisma' },
        { attribute: 'Intellect' },
      ],
      statSelect: [
        { attribute: 'Abilities' },
        { attribute: 'Items' },
        { attribute: 'Spells' },
        { attribute: 'Character Name' },
        { attribute: 'Campaign' },
        { attribute: 'State' },
      ],
      statTextBox: [
        { attribute: 'Level' },
        { attribute: 'XP' },
        { attribute: 'Alignment' },
        { attribute: 'Background' },
        { attribute: 'Speed' },
        { attribute: 'Weight' },
        { attribute: 'Height' },
        { attribute: 'Sex' },
        { attribute: 'ArmourClass' },
        { attribute: 'Location' },
      ],
      statTextBoxDisabled: [
        { modifier: 'Initiative' },
        { modifier: 'Number of Dies' },
        { modifier: 'Strength' },
        { modifier: 'Constitution' },
        { modifier: 'Wisdom' },
        { modifier: 'Hit Points' },
        { modifier: 'Dexterity' },
        { modifier: 'Ability Points' },
        { modifier: 'Charisma' },
        { modifier: 'Intellect' },
        { modifier: 'Proficency' },
      ],
    };
  },
};
</script>
