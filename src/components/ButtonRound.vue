<template>
  <button class="button-round" :class="size" @click.prevent="onClick" :disabled="disabled">
    <b-icon :icon="icon" :size="size" />
  </button>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";

@Component({
  name: "ButtonRound"
})
export default class ButtonRound extends Vue {
  @Prop(String)
  readonly icon!: string;
  @Prop({ default: "is-small" })
  readonly size!: string;
  @Prop({ default: false })
  readonly disabled!: boolean;

  onClick() {
    if (!this.disabled) {
      this.$emit("click");
    }
  }
}
</script>

<style scoped lang="scss">
.button-round {
  border: 1px solid #555;
  background: none;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #555;

  &.is-small {
    width: 24px;
    height: 24px;
  }

  &.is-medium {
    width: 32px;
    height: 32px;
  }

  &.is-large {
    width: 48px;
    height: 48px;
  }
}
</style>
