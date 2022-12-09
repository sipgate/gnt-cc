import {
  faMemory,
  faMicrochip,
  faServer,
} from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useApi } from "../../api";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import PrefixLink from "../../components/PrefixLink";
import QuickInfoBanner from "../../components/QuickInfoBanner/QuickInfoBanner";
import StatusBadge from "../../components/StatusBadge/StatusBadge";
import { prettyPrintMiB } from "../../helpers";
import { useClusterName } from "../../helpers/hooks";
import styles from "./Dashboard.module.scss";

interface StatisticElement {
  count: number;
  memoryTotal: number;
  cpuCount: number;
}
interface StatisticsResponse {
  instances: StatisticElement;
  nodes: StatisticElement;
  master: string;
}

function Dashboard(): ReactElement {
  const clusterName = useClusterName();

  const [apiProps] = useApi<StatisticsResponse>(
    `clusters/${clusterName}/statistics`
  );

  return (
    <ContentWrapper>
      <ApiDataRenderer<StatisticsResponse>
        {...apiProps}
        render={({ master, nodes, instances }) => (
          <>
            <div className={styles.clusterSpecifications}>
              <QuickInfoBanner>
                <QuickInfoBanner.Item
                  icon={faServer}
                  label="Nodes"
                  value={String(nodes.count)}
                />
                <QuickInfoBanner.Item
                  icon={faServer}
                  label="Instances"
                  value={String(instances.count)}
                />
                <QuickInfoBanner.Item
                  icon={faMicrochip}
                  label="Node CPU Cores"
                  value={String(nodes.cpuCount)}
                />
                <QuickInfoBanner.Item
                  icon={faMemory}
                  label="Memory"
                  value={prettyPrintMiB(nodes.memoryTotal)}
                />
              </QuickInfoBanner>
            </div>
            <div className={styles.currentMaster}>
              <StatusBadge>Master</StatusBadge>
              <PrefixLink to={`/nodes/${master}`}>
                <span className={styles.master}>{master}</span>
              </PrefixLink>
            </div>
          </>
        )}
      />
    </ContentWrapper>
  );
}

export default Dashboard;
