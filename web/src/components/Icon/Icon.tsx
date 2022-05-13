import React, { ReactElement } from "react";
import styles from "./Icon.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";

export interface IconProps {
  icon: IconDefinition;
  spin?: boolean;
  className?: string;
}

function Icon({ icon, spin, className }: IconProps): ReactElement {
  return (
    <span className={classNames(className, styles.icon)}>
      <FontAwesomeIcon icon={icon} spin={spin} />
    </span>
  );
}

export default Icon;
