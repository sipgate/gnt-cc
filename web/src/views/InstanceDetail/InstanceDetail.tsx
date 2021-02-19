import { faSkullCrossbones } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntInstance } from "../../api/models";
import Button from "../../components/Button/Button";
import Card from "../../components/cards/Card/Card";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import InstanceConfigurator from "../../components/InstanceConfigurator/InstanceConfigurator";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import PrefixLink from "../../components/PrefixLink";
import Tag from "../../components/Tag/Tag";
import { useClusterName } from "../../helpers/hooks";
import styles from "./InstanceDetail.module.scss";

interface InstanceResponse {
  instance: GntInstance;
}

const InstanceDetail = (): ReactElement => {
  const { instanceName } = useParams<{ instanceName: string }>();
  const clusterName = useClusterName();

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
    <>
      {instance && (
        <div>
          <div className={styles.actions}>
            {instance.offersVnc && (
              <PrefixLink
                className={styles.link}
                to={`/instances/${instance.name}/console`}
              >
                <Button
                  className={styles.action}
                  label="Open Console"
                  primary
                />
              </PrefixLink>
            )}
            <Button className={styles.action} label="Migrate" primary />
            <Button className={styles.action} label="Failover" primary />
            <Button className={styles.action} label="Shutdown" danger />
            <Button
              className={styles.action}
              label="Kill"
              icon={faSkullCrossbones}
              danger
            />
          </div>
          <ContentWrapper>
            <InstanceConfigurator instance={instance} />
            <div className={styles.details}>
              <Card
                title="Nodes"
                subtitle={`Total: ${instance.secondaryNodes.length + 1}`}
              >
                <ul className={styles.nodeList}>
                  <li className={styles.nodeListNode}>
                    <PrefixLink
                      className={styles.link}
                      to={`/nodes/${instance.primaryNode}`}
                    >
                      {instance.primaryNode}
                    </PrefixLink>
                    <Tag label="primary" />
                  </li>
                  {instance.secondaryNodes.map((node) => (
                    <li key={node} className={styles.nodeListNode}>
                      <PrefixLink className={styles.link} to={`/nodes/${node}`}>
                        {node}
                      </PrefixLink>
                    </li>
                  ))}
                </ul>
              </Card>
            </div>
          </ContentWrapper>
        </div>
      )}
      {!instance && <div>Instance not found</div>}
    </>
  );
};

export default InstanceDetail;
