<template>
  <div class="cluster-selector">
    <Dropdown :label="currentClusterName" icon="server">
      <router-link
        class="cluster"
        :class="{ active: cluster.name === currentClusterName }"
        :to="{ params: { cluster: cluster.name } }"
        v-for="cluster in clusters"
        :key="cluster.name"
      >
        <span class="name">{{ cluster.name }}</span>
        <span class="host">{{ cluster.hostname }}</span>
        <span class="description">{{ cluster.description }}</span>
      </router-link>
    </Dropdown>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { State } from "vuex-class";
import { StoreState } from "@/store";
import Params from "@/data/enum/Params";
import Dropdown from "@/components/Dropdown.vue";
import GntCluster from "../model/GntCluster";

@Component({
  name: "ClusterSelector",
  components: { Dropdown }
})
export default class ClusterSelector extends Vue {
  @State((state: StoreState) => state.clusters) clusters!: GntCluster[];

  get currentClusterName(): string {
    return this.$route.params[Params.Cluster] || "";
  }
}
</script>

<style scoped lang="scss">
$colorPrimary: #0f4c5c;

.cluster-selector {
  position: relative;

  .cluster {
    height: 100%;
    display: flex;
    flex-direction: column;
    text-decoration: none;
    padding: 0.5rem 1rem;
    min-width: 260px;
    color: #222;
    opacity: 0.8;
    transition: opacity 0.2s;
    border-left: 6px solid transparent;

    &.active {
      border-color: #e2e2e2;
      opacity: 0.5;
    }

    &:hover {
      border-color: $colorPrimary;
      opacity: 1;

      .name {
        color: $colorPrimary;
      }
    }

    .name {
      font-weight: bold;
    }

    .host,
    .description {
      font-size: 0.8rem;
    }

    .host {
      font-family: monospace;
    }

    .description {
      white-space: initial;
    }
  }
}
</style>
