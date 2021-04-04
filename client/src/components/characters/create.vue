// Page is used for character creation

<template>
<div class="background">

  <template>
  <v-stepper v-model="e1">
    <v-stepper-header>
      <v-stepper-step
        :complete="e1 > 1"
        step="1"
      >
        Character Information
      </v-stepper-step>
      <v-divider></v-divider>

      <v-stepper-step
        :complete="e1 > 2"
        step="2"
      >
        Base Stats
      </v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step step="3">
        Class Specifics
      </v-stepper-step>
    </v-stepper-header>

    <v-stepper-items>
<!-- Page 1 -->
      <v-stepper-content step="1">
        <v-row>
        <v-col cols="12" lg="6">
          <v-card>
            <v-form class="px-3">
              <v-text-field label="Name" v-model="selectedName"></v-text-field>
              <v-card-actions>
                <v-slider v-model="selectedHeight" :max="1000" class="align-center"
                label="Height (cm)">
                  <template v-slot:append>
                    <v-text-field v-model="selectedHeight" class="mt-0 pt-0" type="text">
                    </v-text-field>
                  </template>
                </v-slider>
              </v-card-actions>
              <v-card-actions>
                <v-slider v-model="selectedWeight" :max="1000" class="align-center"
                label="Weight (kg)">
                  <template v-slot:append>
                    <v-text-field v-model="selectedWeight" class="mt-0 pt-0" type="text">
                    </v-text-field>
                  </template>
                </v-slider>
              </v-card-actions>
              <v-card-actions>
                <v-slider v-model="selectedXP" :max="355000" class="align-center"
                label="XP">
                  <template v-slot:append>
                    <v-text-field v-model="selectedXP" class="mt-0 pt-0" type="text">
                    </v-text-field>
                  </template>
                </v-slider>
              </v-card-actions>
            </v-form>
          </v-card>
        </v-col>
        <v-col cols="12" lg="6">
          <v-card>
            <v-form class="px-3">
              <v-card-actions>
                <v-select :items="playerClass" filled label="Class" v-model="selectedClass">
                </v-select>
              </v-card-actions>
              <v-card-actions>
                <v-select :items="race" filled label="Race" v-model="selectedRace">
                </v-select>
              </v-card-actions>
              <v-card-actions>
                <v-select :items="alignment" filled label="Alignment" v-model="selectedAlignment">
                </v-select>
              </v-card-actions>
              <v-card-actions>
                  <v-select :items="sex" filled label="Sex" v-model="selectedSex">
                  </v-select>
              </v-card-actions>
            </v-form>
          </v-card>
        </v-col>
        </v-row>

        <v-container class="characterBackground">
        <v-row>
            <v-col cols="12" lg="12">
              <v-card flat class="text-xs-center" elevation="8" outlined tile>
                <v-card-title>
                  {{statUniqueElements[2].attribute}}
                </v-card-title>
                <v-card-actions>
                  <v-textarea
                    v-model="selectedBackground"
                    filled
                    name="input-7-4"
                    label="Write your character's backstory here."
                  ></v-textarea>
                </v-card-actions>
              </v-card>
            </v-col>
        </v-row>
        </v-container>

        <v-btn
          color="primary"
          @click="e1 = 2"
        >
          Continue
        </v-btn>

        <v-btn text>
          Cancel
        </v-btn>
      </v-stepper-content>

<!-- Page 2 -->
      <v-stepper-content step="2">
        <v-card>
          <v-container class="characterSliders">
                <v-row wrap>
                    <v-col cols="12" lg="3" v-for="row in stats" :key="row.attribute">
                      <v-card flat class="text-xs-center" elevation="8" outlined tile>
                        <v-card-title>
                          {{row.attribute}}
                        </v-card-title>
                        <v-card-actions>
                          <v-slider v-model="row.value" :max="row.max" class="align-center">
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
        </v-card>

        <v-btn
          color="primary"
          @click="e1 = 3"
        >
          Continue
        </v-btn>

        <v-btn text>
          Cancel
        </v-btn>
      </v-stepper-content>

<!-- Page 3 -->
<!-- I want to make it so this page only load relevant content to the character.
This is based on what class was selected in step 1. -->
      <v-stepper-content step="3">
        <v-card>
          <v-row>
        <v-col cols="12" lg="6">
          <v-card>
            <v-form class="px-3">
              <v-card-actions>
                <v-select :items="playerClass" filled label="Conditional Load"
                v-model="selectedClass">
                </v-select>
              </v-card-actions>
              <v-card-actions>
                <v-select :items="race" filled label="Conditional Load" v-model="selectedRace">
                </v-select>
              </v-card-actions>
            </v-form>
          </v-card>
        </v-col>
        </v-row>
        </v-card>

        <v-btn
          color="primary"
          @click="e1 = 1"
        >
          Continue
        </v-btn>

        <v-btn text>
          Cancel
        </v-btn>
      </v-stepper-content>
    </v-stepper-items>
  </v-stepper>
