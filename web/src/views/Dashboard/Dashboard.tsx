import React, { ReactElement } from "react";
import { useApi } from "../../api";
import Card from "../../components/cards/Card/Card";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { convertMBToGB } from "../../helpers";
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

  const stats = [
    {
      title: "Nodes",
      value: data.nodes.count,
    },
    {
      title: "Instances",
      value: data.instances.count,
    },
    {
      title: "Node CPU Cores",
      value: data.nodes.cpuCount,
    },
    {
      title: "Total Node Memory",
      value: `${convertMBToGB(data.nodes.memoryTotal)} GB`,
    },
  ];

  return (
    <>
      <ContentWrapper>
        <div className={styles.statistics}>
          {stats.map((stat) => (
            <Card key={stat.title} title={stat.title}>
              <div className={styles.stat}>
                <span>{stat.value}</span>
              </div>
            </Card>
          ))}
        </div>
      </ContentWrapper>
    </>
  );
}

export default Dashboard;
