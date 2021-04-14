<template>
  <div class="characters">
    <ItemManagement />
  </div>
</template>

<script>
import ItemManagement from '../components/items/ItemManagement.vue';

export default {
  name: 'Items',
  components: {
    ItemManagement,
  },
};
</script>
