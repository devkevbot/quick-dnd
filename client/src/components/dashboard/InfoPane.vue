<template>
  <v-container>
    <!-- Select dropdown used to select a character. -->
    <v-col cols="3" class="my-n3">
      <v-select
        dense
        solo
        hint="Select a character"
        persistent-hint
        :items="characterNames"
        :value="selectedCharacter"
      ></v-select>
    </v-col>
    <v-card>
      <!-- Tabs used to select the information page to view. -->
      <v-tabs
        v-model="tab"
        background-color="primary"
        centered
        dark
        icons-with-text
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
          <v-card flat>
            <v-card-text>
              <component v-bind:is="item.content"></component>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: 'InfoPane',
  components: {},
  data() {
    return {
      tab: null,
      items: [
        /* TODO: import and add components when they are created */
        { tab: 'Spells', content: '', icon: 'mdi-fire' },
        { tab: 'Items', content: '', icon: 'mdi-sword' },
        { tab: 'Campaigns', content: '', icon: 'mdi-castle' },
      ],
      characters: [],
      selectedCharacter: null,
    };
  },
  /* Fetches all of the user's character data when the component loads. */
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
      })
      .catch(() => {
        /* TODO: Add error handling. */
      });
  },
  computed: {
    /* Extract the name property from each character. */
    characterNames() {
      return this.characters.map((x) => x.name);
    },
  },
};
</script>
