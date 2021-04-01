<!-- NavDrawer.vue
Renders a navigation drawer containing links to other areas of the
website.
-->

<template>
  <v-card>
    <v-navigation-drawer app color="secondary" permanent clipped>
      <v-list>
        <v-list-item v-for="item in items" :key="item.title" link>
          <v-list-item-icon>
            <v-icon color="accent">{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="pa-2">
          <v-btn block color="accent"> Logout </v-btn>
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
        { title: 'Home', icon: 'mdi-view-dashboard' },
        { title: 'My profile', icon: 'mdi-account-box' },
        { title: 'Create a Character', icon: 'mdi-gavel' },
        { title: 'Character Managment', icon: 'mdi-gavel' },
        { title: 'Start a Campaign', icon: 'mdi-gavel' },
        { title: 'Campaign Managment', icon: 'mdi-gavel' },
      ],
    };
  },
};
</script>
