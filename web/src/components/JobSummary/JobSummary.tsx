import React, { ReactElement } from "react";
import styles from "./JobSummary.module.scss";

type Props = {
  summary: string;
};

function JobSummary({ summary }: Props): ReactElement {
  const regex = /([A-Z_]+)(?:\((.*)\))?/;

  const matches = summary.match(regex);

  if (!matches) {
    return <></>;
  }

  const jobType = matches[1] || "";
  const jobDetails = matches[2] || "";

  return (
    <span>
      <span className={styles.jobType}>
        {jobType.toLowerCase().replace(/_/g, " ")}
      </span>
      <span className={styles.jobDetails}>{jobDetails}</span>
    </span>
  );
}

export default JobSummary;
