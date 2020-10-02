import React, { ReactElement } from "react";
import styles from "./Icon.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import { classNameHelper } from "../../helpers";

export interface IconProps {
  icon: IconDefinition;
  className?: string;
}

function Icon({ icon, className }: IconProps): ReactElement {
  return (
    <span className={classNameHelper([className || null, styles.icon])}>
      <FontAwesomeIcon icon={icon} />
    </span>
  );
}

export default Icon;
