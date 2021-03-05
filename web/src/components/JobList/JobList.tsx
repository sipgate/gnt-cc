import React, { ReactElement } from "react";
import { IDataTableColumn } from "react-data-table-component";
import { GntJob } from "../../api/models";
import { durationHumanReadable, unixToDate } from "../../helpers/time";
import Badge, { BadgeStatus } from "../Badge/Badge";
import CustomDataTable from "../CustomDataTable/CustomDataTable";
import styles from "./JobList.module.scss";

function prettifySummary(summary: string): ReactElement {
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

const columns: IDataTableColumn<GntJob>[] = [
  {
    name: "ID",
    sortable: true,
    selector: (row) => row.id,
    width: "120px",
  },
  {
    name: "Status",
    sortable: true,
    selector: (row) => row.status,
    width: "120px",
    cell: (row) => {
      const status = getBadgeStatusFromJobStatus(row.status);
      return <Badge status={status}>{row.status}</Badge>;
    },
  },
  {
    name: "Summary",
    sortable: true,
    selector: (row) => row.summary,
    cell: (row) => prettifySummary(row.summary),
  },
  {
    name: "Start",
    sortable: true,
    selector: (row) => row.startedAt,
    format: (row) => {
      if (row.startedAt < 0) {
        return "-";
      }

      return unixToDate(row.startedAt);
    },
    width: "200px",
  },
  {
    name: "Duration",
    sortable: false,
    width: "120px",
    selector: (row) => row.endedAt,
    format: (row) => {
      if (row.startedAt < 0 || row.endedAt < 0) {
        return "-";
      }

      return durationHumanReadable(row.endedAt - row.startedAt);
    },
  },
];

interface Props {
  jobs: GntJob[];
}

function getBadgeStatusFromJobStatus(status: string) {
  if (status === "success") {
    return BadgeStatus.SUCCESS;
  }
  if (["queued", "waiting", "canceling"].includes(status)) {
    return BadgeStatus.WARNING;
  }
  if (status === "error") {
    return BadgeStatus.FAILURE;
  }

  return undefined;
}

function JobList({ jobs }: Props): ReactElement {
  return (
    <div className={styles.jobList}>
      <CustomDataTable<GntJob>
        columns={columns}
        data={jobs}
        keyField="id"
        defaultSortAsc={false}
        defaultSortField="id"
      />
    </div>
  );
}

export default JobList;
