<template>
  <div class="instance-details">
    <section class="instance-header">
      <span class="instance-name">{{ instanceName }}</span>
      <div class="instance-actions">
        <Button label="Migrate" />
        <Button label="Failover" />
        <Button label="Shutdown" type="danger" />
        <Button label="Kill" icon="skull-crossbones" type="danger" />
      </div>
    </section>

    <section class="cards" v-if="instance">
      <CardNumber
        title="memory"
        unit="MB"
        :startValue="instance.memoryTotal"
        :minimumValue="512"
        :maximumValue="8192"
        :step="256"
      />
      <CardNumber
        title="vcpu"
        unit="cores"
        :startValue="instance.cpuCount"
        :minimumValue="1"
        :maximumValue="16"
      />
      <CardNumber
        v-for="disk in instance.disks"
        :key="disk.uid"
        :title="disk.name"
        unit="GB"
        :startValue="disk.size | mbToGb"
        :minimumValue="disk.size | mbToGb"
        :maximumValue="10000"
        :step="5"
      />

      <CardNodes :primaryNode="instance.primaryNode" :secondaryNodes="instance.secondaryNodes" />
    </section>
    <div>
      <Button label="Reset all" type="reset" />
      <Button label="Apply all" type="accept" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Actions, StoreState } from "@/store";
import { State } from "vuex-class";
import GntInstance from "@/model/GntInstance";
import { Watch } from "vue-property-decorator";
import Params from "@/data/enum/Params";
import Button from "@/components/Button.vue";
import Card from "@/components/Card.vue";
import CardNumber from "@/components/CardNumber.vue";
import CardNodes from "@/components/CardNodes.vue";

@Component({
  name: "InstanceDetailView",
  components: { Button, Card, CardNumber, CardNodes }
})
export default class InstanceDetailView extends Vue {
  @State((state: StoreState) => state.instances) allInstances!: Record<string, GntInstance[]>;

  @Watch("instanceName")
  onInstanceNameChanged() {
    this.loadInstance();
  }

  async created() {
    this.loadInstance();
  }

  async loadInstance() {
    await this.$store.dispatch(Actions.LoadInstance, {
      cluster: this.currentCluster,
      instance: this.instanceName
    });
  }

  get instance(): GntInstance | undefined {
    if (!this.allInstances[this.currentCluster]) {
      return undefined;
    }

    return this.allInstances[this.currentCluster].find(
      instance => instance.name === this.instanceName
    );
  }

  get instanceName(): string {
    return this.$route.params[Params.InstanceName];
  }

  get currentCluster(): string {
    return this.$route.params[Params.Cluster];
  }
}
</script>

<style scoped lang="scss">
$spacingOuter: 3rem;
$spacingInner: 1.5rem;

.instance-details {
  section.instance-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: $spacingOuter;
    background: #000;
    color: #fff;

    .instance-name {
      font-size: 2rem;
    }

    .instance-actions {
      display: flex;

      .custom-button {
        margin-left: $spacingInner;
      }
    }
  }

  section.cards {
    display: grid;
    padding: $spacingOuter;
    row-gap: $spacingInner;
    column-gap: $spacingInner;
    grid-template-columns: repeat(4, 1fr);
    grid-template-rows: repeat(2, auto);
    grid-template-areas:
      "stats stats stats stats"
      "quickstats none none none";
  }
}
</style>
