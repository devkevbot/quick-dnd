<!-- Registration.vue
This component is the view shown when a prospective user is on the
registration page.
-->

<template>
  <div class="registration">
    <RegistrationForm />
  </div>
</template>

<script>
import RegistrationForm from '../components/registration/RegistrationForm.vue';

export default {
  components: {
    RegistrationForm,
  },
};
</script>
