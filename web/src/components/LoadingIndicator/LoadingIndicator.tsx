import React, { ReactElement } from "react";
import styles from "./LoadingIndicator.module.scss";
import { faSpinner } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const LoadingIndicator = (): ReactElement => {
  return (
    <div className={styles.loadingIndicator}>
      <FontAwesomeIcon pulse className={styles.icon} icon={faSpinner} />
    </div>
  );
};

export default LoadingIndicator;
