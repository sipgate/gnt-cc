<template>
  <div class="dropdown" :class="{ expanded: isExpanded, 'has-icon': !!icon }" @click.stop="toggle">
    <div class="current">
      <span class="label">{{ label }}</span>
      <b-icon v-if="icon" class="icon" :icon="icon" size="is-small" />
    </div>
    <div class="options">
      <span class="triangle" />

      <div class="options-items">
        <slot />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";

@Component({
  name: "Dropdown"
})
export default class Dropdown extends Vue {
  @Prop({ required: true })
  readonly label!: string;

  @Prop()
  readonly icon!: string;

  isExpanded = false;

  mounted() {
    window.addEventListener("click", this.handleOutsideClick);
  }

  beforeDestroy() {
    window.removeEventListener("click", this.handleOutsideClick);
  }

  selectOption(value: string) {
    this.$emit("change", value);

    this.isExpanded = false;
  }

  toggle() {
    this.isExpanded = !this.isExpanded;
  }

  handleOutsideClick() {
    this.isExpanded = false;
  }
}
</script>

<style lang="scss" scoped>
$colorPrimary: #0f4c5c;
$spacingOuter: 1rem;
$spacingInner: 0.5rem;
$transitionDuration: 0.2s;
$width: 200px;
$borderRadius: 4px;

.dropdown {
  position: relative;
  width: $width;
  height: 48px;
  cursor: pointer;

  &.expanded {
    .current {
      .caret {
        transform: rotate(180deg);
      }
    }

    .options {
      visibility: visible;
      opacity: 1;
      transition: opacity $transitionDuration, visibility 0s;
    }
  }

  &.has-icon {
    .current .label {
      margin: 0 $spacingOuter 0 0;
    }
  }

  .current {
    display: flex;
    align-items: center;
    z-index: 1;
    padding: $spacingOuter;
    background: $colorPrimary;
    border-radius: $borderRadius;
    color: #fff;
    width: 100%;

    .icon {
      color: inherit;
    }

    .label {
      width: 100%;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
      color: inherit;
      margin: 0;
    }
  }

  .options {
    position: absolute;
    right: 0;
    visibility: hidden;
    top: calc(100% + 24px);
    min-width: $width;
    border-top: 0;
    background: #fff;
    box-shadow: 0 0 1rem rgba(#222, 0.4);
    opacity: 0;
    border-radius: $borderRadius;
    transition: opacity $transitionDuration, visibility 0s $transitionDuration;

    .triangle {
      width: $spacingOuter;
      height: $spacingOuter;
      background: #fff;
      position: absolute;
      top: $spacingOuter * -0.5;
      right: $spacingOuter;
      transform: rotate(45deg);
    }

    .options-items {
      overflow: hidden;
      border-radius: $borderRadius;
    }
  }
}
</style>
