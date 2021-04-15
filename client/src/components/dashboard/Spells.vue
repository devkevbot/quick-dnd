<template>
  <div class="spells">
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
              <b>Spell stats</b>
            </v-card-subtitle>

            <!-- Display spell stats in a table to showcase the
            aggregration query requirement. -->
            <v-simple-table>
              <thead>
                <tr>
                  <th class="primary--text display-1">School</th>
                  <th class="primary--text display-1">Spell count</th>
                </tr>
              </thead>

              <tbody>
                <tr v-if="data.spells.length === 0">
                  <td>This character has no spells!</td>
                </tr>

                <tr
                  v-else
                  v-for="s in data.spellCountPerSchool"
                  :key="s.school"
                >
                  <td class="subtitle-1">
                    {{ s.school }}
                  </td>
                  <td class="subtitle-1">
                    {{ s.count }}
                  </td>
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
          v-for="spell in data.spells"
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
      type: Object,
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
