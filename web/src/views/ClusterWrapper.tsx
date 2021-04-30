import React, { ReactElement } from "react";
import { useApi } from "../api";
import { GntCluster } from "../api/models";
import LoadingIndicator from "../components/LoadingIndicator/LoadingIndicator";
import { Switch, useParams, Redirect, useRouteMatch } from "react-router-dom";
import { AuthenticatedRoute } from "../App";
import InstanceDetail from "./InstanceDetail/InstanceDetail";
import Dashboard from "./Dashboard/Dashboard";
import NodeList from "./NodeList/NodeList";
import Navbar from "../components/Navbar/Navbar";
import ClusterNotFound from "../components/ClusterNotFound/ClusterNotFound";
import NodeDetail from "./NodeDetail/NodeDetail";
import Instances from "./Instances/Instances";
import InstanceConsole from "./InstanceConsole/InstanceConsole";
import Jobs from "./Jobs/Jobs";
import Breadcrumbs from "../components/Breadcrumbs/Breadcrumbs";
import JobDetail from "./JobDetail/JobDetail";

interface ClusterResponse {
  clusters: GntCluster[];
}

const ClusterWrapper = (): ReactElement => {
  const [
    { data: clusterData, isLoading: clustersLoading, error: clusterLoadError },
  ] = useApi<ClusterResponse>("clusters");

  const { clusterName } = useParams<{ clusterName: string }>();
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
            <ClusterNotFound clusters={clusterData.clusters} />
          )}

          {clusterExists && (
            <>
              <Navbar clusters={clusterData.clusters} />
              <Breadcrumbs />
              <Switch>
                <AuthenticatedRoute
                  exact
                  path={`${path}/`}
                  component={Dashboard}
                />
                <AuthenticatedRoute
                  path={`${path}/instances/:instanceName/console`}
                  component={InstanceConsole}
                />
                <AuthenticatedRoute
                  path={`${path}/instances/:instanceName`}
                  component={InstanceDetail}
                />
                <AuthenticatedRoute
                  path={`${path}/instances`}
                  component={Instances}
                />
                <AuthenticatedRoute
                  path={`${path}/nodes/:nodeName`}
                  component={NodeDetail}
                />
                <AuthenticatedRoute
                  path={`${path}/nodes`}
                  component={NodeList}
                />
                <AuthenticatedRoute
                  path={`${path}/jobs/:jobID`}
                  component={JobDetail}
                />
                <AuthenticatedRoute path={`${path}/jobs`} component={Jobs} />
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
