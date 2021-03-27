<!-- Characters.vue
This component hold the temporary UI layout for character selection.
There are currently 4 types of cards: Sliders, Selects, TextBoxes, and TextBoxesDisabled.
Currently a work in progress.
-->

<template>
<div class="characters">

<template>
   <div class="stat-card">
     <v-container class="mx-2">
       <v-row>
         <v-col cols="4">
           <v-text-field
             v-model="characterID"
             outlined
             label="Search for a character by ID"
             placeholder="Enter a numeric character ID"
             type="text"
             append-icon="mdi-send"
             @click:append="fetchCharacter(characterID)"
             @keyup.enter="fetchCharacter(characterID)"
           >
           </v-text-field>
         </v-col>
       </v-row>
       <v-row>
         <v-card flat>
           <v-card-title>Results:</v-card-title>
           <v-card-text>
             <pre>{{ characterDataJSON }}</pre>
             <pre>{{ characterData.id }}</pre>
           </v-card-text>
         </v-card>
       </v-row>
     </v-container>
   </div>
 </template>

  <v-btn>
    <h3> SAVE </h3>
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
                <v-text-field v-model="statUniqueElements[1].value" label="TEMP" solo disabled>
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
                <v-text-field v-model="statUniqueElements[3].value" label="5000" filled>
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
              <v-text-field v-model="row.value" label="Temp" solo disabled>
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
                  v-model="statUniqueElements[2].value"
                  filled
                  name="input-7-4"
                  label="Filled textarea"
                  value="statUniqueElements[2].value will go here"
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
                <v-slider v-model="row.value" :max="100" class="align-center">
                  <template v-slot:append>
                    <v-text-field v-model="row.value" class="mt-0 pt-0" type="text">
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
              <v-text-field v-model="row.value" label="TEMP" solo disabled>
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
      characterID: undefined,
      characterDataJSON: {},
      characterData: {},
    };
  },
  computed: {
    characterTraits() {
      return [
        { attribute: 'Race', value: this.characterData.race },
        { attribute: 'Class', value: this.characterData.class },
        { attribute: 'Weight', value: this.characterData.weight },
        { attribute: 'Height', value: this.characterData.height },
        { attribute: 'Sex', value: this.characterData.sex },
        { attribute: 'Armour Class', value: 'TEMP' },
      ];
    },
    statSlider() {
      return [
        { attribute: 'Strength', value: this.characterData.strength },
        { attribute: 'Constitution', value: this.characterData.constitution },
        { attribute: 'Wisdom', value: this.characterData.wisdom },
        { attribute: 'Hit Points', value: this.characterData.hp_max },
        { attribute: 'Dexterity', value: this.characterData.dexterity },
        { attribute: 'Ability Points', value: this.characterData.ability_points },
        { attribute: 'Charisma', value: this.characterData.charisma },
        { attribute: 'Intellect', value: this.characterData.intelligence },
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
    statModifiers() {
      return [
        { modifier: 'Initiative', value: 'TEMP' },
        { modifier: 'Number of Dies', value: 'TEMP' },
        { modifier: 'Strength', value: Math.floor((this.characterData.strength - 10) / 2) },
        { modifier: 'Constitution', value: Math.floor((this.characterData.constitution - 10) / 2) },
        { modifier: 'Wisdom', value: Math.floor((this.characterData.wisdom - 10) / 2) },
        { modifier: 'Hit Points', value: 'TEMP' },
        { modifier: 'Dexterity', value: Math.floor((this.characterData.dexterity - 10) / 2) },
        { modifier: 'Ability Points', value: 'TEMP' },
        { modifier: 'Charisma', value: Math.floor((this.characterData.charisma - 10) / 2) },
        { modifier: 'Intellect', value: Math.floor((this.characterData.intelligence - 10) / 2) },
        { modifier: 'Proficency', value: (Math.ceil((this.level) / 4) + 1) },
        { modifier: 'Speed', value: this.characterData.speed },
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
      } else { tempLevel = 20; }
      return tempLevel;
    },
  },
  methods: {
    async fetchCharacter(id) {
      const numericID = parseInt(id, 10);
      const requestURI = `/character/${numericID}`;
      const method = 'GET';

      await this.$http({ url: requestURI, data: null, method })
        .then((resp) => {
          this.characterDataJSON = JSON.stringify(resp.data.data, null, '\t');
          // this.characterData = JSON.parse(this.characterDataJSON);
          this.characterData = resp.data.data;
        })
        .catch(() => {
          const msg = { error: 'character not found' };
          this.characterDataJSON = JSON.stringify(msg, null, '\t');
        });
    },
  },
};
</script>
