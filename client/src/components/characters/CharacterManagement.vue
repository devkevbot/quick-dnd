<!-- CharacterManagement.vue-->

<template>
  <v-container>
    <v-card>
      <!-- Tabs used to select the information page to view. -->
      <v-tabs v-model="tab" background-color="primary" grow dark>
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
              <component v-bind:is="item.content"> </component>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
    </v-card>
  </v-container>
</template>

<script>
import CreationStepper from './creation/CreationStepper.vue';

export default {
  components: {
    CreationStepper,
  },
  data() {
    return {
      tab: null,
      items: [
        {
          tab: 'Manage Existing Character',
          content: '',
          icon: 'mdi-book-open-outline',
        },
        {
          tab: 'Create a New Character',
          content: CreationStepper,
          icon: 'mdi-plus',
        },
      ],
      characterData: [],
    };
  },
};
</script>
