import React, { ReactElement } from "react";
import {
  Redirect,
  Route,
  useLocation,
  useParams,
  useRouteMatch,
} from "react-router-dom";
import { useApi } from "../../api";
import { GntInstance, GntNode } from "../../api/models";
import { AuthenticatedRoute } from "../../App";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
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

  const [{ data, isLoading, error }] = useApi<NodeResponse>(
    `clusters/${clusterName}/nodes/${nodeName}`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

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

      <div className={styles.content}>
        <AuthenticatedRoute path={`${path}/primary-instances`}>
          <NodePrimaryInstances instances={data.primaryInstances} />
        </AuthenticatedRoute>

        <AuthenticatedRoute path={`${path}/secondary-instances`}>
          <NodeSecondaryInstances instances={data.secondaryInstances} />
        </AuthenticatedRoute>

        <AuthenticatedRoute path={path} exact>
          <Redirect to={`${url}/primary-instances`} />
        </AuthenticatedRoute>
      </div>
    </ContentWrapper>
  );
};

export default NodeDetail;
