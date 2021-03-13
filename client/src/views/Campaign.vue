<!-- Campaign.vue
This component is the view that shows campaign data, state, and
management options. Assumes the user is authenticated.
-->

<template>
  <div class="campaign">
    <Timeline />
  </div>
</template>

<script>
import Timeline from '../components/campaign/Timeline.vue';

export default {
  name: 'Campaign',
  components: {
    Timeline,
  },
};
</script>
