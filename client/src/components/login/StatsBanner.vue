<template>
  <div class="stats-banner">
    <v-card>
      <v-card-title class="primary display-1 justify-center white--text">
        {{ statSentence }}
      </v-card-title>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'StatsBanner',
  data() {
    return {
      stats: null,
    };
  },
  async created() {
    const requestURI = 'stat';
    const method = 'GET';

    await this.$http({
      url: requestURI,
      data: null,
      method,
    })
      .then((resp) => {
        this.stats = resp.data.data;
      })
      .catch(() => {
        /* TODO: Add error handling. */
      });
  },
  computed: {
    statSentence() {
      if (this.stats === null) return 'Sign up today!';
      return `Sign up today and join over ${this.stats.num_player_account} registered players and ${this.stats.num_character_created} created characters.`;
    },
  },
};
</script>
