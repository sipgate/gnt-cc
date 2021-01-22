import React, { ReactElement } from "react";
import { useApi } from "../../api";
import { GntJob } from "../../api/models";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import Hero from "../../components/Hero/Hero";
import JobList from "../../components/JobList/JobList";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { useClusterName } from "../../helpers/hooks";

interface JobResponse {
  jobs: GntJob[];
}

const Jobs = (): ReactElement => {
  const clusterName = useClusterName();
  const [{ data, isLoading, error }] = useApi<JobResponse>(
    `clusters/${clusterName}/jobs`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  const renderContent = (): ReactElement => {
    if (isLoading) {
      return <LoadingIndicator />;
    }

    if (!data) {
      return <div>Failed to load: {error}</div>;
    }

    return <JobList jobs={data.jobs} />;
  };

  return (
    <>
      <Hero title="Jobs" />
      <ContentWrapper>{renderContent()}</ContentWrapper>
    </>
  );
};

export default Jobs;
