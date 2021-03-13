<!-- Login.vue
This component is the view shown when a user is on the login page.
-->

<template>
  <div class="login">
    <LoginForm />
  </div>
</template>

<script>
import LoginForm from '../components/login/LoginForm.vue';

export default {
  components: {
    LoginForm,
  },
};
</script>
