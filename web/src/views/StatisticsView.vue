<template>
  <div class="statistics">
    <CardStatistic
      v-for="stat in stats"
      :key="stat.title"
      :title="stat.title"
      :value="stat.value"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import CardStatistic from "@/components/CardStatistic.vue";
import Component from "vue-class-component";

@Component({
  name: "Statistics",
  components: {
    CardStatistic
  }
})
export default class StatisticsView extends Vue {
  stats = [
    {
      title: "Nodes",
      value: "12"
    },
    {
      title: "Instances",
      value: "342"
    },
    {
      title: "CPU Cores",
      value: "77"
    },
    {
      title: "Disk Space",
      value: "5 TB"
    },
    {
      title: "Memory",
      value: "256 GB"
    }
  ];
}
</script>

<style scoped>
.statistics {
  padding: 2rem;
  display: grid;
  grid-gap: 1rem;
  grid-template-columns: repeat(6, 1fr);
  grid-template-rows: auto;
}
</style>
