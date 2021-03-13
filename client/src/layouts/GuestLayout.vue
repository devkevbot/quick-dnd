<!-- GuestLayout.vue
This file is one of several layouts used in the website. This layout
should be shown to unauthenticated users (users who are not signed-in).
-->

<template>
  <div class="guest-layout">
    <v-main>
      <v-container fluid>
        <slot />
      </v-container>
    </v-main>
    <Footer />
  </div>
</template>

<script>
import Footer from './Footer.vue';

export default {
  name: 'GuestLayout',
  components: {
    Footer,
  },
};
</script>
