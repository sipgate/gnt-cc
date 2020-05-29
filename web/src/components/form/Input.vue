<template>
  <div
    class="custom-input"
    :class="{
      'has-error': hasError,
      focused: isFocused,
      'has-been-focused': hasBeenFocused
    }"
    @keyup.enter="blur"
  >
    <div class="textarea-height-fix" v-if="type === 'textarea'" v-html="fixedValue" />

    <textarea
      v-if="type === 'textarea'"
      :autofocus="autofocus"
      :value="value"
      :maxlength="maximumLength"
      :placeholder="isFocused ? '' : label"
      @input="$emit('input', $event.target.value)"
      @focus="focus"
      @blur="blur"
    />

    <input
      v-else
      :type="type"
      :autofocus="autofocus"
      :value="value"
      :maxlength="maximumLength"
      :placeholder="isFocused ? '' : label"
      @input="$emit('input', $event.target.value)"
      @focus="focus"
      @blur="blur"
    />
    <label>{{ label }}</label>

    <div class="error-message">
      <span>{{ error }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";
import Icon from "@/components/Icon.vue";

@Component({
  name: "Input",
  components: { Icon }
})
export default class Input extends Vue {
  @Prop({
    default: "text",
    validator: val =>
      [
        "text",
        "email",
        "password",
        "search",
        "url",
        "number",
        "date",
        "textarea"
      ].includes(val)
  })
  readonly type!: string;

  @Prop({ required: true })
  readonly label!: string;

  @Prop({ required: true })
  readonly value!: string;

  @Prop({ default: false })
  readonly autofocus!: boolean;

  @Prop({ default: 255 })
  readonly maximumLength!: boolean;

  @Prop({ default: "" })
  readonly error!: string;

  userFocused = false;
  autoFocused = false;
  hasBeenFocused = false;

  mounted() {
    this.autoFocused = !this.isEmpty;
  }

  updated() {
    this.autoFocused = !this.isEmpty;
  }

  get isEmpty() {
    return this.value.length === 0;
  }

  get hasError() {
    return this.error.length > 0;
  }

  get fixedValue() {
    return encodeURIComponent(this.value).replace(/%0/g, "<br/>");
  }

  get isFocused() {
    return this.userFocused || this.autoFocused;
  }

  focus() {
    this.$emit("focus");
    this.userFocused = true;
  }

  blur() {
    this.$emit("blur");
    if (!this.hasBeenFocused) {
      this.hasBeenFocused = true;
    }
    if (this.isEmpty) {
      this.userFocused = false;
      this.autoFocused = false;
    }
  }
}
</script>

<style scoped lang="scss">
// TODO: use project vars
$colorRed: #e74c3c;
$colorBlue: #435f7a;
$colorDarkBlue: #34495e;
$colorGreen: #27ae60;
$colorDarkGrey: #333;
$colorGrey: #727272;
$colorLightGrey: #e2e2e2;
$colorWhite: #fff;
$colorBlack: #000;
$transitionDuration: 0.1s;

.custom-input {
  // TODO: change class to 'input' once buefy is gone
  position: relative;
  display: inline-block;
  color: $colorDarkGrey;
  width: 100%;
  max-width: 100%;
  margin: 0.5rem 0;
  border: 1px solid $colorLightGrey;
  border-radius: 3px;
  box-shadow: inset 0 1px 3px 0 rgba(0, 0, 0, 0.04);
  overflow: hidden;

  &.focused {
    border-color: $colorBlue;

    input,
    textarea {
      transform: translateY(0.3rem);
    }

    label {
      opacity: 1;
      transform: translateY(0);
    }
  }

  &.has-error {
    border-color: $colorRed;

    .error-message {
      opacity: 1;
      transform: translateY(0);
    }
  }

  /* &.has-error {
    .error-icon,
    .error-message {
      opacity: 1;
    }

    &.light,
    &.dark {
      input,
      textarea {
        border-color: $colorRed;
      }
    }
  } */

  /*  .error-icon,
  .error-message {
    color: $colorRed;
    opacity: 0;
    transition: opacity 0.2s;
  }

  .error-icon {
    position: absolute;
    font-size: 18px;
    top: 10px;
    right: 4px;
  }

  .error-message {
    position: absolute;
    width: 100%;
    padding: 4px 0;
    font-size: 0.9em;

    span {
      display: inline-block;
      white-space: normal;
      min-width: 100%;
      width: 0;
    }
  } */

  input,
  textarea {
    color: $colorGrey;
    padding: 0.75rem 0.5rem 0.5rem 0.5rem;
    display: block;
    background: transparent;
    transition: all $transitionDuration;
    width: 100%;
    min-width: 100%;
    max-width: 100%;
    height: 48px;
    font-size: 1rem;
    border: 0;

    &:focus,
    &:active {
      outline: none;
      box-shadow: 0;
    }
  }

  textarea {
    height: 100%;
    resize: none;
    overflow: hidden;
    overflow-y: auto;
    position: absolute;
    top: 0;
    box-sizing: border-box;
  }

  .textarea-height-fix {
    max-height: 120px;
    overflow: hidden;
    min-height: 38px;
    padding: 11px 5px; // use slightly more padding than textarea to hide scrollbar
    box-sizing: border-box;
    visibility: hidden;
    width: 100%;
  }

  label,
  .error-message {
    position: absolute;
    opacity: 0;
    font-size: 0.75rem;
    top: 0.1rem;
    font-weight: bold;
    transform: translateY(-0.25rem);
    transition: all $transitionDuration;
  }

  label {
    left: 0.5rem;
    color: $colorBlue;
  }

  .error-message {
    right: 0.5rem;
    color: $colorRed;
  }
}
</style>
