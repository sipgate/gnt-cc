import React, { ReactElement, PropsWithChildren } from "react";
import styles from "./Card.module.scss";
import { classNameHelper } from "../../../helpers";

interface Props {
  title?: string;
  subtitle?: string;
  noHorizontalPadding?: boolean;
  leftTitleSlot?: () => ReactElement;
  rightTitleSlot?: () => ReactElement;
  className?: string;
}

function Card({
  children,
  title,
  subtitle,
  noHorizontalPadding,
  leftTitleSlot,
  rightTitleSlot,
  className,
}: PropsWithChildren<Props>): ReactElement {
  const classNames = classNameHelper([
    className || null,
    styles.card,
    {
      [styles.noHorizontalPadding]: !!noHorizontalPadding,
    },
  ]);

  return (
    <div className={classNames}>
      <div className={styles.cardHeader}>
        {leftTitleSlot && leftTitleSlot()}
        <span className={styles.cardTitle}>{title}</span>
        {rightTitleSlot && rightTitleSlot()}
      </div>
      <div className={styles.cardBody}>{children}</div>
      <div className={styles.cardFooter}>
        <span className={styles.cardTitle}>{subtitle}</span>
      </div>
    </div>
  );
}

export default Card;
