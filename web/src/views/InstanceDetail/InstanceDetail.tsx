import { faHdd, faNetworkWired } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntInstance } from "../../api/models";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import Card from "../../components/Card/Card";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import InstanceActions from "../../components/InstanceActions/InstanceActions";
import InstanceBanner from "../../components/InstanceBanner/InstanceBanner";
import StatusBadge, {
  BadgeStatus,
} from "../../components/StatusBadge/StatusBadge";
import { useClusterName } from "../../helpers/hooks";
import { prettyPrintMiB } from "../../helpers/numbers";
import styles from "./InstanceDetail.module.scss";

type InstanceResponse = {
  instance: GntInstance;
};

const InstanceDetail = (): ReactElement => {
  const { instanceName } = useParams<{ instanceName: string }>();
  const clusterName = useClusterName();

  const [apiProps] = useApi<InstanceResponse>(
    `clusters/${clusterName}/instances/${instanceName}`
  );

  return (
    <ContentWrapper>
      <ApiDataRenderer<InstanceResponse>
        {...apiProps}
        render={({ instance }) => {
          const hostname = instance.name.split(".")[0];

          return (
            <>
              <header className={styles.header}>
                <h1>{hostname}</h1>
                {instance.isRunning ? (
                  <StatusBadge status={BadgeStatus.SUCCESS}>Online</StatusBadge>
                ) : (
                  <StatusBadge status={BadgeStatus.FAILURE}>
                    Offline
                  </StatusBadge>
                )}

                <InstanceActions
                  clusterName={clusterName}
                  instance={instance}
                />
              </header>

              <InstanceBanner instance={instance} />

              <div className={styles.cards}>
                <Card icon={faNetworkWired} title="Networking">
                  {instance.nics.map(({ name, mode, bridge, mac }) => (
                    <div className={styles.nic} key={name}>
                      <p>
                        <b>{name}</b>
                      </p>
                      <p>
                        {mode}: {mode === "bridged" ? bridge : ""} • {mac}
                      </p>
                    </div>
                  ))}
                </Card>
                <Card icon={faHdd} title="Disks">
                  {instance.disks.map(({ name, template, capacity }) => (
                    <div className={styles.disk} key={name}>
                      <p>
                        <b>{name}</b>
                      </p>
                      <p>
                        {template} • {prettyPrintMiB(capacity)}
                      </p>
                    </div>
                  ))}
                </Card>
              </div>
            </>
          );
        }}
      />
    </ContentWrapper>
  );
};

export default InstanceDetail;
