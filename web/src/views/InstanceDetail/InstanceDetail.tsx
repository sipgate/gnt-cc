import React, { ReactElement } from "react";
import styles from "./InstanceDetail.module.scss";
import { useParams, Link } from "react-router-dom";
import { GntInstance } from "../../api/models";
import { useApi } from "../../api";
import Button from "../../components/Button/Button";
import { faSkullCrossbones } from "@fortawesome/free-solid-svg-icons";
import Card from "../../components/cards/Card/Card";
import Tag from "../../components/Tag/Tag";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";

interface InstanceResponse {
  instance: GntInstance;
}

const InstanceDetail = (): ReactElement => {
  const { instanceName, clusterName } = useParams();

  if (!clusterName) {
    throw new Error("cluster not found");
  }

  const [{ data, isLoading, error }] = useApi<InstanceResponse>(
    `clusters/${clusterName}/instances/${instanceName}`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  const { instance } = data;

  return (
    <div className={styles.instanceDetail}>
      {instance && (
        <div>
          <div className={styles.hero}>
            <h1 className={styles.title}>{instance.name}</h1>
            <div className={styles.actions}>
              <Button className={styles.action} label="Migrate" />
              <Button className={styles.action} label="Failover" />
              <Button className={styles.action} label="Shutdown" />
              <Button
                className={styles.action}
                label="Kill"
                icon={faSkullCrossbones}
              />
            </div>
          </div>
          <div className={styles.details}>
            <Card
              title="Nodes"
              subtitle={`Total: ${instance.secondaryNodes.length + 1}`}
            >
              <ul className={styles.nodeList}>
                <li className={styles.nodeListNode}>
                  <Link to={`/nodes/${instance.primaryNode}`}>
                    {instance.primaryNode}
                  </Link>
                  <Tag label="primary" />
                </li>
                {instance.secondaryNodes.map((node) => (
                  <li key={node} className={styles.nodeListNode}>
                    <Link to={`/nodes/${node}`}>{node}</Link>
                  </li>
                ))}
              </ul>
            </Card>
          </div>
        </div>
      )}
      {!instance && <div>Instance not found</div>}
    </div>
  );
};

export default InstanceDetail;
