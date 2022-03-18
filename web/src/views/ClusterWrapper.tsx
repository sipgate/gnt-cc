import React, { ReactElement } from "react";
import { Navigate, Outlet, useParams } from "react-router-dom";
import { useApi } from "../api";
import { GntCluster } from "../api/models";
import Breadcrumbs from "../components/Breadcrumbs/Breadcrumbs";
import ClusterNotFound from "../components/ClusterNotFound/ClusterNotFound";
import LoadingIndicator from "../components/LoadingIndicator/LoadingIndicator";
import Navbar from "../components/Navbar/Navbar";

interface ClusterResponse {
  clusters: GntCluster[];
}

const ClusterWrapper = (): ReactElement => {
  const [
    { data: clusterData, isLoading: clustersLoading, error: clusterLoadError },
  ] = useApi<ClusterResponse>("clusters");

  const { clusterName } = useParams<{ clusterName: string }>();

  const clusterExists =
    clusterData &&
    clusterData.clusters.find((cluster) => cluster.name === clusterName) !==
      undefined;

  return (
    <>
      {clusterData && clusterData.clusters.length > 0 && (
        <>
          {!clusterName && <Navigate to={`${clusterData.clusters[0].name}`} />}

          {!clusterExists && clusterName && (
            <ClusterNotFound clusters={clusterData.clusters} />
          )}

          {clusterExists && (
            <>
              <Navbar clusters={clusterData.clusters} />
              <Breadcrumbs />
              <Outlet />
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
