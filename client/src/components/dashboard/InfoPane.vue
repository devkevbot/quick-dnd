<template>
  <v-container>
    <!-- Select dropdown used to select a character. -->
    <v-col cols="12" md="6" lg="4" class="my-n3">
      <v-row class="ml-n3 mb-1">
        <p class="display-1 primary--text">Select your character</p>
      </v-row>

      <v-select
        class="ml-n3"
        solo
        prepend-inner-icon="mdi-account-search"
        :items="characterNames"
        v-model="selectedCharName"
        @change="fetchTabData(tab)"
      ></v-select>
    </v-col>
    <v-card>
      <!-- Tabs used to select the information page to view. -->
      <v-tabs
        v-model="tab"
        background-color="primary"
        grow
        dark
        @change="fetchTabData(tab)"
      >
        <!-- The individual tab components. -->
        <v-tab v-for="item in items" :key="item.tab">
          {{ item.tab }}
          <v-icon class="ml-2">{{ item.icon }}</v-icon>
        </v-tab>
      </v-tabs>

      <!-- Conditionally renders a component based on what tab is
      selected by the user. -->
      <v-tabs-items v-model="tab">
        <v-tab-item v-for="item in items" :key="item.tab">
          <v-card flat class="secondary">
            <v-card-text>
              <component v-bind:is="item.content" :data="data"></component>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
    </v-card>
  </v-container>
</template>

<script>
import Spells from './Spells.vue';

export default {
  name: 'InfoPane',
  components: {
    Spells,
  },
  data() {
    return {
      tab: null,
      items: [
        /* TODO: import and add components when they are created */
        {
          tab: 'Spells',
          content: Spells,
          icon: 'mdi-fire',
        },
        {
          tab: 'Items',
          content: '',
          icon: 'mdi-sword',
        },
        {
          tab: 'Campaigns',
          content: '',
          icon: 'mdi-castle',
        },
      ],

      characters: [],
      selectedCharName: null,

      data: [],
      spells: [],
    };
  },
  /**
   * Fetches all of the user's character data when the component loads.
   */
  async created() {
    const requestURI = 'auth/character/me';
    const method = 'GET';

    await this.$http({
      url: requestURI,
      data: null,
      method,
    })
      .then((resp) => {
        this.characters = resp.data.data.characters;
        if (this.characters.length > 0) {
          const targetChar = this.characters[0];
          this.selectedCharName = targetChar.name;
          this.fetchCharacterSpells(targetChar.id);
        }
      })
      .catch(() => {
        /* TODO: Add error handling. */
      });
  },
  computed: {
    /**
     * @returns {Array<String>} - A list of character names that the
     * in user has created.
     */
    characterNames() {
      return this.characters.map((x) => x.name);
    },
    /**
     * @returns {String} - The ID of the character currently selected by
     * the user.
     */
    selectedCharacterID() {
      if (this.selectedCharName === null) return '';

      const target = this.selectedCharName;
      return this.characters.filter((x) => x.name === target)[0].id;
    },
  },
  methods: {
    /** Fetches data dependent on which tab the user has selected. For
     * example, if the user has selected the "Spells" tab, spell data
     * will be fetched.
     * @param {String} tabNumber - The number of the item tab to fetch data for.
     */
    async fetchTabData(tabNumber) {
      if (tabNumber === null || this.selectedCharName === null) return;

      const selected = this.items[tabNumber];

      switch (selected.tab) {
        case 'Spells':
          await this.fetchCharacterSpells(this.selectedCharacterID);
          break;
        /* TODO: create item fetching API */
        case 'Items':
          break;
        /* TODO: create campaign fetching API */
        case 'Campaigns':
          break;
        default:
          break;
      }
    },
    /**
     * @param {String} charID - The character ID which can help to identify the spells
     * to fetch.
     */
    async fetchCharacterSpells(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/spell`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.spells = resp.data.data.spells;
          this.data = this.spells;
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
  },
};
</script>
