<template>
  <div class="dashboard">
    <InfoPane />
  </div>
</template>

<script>
import InfoPane from '../components/dashboard/InfoPane.vue';

export default {
  name: 'Dashboard',
  components: {
    InfoPane,
  },
};
</script>
