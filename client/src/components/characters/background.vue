// Compnent for inputing a character's background

<template>
<div class="background">

  <v-btn>
    <h3> Save Changes</h3>
  </v-btn>

    <v-container class="characterBackground">
      <v-row>
          <v-col cols="12" lg="12">
            <v-card flat class="text-xs-center" elevation="8" outlined tile>
              <v-card-title>
                {{statUniqueElements[2].attribute}}
              </v-card-title>
              <v-card-actions>
                <v-textarea
                  v-model="statUniqueElements[2].value"
                  filled
                  name="input-7-4"
                  label="Filled textarea"
                  value="statUniqueElements[2].value will go here"
                ></v-textarea>
              </v-card-actions>
            </v-card>
          </v-col>
      </v-row>
  </v-container>

</div>
</template>

<script>
export default {
  props: {
    characterData: {
      type: Array,
      required: false,
    },
  },
  name: 'background',
  computed: {
    statUniqueElements() {
      return [
        { attribute: 'Location', value: 'TEMP' },
        { attribute: 'Level', value: this.level },
        { attribute: 'Character Background', value: this.characterData.background },
        { attribute: 'XP', value: this.characterData.xp_points },
      ];
    },
  },
};
</script>
