<template>
  <div class="campaigns">
    <v-container>
      <v-row wrap>
        <v-col
          cols="12"
          lg="6"
          v-for="campaign in data.campaigns"
          :key="campaign.id"
        >
          <v-card
            class="d-flex flex-column"
            elevation="2"
            outlined
            tile
            height="100%"
          >
            <v-card-subtitle class="primary white--text mb-2">
              <b>Dungeon master: {{ campaign.dungeon_master }}</b>
            </v-card-subtitle>

            <v-card-text class="display-1 primary--text">
              <div class="truncate">
                {{ campaign.name }}
              </div>
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text class="text--primary">
              <b>Location:</b> {{ campaign.current_location }}
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text class="text--primary">
              <b>State:</b> {{ campaign.state }}
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
  name: 'Campaigns',
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
