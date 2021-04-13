<!-- AddCharacterDialog.vue
This component is the table that shows all the character that
are currently added to a campaign.
-->

<template>
  <div class="add-character-table">
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">ID</th>
            <th class="text-left">Name</th>
            <th class="text-left">Owner</th>
            <th class="text-left">Remove</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="items.length === 0">
            <td>No data</td>
          </tr>

          <tr v-else v-for="item in items" :key="item.id">
            <td>{{ item.id }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.owner }}</td>
            <td>
              <v-icon class="pl-2" @click="removeItem(item.id)">
                mdi-delete
              </v-icon>
            </td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  name: 'AddCharacterTable',
  data() {
    return {
      items: [],
    };
  },
  methods: {
    addItem(iId, iName, iOwner) {
      this.items.push({ id: iId, name: iName, owner: iOwner });
    },
    removeItem(id) {
      this.items = this.items.filter((x) => x.id !== id);
    },
  },
  computed: {
    getItems() {
      return this.items;
    },
    getIDs() {
      return this.items.map(({ id }) => parseInt(id, 10));
    },
  },
};
</script>