</template>

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
  name: 'background',
  data() {
    return {
      e1: 1,
      selectedName: '',
      selectedHeight: '',
      selectedWeight: '',
      selectedXP: '',
      selectedClass: '',
      selectedRace: '',
      selectedAlignment: '',
      slecetedSex: '',
      slecetedBackground: '',
      playerClass: [
        'Barbarian',
        'Bard',
        'Cleric',
        'Druid',
        'Fighter',
        'Monk',
        'Paladin',
        'Ranger',
        'Rogue',
        'Sorcerer',
        'Warlock',
        'Wizard',
        'None',
      ],
      race: [
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
      alignment: [
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
      sex: [
        'Male',
        'Female',
        'Other',
      ],
      stats: [
        { attribute: 'Strength', value: 0, max: 30 },
        { attribute: 'Constitution', value: 0, max: 30 },
        { attribute: 'Wisdom', value: 0, max: 30 },
        { attribute: 'Hit Points', value: 0, max: 440 },
        { attribute: 'Dexterity', value: 0, max: 30 },
        { attribute: 'Ability Points', value: 0, max: 180 },
        { attribute: 'Charisma', value: 0, max: 30 },
        { attribute: 'Intellect', value: 0, max: 30 },
      ],
      barbarianPrimalPath: [
        'Ancestral Guardian',
        'Battlerager',
        'Beast',
        'Berserker',
        'Storm Herald',
        'Totem Warrior',
        'Wild Magic',
        'Zealot',
      ],
      bardCollege: [
        'Creation',
        'Eloquence',
        'Glamour',
        'Lore',
        'Swords',
        'Valor',
        'Whispers',
      ],
      clericDomain: [
        'Arcana',
        'Death',
        'Forge',
        'Grave',
        'Knowledge',
        'Life',
        'Light',
        'Nature',
        'Order',
        'Peace',
        'Tempest',
        'Trickery',
        'Twilight',
        'War',
      ],
      druidCircle: [
        'Dreams',
        'The Land',
        'The Moon',
        'The Shepherd',
        'The Spores',
        'The Stars',
        'Wildfire',
      ],
      fighterArchetype: [
        'Arcane Archer',
        'Banneret',
        'Battle Master',
        'Cavalier',
        'Champion',
        'Echo Knight',
        'Eldritch Knight',
        'Psi Warrior',
        'Rune Knight',
        'Samurai',
      ],
      fightingStyle: [
        'Archery',
        'Blind Fighting',
        'Defense',
        'Druidic Warrior',
        'Dueling',
        'Great Weapon Fighting',
        'Interception',
        'Protection',
        'Superior Technique',
        'Thrown Weapon Fighting',
        'Two-Weapon Fighting',
        'Unarmed Fighting',
      ],
      monkTradition: [
        'Astral Self',
        'Drunken Master',
        'Four Elements',
        'Kensei',
        'Long Death',
        'Mercy',
        'Open Hand',
        'Shadow',
        'Sun Soul',
      ],
      palidinOath: [
        'The Ancients',
        'Conquest',
        'The Crown',
        'Devotion',
        'Glory',
        'Redemption',
        'Vengeance',
        'The Watchers',
        'Oathbreaker',
      ],
      rangerConclave: [
        'Beast Master',
        'Fey Wanderer',
        'Gloom Stalker',
        'Horizon Walker',
        'Hunter',
        'Monster Slayer',
        'Swarmkeeper',
      ],
      rangerFavouredEnemy: '',
      rogueArchetype: [
        'Arcane Trickster',
        'Assassin',
        'Inquisitive',
        'Mastermind',
        'Phantom',
        'Scout',
        'Soulknife',
        'Swashbuckler',
        'Thief',
      ],
      sorcererOrigin: [
        'Aberrant Mind',
        'Clockwork Soul',
        'Draconic Bloodline',
        'Divine Soul',
        'Shadow',
        'Storm',
        'Wild Magic',
      ],
      warlockPatron: [
        'Archfey',
        'Celestial',
        'Fathomless',
        'Fiend',
        'Genie',
        'Great Old One',
        'Hexblade',
        'Undying',
      ],
      wizzardTradition: [
        'Abjuration',
        'Bladesinging',
        'Chronurgy',
        'Conjuration',
        'Divination',
        'Enchantment',
        'Evocation',
        'Graviturgy',
        'Illusion',
        'Necromancy',
        'Order of Scribes',
        'Transmutation',
        'War Magic',
      ],
    };
  },
  computed: {
    statUniqueElements() {
      return [
        { attribute: 'Location', value: '' },
        { attribute: 'Level', value: '' },
        { attribute: 'Character Background', value: '' },
        { attribute: 'XP', value: '' },
      ];
    },
  },
};
</script>
