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
          <div
            v-for="attribute in selectedClassAttributes"
            :key="attribute.name"
          >
            <!-- Values for each attribute. -->
            <v-select
              :label="attribute.name"
              :items="attribute.value"
              required
            ></v-select>
          </div>
        </span>
      </v-col>
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
          attributes: [
            {
              name: 'Primal Path',
              value: [
                'Ancestral Guardian',
                'Battlerager',
                'Beast',
                'Berserker',
                'Storm Herald',
                'Totem Warrior',
                'Wild Magic',
                'Zealot',
              ],
            },
          ],
        },
        {
          name: 'Bard',
          attributes: [
            {
              name: 'College',
              value: [
                'Creation',
                'Eloquence',
                'Glamour',
                'Lore',
                'Swords',
                'Valor',
                'Whispers',
              ],
            },
          ],
        },
        {
          name: 'Cleric',
          attributes: [
            {
              name: 'Domain',
              value: [
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
            },
          ],
        },
        {
          name: 'Druid',
          attributes: [
            {
              name: 'Circle',
              value: [
                'Dreams',
                'The Land',
                'The Moon',
                'The Shepherd',
                'The Spores',
                'The Stars',
                'Wildfire',
              ],
            },
          ],
        },
        {
          name: 'Fighter',
          attributes: [
            {
              name: 'Archetype',
              value: [
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
            },
            {
              name: 'Fighting Style',
              value: [
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
            },
          ],
        },
        {
          name: 'Monk',
          attributes: [
            {
              name: 'Tradition',
              value: [
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
            },
          ],
        },
        {
          name: 'Paladin',
          attributes: [
            {
              name: 'Oath',
              value: [
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
            },
            {
              name: 'Fighting Style',
              value: [
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
            },
          ],
        },
        {
          /* TODO: insert "favored enemy" attribute. */
          name: 'Ranger',
          attributes: [
            {
              name: 'Conclave',
              value: [
                'Beast Master',
                'Fey Wanderer',
                'Gloom Stalker',
                'Horizon Walker',
                'Hunter',
                'Monster Slayer',
                'Swarmkeeper',
              ],
            },
          ],
        },
        {
          name: 'Rogue',
          attributes: [
            {
              name: 'Archetype',
              value: [
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
            },
          ],
        },
        {
          name: 'Sorcerer',
          attributes: [
            {
              name: 'Origin',
              value: [
                'Aberrant Mind',
                'Clockwork Soul',
                'Draconic Bloodline',
                'Divine Soul',
                'Shadow',
                'Storm',
                'Wild Magic',
              ],
            },
          ],
        },
        {
          name: 'Warlock',
          attributes: [
            {
              name: 'Patron',
              value: [
                'Archfey',
                'Celestial',
                'Fathomless',
                'Fiend',
                'Genie',
                'Great Old One',
                'Hexblade',
                'Undying',
              ],
            },
          ],
        },
        {
          name: 'Wizard',
          attributes: [
            {
              name: 'Arcane Tradition',
              value: [
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
            },
          ],
        },
      ],
    };
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
     * @returns {Array<Object>} A list of attribute objects for the
     * class that is selected.
     */
    selectedClassAttributes() {
      const c = this.classOptions.filter((x) => x.name === this.selectedClass);
      return c.map((x) => x.attributes)[0];
    },
  },
};
</script>
