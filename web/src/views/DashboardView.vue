<template>
  <div>
    <NavBar />
    <router-view />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import NavBar from "@/components/NavBar.vue";
import { Actions } from "@/store";
import Params from "@/data/enum/Params";
import GntCluster from "../model/GntCluster";

@Component({
  name: "DashboardView",
  components: { NavBar }
})
export default class DashboardView extends Vue {
  async created() {
    const clusters: GntCluster[] = await this.$store.dispatch(Actions.LoadClusters);

    if (
      typeof this.$route.params[Params.Cluster] === "undefined" &&
      typeof clusters !== "undefined" &&
      clusters.length > 0
    ) {
      await this.$router.replace({
        params: {
          cluster: clusters[0].name
        }
      });
    }
  }
}
</script>

<style scoped></style>
