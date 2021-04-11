<template>
  <v-form>
    <v-container>
      <p class="headline primary--text">Select character class</p>
      <v-col cols="12" md="6">
        <!-- Display a list of classes that the user can select from. -->
        <v-select
          :items="listOfClassNames"
          v-model="selectedClass"
          required
        ></v-select>

        <!-- Once the user has selected a class, let them choose the
        class-specific attributes. -->
        <span v-if="selectedClass">
          <!-- Values for each attribute. -->
          <v-select
            v-model="classSpecificAttrs.selected"
            :label="classSpecificAttrs.name"
            :items="classSpecificAttrs.values"
            required
          ></v-select>
        </span>
      </v-col>

      <v-btn text @click="$emit('cancel')">Back</v-btn>
      <v-btn class="primary" @click="onComplete">Create</v-btn>
    </v-container>
  </v-form>
</template>

<script>
export default {
  name: 'Class',
  data() {
    return {
      selectedClass: null,
      /* Data for each class, including any attributes relevant to that
      class. */
      classOptions: [
        {
          name: 'Barbarian',
          attributes: {
            name: 'Primal Path',
            dataName: 'primal_path',
            values: [
              'Ancestral Guardian',
              'Battlerager',
              'Beast',
              'Berserker',
              'Storm Herald',
              'Totem Warrior',
              'Wild Magic',
              'Zealot',
            ],
            selected: null,
          },
        },
        {
          name: 'Bard',
          attributes: {
            name: 'College',
            dataName: 'college',
            values: [
              'Creation',
              'Eloquence',
              'Glamour',
              'Lore',
              'Swords',
              'Valor',
              'Whispers',
            ],
            selected: null,
          },
        },
        {
          name: 'Cleric',
          attributes: {
            name: 'Domain',
            dataName: 'divine_domain',
            values: [
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
            selected: null,
          },
        },
        {
          name: 'Druid',
          attributes: {
            name: 'Circle',
            dataName: 'druid_circle',
            values: [
              'Dreams',
              'The Land',
              'The Moon',
              'The Shepherd',
              'The Spores',
              'The Stars',
              'Wildfire',
            ],
            selected: null,
          },
        },
        {
          name: 'Fighter',
          attributes: {
            name: 'Archetype',
            dataName: 'archetype',
            values: [
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
            selected: null,
          },
        },
        {
          name: 'Monk',
          attributes: {
            name: 'Tradition',
            dataName: 'monastic_tradition',
            values: [
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
            selected: null,
          },
        },
        {
          name: 'Paladin',
          attributes: {
            name: 'Oath',
            dataName: 'oath',
            values: [
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
            selected: null,
          },
        },
        {
          name: 'Ranger',
          attributes: {
            name: 'Conclave',
            dataName: 'conclave',
            values: [
              'Beast Master',
              'Fey Wanderer',
              'Gloom Stalker',
              'Horizon Walker',
              'Hunter',
              'Monster Slayer',
              'Swarmkeeper',
            ],
            selected: null,
          },
        },
        {
          name: 'Rogue',
          attributes: {
            name: 'Archetype',
            dataName: 'archetype',
            values: [
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
            selected: null,
          },
        },
        {
          name: 'Sorcerer',
          attributes: {
            name: 'Origin',
            dataName: 'sorcerous_origin',
            values: [
              'Aberrant Mind',
              'Clockwork Soul',
              'Draconic Bloodline',
              'Divine Soul',
              'Shadow',
              'Storm',
              'Wild Magic',
            ],
            selected: null,
          },
        },
        {
          name: 'Warlock',
          attributes: {
            name: 'Patron',
            dataName: 'patron',
            values: [
              'Archfey',
              'Celestial',
              'Fathomless',
              'Fiend',
              'Genie',
              'Great Old One',
              'Hexblade',
              'Undying',
            ],
            selected: null,
          },
        },
        {
          name: 'Wizard',
          attributes: {
            name: 'Arcane Tradition',
            dataName: 'arcane_tradition',
            values: [
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
            selected: null,
          },
        },
      ],
    };
  },
  methods: {
    /**
     * Emits the up-to-date class data entered by the user.
     * Should be used a click handler for the "next step" button.
     */
    onComplete() {
      const data = this.prepareDataForEmit();
      this.$emit('complete', data);
    },
    /**
     * @returns {Object} - Formatted class data that is ready to emit.
     */
    prepareDataForEmit() {
      const data = {
        className: this.selectedClass,
        classAttribute: this.classSpecificAttrs.selected,
      };
      return data;
    },
  },
  computed: {
    /**
     * @returns {Array<String>} A list of class names that the user can
     * select from.
     */
    listOfClassNames() {
      return this.classOptions.map((x) => x.name);
    },
    /**
     * @returns {Array<Object>} The class-specific attributes for the
     * class that the user has selected.
     */
    classSpecificAttrs() {
      const c = this.classOptions.filter((x) => x.name === this.selectedClass);
      return c.map((x) => x.attributes)[0];
    },
  },
};
</script>
