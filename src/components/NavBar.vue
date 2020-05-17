<template>
  <div id="nav">
    <b-navbar>
      <template slot="brand">
        <b-navbar-item tag="router-link" to="/">
          <img class="brand-logo" src="../assets/ganeti_logo.svg" />
        </b-navbar-item>
      </template>
      <template slot="start">
        <b-navbar-item tag="router-link" :to="links.statistics" class="needs-exact-match">
          Statistics
        </b-navbar-item>
        <b-navbar-item tag="router-link" :to="links.instances">
          Instances
        </b-navbar-item>
        <b-navbar-item tag="router-link" :to="links.jobs">
          Jobs
        </b-navbar-item>
        <b-navbar-item tag="router-link" :to="links.nodes">
          Nodes
        </b-navbar-item>
      </template>
      <template slot="end">
        <b-navbar-item tag="div">
          <ClusterSelector />
        </b-navbar-item>
      </template>
    </b-navbar>
  </div>
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

<style scoped>
#nav {
  padding: 30px;
  background: #ffffff;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active.needs-exact-match {
  color: #42b983;
}

#nav a.router-link-active:not(.needs-exact-match) {
  color: #42b983;
}

.brand-logo {
  max-height: 3rem;
}
</style>
