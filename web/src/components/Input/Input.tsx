import React, { ChangeEvent, useState, ReactElement } from "react";
import styles from "./Input.module.scss";
import { classNameHelper } from "../../helpers";

type InputType = "text" | "email" | "password";

interface Props {
  type: InputType;
  label: string;
  value: string;
  name: string;
  error?: string | false;
  onBlur?: (event: ChangeEvent<Element>) => void;
  onChange?: (event: ChangeEvent<Element>) => void;
}

const Input = ({
  error,
  type,
  label,
  name,
  onBlur,
  onChange,
  value,
}: Props): ReactElement => {
  const [focused, setFocused] = useState(false);

  const handleInputFocus = () => setFocused(true);

  const handleInputBlur = (event: ChangeEvent<Element>) => {
    setFocused(false);

    if (onBlur) {
      onBlur(event);
    }
  };

  const classNames = classNameHelper([
    styles.inputWrapper,
    {
      [styles.hasError]: !!error,
      [styles.isFocused]: !!focused,
      [styles.hasContent]: !!value,
    },
  ]);

  return (
    <div className={classNames}>
      <input
        className={styles.input}
        id={name}
        name={name}
        type={type}
        onChange={onChange}
        onFocus={handleInputFocus}
        onBlur={handleInputBlur}
        placeholder={label}
      />

      <label className={styles.label} htmlFor={name}>
        {label}
      </label>

      <div className={styles.error}>
        <span>{error}</span>
      </div>
    </div>
  );
};

export default Input;
