<!-- AppBar.vue
Renders a simple app bar for the website.
-->

<template>
  <v-app-bar app color="primary" dark clipped-left flat>
    <v-icon class="mr-4">mdi-home</v-icon>
    <v-toolbar-title>QuickDND</v-toolbar-title>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script>
export default {
  name: 'AppBar',
};
</script>
