import { faWrench } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntJobLogEntry, GntJobWithLog } from "../../api/models";
import Card from "../../components/Card/Card";
import CardGrid from "../../components/CardGrid/CardGrid";
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

function LogEntryCard({ message, startedAt }: GntJobLogEntry): ReactElement {
  return (
    <Card icon={faWrench} title={message}>
      <div className={styles.text}>
        <JobStartedAt timestamp={startedAt} />
      </div>
    </Card>
  );
}

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
      <CardGrid.Section headline="Log">
        {log.map((entry) => (
          <LogEntryCard key={entry.serial} {...entry} />
        ))}
      </CardGrid.Section>
    </ContentWrapper>
  );
}
