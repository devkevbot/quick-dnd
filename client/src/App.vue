<template>
  <v-app>
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

    <v-app-bar app color="primary" dark clipped-left flat>
      <v-icon class="mr-4">mdi-home</v-icon>
      <v-toolbar-title>QuickDND</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>

    <v-footer app color="primary" padless>
      <v-row justify="center" no-gutters>
        <v-col class="primary py-4 text-center white--text" cols="12">
          {{ new Date().getFullYear() }} — <strong>QuickDND</strong>
        </v-col>
      </v-row>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      items: [
        { title: 'Home', icon: 'mdi-view-dashboard' },
        { title: 'My profile', icon: 'mdi-account-box' },
        { title: 'Create a character', icon: 'mdi-gavel' },
      ],
    };
  },
};
</script>
