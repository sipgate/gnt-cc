import React, { PropsWithChildren, ReactElement } from "react";
import classNames from "classnames";
import styles from "./Badge.module.scss";

export enum BadgeStatus {
  SUCCESS,
  FAILURE,
  WARNING,
}

type Props = {
  status?: BadgeStatus;
};

function Badge({ children, status }: PropsWithChildren<Props>): ReactElement {
  return (
    <span
      className={classNames(styles.badge, {
        [styles.success]: status === BadgeStatus.SUCCESS,
        [styles.warning]: status === BadgeStatus.WARNING,
        [styles.failure]: status === BadgeStatus.FAILURE,
      })}
    >
      {children}
    </span>
  );
}

export default Badge;
