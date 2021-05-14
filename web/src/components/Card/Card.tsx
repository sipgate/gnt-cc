import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import classNames from "classnames";
import React, { PropsWithChildren, ReactElement } from "react";
import Icon from "../Icon/Icon";
import styles from "./Card.module.scss";

type Props = {
  icon: IconDefinition;
  title: string;
  badge?: ReactElement;
};

function Card({
  icon,
  badge,
  title,
  children,
}: PropsWithChildren<Props>): ReactElement {
  return (
    <div
      className={classNames(styles.root, {
        [styles.hasBody]: !!children,
      })}
    >
      <Icon className={styles.icon} icon={icon} />
      <span className={styles.title}>{title}</span>
      {badge}
      {children && <div className={styles.body}>{children}</div>}
    </div>
  );
}

export default Card;
