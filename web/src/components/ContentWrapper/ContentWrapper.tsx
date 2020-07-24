import React, { ReactElement, PropsWithChildren } from "react";
import styles from "./ContentWrapper.module.scss";

const ContentWrapper = ({
  children,
}: PropsWithChildren<unknown>): ReactElement => {
  return <div className={styles.contentWrapper}>{children}</div>;
};

export default ContentWrapper;
