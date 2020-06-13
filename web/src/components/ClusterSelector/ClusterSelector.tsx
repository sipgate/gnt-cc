import React, { ReactElement } from "react";
import styles from "./ClusterSelector.module.scss";
import Dropdown from "../Dropdown/Dropdown";
import { faServer } from "@fortawesome/free-solid-svg-icons";
import { GntCluster } from "../../api/models";
import { useParams, useHistory } from "react-router-dom";
import { classNameHelper } from "../../helpers";

interface Props {
  clusters: GntCluster[];
}

const ClusterSelector = ({ clusters }: Props): ReactElement => {
  const { clusterName } = useParams();
  const history = useHistory();

  if (!clusterName) {
    throw new Error("cluster not found");
  }

  const selectCluster = (name: string): void => {
    // TODO: this is rather ugly and error-prone right now
    history.push(history.location.pathname.replace(clusterName, name));
  };

  return (
    <div className={styles.clusterSelector}>
      <Dropdown label={clusterName} icon={faServer}>
        {clusters.map((cluster) => (
          <div
            className={classNameHelper([
              styles.cluster,
              { [styles.selected]: cluster.name === clusterName },
            ])}
            key={cluster.name}
            onClick={() => selectCluster(cluster.name)}
          >
            <span className={styles.name}>{cluster.name}</span>
            <span className={styles.hostname}>{cluster.hostname}</span>
            <span className={styles.description}>{cluster.description}</span>
          </div>
        ))}
      </Dropdown>
    </div>
  );
};

export default ClusterSelector;
