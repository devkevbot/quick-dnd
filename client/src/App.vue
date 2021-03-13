<!-- App.vue
This component is the starting place for all other component rendering.
Layouts are computed and shown. The main content of the page is
determined from the current page route and shown by <router-view>.
-->

<template>
  <v-app id="app">
    <component :is="layout">
      <router-view />
    </component>
  </v-app>
</template>

<script>
import MainLayout from './layouts/MainLayout.vue';
import GuestLayout from './layouts/GuestLayout.vue';

export default {
  name: 'App',
  components: {
    MainLayout,
    GuestLayout,
  },
  computed: {
    layout() {
      return this.$route.meta.layout === 'Main' ? MainLayout : GuestLayout;
    },
  },
};
</script>
