import React, { ReactElement } from "react";
import styles from "./Button.module.scss";
import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import Icon from "../Icon/Icon";
import { classNameHelper } from "../../helpers";

export interface ButtonProps {
  onClick?: () => void;
  disabled?: boolean;
  label?: string;
  icon?: IconDefinition;
  type?: "button" | "submit";
  className?: string;
}

function Button({
  onClick,
  disabled,
  icon,
  label,
  type,
  className,
}: ButtonProps): ReactElement {
  const handleClick = () => {
    if (!disabled && onClick) {
      onClick();
    }
  };

  const buttonClassNames = classNameHelper([
    className || null,
    {
      [styles.button]: true,
      [styles.hasSpacing]: !!icon && !!label,
    },
  ]);

  return (
    <button
      type={type || "button"}
      onClick={handleClick}
      className={buttonClassNames}
      disabled={disabled}
    >
      {icon && <Icon icon={icon} />}
      {label && <span className={styles.label}>{label}</span>}
    </button>
  );
}

export default Button;
