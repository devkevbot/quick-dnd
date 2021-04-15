<template>
  <div class="items">
    <v-container>
      <v-row wrap>
        <v-col cols="6">
          <v-card
            class="d-flex flex-column"
            elevation="2"
            outlined
            tile
            height="100%"
          >
            <v-card-subtitle class="primary white--text mb-2">
              <b>Item stats</b>
            </v-card-subtitle>

            <!-- Display item stats in a table to showcase the
            aggregration query requirement. -->
            <v-simple-table>
              <thead>
                <tr>
                  <th class="primary--text display-1">Stat</th>
                  <th class="primary--text display-1">Value</th>
                </tr>
              </thead>

              <tbody>
                <tr v-if="data.items.length === 0">
                  <td>This character has no items!</td>
                </tr>

                <tr v-else>
                  <td class="subtitle-1">Total item weight</td>
                  <td class="subtitle-1">{{ data.totalItemWeight }} lbs</td>
                </tr>
              </tbody>
            </v-simple-table>
          </v-card>
        </v-col>
      </v-row>

      <v-row wrap>
        <v-col
          cols="12"
          md="6"
          lg="3"
          v-for="item in data.items"
          :key="item.item_name"
        >
          <v-card
            class="d-flex flex-column"
            elevation="2"
            outlined
            tile
            height="100%"
          >
            <v-card-subtitle class="primary white--text mb-2">
              <b>{{ item.type }} - {{ item.rarity }}</b>
            </v-card-subtitle>

            <v-card-text class="display-1 primary--text">
              <div class="truncate">
                {{ item.item_name }}
              </div>
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text>
              Weight: {{ item.weight }}lbs <br />
              Value: {{ item.gold_value }} gold <br />
              Quantity: {{ item.quantity }}
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text class="text--primary">
              {{ item.description }}
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  props: {
    data: {
      type: Object,
      required: false,
    },
  },
  name: 'Items',
};
</script>

<style scoped>
.truncate {
  white-space: nowrap;
  word-break: normal;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
