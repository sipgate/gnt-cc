import React, { PropsWithChildren, ReactElement } from "react";
import classNames from "classnames";
import styles from "./Badge.module.scss";

export enum BadgeStatus {
  SUCCESS,
  FAILURE,
  WARNING,
  PRIMARY,
}

type Props = {
  status?: BadgeStatus;
  className?: string;
};

function Badge({
  children,
  status,
  className,
}: PropsWithChildren<Props>): ReactElement {
  return (
    <span
      className={classNames(
        styles.badge,
        {
          [styles.success]: status === BadgeStatus.SUCCESS,
          [styles.warning]: status === BadgeStatus.WARNING,
          [styles.failure]: status === BadgeStatus.FAILURE,
          [styles.primary]: status === BadgeStatus.PRIMARY,
        },
        className
      )}
    >
      {children}
    </span>
  );
}

export default Badge;
