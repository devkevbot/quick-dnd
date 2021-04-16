<!-- Campaign.vue
This component is the view shown when a user is on the Campaign creation page.
-->

<template>
  <div class="campaign">
    <CampaignManagement />
  </div>
</template>

<script>
import CampaignManagement from '../components/campaign/CampaignManagement.vue';

export default {
  name: 'Campaign',
  components: {
    CampaignManagement,
  },
};
</script>
