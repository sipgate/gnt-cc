import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import React, { PropsWithChildren, ReactElement } from "react";
import Icon from "../Icon/Icon";

import styles from "./QuickInfoBanner.module.scss";

type ItemProps = {
  value: string;
  label: string;
  icon: IconDefinition;
};

function Item({ value, label, icon }: ItemProps): ReactElement {
  return (
    <div>
      <span className={styles.value}>{value}</span>
      <span className={styles.label}>
        <Icon className={styles.icon} icon={icon} />
        {label}
      </span>
    </div>
  );
}

function QuickInfoBanner({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  return <div className={styles.root}>{children}</div>;
}

QuickInfoBanner.Item = Item;

export default QuickInfoBanner;
