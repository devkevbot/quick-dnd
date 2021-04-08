<!-- CharacterManagement.vue
This is the page which houses the components for character managment.
Here you are able to create, update, and delete character data.
-->

<template>
  <div class="characters">
    <template>
      <v-container>
        <!-- Select dropdown used to select a character. -->
        <v-col cols="12" md="6" lg="4" class="my-n3">
          <v-row class="ml-n3 mb-1">
            <p class="display-1 primary--text">Select your character</p>
          </v-row>
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
                  <component
                    v-bind:is="item.content"
                    :characterData="characterData"
                  >
                  </component>
                </v-card-text>
              </v-card>
            </v-tab-item>
          </v-tabs-items>
        </v-card>
      </v-container>
    </template>

    <!-- Temp in until character dropdown is compeletly implemented -->
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
                <pre>{{ characterNames }}</pre>
              </v-card-text>
            </v-card>
          </v-row>
        </v-container>
      </div>
    </template>
  </div>
</template>

<script>
import PrimaryStats from './MainDetails.vue';
import StatModifiers from './StatModifiers.vue';
import ModifyBackground from './ModifyBackground.vue';
import CreationStepper from './CreationStepper.vue';

export default {
  components: {
    PrimaryStats,
    StatModifiers,
    ModifyBackground,
    CreationStepper,
  },
  data() {
    return {
      characterID: undefined,
      characterDataJSON: {},
      characters: [],
      selectedCharName: null,
      tab: null,
      e1: 1,
      items: [
        {
          tab: 'Main Stats',
          content: PrimaryStats,
          icon: 'mdi-fire',
        },
        {
          tab: 'Modifiers',
          content: StatModifiers,
          icon: 'mdi-sword',
        },
        {
          tab: 'Character Background',
          content: ModifyBackground,
          icon: 'mdi-castle',
        },
        {
          tab: 'Create a New Character',
          content: CreationStepper,
          icon: 'mdi-fire',
        },
      ],
      characterData: [],
    };
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
    async fetchCharacter(id) {
      const numericID = parseInt(id, 10);
      const requestURI = `/character/${numericID}`;
      const method = 'GET';
      await this.$http({ url: requestURI, data: null, method })
        .then((resp) => {
          this.characterDataJSON = JSON.stringify(resp.data.data, null, '\t');
          // this.characterData = JSON.parse(this.characterDataJSON);
          this.characterData = resp.data.data;
          this.characterDataTemp = this.characterData;
        })
        .catch(() => {
          const msg = { error: 'character not found' };
          this.characterDataJSON = JSON.stringify(msg, null, '\t');
        });
    },
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
};
</script>
