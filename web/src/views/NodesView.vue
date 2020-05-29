<template>
  <div class="nodes">
    <b-table :data="nodes" :columns="columns" default-sort-direction="asc" sort-icon-size="small" />
  </div>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { Actions, StoreState } from "@/store";
import { State } from "vuex-class";
import { Watch } from "vue-property-decorator";
import GntNode from "@/model/GntNode";
import Params from "@/data/enum/Params";

@Component<NodesView>({
  name: "NodesView"
})
export default class NodesView extends Vue {
  @State((state: StoreState) => state.nodes) allNodes!: Record<string, GntNode[]>;

  @Watch("currentCluster")
  onCurrentClusterChanged() {
    this.loadNodes();
  }

  async created() {
    if (this.currentCluster.length > 0) {
      this.loadNodes();
    }
  }

  async loadNodes() {
    await this.$store.dispatch(Actions.LoadNodes, {
      cluster: this.currentCluster
    });
  }

  get currentCluster(): string {
    return this.$route.params[Params.Cluster];
  }

  get nodes(): GntNode[] {
    return this.allNodes[this.currentCluster];
  }

  get columns() {
    return [
      {
        field: "name",
        label: "Name",
        searchable: true,
        sortable: true
      },
      {
        field: "mtotal",
        label: "Total Memory",
        searchable: false,
        sortable: true
      },
      {
        field: "mfree",
        label: "Free Memory",
        searchable: false,
        sortable: true
      }
    ];
  }
}
</script>

<style lang="scss" scoped></style>
