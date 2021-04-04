<!-- Login.vue
This component is the view shown when a user is on the login page.
-->

<template>
  <div class="login">
    <StatsBanner />
    <LoginForm />
  </div>
</template>

<script>
import StatsBanner from '../components/login/StatsBanner.vue';
import LoginForm from '../components/login/LoginForm.vue';

export default {
  components: {
    StatsBanner,
    LoginForm,
  },
};
</script>
