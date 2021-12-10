import {
  faHdd,
  faMemory,
  faMicrochip,
  faNetworkWired,
  faServer,
  faTag,
  faTerminal,
} from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement, useContext } from "react";
import { useParams } from "react-router-dom";
import { HttpMethod, useApi } from "../../api";
import { GntDisk, GntInstance, GntNic, JobIdResponse } from "../../api/models";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import Button from "../../components/Button/Button";
import Card from "../../components/Card/Card";
import CardGrid from "../../components/CardGrid/CardGrid";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import PrefixLink from "../../components/PrefixLink";
import QuickInfoBanner from "../../components/QuickInfoBanner/QuickInfoBanner";
import StatusBadge, {
  BadgeStatus,
} from "../../components/StatusBadge/StatusBadge";
import JobWatchContext from "../../contexts/JobWatchContext";
import { useClusterName } from "../../helpers/hooks";
import { prettyPrintMiB } from "../../helpers/numbers";
import styles from "./InstanceDetail.module.scss";

function DiskCard({ name, capacity, template }: GntDisk): ReactElement {
  return (
    <Card
      icon={faHdd}
      title={name}
      badge={<StatusBadge>{template}</StatusBadge>}
    >
      <p className={styles.diskCapacity}>{prettyPrintMiB(capacity)}</p>
    </Card>
  );
}

function TagCard({ tag }: { tag: string }): ReactElement {
  return <Card icon={faTag} title={tag} />;
}

type NodeCardProps = {
  name: string;
  primary?: boolean;
};

function NodeCard({ name, primary }: NodeCardProps): ReactElement {
  return (
    <Card
      icon={faServer}
      title={name}
      badge={
        primary ? (
          <StatusBadge status={BadgeStatus.PRIMARY}>Primary</StatusBadge>
        ) : undefined
      }
    />
  );
}

function NicCard({ name, mode, mac, bridge, vlan }: GntNic): ReactElement {
  return (
    <Card
      icon={faNetworkWired}
      title={name}
      badge={<StatusBadge>{mode}</StatusBadge>}
    >
      {bridge.length && <p className={styles.nicBridge}>{bridge}</p>}
      {vlan && (
        <p className={styles.nicVlan}>
          Vlan: <b>{vlan}</b>
        </p>
      )}
      <p className={styles.nicMac}>{mac}</p>
    </Card>
  );
}

type InstanceResponse = {
  instance: GntInstance;
};

const InstanceDetail = (): ReactElement => {
  const { trackJob } = useContext(JobWatchContext);
  const { instanceName } = useParams<{ instanceName: string }>();
  const clusterName = useClusterName();

  const [apiProps] = useApi<InstanceResponse>(
    `clusters/${clusterName}/instances/${instanceName}`
  );

  const [, executeRestart] = useApi<JobIdResponse>(
    {
      slug: `/clusters/${clusterName}/instances/${instanceName}/start`,
      method: HttpMethod.Post,
    },
    {
      manual: true,
    }
  );

  const [, executeStart] = useApi<JobIdResponse>(
    {
      slug: `/clusters/${clusterName}/instances/${instanceName}/reboot`,
      method: HttpMethod.Post,
    },
    {
      manual: true,
    }
  );

  const [, executeShutdown] = useApi<JobIdResponse>(
    {
      slug: `/clusters/${clusterName}/instances/${instanceName}/shutdown`,
      method: HttpMethod.Post,
    },
    {
      manual: true,
    }
  );

  async function restartInstance() {
    const response = await executeRestart();

    if (typeof response === "string") {
      console.error(response);
      return;
    }

    trackJob(response.jobId);
  }

  async function startInstance() {
    const response = await executeStart();

    if (typeof response === "string") {
      console.error(response);
      return;
    }

    trackJob(response.jobId);
  }

  async function shutdownInstance() {
    const response = await executeShutdown();

    if (typeof response === "string") {
      console.error(response);
      return;
    }

    trackJob(response.jobId);
  }

  return (
    <ContentWrapper>
      <ApiDataRenderer<InstanceResponse>
        {...apiProps}
        render={({ instance }) => {
          const totalStorage = instance.disks
            .map(({ capacity }) => capacity)
            .reduce((prev, cur) => prev + cur);

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

                <div className={styles.actions}>
                  <Button onClick={startInstance} label="Start" />
                  <Button onClick={restartInstance} label="Restart" />
                  <Button onClick={shutdownInstance} label="Shutdown" />

                  {instance.offersVnc && (
                    <PrefixLink to={`/instances/${instance.name}/console`}>
                      <Button label="Open Console" icon={faTerminal} />
                    </PrefixLink>
                  )}
                </div>
              </header>

              <QuickInfoBanner>
                <QuickInfoBanner.Item
                  icon={faMicrochip}
                  label="vCPUs"
                  value={instance.cpuCount.toString()}
                />
                <QuickInfoBanner.Item
                  icon={faMemory}
                  label="Memory"
                  value={prettyPrintMiB(instance.memoryTotal)}
                />
                <QuickInfoBanner.Item
                  icon={faHdd}
                  label="Storage"
                  value={prettyPrintMiB(totalStorage)}
                />
              </QuickInfoBanner>

              <CardGrid>
                <CardGrid.Section headline="Disks">
                  {instance.disks.map((disk) => (
                    <DiskCard key={disk.name} {...disk} />
                  ))}
                </CardGrid.Section>
                <CardGrid.Section headline="Networking">
                  {instance.nics.map((nic) => (
                    <NicCard key={nic.name} {...nic} />
                  ))}
                </CardGrid.Section>
                <CardGrid.Section headline="Nodes">
                  <NodeCard
                    key={instance.primaryNode}
                    name={instance.primaryNode}
                    primary
                  />
                  {instance.secondaryNodes.map((node) => (
                    <NodeCard key={node} name={node} />
                  ))}
                </CardGrid.Section>
                <CardGrid.Section headline="Tags">
                  {instance.tags.map((tag) => (
                    <TagCard key={tag} tag={tag} />
                  ))}
                </CardGrid.Section>
              </CardGrid>
            </>
          );
        }}
      />
    </ContentWrapper>
  );
};

export default InstanceDetail;
