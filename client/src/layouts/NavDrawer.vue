<!-- NavDrawer.vue
Renders a navigation drawer containing links to other areas of the
website.
-->

<template>
  <v-card>
    <v-navigation-drawer app permanent clipped class="accent">
      <v-list>
        <v-list-item v-for="item in items" :key="item.title" link :to="item.to">
          <v-list-item-icon>
            <v-icon color="primary">{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="pa-2">
          <v-btn block color="primary"> Logout </v-btn>
        </div>
      </template>
    </v-navigation-drawer>
  </v-card>
</template>

<script>
export default {
  name: 'NavDrawer',
  data() {
    return {
      items: [
        /* TODO: update links when they are created */
        { title: 'Dashboard', icon: 'mdi-view-dashboard', to: '/dashboard' },
        { title: 'My profile', icon: 'mdi-account-box', to: '' },
        { title: 'Create a character', icon: 'mdi-gavel', to: '/characters' },
        { title: 'Character Managment', icon: 'mdi-gavel', to: '/characters' },
        { title: 'Start a campaign', icon: 'mdi-castle', to: '/campaign' },
      ],
    };
  },
};
</script>
