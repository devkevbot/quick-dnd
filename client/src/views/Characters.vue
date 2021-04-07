<template>
  <div class="characters">
    <EditExisting />
  </div>
</template>

<script>
import EditExisting from '../components/characters/EditExisting.vue';

export default {
  name: 'Characters',
  components: {
    EditExisting,
  },
};
</script>
