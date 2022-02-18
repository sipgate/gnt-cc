import {
  faCheck,
  faExclamation,
  faEye,
  faEyeSlash,
} from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";
import React, { ReactElement, useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { buildApiUrl } from "../../api";
import { GntJob, GntJobWithLog } from "../../api/models";
import JobWatchContext from "../../contexts/JobWatchContext";
import Dropdown, { Alignment } from "../Dropdown/Dropdown";
import Icon from "../Icon/Icon";
import JobSummary from "../JobSummary/JobSummary";
import {
  getFinishedJobs,
  getUnfinishedJobs,
  getWatcherStatus,
  groupJobIdsByCluster,
  joinJobListsUnique,
  sortJobs,
  WatcherStatus,
} from "./helpers";
import styles from "./JobWatcher.module.scss";

interface JobResponse {
  numberOfJobs: number;
  jobs: GntJobWithLog[];
}

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

function JobWatcher(): ReactElement | null {
  const { trackedJobs, untrackJob } = useContext(JobWatchContext);
  const [jobs, setJobs] = useState<GntJob[]>([]);

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

    async function loadAllJobs() {
      const jobIdsByCluster = groupJobIdsByCluster(trackedJobs);

      const requests: Promise<Response>[] = [];

      jobIdsByCluster.forEach((ids, clusterName) => {
        requests.push(
          fetch(
            buildApiUrl(
              `clusters/${clusterName}/jobs/many?ids=${ids.join(",")}`
            )
          )
        );
      });

      const responses = await Promise.all(requests);

      for (const response of responses) {
        const result = await processResponse(response);

        if (typeof result === "string") {
          console.warn(result);
        } else {
          // remove finished jobs from track list,
          // but keep in jobs list
          getFinishedJobs(result.jobs).forEach(untrackJob);
          setJobs((jobs) => joinJobListsUnique(jobs, result.jobs));
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

  const watcherStatus = getWatcherStatus(jobs);
  const unfinishedJobs = getUnfinishedJobs(jobs);
  const sortedJobs = sortJobs(jobs);

  return (
    <section className={styles.root}>
      <Dropdown icon={faEye} align={Alignment.CENTER}>
        {sortedJobs.map((job) => (
          <div
            className={classNames(styles.job, getJobStatusStyles(job.status))}
            key={job.id}
          >
            <div className={styles.actions}>
              <button
                className={styles.untrackButton}
                onClick={(ev) => {
                  ev.stopPropagation();
                  untrackJob(job);
                }}
                title={`Job status: ${job.status}`}
              >
                <Icon icon={faEyeSlash} />
              </button>
            </div>
            <div className={styles.content}>
              <Link to={`/${job.clusterName}/jobs/${job.id}`}>
                <JobSummary summary={job.summary} />
              </Link>
              <div className={styles.step}></div>
            </div>
          </div>
        ))}
      </Dropdown>
      {watcherStatus === WatcherStatus.InProgress && (
        <span className={styles.count}>{unfinishedJobs.length}</span>
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
