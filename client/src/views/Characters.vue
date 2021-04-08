<template>
  <div class="characters">
    <CharacterManagement />
  </div>
</template>

<script>
import CharacterManagement from '../components/characters/CharacterManagement.vue';

export default {
  name: 'Characters',
  components: {
    CharacterManagement,
  },
};
</script>
