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

    <!-- Displays notifications to the user. -->
    <v-snackbar
      v-model="notification.show"
      :color="notification.color"
      :timeout="notification.timeout"
      bottom
      rounded
      text
    >
      {{ notification.message }}
    </v-snackbar>
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
  /**
   * Listens for changes to the notifications state and re-renders the
   * `v-snackbar` component when they occur.
   */
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type !== 'notifications/updateState') return;
      this.notification.message = state.notifications.message;
      this.notification.color = state.notifications.color;
      this.notification.timeout = state.notifications.timeout;
      this.notification.show = true;
    });
  },
  data() {
    return {
      notification: {
        show: false,
        color: '',
        timeout: -1,
        message: '',
      },
    };
  },
  computed: {
    layout() {
      return this.$route.meta.layout === 'Main' ? MainLayout : GuestLayout;
    },
  },
};
</script>
