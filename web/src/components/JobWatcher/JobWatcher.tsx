import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";
import React, { ReactElement, useContext, useEffect } from "react";
import { useApi } from "../../api";
import { GntJobWithLog } from "../../api/models";
import JobWatchContext from "../../contexts/JobWatchContext";
import { useClusterName } from "../../helpers/hooks";
import Dropdown, { Alignment } from "../Dropdown/Dropdown";
import Icon from "../Icon/Icon";
import JobSummary from "../JobSummary/JobSummary";
import PrefixLink from "../PrefixLink";
import styles from "./JobWatcher.module.scss";

interface JobResponse {
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

function getJobStatusTitle(status: string) {
  return `Job status: ${status}`;
}

function JobWatcher(): ReactElement | null {
  const clusterName = useClusterName();
  const { trackedJobs, untrackJob } = useContext(JobWatchContext);

  const [{ data, error }, loadJobs] = useApi<JobResponse>(
    `clusters/${clusterName}/jobs/many?ids=${trackedJobs.join(",")}`,
    { manual: true }
  );

  useEffect(() => {
    function updateJobs() {
      if (trackedJobs.length > 0) {
        loadJobs();
      }
    }

    const interval = setInterval(updateJobs, 2000);
    updateJobs();

    return () => {
      clearInterval(interval);
    };
  }, [trackedJobs]);

  if (error) {
    return <span>Error</span>;
  }

  if (trackedJobs.length === 0) {
    return null;
  }

  const jobs = data
    ? data.jobs.filter((job) => trackedJobs.includes(job.id)).reverse()
    : [];

  return (
    <section className={styles.root}>
      <Dropdown icon={faEye} align={Alignment.CENTER}>
        {jobs.map((job) => (
          <div
            className={classNames(styles.job, getJobStatusStyles(job.status))}
            key={job.id}
          >
            <div className={styles.actions}>
              <button
                className={styles.untrackButton}
                onClick={() => untrackJob(job.id)}
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
      <span className={styles.count}>{trackedJobs.length}</span>
    </section>
  );
}

export default JobWatcher;
