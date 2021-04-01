<template>
  <div class="spells">
    <v-container>
      <v-row wrap>
        <v-col
          cols="12"
          md="6"
          lg="2"
          v-for="spell in data"
          :key="spell.spell_name"
        >
          <v-card
            class="d-flex flex-column"
            elevation="2"
            outlined
            tile
            height="100%"
          >
            <v-card-subtitle class="primary white--text mb-2">
              <b>{{ spell.school }} - Level {{ spell.level }}</b>
            </v-card-subtitle>

            <v-card-text class="display-1 primary--text">
              <div class="truncate">
                {{ spell.spell_name }}
              </div>
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text>
              Range: {{ spell.range }}ft <br />
              Duration: {{ spell.duration }}s <br />
              Casting Time: {{ spell.casting_time }}s
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text class="text--primary">
              {{ spell.description }}
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
      type: Array,
      required: false,
    },
  },
  name: 'Spells',
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
