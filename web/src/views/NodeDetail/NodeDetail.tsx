import React, { ReactElement } from "react";
import { Outlet, useLocation, useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntInstance, GntNode } from "../../api/models";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import TabBar from "../../components/TabBar/TabBar";
import { useClusterName } from "../../helpers/hooks";
import styles from "./NodeDetail.module.scss";

interface NodeResponse {
  node: GntNode;
  primaryInstances: GntInstance[];
  secondaryInstances: GntInstance[];
}

const NodeDetail = (): ReactElement => {
  const { nodeName } = useParams<{ nodeName: string }>();
  const clusterName = useClusterName();
  const { pathname } = useLocation();

  const [apiProps] = useApi<NodeResponse>(
    `clusters/${clusterName}/nodes/${nodeName}`
  );

  return (
    <ContentWrapper>
      <div className={styles.tabBarWrapper}>
        <TabBar>
          <TabBar.Tab
            to={`primary-instances`}
            label="Primary Instances"
            isActive={pathname.includes("primary-instances")}
          />
          <TabBar.Tab
            to={`secondary-instances`}
            label="Secondary Instances"
            isActive={pathname.includes("secondary-instances")}
          />
        </TabBar>
      </div>

      <ApiDataRenderer<NodeResponse>
        {...apiProps}
        render={({ primaryInstances, secondaryInstances }) => (
          <div className={styles.content}>
            <Outlet
              context={{
                primaryInstances,
                secondaryInstances,
              }}
            />
          </div>
        )}
      />
    </ContentWrapper>
  );
};

export default NodeDetail;
