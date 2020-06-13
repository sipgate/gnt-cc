import React, { ReactElement } from "react";
import styles from "./Icon.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconDefinition } from "@fortawesome/free-solid-svg-icons";

export interface IconProps {
  icon: IconDefinition;
}

function Icon({ icon }: IconProps): ReactElement {
  return (
    <span className={styles.icon}>
      <FontAwesomeIcon icon={icon} />
    </span>
  );
}

export default Icon;
