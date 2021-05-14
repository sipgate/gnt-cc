import { faThumbtack } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useApi } from "../../api";
import { GntJobLogEntry, GntJobWithLog } from "../../api/models";
import Card from "../../components/Card/Card";
import CardGrid from "../../components/CardGrid/CardGrid";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import JobStatus from "../../components/JobStatus";
import JobSummary from "../../components/JobSummary/JobSummary";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { useClusterName } from "../../helpers/hooks";
import { durationHumanReadable, unixToDate } from "../../helpers/time";
import "./JobDetail.module.scss";

type JobResponse = {
  job: GntJobWithLog;
};

function LogEntryCard({
  message,
  startedAt,
  duration,
}: GntJobLogEntry): ReactElement {
  return (
    <Card icon={faThumbtack} title={message}>
      <p>Started at: {unixToDate(startedAt)}</p>
      {/* TODO: is duration given in milliseconds? */}
      <p>Duration: {durationHumanReadable(duration / 1000)}</p>
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
      <header>
        <h1>{id}</h1>
        <JobStatus status={status} />
      </header>
      <JobSummary summary={summary} />
      <CardGrid>
        <CardGrid.Section headline="Log">
          {log.map((entry) => (
            <LogEntryCard key={entry.serial} {...entry} />
          ))}
        </CardGrid.Section>
      </CardGrid>
    </ContentWrapper>
  );
}
