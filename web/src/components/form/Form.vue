<template>
  <div class="form">
    <form @submit.prevent="handleSubmit">
      <div class="input-group" v-for="field in fields" :key="field.name">
        <Checkbox
          v-if="field.type === 'checkbox'"
          :label="field.label"
          v-model="values[field.name]"
          @input="validateInput(field)"
          :error="getErrorForField(field)"
        />

        <Input
          v-else
          :label="field.label"
          v-model="values[field.name]"
          :type="field.type"
          :maximumLength="field.maximumLength"
          :error="getErrorForField(field)"
          :onBlur="() => validateInput(field)"
          @input="handleInputChange(field)"
        />
      </div>

      <div class="form-submit">
        <Button :label="submitLabel" />
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";
import Input from "@/components/form/Input.vue";
import Button from "@/components/Button.vue";
import Checkbox from "@/components/form/Checkbox.vue";

export enum FormFieldType {
  text = "text",
  email = "email",
  password = "password",
  search = "search",
  url = "url",
  number = "number",
  textarea = "textarea",
  checkbox = "checkbox"
}

export interface FormFieldValidator {
  message: string;
  validate: <T>(value: T) => boolean;
}

export interface FormField {
  name: string;
  label: string;
  errorOverride: string;
  type: FormFieldType;
  required: boolean;
  maximumLength: number;
  validators: FormFieldValidator[];
}

@Component({
  name: "Form",
  components: { Input, Button, Checkbox }
})
export default class Form extends Vue {
  @Prop({ default: "Submit" })
  submitLabel!: string;

  @Prop({ required: true })
  values!: { [field: string]: string | boolean };

  @Prop({ required: true })
  fields!: FormField[];

  @Prop({ default: false })
  busy!: boolean;

  errors: { [field: string]: string[] } = {};

  @Watch("fields")
  onFieldsChange() {
    this.validateAllFields();
  }

  get canSubmit(): boolean {
    return this.fields.findIndex(field => !this.validateField(field)) === -1;
  }

  validateAllFields(): void {
    this.fields.forEach(this.validateField);
  }

  validateField(field: FormField): boolean {
    this.clearErrors(field);
    if (field.required && this.isFieldEmpty(field)) {
      this.addError(field, "This field is required.");
    }
    if (field.validators) {
      field.validators.forEach(validator => {
        if (validator.validate(this.values[field.name]) !== true) {
          this.addError(field, validator.message);
        }
      });
    }
    return this.errors[field.name].length === 0;
  }

  handleInputChange(field: FormField): void {
    if (this.hasErrors(field)) {
      this.validateField(field);
    }
  }

  validateInput(field: FormField): void {
    this.validateField(field);
  }

  handleSubmit(): void {
    this.validateAllFields();
    if (!this.hasErrors()) {
      this.$emit("submit");
    }
  }

  isFieldEmpty(field: FormField): boolean {
    const value = this.values[field.name];

    if (typeof value === "boolean") {
      return value;
    }

    return value.length === 0 || !value.trim();
  }

  addError(field: FormField, errorMessage: string) {
    if (!this.errors[field.name]) {
      Vue.set(this.errors, field.name, []);
    }
    this.errors[field.name].push(errorMessage);
  }

  getErrorForField(field: FormField): string {
    if (field.errorOverride && field.errorOverride.length) {
      return field.errorOverride;
    }
    return this.errors[field.name] && this.errors[field.name].length
      ? this.errors[field.name][0]
      : "";
  }

  clearErrors(field: FormField): void {
    if (field) {
      Vue.set(this.errors, field.name, []);
    } else {
      this.errors = {};
    }
  }

  hasErrors(field: FormField | null = null): boolean {
    if (field) {
      return this.errors[field.name] && this.errors[field.name].length > 0;
    }
    return Object.keys(this.errors).some(key => !!this.errors[key].length);
  }
}
</script>

<style lang="scss">
.form {
  padding: 1rem;
  min-width: 300px;

  .form-submit {
    padding-top: 1rem;
    display: flex;
    flex-direction: row-reverse;
  }
}
</style>