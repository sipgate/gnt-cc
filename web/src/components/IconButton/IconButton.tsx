import React, { ReactElement } from "react";
import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import Icon from "../Icon/Icon";
import styles from "./IconButton.module.scss";

type Props = {
  icon: IconDefinition;
  onClick?: () => void;
};

export default function IconButton({ icon, onClick }: Props): ReactElement {
  return (
    <button className={styles.root} onClick={onClick}>
      <Icon className={styles.icon} icon={icon} />
    </button>
  );
}
