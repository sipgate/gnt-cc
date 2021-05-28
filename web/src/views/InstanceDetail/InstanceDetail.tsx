import {
  faHdd,
  faMemory,
  faMicrochip,
  faNetworkWired,
  faServer,
  faTag,
  faTerminal,
} from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntDisk, GntInstance, GntNic } from "../../api/models";
import Badge, { BadgeStatus } from "../../components/Badge/Badge";
import Button from "../../components/Button/Button";
import Card from "../../components/Card/Card";
import CardGrid from "../../components/CardGrid/CardGrid";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import PrefixLink from "../../components/PrefixLink";
import QuickInfoBanner from "../../components/QuickInfoBanner/QuickInfoBanner";
import { useClusterName } from "../../helpers/hooks";
import { prettyPrintMiB } from "../../helpers/numbers";
import styles from "./InstanceDetail.module.scss";

function DiskCard({ name, capacity, template }: GntDisk): ReactElement {
  return (
    <Card icon={faHdd} title={name} badge={<Badge>{template}</Badge>}>
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
          <Badge status={BadgeStatus.PRIMARY}>Primary</Badge>
        ) : undefined
      }
    />
  );
}

function NicCard({ name, mode, mac, bridge }: GntNic): ReactElement {
  return (
    <Card icon={faNetworkWired} title={name} badge={<Badge>{mode}</Badge>}>
      {bridge.length && <p className={styles.nicBridge}>{bridge}</p>}
      <p className={styles.nicMac}>{mac}</p>
    </Card>
  );
}

type InstanceResponse = {
  instance: GntInstance;
};

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

  const totalStorage = instance.disks
    .map(({ capacity }) => capacity)
    .reduce((prev, cur) => prev + cur);

  const hostname = instance.name.split(".")[0];

  return (
    <>
      {instance && (
        <ContentWrapper>
          <header className={styles.header}>
            <h1>{hostname}</h1>
            {instance.isRunning ? (
              <Badge status={BadgeStatus.SUCCESS}>Online</Badge>
            ) : (
              <Badge status={BadgeStatus.FAILURE}>Offline</Badge>
            )}
            {instance.offersVnc && (
              <PrefixLink
                className={styles.consoleLink}
                to={`/instances/${instance.name}/console`}
              >
                <Button label="Open Console" icon={faTerminal} />
              </PrefixLink>
            )}
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
        </ContentWrapper>
      )}
      {!instance && <div>Instance not found</div>}
    </>
  );
};

export default InstanceDetail;
