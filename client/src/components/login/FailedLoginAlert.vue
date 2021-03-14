<template>
  <v-alert :value="show" type="error">{{ errorText }}</v-alert>
</template>

<script>
export default {
  name: 'FailedLoginAlert',
  props: {
    errorText: {
      required: true,
      type: String,
    },
    show: {
      required: true,
      type: Boolean,
      default: false,
    },
  },
};
</script>

<style>
</style>
