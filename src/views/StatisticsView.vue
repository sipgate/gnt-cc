<template>
  <div class="statistics">
    <ClusterStatsDigits />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ClusterStatsDigits from "@/components/ClusterStatsDigits.vue";
import Component from "vue-class-component";

@Component({
  name: "Statistics",
  components: {
    ClusterStatsDigits
  }
})
export default class StatisticsView extends Vue {}
</script>

<style scoped>
.statistics {
}
</style>
