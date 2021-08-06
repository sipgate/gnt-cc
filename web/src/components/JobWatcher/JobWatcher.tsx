import React, { ReactElement, useContext } from "react";
import JobWatchContext from "../../contexts/JobWatchContext";
import styles from "./JobWatcher.module.scss";

function JobWatcher(): ReactElement {
  const { trackedJobs } = useContext(JobWatchContext);

  return (
    <section className={styles.root}>
      {trackedJobs.map((jobID) => (
        <p>{jobID}</p>
      ))}
    </section>
  );
}

export default JobWatcher;
