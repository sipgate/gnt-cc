import {
  faHdd,
  faMemory,
  faMicrochip,
  faNetworkWired,
  faTerminal,
  IconDefinition,
} from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntDisk, GntInstance, GntNic } from "../../api/models";
import Badge, { BadgeStatus } from "../../components/Badge/Badge";
import Button from "../../components/Button/Button";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import Icon from "../../components/Icon/Icon";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import PrefixLink from "../../components/PrefixLink";
import { useClusterName } from "../../helpers/hooks";
import { prettyPrintMiB } from "../../helpers/numbers";
import styles from "./InstanceDetail.module.scss";

type QuickInfoItemProps = {
  value: string;
  label: string;
  icon: IconDefinition;
};

function QuickInfoItem({
  value,
  label,
  icon,
}: QuickInfoItemProps): ReactElement {
  return (
    <div>
      <span className={styles.value}>{value}</span>
      <span className={styles.label}>
        <Icon icon={icon} />
        {label}
      </span>
    </div>
  );
}

function DiskCard({ name, capacity, template }: GntDisk): ReactElement {
  return (
    <div className={styles.diskCard}>
      <Icon icon={faHdd} />
      <div>
        <p>{name}</p>
        <p className={styles.diskCapacity}>{prettyPrintMiB(capacity)}</p>
      </div>
      <Badge className={styles.badge}>{template}</Badge>
    </div>
  );
}

function NicCard({ name, mode, mac, bridge }: GntNic): ReactElement {
  return (
    <div className={styles.nicCard}>
      <Icon icon={faNetworkWired} />
      <div>
        <p>{name}</p>
        {bridge.length && <p className={styles.nicBridge}>{bridge}</p>}
        <p className={styles.nicMac}>{mac}</p>
      </div>
      <Badge className={styles.badge}>{mode}</Badge>
    </div>
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
          <header>
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
          <div className={styles.quickInfo}>
            <QuickInfoItem
              icon={faMicrochip}
              label="vCPUs"
              value={instance.cpuCount.toString()}
            />
            <QuickInfoItem
              icon={faMemory}
              label="Memory"
              value={prettyPrintMiB(instance.memoryTotal)}
            />
            <QuickInfoItem
              icon={faHdd}
              label="Storage"
              value={prettyPrintMiB(totalStorage)}
            />
          </div>

          <div className={styles.sections}>
            <section>
              <h2>Disks</h2>
              {instance.disks.map((disk) => (
                <DiskCard key={disk.name} {...disk} />
              ))}
            </section>
            <section>
              <h2>Networking</h2>
              {instance.nics.map((nic) => (
                <NicCard key={nic.name} {...nic} />
              ))}
            </section>
          </div>
        </ContentWrapper>
      )}
      {!instance && <div>Instance not found</div>}
    </>
  );
};

export default InstanceDetail;
