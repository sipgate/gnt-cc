<template>
  <span class="icon">
    <FontAwesomeIcon class="svg-icon" :icon="icon" />
  </span>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

@Component({
  name: "Icon",
  components: { FontAwesomeIcon }
})
export default class Icon extends Vue {
  @Prop({ required: true })
  readonly icon!: string;
}
</script>

<style lang="scss" scoped>
.icon {
  width: 1rem;
  height: 1rem;
  flex-grow: 0;
  flex-shrink: 0;

  .svg-icon {
    width: 0.875em;
  }
}
</style>
