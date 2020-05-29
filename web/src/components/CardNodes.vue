<template>
  <Card class="card-nodes" title="Nodes" :subtitle="subtitle">
    <div class="node">
      {{ primaryNode }}
      <b-tag rounded>primary</b-tag>
    </div>
    <div class="node" v-for="node in secondaryNodes" :key="node">
      <span>{{ node }}</span>
    </div>
  </Card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";
import Card from "@/components/Card.vue";
import GntNode from "../model/GntNode";

@Component<CardNodes>({
  name: "CardNodes",
  components: { Card }
})
export default class CardNodes extends Vue {
  $refs!: {
    input: HTMLInputElement;
  };

  @Prop()
  primaryNode!: GntNode;

  @Prop()
  secondaryNodes!: GntNode[];

  get subtitle() {
    return `Total: ${this.secondaryNodes.length + 1}`;
  }
}
</script>

<style lang="scss" scoped>
.card-nodes {
  font-size: 1rem;
  text-align: left;
  margin: 0.5rem 0;

  .node {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    line-height: 2.5rem;
    border-bottom: 1px solid rgba(#000, 0.1);
  }
}
</style>
