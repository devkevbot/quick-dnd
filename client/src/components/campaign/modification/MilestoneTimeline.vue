<template>
  <v-container fluid>
    <v-timeline>
      <v-timeline-item
        v-for="(milestone, index) in milestones"
        :key="index"
        large
      >
        <v-card dark color="primary">
          <v-card-title class="subtitle-1">
            Milestone {{ index + 1 }}
          </v-card-title>

          <v-card-text class="white text--primary pt-2">
            {{ milestone }}
          </v-card-text>
        </v-card>
      </v-timeline-item>
    </v-timeline>
  </v-container>
</template>

<script>
export default {
  name: 'MilestoneTimeline',
  props: {
    milestones: {
      type: Array,
      required: false,
    },
  },
};
</script>
