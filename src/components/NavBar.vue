<template>
  <nav class="navbar">
    <div class="begin">
      <router-link class="logo-container" to="/">
        <img class="logo" src="../assets/ganeti_logo.svg" />
      </router-link>
      <div class="items">
        <router-link class="item needs-exact-match" :to="links.statistics">
          Statistics
        </router-link>
        <router-link class="item" :to="links.instances">
          Instances
        </router-link>
        <router-link class="item" :to="links.jobs">
          Jobs
        </router-link>
        <router-link class="item" :to="links.nodes">
          Nodes
        </router-link>
      </div>
    </div>
    <div class="end">
      <ClusterSelector />
    </div>
  </nav>
</template>

<script lang="ts">
import PageNames from "@/data/enum/PageNames";
import Vue from "vue";
import Component from "vue-class-component";
import ClusterSelector from "@/components/ClusterSelector.vue";
import Params from "@/data/enum/Params";

@Component({
  name: "NavBar",
  components: { ClusterSelector }
})
export default class NavBar extends Vue {
  get currentCluster(): string {
    return this.$route.params[Params.Cluster];
  }

  get links() {
    return {
      statistics: {
        name: PageNames.Statistics,
        params: {
          cluster: this.currentCluster
        }
      },
      instances: {
        name: PageNames.InstancesList,
        params: {
          cluster: this.currentCluster
        }
      },
      jobs: {
        name: PageNames.Jobs,
        params: {
          cluster: this.currentCluster
        }
      },
      nodes: {
        name: PageNames.Nodes,
        params: {
          cluster: this.currentCluster
        }
      }
    };
  }
}
</script>

<style scoped lang="scss">
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 2rem;
  background: #fff;

  .logo-container {
    display: flex;
    justify-content: center;
    align-items: center;

    .logo {
      width: 3rem;
      height: 3rem;
    }
  }

  .begin {
    display: flex;
    align-items: center;

    .items {
      display: flex;
      align-items: center;
      margin-left: 2rem;

      .item {
        display: block;
        font-weight: bold;
        color: #2c3e50;
        padding: 1rem;

        &.router-link-exact-active.needs-exact-match {
          color: #42b983;
        }

        &.router-link-active:not(.needs-exact-match) {
          color: #42b983;
        }
      }
    }
  }
}
</style>
