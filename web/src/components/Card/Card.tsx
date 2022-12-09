import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import React, { PropsWithChildren, ReactElement } from "react";
import Icon from "../Icon/Icon";
import styles from "./Card.module.scss";

type Props = {
  icon: IconDefinition;
  title: string;
};

function Card({
  icon,
  title,
  children,
}: PropsWithChildren<Props>): ReactElement {
  return (
    <div className={styles.root}>
      <header>
        <Icon className={styles.icon} icon={icon} />
        <span className={styles.title}>{title}</span>
      </header>
      <section>{children}</section>
    </div>
  );
}

export default Card;
