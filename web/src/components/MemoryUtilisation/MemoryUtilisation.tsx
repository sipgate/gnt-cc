import React, { ReactElement } from "react";
import styles from "./MemoryUtilisation.module.scss";
import { classNameHelper } from "../../helpers";

interface Props {
  memoryTotal: number;
  memoryInUse: number;
  memoryUnit: "KB" | "MB" | "GB" | "TB";
}

function MemoryUtilisation({
  memoryInUse,
  memoryTotal,
  memoryUnit,
}: Props): ReactElement {
  const usagePercent = (memoryInUse / memoryTotal) * 100;
  const classNames = classNameHelper([
    styles.wrapper,
    {
      [styles.ok]: usagePercent <= 70,
      [styles.warn]: usagePercent > 70 && usagePercent <= 90,
      [styles.critical]: usagePercent > 90,
    },
  ]);

  return (
    <div className={classNames}>
      <span className={styles.inUse}>{memoryInUse}</span>
      <span className={styles.separator}>/</span>
      <span className={styles.total}>{memoryTotal}</span>
      <span className={styles.unit}>{memoryUnit}</span>
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
