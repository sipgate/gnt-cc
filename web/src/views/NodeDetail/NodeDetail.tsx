import React, { ReactElement } from "react";
import {
  Redirect,
  useLocation,
  useParams,
  useRouteMatch,
} from "react-router-dom";
import { useApi } from "../../api";
import { GntInstance, GntNode } from "../../api/models";
import { AuthenticatedRoute } from "../../App";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import TabBar from "../../components/TabBar/TabBar";
import { useClusterName } from "../../helpers/hooks";
import NodePrimaryInstances from "../NodePrimaryInstances/NodePrimaryInstances";
import NodeSecondaryInstances from "../NodeSecondaryInstances/NodeSecondaryInstances";
import styles from "./NodeDetail.module.scss";

interface NodeResponse {
  node: GntNode;
  primaryInstances: GntInstance[];
  secondaryInstances: GntInstance[];
}

const NodeDetail = (): ReactElement => {
  const { nodeName } = useParams<{ nodeName: string }>();
  const clusterName = useClusterName();
  const { url, path } = useRouteMatch();
  const { pathname } = useLocation();

  const [apiProps] = useApi<NodeResponse>(
    `clusters/${clusterName}/nodes/${nodeName}`
  );

  return (
    <ContentWrapper>
      <div className={styles.tabBarWrapper}>
        <TabBar>
          <TabBar.Tab
            to={`${url}/primary-instances`}
            label="Primary Instances"
            isActive={pathname.includes("primary-instances")}
          />
          <TabBar.Tab
            to={`${url}/secondary-instances`}
            label="Secondary Instances"
            isActive={pathname.includes("secondary-instances")}
          />
        </TabBar>
      </div>

      <ApiDataRenderer<NodeResponse>
        {...apiProps}
        render={({ primaryInstances, secondaryInstances }) => (
          <div className={styles.content}>
            <AuthenticatedRoute path={`${path}/primary-instances`}>
              <NodePrimaryInstances instances={primaryInstances} />
            </AuthenticatedRoute>

            <AuthenticatedRoute path={`${path}/secondary-instances`}>
              <NodeSecondaryInstances instances={secondaryInstances} />
            </AuthenticatedRoute>

            <AuthenticatedRoute path={path} exact>
              <Redirect to={`${url}/primary-instances`} />
            </AuthenticatedRoute>
          </div>
        )}
      />
    </ContentWrapper>
  );
};

export default NodeDetail;
