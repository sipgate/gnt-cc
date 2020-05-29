<template>
  <div class="card" :class="{ 'no-horizontal-padding': noHorizontalPadding }">
    <div class="card-title">
      <slot name="title-left" />
      <span class="card-label">{{ title }}</span>
      <slot name="title-right" />
    </div>
    <div class="card-body">
      <slot />
    </div>
    <div class="card-subtitle">
      <span class="card-label">{{ subtitle }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";

@Component({
  name: "Card"
})
export default class Card extends Vue {
  @Prop(String)
  readonly title!: string;

  @Prop(String)
  readonly subtitle!: string;

  @Prop({ default: false })
  readonly noHorizontalPadding!: boolean;
}
</script>

<style scoped lang="scss">
div.card {
  border-radius: 4px;
  padding: 0 30px;
  flex: 1 0 auto;
  box-sizing: border-box;
  border: 1px solid #eaedf3;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  &.no-horizontal-padding {
    padding: 0;
  }

  .card-title,
  .card-subtitle {
    height: 48px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.75rem;

    .card-label {
      display: block;
      width: 100%;
      text-align: center;
      font-weight: 500;
      color: #9ea0a5;
      text-transform: uppercase;
    }
  }

  div.card-title {
    justify-content: space-between;
  }
  div.card-body {
    margin: 12px 0;
  }
  div.card-subtitle {
    font-size: 0.75rem;
  }
}
</style>
