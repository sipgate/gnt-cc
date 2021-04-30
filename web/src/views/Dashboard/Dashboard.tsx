import {
  faMemory,
  faMicrochip,
  faServer,
} from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useApi } from "../../api";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import QuickInfoBanner from "../../components/QuickInfoBanner/QuickInfoBanner";
import { prettyPrintMiB } from "../../helpers";
import { useClusterName } from "../../helpers/hooks";

interface StatisticElement {
  count: number;
  memoryTotal: number;
  cpuCount: number;
}
interface StatisticsResponse {
  instances: StatisticElement;
  nodes: StatisticElement;
}

function Dashboard(): ReactElement {
  const clusterName = useClusterName();
  const [{ data, isLoading, error }] = useApi<StatisticsResponse>(
    `clusters/${clusterName}/statistics`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  return (
    <>
      <ContentWrapper>
        <QuickInfoBanner>
          <QuickInfoBanner.Item
            icon={faServer}
            label="Nodes"
            value={String(data.nodes.count)}
          />
          <QuickInfoBanner.Item
            icon={faServer}
            label="Instances"
            value={String(data.instances.count)}
          />
          <QuickInfoBanner.Item
            icon={faMicrochip}
            label="Node CPU Cores"
            value={String(data.nodes.cpuCount)}
          />
          <QuickInfoBanner.Item
            icon={faMemory}
            label="Memory"
            value={prettyPrintMiB(data.nodes.memoryTotal)}
          />
        </QuickInfoBanner>
      </ContentWrapper>
    </>
  );
}

export default Dashboard;
