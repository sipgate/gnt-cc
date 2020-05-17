<template>
  <button class="button-round" @click.prevent="onClick" :disabled="disabled">
    <Icon :icon="icon" />
  </button>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";
import Icon from "@/components/Icon.vue";

@Component({
  name: "ButtonRound",
  components: { Icon }
})
export default class ButtonRound extends Vue {
  @Prop(String)
  readonly icon!: string;
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
  width: 24px;
  height: 24px;
}
</style>
