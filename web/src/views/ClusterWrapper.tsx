import React, { ReactElement } from "react";
import { useApi } from "../api";
import { GntCluster } from "../api/models";
import LoadingIndicator from "../components/LoadingIndicator/LoadingIndicator";
import { Switch, useParams, Redirect, useRouteMatch } from "react-router-dom";
import { AuthenticatedRoute } from "../App";
import InstanceDetail from "./InstanceDetail/InstanceDetail";
import Dashboard from "./Dashboard/Dashboard";
import InstanceList from "./InstanceList/InstanceList";
import NodeList from "./NodeList/NodeList";
import Navbar from "../components/Navbar/Navbar";
import ClusterSelector from "../components/ClusterSelector/ClusterSelector";

interface ClusterResponse {
  clusters: GntCluster[];
}

const ClusterWrapper = (): ReactElement => {
  const [
    { data: clusterData, isLoading: clustersLoading, error: clusterLoadError },
  ] = useApi<ClusterResponse>("clusters");

  const { clusterName } = useParams();
  const { path } = useRouteMatch();

  const clusterExists =
    clusterData &&
    clusterData.clusters.find((cluster) => cluster.name === clusterName) !==
      undefined;

  return (
    <>
      {clusterData && clusterData.clusters.length > 0 && (
        <>
          {!clusterName && <Redirect to={`${clusterData.clusters[0].name}`} />}

          {!clusterExists && clusterName && (
            <div>
              <p>Cluster not found, please choose a different one.</p>
              <ClusterSelector clusters={clusterData.clusters} />
            </div>
          )}

          {clusterExists && (
            <>
              <Navbar clusters={clusterData.clusters} />
              <Switch>
                <AuthenticatedRoute
                  exact
                  path={`${path}/`}
                  component={Dashboard}
                />
                <AuthenticatedRoute
                  path={`${path}/instances/:instanceName`}
                  component={InstanceDetail}
                />

                <AuthenticatedRoute
                  path={`${path}/instances`}
                  component={InstanceList}
                />
                <AuthenticatedRoute
                  path={`${path}/nodes`}
                  component={NodeList}
                />
              </Switch>
            </>
          )}
        </>
      )}
      {clustersLoading && <LoadingIndicator />}
      {clusterLoadError && <div>API Error: {clusterLoadError}</div>}
    </>
  );
};

export default ClusterWrapper;
