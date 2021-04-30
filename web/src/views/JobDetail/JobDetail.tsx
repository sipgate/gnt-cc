import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { useClusterName } from "../../helpers/hooks";

export default function JobDetail(): ReactElement {
  const clusterName = useClusterName();
  const { jobID } = useParams<{ jobID: string }>();
  const [{ data, isLoading, error }] = useApi<unknown>(
    `clusters/${clusterName}/jobs/${jobID}`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  return <>success</>;
}
