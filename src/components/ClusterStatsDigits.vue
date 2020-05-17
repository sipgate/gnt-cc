<template>
  <div class="about">
    <div class="tile is-ancestor">
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">Nodes</p>
          <p class="title information">12</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">Instances</p>
          <p class="title information">184</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">CPU Cores</p>
          <p class="title information">584</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">Disk Space</p>
          <p class="title information">24 TB</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">Memory</p>
          <p class="title information">2 TB</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">Inst. per Node</p>
          <p class="title information">10 / 15 / 20</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

@Component({
  name: "ClusterStatsDigits"
})
export default class ClusterStatsDigits extends Vue {}
</script>

<style scoped>
.tile {
  margin: 5px;
}

.information {
  font-weight: normal;
}
</style>
