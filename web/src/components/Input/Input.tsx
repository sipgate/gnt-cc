import React, { ChangeEvent, useState, ReactElement } from "react";
import styles from "./Input.module.scss";
import { classNameHelper } from "../../helpers";

type InputType = "text" | "email" | "password" | "search";

interface Props {
  type: InputType;
  label: string;
  value: string;
  name: string;
  error?: string | false;
  onBlur?: (event: ChangeEvent<HTMLInputElement>) => void;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  className?: string;
}

const Input = ({
  error,
  type,
  label,
  name,
  onBlur,
  onChange,
  value,
  className,
}: Props): ReactElement => {
  const [focused, setFocused] = useState(false);

  const handleInputFocus = () => setFocused(true);

  const handleInputBlur = (event: ChangeEvent<HTMLInputElement>) => {
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
    className || null,
  ]);

  return (
    <div className={classNames}>
      <input
        className={styles.input}
        id={name}
        name={name}
        type={type}
        value={value}
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
