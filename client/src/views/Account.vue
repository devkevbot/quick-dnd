<template>
  <div class="account">
    <ProfileCard />
  </div>
</template>

<script>
import ProfileCard from '../components/account/ProfileCard.vue';

export default {
  name: 'Account',
  components: {
    ProfileCard,
  },
};
</script>
