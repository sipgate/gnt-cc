import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntJobWithLog } from "../../api/models";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import JobStartedAt from "../../components/JobStartedAt";
import JobStatus from "../../components/JobStatus";
import JobSummary from "../../components/JobSummary/JobSummary";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { useClusterName } from "../../helpers/hooks";
import styles from "./JobDetail.module.scss";

type JobResponse = {
  job: GntJobWithLog;
};

export default function JobDetail(): ReactElement {
  const clusterName = useClusterName();
  const { jobID } = useParams<{ jobID: string }>();
  const [{ data, isLoading, error }] = useApi<JobResponse>(
    `clusters/${clusterName}/jobs/${jobID}`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  const { id, status, summary, log } = data.job;

  return (
    <ContentWrapper>
      <header className={styles.header}>
        <h1>{id}</h1>
        <JobStatus status={status} />
      </header>
      <JobSummary summary={summary} />
      <div className={styles.log}>
        <h3>Log</h3>
        <div className={styles.console}>
          {log.map((entry) => (
            <div key={entry.serial}>
              <JobStartedAt timestamp={entry.startedAt} />{" "}
              <span>{entry.message}</span>
            </div>
          ))}
        </div>
      </div>
    </ContentWrapper>
  );
}
