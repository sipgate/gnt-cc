import React, { ReactElement } from "react";
import styles from "./InstanceDetail.module.scss";
import { useParams } from "react-router-dom";
import { GntInstance } from "../../api/models";
import { useApi } from "../../api";
import Button from "../../components/Button/Button";
import { faSkullCrossbones } from "@fortawesome/free-solid-svg-icons";
import Card from "../../components/cards/Card/Card";
import Tag from "../../components/Tag/Tag";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { useClusterName } from "../../helpers/hooks";
import PrefixLink from "../../components/PrefixLink";
import InstanceConfigurator from "../../components/InstanceConfigurator/InstanceConfigurator";
import Hero from "../../components/Hero/Hero";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";

interface InstanceResponse {
  instance: GntInstance;
}

const InstanceDetail = (): ReactElement => {
  const { instanceName } = useParams();
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
          <Hero title={instance.name}>
            <div className={styles.actions}>
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
          </Hero>
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
