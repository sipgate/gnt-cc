<template>
  <div class="checkbox">
    <label>
      <Icon v-show="!value" icon="square" />
      <Icon v-show="value" icon="check-square" />

      <input type="checkbox" :checked="value" :required="required" @change="handleChange" />
      <span class="label">{{ label }}</span>
    </label>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Prop, Component } from "vue-property-decorator";
import Icon from "@/components/Icon.vue";

@Component({
  name: "Checkbox",
  components: { Icon }
})
export default class Checkbox extends Vue {
  @Prop()
  label!: string;

  @Prop({ default: false })
  required!: boolean;

  @Prop({ required: true })
  value!: boolean;

  handleChange(event: Event): void {
    const target = event.target as HTMLInputElement;
    this.$emit("input", target.checked);
  }
}
</script>

<style lang="scss">
.checkbox {
  position: relative;

  label {
    display: flex;
    color: gray;
    align-items: center;
    cursor: pointer;

    .label {
      display: block;
      height: 100%;
      margin: 0.2rem 0 0 0.5rem;
    }

    .icon {
      width: 1.5rem;
      height: 1.5rem;
    }

    input {
      visibility: hidden;
      display: none;
    }
  }
}
</style>