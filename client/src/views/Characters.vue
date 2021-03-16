<template>
<div class="characters">
    <v-container class="characterStats">
        <v-row wrap>
            <v-col cols="12" md="6" lg="2" v-for="row in stat_sliders" :key="row.attribute">
              <v-card flat class="text-xs-center" elevation="8" outlined tile>
                <v-card-title>
                  {{row.attribute}}
                </v-card-title>
                <v-card-subtitle>
                  <!-- <div class="subtitle-1">Subtext</div> -->
                </v-card-subtitle>
                <v-card-text>
                  <!--Text -->
                </v-card-text>
                <v-card-actions>
                  <v-slider v-model="blue" :max="100" class="align-center">
                    <template v-slot:append>
                      <v-text-field
                        v-model="blue"
                        class="mt-0 pt-0"
                        type="number"
                        style="width: 60px"
                      >
                    </v-text-field>
                    </template>
                  </v-slider>
                </v-card-actions>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
export default {
  data() {
    return {
      stat_sliders: [
        { attribute: 'Strength' },
        { attribute: 'Contribution' },
        { attribute: 'Wisdom' },
        { attribute: 'HP_Points' },
        { attribute: 'Dexterity' },
        { attribute: 'Ability_Points' },
        { attribute: 'Charisma' },
        { attribute: 'Intellect' },
      ],
    };
  },
};
</script>
