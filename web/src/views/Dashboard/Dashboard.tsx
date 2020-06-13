import React, { ReactElement } from "react";
import styles from "./Dashboard.module.scss";
import Card from "../../components/cards/Card/Card";

const dummyStats = [
  {
    title: "Nodes",
    value: "12",
  },
  {
    title: "Instances",
    value: "342",
  },
  {
    title: "CPU Cores",
    value: "77",
  },
  {
    title: "Disk Space",
    value: "5 TB",
  },
  {
    title: "Memory",
    value: "256 GB",
  },
];

function Dashboard(): ReactElement {
  return (
    <div className={styles.dashboard}>
      <div className={styles.statistics}>
        {dummyStats.map((stat) => (
          <Card key={stat.title} title={stat.title}>
            <div className={styles.stat}>
              <span>{stat.value}</span>
            </div>
          </Card>
        ))}
      </div>
    </div>
  );
}

export default Dashboard;
