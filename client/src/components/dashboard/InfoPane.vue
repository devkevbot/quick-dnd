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
import { mapActions } from 'vuex';
import Spells from './Spells.vue';
import Items from './Items.vue';
import Campaigns from './Campaigns.vue';

export default {
  name: 'InfoPane',
  components: {
    Spells,
    Items,
  },
  data() {
    return {
      tab: null,
      items: [
        {
          tab: 'Spells',
          content: Spells,
          icon: 'mdi-fire',
        },
        {
          tab: 'Items',
          content: Items,
          icon: 'mdi-sword',
        },
        {
          tab: 'Campaigns',
          content: Campaigns,
          icon: 'mdi-castle',
        },
      ],

      characters: [],
      selectedCharName: null,

      data: {
        items: [],
        itemStats: {},
        spells: [],
        spellCountPerSchool: [],
        campaigns: [],
      },
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
      .then(async (resp) => {
        const fetchedCharacters = resp.data.data.characters;
        if (fetchedCharacters === null) return;

        this.characters = fetchedCharacters;
        const targetChar = this.characters[0];
        this.selectedCharName = targetChar.name;

        await this.fetchSpellData(targetChar.id);
        await this.fetchSpellStats(targetChar.id);
      })
      .catch((err) => {
        let message = 'Could not fetch characters. Please try again.';
        if (err.response) {
          message = `Error: ${err.response.data.message}. Please try again.`;
        }
        this.display({
          message,
          color: 'error',
          timeout: 10000,
        });
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
    ...mapActions({
      display: 'notifications/display',
    }),
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
          await this.fetchSpellData(this.selectedCharacterID);
          await this.fetchSpellStats(this.selectedCharacterID);
          break;
        case 'Items':
          await this.fetchItemData(this.selectedCharacterID);
          await this.fetchItemStats(this.selectedCharacterID);
          break;
        case 'Campaigns':
          await this.fetchCampaignData(this.selectedCharacterID);
          break;
        default:
          break;
      }
    },
    /**
     * @param {String} charID - The character ID which can help to identify the spells
     * to fetch.
     */
    async fetchSpellData(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/spell`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.data.spells = resp.data.data.spells ?? [];
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * @param {String} charID - The character ID which can help to
     * identify the character data to fetch.
     */
    async fetchSpellStats(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/spell/count-per-school`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.data.spellCountPerSchool = resp.data.data.spells_count ?? [];
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * @param {String} charID - The character ID which can help to identify the items
     * to fetch.
     */
    async fetchItemData(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/item`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.data.items = resp.data.data.items ?? [];
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * @param {String} charID - The character ID which can help to
     * identify the character data to fetch.
     */
    async fetchItemStats(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/item/stats`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.data.itemStats = resp.data.data.stats ?? {};
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * @param {String} charID - The character ID which can help to
     * identify the campaign data to fetch.
     */
    async fetchCampaignData(charID) {
      const integerID = parseInt(charID, 10);

      const requestURI = `auth/character/${integerID}/campaign`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.data.campaigns = resp.data.data.campaigns ?? [];
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
  },
};
</script>
