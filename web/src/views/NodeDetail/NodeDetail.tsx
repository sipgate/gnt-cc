import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useClusterName } from "../../helpers/hooks";
import { GntNode } from "../../api/models";
import { useApi } from "../../api";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import styles from "./NodeDetail.module.scss";

interface NodeResponse {
  node: GntNode;
}

const NodeDetail = (): ReactElement => {
  const { nodeName } = useParams();
  const clusterName = useClusterName();

  const [{ data, isLoading, error }] = useApi<NodeResponse>(
    `clusters/${clusterName}/nodes/${nodeName}`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  const { node } = data;

  return (
    <div className={styles.nodeDetail}>
      <div className={styles.hero}>
        <h1 className={styles.title}>{node.name}</h1>
      </div>
    </div>
  );
};

export default NodeDetail;
