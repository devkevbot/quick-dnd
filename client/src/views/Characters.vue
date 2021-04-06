<template>
  <div class="characters">
    <CreationStepper />
  </div>
</template>

<script>
import CreationStepper from '../components/characters/CreationStepper.vue';

export default {
  name: 'Characters',
  components: {
    CreationStepper,
  },
};
</script>
