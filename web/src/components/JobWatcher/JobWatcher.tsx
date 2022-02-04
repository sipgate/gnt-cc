import {
  faCheck,
  faExclamation,
  faEye,
  faEyeSlash,
} from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";
import React, { ReactElement, useContext, useEffect, useState } from "react";
import { buildApiUrl } from "../../api";
import AuthContext from "../../api/AuthContext";
import { GntJob, GntJobWithLog } from "../../api/models";
import JobWatchContext, { TrackedJob } from "../../contexts/JobWatchContext";
import Dropdown, { Alignment } from "../Dropdown/Dropdown";
import Icon from "../Icon/Icon";
import JobSummary from "../JobSummary/JobSummary";
import PrefixLink from "../PrefixLink";
import styles from "./JobWatcher.module.scss";

interface JobResponse {
  cluster: string;
  numberOfJobs: number;
  jobs: GntJobWithLog[];
}

enum WatcherStatus {
  InProgress,
  Succeeded,
  HasFailures,
}

type GntJobWithClusterName = {
  job: GntJob;
  clusterName: string;
};

function getJobStatusStyles(status: string) {
  if (status === "success") {
    return styles.success;
  }
  if (["queued", "waiting", "canceling"].includes(status)) {
    return styles.pending;
  }
  if (status === "error") {
    return styles.error;
  }

  return styles.running;
}

function getJobStatusTitle(status: string) {
  return `Job status: ${status}`;
}

function isInProgress(job: GntJob): boolean {
  return job.status !== "error" && job.status !== "success";
}

function getWatcherStatus(jobs: GntJob[]): WatcherStatus {
  for (const job of jobs) {
    if (isInProgress(job)) {
      return WatcherStatus.InProgress;
    }
  }

  for (const job of jobs) {
    if (job.status === "error") {
      return WatcherStatus.HasFailures;
    }
  }

  return WatcherStatus.Succeeded;
}

function initJobsRequest(
  clusterName: string,
  jobIds: number[],
  authToken: string | null
): Promise<Response> {
  return fetch(
    buildApiUrl(`clusters/${clusterName}/jobs/many?ids=${jobIds.join(",")}`),
    {
      headers: {
        Authorization: `Bearer ${authToken}`,
      },
    }
  );
}

function createClusterJobsMap(
  trackedJobs: TrackedJob[]
): Record<string, number[]> {
  const map: Record<string, number[]> = {};

  for (const { clusterName, id } of trackedJobs) {
    if (!map[clusterName]) {
      map[clusterName] = [];
    }

    map[clusterName].push(id);
  }

  return map;
}

function JobWatcher(): ReactElement | null {
  const { authToken } = useContext(AuthContext);
  const { trackedJobs, untrackJob } = useContext(JobWatchContext);
  const [jobs, setJobs] = useState<GntJobWithClusterName[]>([]);

  useEffect(() => {
    async function processResponse(
      response: Response
    ): Promise<JobResponse | string> {
      if (response.status !== 200) {
        const body = await response.text();
        return body || "unknown error";
      }

      const body = await response.json();
      return body as JobResponse;
    }

    function processJobs(clusterName: string, jobs: GntJob[]) {
      const newJobs: GntJobWithClusterName[] = jobs.map((job) => ({
        clusterName,
        job,
      }));

      // remove finidhed jobs from track list,
      // but keep in jobs list
      for (const { job, clusterName } of newJobs) {
        if (!isInProgress(job)) {
          untrackJob({ clusterName, id: job.id });
        }
      }

      setJobs((jobs) => [
        ...jobs.filter(
          (jobA) =>
            !newJobs.find(
              (jobB) =>
                jobB.clusterName === jobA.clusterName &&
                jobB.job.id === jobA.job.id
            )
        ),
        ...newJobs,
      ]);
    }

    async function loadAllJobs() {
      const clusterJobs = createClusterJobsMap(trackedJobs);

      const requests = Object.keys(clusterJobs).map((clusterName) =>
        initJobsRequest(clusterName, clusterJobs[clusterName], authToken)
      );

      const responses = await Promise.all(requests);

      for (const response of responses) {
        const result = await processResponse(response);

        if (typeof result === "string") {
          console.warn(result);
        } else {
          processJobs(result.cluster, result.jobs);
        }
      }
    }

    function checkForUpdates() {
      if (trackedJobs.length > 0) {
        loadAllJobs();
      }
    }

    const interval = setInterval(checkForUpdates, 2000);
    checkForUpdates();

    return () => {
      clearInterval(interval);
    };
  }, [trackedJobs]);

  if (jobs.length === 0) {
    return null;
  }

  const watcherStatus = getWatcherStatus(jobs.flatMap(({ job }) => job));

  const unfinishedJobsCount = jobs.filter(
    ({ job }) => job.status !== "success" && job.status !== "error"
  ).length;

  const sortedJobs = jobs.sort((a, b) => {
    if (isInProgress(a.job) && !isInProgress(b.job)) {
      return -1;
    }

    if (!isInProgress(a.job) && isInProgress(b.job)) {
      return 1;
    }

    return 0;
  });

  return (
    <section className={styles.root}>
      <Dropdown icon={faEye} align={Alignment.CENTER}>
        {sortedJobs.map(({ job, clusterName }) => (
          <div
            className={classNames(styles.job, getJobStatusStyles(job.status))}
            key={job.id}
          >
            <div className={styles.actions}>
              <button
                className={styles.untrackButton}
                onClick={(ev) => {
                  ev.stopPropagation();
                  untrackJob({ clusterName, id: job.id });
                }}
                title={getJobStatusTitle(job.status)}
              >
                <Icon icon={faEyeSlash} />
              </button>
            </div>
            <div className={styles.content}>
              <PrefixLink to={`/jobs/${job.id}`}>
                <JobSummary summary={job.summary} />
              </PrefixLink>
              <div className={styles.step}></div>
            </div>
          </div>
        ))}
      </Dropdown>
      {watcherStatus === WatcherStatus.InProgress && (
        <span className={styles.count}>{unfinishedJobsCount}</span>
      )}
      {watcherStatus === WatcherStatus.Succeeded && (
        <span className={styles.successIndicator}>
          <Icon icon={faCheck} />
        </span>
      )}
      {watcherStatus === WatcherStatus.HasFailures && (
        <span className={styles.failureIndicator}>
          <Icon icon={faExclamation} />
        </span>
      )}
    </section>
  );
}

export default JobWatcher;
