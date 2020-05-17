<template>
  <div class="instances">
    <b-table :data="instances" default-sort-direction="asc" sort-icon-size="small">
      <template slot-scope="data">
        <b-table-column label="Name" field="name" searchable sortable>
          <router-link :to="createInstanceLink(data.row.name)">
            {{ data.row.name }}
          </router-link>
        </b-table-column>
        <b-table-column label="Primary Node" field="primaryNode" searchable sortable>
          {{ data.row.primaryNode }}
        </b-table-column>
        <b-table-column label="Secondary Node(s)" field="secondaryNodes" searchable sortable>
          {{ data.row.secondaryNodes[0] }}
          <b-tag rounded v-if="data.row.secondaryNodes.length > 1">
            +{{ data.row.secondaryNodes.length - 1 }}
          </b-tag>
        </b-table-column>
        <b-table-column label="vCPUs" field="cpuCount" sortable>
          {{ data.row.cpuCount }}
        </b-table-column>
        <b-table-column label="Memory" field="memoryTotal" sortable>
          {{ data.row.memoryTotal }} MB
        </b-table-column>
      </template>
    </b-table>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Actions, StoreState } from "@/store";
import { State } from "vuex-class";
import GntInstance from "@/model/GntInstance";
import { Watch } from "vue-property-decorator";
import { Location } from "vue-router";
import PageNames from "@/data/enum/PageNames";
import Params from "@/data/enum/Params";

@Component({
  name: "InstancesListView.vue"
})
export default class InstancesListView extends Vue {
  @State((state: StoreState) => state.instances) allInstances!: Record<string, GntInstance[]>;

  @Watch("currentCluster")
  onCurrentClusterChanged() {
    this.loadInstances();
  }

  async created() {
    if (this.currentCluster.length > 0) {
      this.loadInstances();
    }
  }

  async loadInstances() {
    await this.$store.dispatch(Actions.LoadInstances, {
      cluster: this.currentCluster
    });
  }

  get instances(): GntInstance[] {
    return this.allInstances[this.currentCluster];
  }

  get currentCluster(): string {
    return this.$route.params[Params.Cluster];
  }

  createInstanceLink(name: string): Location {
    return {
      name: PageNames.InstancesDetail,
      params: {
        [Params.Cluster]: this.currentCluster,
        [Params.InstanceName]: name
      }
    };
  }
}
</script>
