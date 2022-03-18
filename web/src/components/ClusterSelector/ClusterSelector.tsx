import React, { ReactElement } from "react";
import styles from "./ClusterSelector.module.scss";
import Dropdown from "../Dropdown/Dropdown";
import { faServer } from "@fortawesome/free-solid-svg-icons";
import { GntCluster } from "../../api/models";
import { useLocation, useNavigate } from "react-router-dom";
import classNames from "classnames";
import { useClusterName } from "../../helpers/hooks";

interface Props {
  clusters: GntCluster[];
}

const ClusterSelector = ({ clusters }: Props): ReactElement => {
  const clusterName = useClusterName();
  const navigate = useNavigate();
  const location = useLocation();

  const selectCluster = (name: string): void => {
    const { pathname } = location;

    const parts = pathname.split("/");
    const slug = parts[2] ? `/${parts[2]}` : "";

    navigate(`/${name}${slug}`);
  };

  return (
    <div className={styles.clusterSelector}>
      <Dropdown label={clusterName} icon={faServer}>
        {clusters.map((cluster) => (
          <div
            className={classNames(styles.cluster, {
              [styles.selected]: cluster.name === clusterName,
            })}
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
