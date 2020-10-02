import React, { ReactElement } from "react";
import styles from "./Button.module.scss";
import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import Icon from "../Icon/Icon";
import { classNameHelper } from "../../helpers";

export interface ButtonProps {
  onClick?: (event: React.MouseEvent) => void;
  disabled?: boolean;
  label?: string;
  icon?: IconDefinition;
  type?: "button" | "submit";
  className?: string;
  round?: boolean;
  small?: boolean;
  primary?: boolean;
  danger?: boolean;
  href?: string;
}

function Button({
  onClick,
  disabled,
  icon,
  label,
  type,
  className,
  round,
  small,
  primary,
  danger,
  href,
}: ButtonProps): ReactElement {
  const handleClick = (event: React.MouseEvent) => {
    if (!disabled && onClick) {
      onClick(event);
    }
  };

  const buttonClassNames = classNameHelper([
    className || null,
    {
      [styles.button]: true,
      [styles.hasSpacing]: !!icon && !!label,
      [styles.hasLabel]: !!label,
      [styles.isRound]: !!round,
      [styles.isSmall]: !!small,
      [styles.primary]: !!primary,
      [styles.danger]: !!danger,
    },
  ]);

  if (href) {
    return (
      <a href={href} onClick={handleClick} className={buttonClassNames}>
        {icon && <Icon icon={icon} />}
        {label && <span className={styles.label}>{label}</span>}
      </a>
    );
  }

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
