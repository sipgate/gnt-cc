import React, { ReactElement, useContext } from "react";
import styles from "./Dashboard.module.scss";
import Card from "../../components/cards/Card/Card";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import AuthContext from "../../api/AuthContext";
import Hero from "../../components/Hero/Hero";
import { useApi } from "../../api";
import { useClusterName } from "../../helpers/hooks";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { convertMBToGB } from "../../helpers";

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
  const authContext = useContext(AuthContext);
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
      <Hero title={`Welcome back, ${authContext.username}!`} />
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
