<template>
  <div class="characters">
    <SpellManagement />
  </div>
</template>

<script>
import SpellManagement from '../components/spells/SpellManagement.vue';

export default {
  name: 'Spells',
  components: {
    SpellManagement,
  },
};
</script>
