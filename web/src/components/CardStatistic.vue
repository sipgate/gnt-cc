<template>
  <Card class="card-statistic" :title="title">
    <div class="hero">
      <span>{{ value }}</span>
    </div>
  </Card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";
import Card from "@/components/Card.vue";

@Component({
  name: "CardStatistic",
  components: { Card }
})
export default class CardStatistic extends Vue {
  @Prop({ required: true })
  value!: string;

  @Prop({ required: true })
  title!: string;
}
</script>

<style lang="scss" scoped>
.card-statistic {
  .hero {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 4rem;
    color: #a9a9a9;
    font-weight: 800;
  }
}
</style>
