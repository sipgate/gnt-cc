import React, { ReactElement, useContext } from "react";
import styles from "./Dashboard.module.scss";
import Card from "../../components/cards/Card/Card";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import AuthContext from "../../api/AuthContext";
import Hero from "../../components/Hero/Hero";

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
  const authContext = useContext(AuthContext);

  return (
    <>
      <Hero title={`Welcome back, ${authContext.username}!`} />
      <ContentWrapper>
        <div className={styles.statistics}>
          {dummyStats.map((stat) => (
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
