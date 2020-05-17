<template>
  <Card class="card-number" :title="title" :subtitle="unit" :noHorizontalPadding="true">
    <Button
      slot="title-left"
      class="interaction"
      :class="{ hidden: !isDirty && isValid }"
      @click="onReset"
      type="reset"
      icon="undo"
    />

    <Button
      slot="title-right"
      class="interaction"
      :class="{ hidden: !isDirty || !isValid }"
      @click="onAccept"
      type="accept"
      icon="check"
    />

    <div class="content">
      <div class="button-container">
        <ButtonRound
          v-show="!decreaseDisabled"
          class="button decrease"
          icon="minus"
          @click="onDecrease"
        />
      </div>

      <span class="value" v-if="!isEditable">{{ currentValue }}</span>

      <form class="number-form" @submit.prevent="onAccept" v-if="isEditable">
        <input
          @click="onInputClick"
          ref="input"
          class="number-input"
          :class="{ invalid: !isValid }"
          type="text"
          maxlength="6"
          v-model="inputValue"
        />
      </form>

      <div class="button-container">
        <ButtonRound
          v-show="!increaseDisabled"
          class="button increase"
          icon="plus"
          @click="onIncrease"
        />
      </div>
    </div>
  </Card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";
import Button from "@/components/Button.vue";
import ButtonRound from "@/components/ButtonRound.vue";
import Card from "@/components/Card.vue";

@Component<CardNumber>({
  name: "CardNumber",
  components: { Button, ButtonRound, Card }
})
export default class CardNumber extends Vue {
  $refs!: {
    input: HTMLInputElement;
  };

  @Prop(String)
  title!: string;

  @Prop(String)
  unit!: string;

  @Prop(Number)
  startValue!: number;

  @Prop(Number)
  minimumValue!: number;

  @Prop(Number)
  maximumValue!: number;

  @Prop({ default: true })
  isEditable!: boolean;

  @Prop({ default: 1 })
  step!: number;

  @Watch("inputValue")
  onInputValueChange(value: string) {
    if (this.isValid) {
      this.currentValue = Number(value);
    }
  }

  currentValue: number = this.startValue;

  inputValue = this.startValue.toString();

  get isDirty() {
    return this.currentValue !== this.startValue;
  }

  get decreaseDisabled() {
    return this.currentValue === this.minimumValue;
  }

  get increaseDisabled() {
    return this.currentValue === this.maximumValue;
  }

  get isValid(): boolean {
    return (
      !/[^\d]/g.test(this.inputValue) &&
      Number(this.inputValue) <= this.maximumValue &&
      Number(this.inputValue) >= this.minimumValue
    );
  }

  onIncrease() {
    this.currentValue = Math.min(this.currentValue + this.step, this.maximumValue);
    this.inputValue = this.currentValue.toString();
  }

  onDecrease() {
    this.currentValue = Math.max(this.currentValue - this.step, this.minimumValue);
    this.inputValue = this.currentValue.toString();
  }

  onAccept() {
    if (this.isValid) {
      alert("This will do something in the future");
    }
  }

  onReset() {
    this.currentValue = this.startValue;
    this.inputValue = this.startValue.toString();
  }

  onInputClick() {
    this.$refs.input.select();
  }
}
</script>

<style lang="scss" scoped>
.card-number {
  overflow: hidden;
  transition: background-color 0.2s;

  &:hover .content .button:not(:disabled) {
    opacity: 1;
  }

  .interaction {
    flex-shrink: 0;
    flex-grow: 0;
    width: 64px;

    &.hidden {
      transform: translateY(-100%);
    }
  }

  .content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 2.25rem;

    .value,
    .number-input {
      width: 100%;
      display: inline-block;
      font: inherit;
      color: inherit;
      border: 0;
      border-bottom: 1px solid transparent;
      text-align: center;
    }

    .number-form {
      width: calc(100% - 128px);
    }

    .number-input {
      background: none;
      position: relative;

      &.invalid {
        border-color: #ff3860;
      }
    }

    .button-container {
      width: 64px;
    }

    .button-container {
      display: flex;
      justify-content: center;
      align-items: center;
    }

    .button {
      flex-grow: 0;
      flex-shrink: 0;
      opacity: 0.5;
      transition: opacity 0.2s;
    }
  }
}
</style>
