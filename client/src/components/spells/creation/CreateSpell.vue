<template>
  <v-form>
    <v-container>
      <!-- Character selection dropdown -->
      <v-row>
        <v-col cols="4">
          <p class="headline primary--text">Select your character</p>
          <v-select
            solo
            prepend-inner-icon="mdi-account-search"
            :items="characterNames"
            v-model="selectedCharName"
          >
            <!-- Button used to refresh the character list. -->
            <template v-slot:append-outer>
              <v-btn
                color="primary"
                fab
                small
                class="mt-n2"
                @click="fetchUserCharacters"
              >
                <v-icon>mdi-refresh</v-icon>
              </v-btn>
            </template>
          </v-select>
        </v-col>
      </v-row>

      <v-card class="pl-2">
        <!-- Spell name, shcool, and concentration. -->
        <p class="headline primary--text">Spell Information</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="spell_name"
              label="Enter name"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model="school"
              :items="schoolOptions"
              label="Choose School"
              outlined
              required
            >
            </v-select>
          </v-col>

          <v-col cols="12" md="2">
            <v-select
              v-model.trim="concentration"
              :items="concentrationOptions"
              label="Choose Concentration"
              outlined
              required
            >
            </v-select>
          </v-col>
        </v-row>

        <!-- Spell cast time, duration, and range. -->
        <p class="headline primary--text">Values</p>
        <v-row>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="casting_time"
              label="Enter cast time"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="duration"
              label="Enter duration"
              suffix="s"
              outlined
              required
            >
            </v-text-field>
          </v-col>

          <v-col cols="12" md="2">
            <v-text-field
              v-model.trim="range"
              label="Enter Range"
              suffix="ft"
              outlined
              required
            >
            </v-text-field>
          </v-col>
        </v-row>

        <!-- Spell description. -->
        <p class="headline primary--text">Description</p>
        <v-row>
          <v-col cols="12" md="6">
            <v-textarea
              v-model.trim="description"
              label="Enter description"
              placeholder=""
              outlined
              clearable
              no-resize
              required
              counter="250"
            >
            </v-textarea>
          </v-col>
        </v-row>

        <v-btn class="primary" @click="onComplete">Create</v-btn>
      </v-card>
    </v-container>
  </v-form>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'CreateSpell',
  data() {
    return {
      characters: [],
      selectedCharName: null,

      spellName: '',
      level: '',
      school: '',
      concentration: '',
      description: '',
      castingTime: '',
      range: '',
      duration: '',

      schoolOptions: [
        'Abjuration',
        'Conjuration',
        'Divination',
        'Enchantment',
        'Evocation',
        'Illusion',
        'Necromancy',
        'Transmutation',
      ],
      concentrationOptions: ['True', 'False'],
    };
  },
  /**
   * Fetches all of the user's character data when the component loads.
   * Does not handle fetch errors.
   * Not sure how to handle spell loading.
   */
  async created() {
    await this.fetchUserCharacters();
  },
  computed: {
    /**
     * @returns {Array<String>} - The names of the user's characters (if
     * any).
     */
    characterNames() {
      return this.characters.map((x) => x.name);
    },
    /**
     * @returns {String} - The character currently selected by the user.
     */
    selectedCharacter() {
      if (this.selectedCharName === null) return '';

      const target = this.selectedCharName;
      return this.characters.filter((x) => x.name === target)[0];
    },
  },
  methods: {
    ...mapActions({
      display: 'notifications/display',
    }),
    async fetchUserCharacters() {
      const requestURI = 'auth/character/me';
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.characters = resp.data.data.characters;
          if (this.characters) {
            const targetChar = this.characters[0];
            this.selectedCharName = targetChar.name;
          }
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * Updates the current character spell data. Triggered when the
     * 'Create Button' has been clicked.
     * @param {{
     *  spellName: String
     *  level: String
     *  school: String
     *  concentration: String
     *  description: String
     *  castingTime: String
     *  range: String
     *  duration: String
     *  }} spell
     */
    updateCharacterSpellData({
      charID,
      spellName,
      level,
      school,
      concentration,
      description,
      castingTime,
      range,
      duration,
    }) {
      this.spell.characterId = charID;
      this.spell.spellName = spellName;
      this.spell.level = parseInt(level, 10);
      this.spell.school = school;
      this.spell.concentration = concentration;
      this.spell.description = description;
      this.spell.castingTime = parseInt(castingTime, 10);
      this.spell.range = parseInt(range, 10);
      this.spell.duration = parseInt(duration, 10);
    },
  },
};
</script>
