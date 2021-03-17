<!-- Characters.vue
This component hold the temporary UI layout for character selection.
There are currently 4 types of cards: Sliders, Selects, TextBoxes, and TextBoxesDisabled.
Currently a work in progress.
-->

<template>
<div class="characters">

  <v-btn>
    Temp Save Button
  </v-btn>

  <v-container class="campaign">
      <v-row>
          <v-col cols="12" lg="6">
            <v-card flat class="text-xs-center" elevation="8" outlined tile >
              <v-card-title>
                {{statSelect[0].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-select :items="items" filled label="Filled style">
                  </v-select>
                </v-card-actions>
              </v-card>
          </v-col>
          <v-col cols="12" lg="3">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statSelect[1].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-select :items="items" filled label="Filled style">
                  </v-select>
                </v-card-actions>
              </v-card>
          </v-col>
          <v-col cols="12" lg="3">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[0].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-text-field label="Temp" solo >
                  </v-text-field>
                </v-card-actions>
              </v-card>
          </v-col>
      </v-row>
  </v-container>

  <v-container class="characterSelect">
      <v-row>
          <v-col cols="12" lg="3">
            <v-card flat class="text-xs-center" elevation="8" outlined tile >
              <v-card-title>
                {{statSelect[2].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-select :items="items" filled label="Filled style">
                  </v-select>
                </v-card-actions>
              </v-card>
          </v-col>
          <v-col cols="12" lg="2">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[1].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-text-field label="Level 24" solo disabled>
                  </v-text-field>
                </v-card-actions>
              </v-card>
          </v-col>
          <v-col cols="12" lg="1">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[3].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-text-field label="5000" filled>
                  </v-text-field>
                </v-card-actions>
              </v-card>
          </v-col>
          <v-col cols="12" lg="2">
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
          <v-col cols="12" lg="2">
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
          <v-col cols="12" lg="2">
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
                  <v-text-field label="Temp" solo disabled>
                  </v-text-field>
              </v-card-actions>
            </v-card>
        </v-col>
    </v-row>
  </v-container>

    <v-container class="characterBackground">
      <v-row>
          <v-col cols="12" lg="12">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[2].attribute}}
              </v-card-title>
                <v-card-actions>
                  <v-textarea
                    filled
                    name="input-7-4"
                    label="Filled textarea"
                    value="Filled text."
                  ></v-textarea>
                </v-card-actions>
              </v-card>
          </v-col>
      </v-row>
  </v-container>

  <v-container class="characterSliders">
      <v-row wrap>
          <v-col cols="12" lg="3" v-for="row in statSlider" :key="row.attribute">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{row.attribute}}
              </v-card-title>
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

  <v-container class="characterModifiers">
    <v-row wrap>
        <v-col cols="12" md="6" lg="2" v-for="row in statModifiers" :key="row.modifier">
          <v-card flat class="text-xs-center" elevation="8" outlined tile>
            <v-card-title>
              {{row.modifier}}
            </v-card-title>
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
        { attribute: 'Campaign' },
        { attribute: 'State' },
        { attribute: 'Character Name' },
        { attribute: 'Abilities' },
        { attribute: 'Items' },
        { attribute: 'Spells' },
        { attribute: 'Alignment' },
      ],
      characterTraits: [
        { attribute: 'Race' },
        { attribute: 'Class' },
        { attribute: 'Weight' },
        { attribute: 'Height' },
        { attribute: 'Sex' },
        { attribute: 'Armour Class' },
      ],
      statModifiers: [
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
        { modifier: 'Speed' },
      ],
      statUniqueElements: [
        { attribute: 'Location' },
        { attribute: 'Level' },
        { attribute: 'Character Background' },
        { attribute: 'XP' },
      ],
    };
  },
};
</script>
