import React, { PropsWithChildren, ReactElement } from "react";
import { classNameHelper } from "../../helpers";
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
      className={classNameHelper([
        styles.badge,
        {
          [styles.success]: status === BadgeStatus.SUCCESS,
          [styles.warning]: status === BadgeStatus.WARNING,
          [styles.failure]: status === BadgeStatus.FAILURE,
        },
      ])}
    >
      {children}
    </span>
  );
}

export default Badge;
