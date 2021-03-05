import React, { ReactElement } from "react";
import styles from "./Icon.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";

export interface IconProps {
  icon: IconDefinition;
  className?: string;
}

function Icon({ icon, className }: IconProps): ReactElement {
  return (
    <span className={classNames(className, styles.icon)}>
      <FontAwesomeIcon icon={icon} />
    </span>
  );
}

export default Icon;
