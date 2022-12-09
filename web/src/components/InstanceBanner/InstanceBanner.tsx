import {
  faComputer,
  faHdd,
  faMemory,
  faMicrochip,
  faTags,
} from "@fortawesome/free-solid-svg-icons";
import React from "react";
import { GntInstance } from "../../api/models";
import { prettyPrintMiB } from "../../helpers";
import Icon from "../Icon/Icon";
import PrefixLink from "../PrefixLink";
import QuickInfoBanner from "../QuickInfoBanner/QuickInfoBanner";
import StatusBadge, { BadgeStatus } from "../StatusBadge/StatusBadge";
import styles from "./InstanceBanner.module.scss";

type Props = {
  instance: GntInstance;
};

function InstanceBanner({ instance }: Props) {
  const totalStorage = instance.disks
    .map(({ capacity }) => capacity)
    .reduce((prev, cur) => prev + cur);

  return (
    <div className={styles.root}>
      <div className={styles.specifications}>
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
        <div className={styles.divider}></div>
        <div className={styles.nodes}>
          <h3>Nodes</h3>
          <div className={styles.node} key={instance.primaryNode}>
            <PrefixLink to={`/nodes/${instance.primaryNode}`}>
              {instance.primaryNode}
            </PrefixLink>

            <StatusBadge status={BadgeStatus.PRIMARY}>Primary</StatusBadge>
          </div>
          {instance.secondaryNodes.map((node) => (
            <div className={styles.node} key={node}>
              <PrefixLink to={`/nodes/${node}`}>{node}</PrefixLink>
            </div>
          ))}
        </div>
      </div>
      <footer>
        <div className={styles.osName} title="Operating System">
          <Icon icon={faComputer}></Icon>
          {instance.OS}
        </div>
        <div className={styles.tags}>
          <Icon icon={faTags}></Icon>
          {instance.tags.map((tag) => (
            <StatusBadge key={tag}>{tag}</StatusBadge>
          ))}
        </div>
      </footer>
    </div>
  );
}

export default InstanceBanner;
