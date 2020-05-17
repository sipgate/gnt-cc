<template>
  <button
    class="custom-button"
    @click="onClick"
    :class="{ 'has-label': !!label, 'has-icon': !!icon, [type]: !!type }"
  >
    <b-icon :icon="icon" size="is-small" v-if="icon" />
    <span class="custom-label" v-if="label">{{ label }}</span>
  </button>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";

@Component({
  name: "Button"
})
export default class Button extends Vue {
  @Prop({ default: false })
  readonly disabled!: boolean;

  @Prop()
  readonly label!: string;

  @Prop()
  readonly icon!: string;

  @Prop()
  readonly type!: string;

  onClick() {
    if (!this.disabled) {
      this.$emit("click");
    }
  }
}
</script>

<style lang="scss" scoped>
.custom-button {
  border: 0;
  transition: transform 0.2s, opacity 0.2s;
  cursor: pointer;
  opacity: 0.8;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 0.75rem;
  height: 48px;
  background: #76a7fa;
  color: #fff;
  border-radius: 2px;

  &:hover {
    opacity: 1;
  }

  &.has-label.has-icon {
    .icon {
      margin-right: 0.75rem;
    }
  }

  &.accept {
    background: #42b983;
    color: #fff;
  }

  &.reset {
    background: #eaedf3;
    color: #555;
  }

  &.danger {
    background: red;
    color: white;
  }
}
</style>
