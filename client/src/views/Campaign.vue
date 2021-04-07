<!-- Campaign.vue
This component is the view shown when a user is on the Campaign creation page.
-->

<template>
  <div class="campaign">
    <CampaignForm />
  </div>
</template>

<script>
import CampaignForm from '../components/campaign/CampaignForm.vue';

export default {
  name: 'Campaign',
  components: {
    CampaignForm,
  },
};
</script>
