<!-- MainLayout.vue
This file is one of several layouts used in the website. This layout
should be shown to authenticated users (users who are signed-in).
-->

<template>
  <div class="main-layout">
    <NavDrawer />
    <AppBar />
    <v-main>
      <v-container fluid>
        <slot />
      </v-container>
    </v-main>
    <Footer />
  </div>
</template>

<script>
import NavDrawer from './NavDrawer.vue';
import AppBar from './AppBar.vue';
import Footer from './Footer.vue';

export default {
  name: 'MainLayout',
  components: {
    NavDrawer,
    AppBar,
    Footer,
  },
};
</script>
