import classNames from "classnames";
import React, { ReactElement } from "react";
import styles from "./MemoryUtilisation.module.scss";

interface Props {
  memoryTotal: string;
  memoryInUse: string;
  usagePercent: number;
}

function MemoryUtilisation({
  memoryInUse,
  memoryTotal,
  usagePercent,
}: Props): ReactElement {
  return (
    <div
      className={classNames(styles.wrapper, {
        [styles.ok]: usagePercent <= 70,
        [styles.warn]: usagePercent > 70 && usagePercent <= 90,
        [styles.critical]: usagePercent > 90,
      })}
    >
      <span className={styles.inUse}>{memoryInUse}</span>
      <span className={styles.separator}>/</span>
      <span className={styles.total}>{memoryTotal}</span>
      <span className={styles.indicator}>
        <span
          className={styles.progress}
          style={{ width: `${usagePercent}%` }}
        />
      </span>
    </div>
  );
}

export default MemoryUtilisation;
