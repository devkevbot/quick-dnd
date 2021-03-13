<!-- Footer.vue
Renders a basic footer with the curreny year and website name.
-->

<template>
  <v-footer app color="primary" padless>
    <v-row justify="center" no-gutters>
      <v-col class="primary py-4 text-center white--text" cols="12">
        {{ new Date().getFullYear() }} — <strong>QuickDND</strong>
      </v-col>
    </v-row>
  </v-footer>
</template>

<script>
export default {
  name: 'Footer',
};
</script>
