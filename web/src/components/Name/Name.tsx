import React, { PropsWithChildren, ReactElement } from "react";
import styles from "./Name.module.scss";

function Name({ children }: PropsWithChildren<unknown>): ReactElement {
  return <b className={styles.name}>{children}</b>;
}

export default Name;
