import React, { ReactElement } from "react";
import styles from "./ClusterNotFound.module.scss";
import ClusterSelector from "../ClusterSelector/ClusterSelector";
import { GntCluster } from "../../api/models";

interface Props {
  clusters: GntCluster[];
}

const ClusterNotFound = ({ clusters }: Props): ReactElement => {
  return (
    <div className={styles.clusterNotFound}>
      <header className={styles.header}>
        <h1>Cluster not found</h1>
        <p>Please choose a different one</p>
      </header>
      <ClusterSelector clusters={clusters} />
    </div>
  );
};

export default ClusterNotFound;
