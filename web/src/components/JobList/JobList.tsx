import { faEye } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement, useContext, useMemo } from "react";
import { IDataTableColumn } from "react-data-table-component";
import { GntJob } from "../../api/models";
import JobWatchContext from "../../contexts/JobWatchContext";
import { useClusterName } from "../../helpers/hooks";
import { durationHumanReadable } from "../../helpers/time";
import CustomDataTable from "../CustomDataTable/CustomDataTable";
import IconButton from "../IconButton/IconButton";
import JobStartedAt from "../JobStartedAt";
import JobStatus from "../JobStatus";
import JobSummary from "../JobSummary/JobSummary";
import PrefixLink from "../PrefixLink";
import styles from "./JobList.module.scss";

interface Props {
  jobs: GntJob[];
}

function JobList({ jobs }: Props): ReactElement {
  const clusterName = useClusterName();
  const { trackJob } = useContext(JobWatchContext);

  const columns: IDataTableColumn<GntJob>[] = useMemo(
    () => [
      {
        name: "ID",
        sortable: true,
        selector: (row) => row.id,
        cell: (row) => <PrefixLink to={`/jobs/${row.id}`}>{row.id}</PrefixLink>,
        width: "120px",
      },
      {
        name: "Status",
        sortable: true,
        selector: (row) => row.status,
        width: "120px",
        cell: (row) => <JobStatus status={row.status} />,
      },
      {
        name: "Summary",
        sortable: true,
        selector: (row) => row.summary,
        cell: (row) => (
          <PrefixLink to={`/jobs/${row.id}`}>
            <JobSummary summary={row.summary} />
          </PrefixLink>
        ),
      },
      {
        name: "Start",
        sortable: true,
        selector: (row) => row.startedAt,
        format: (row) => <JobStartedAt timestamp={row.startedAt} />,
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
      {
        name: "Watch",
        sortable: false,
        width: "120px",
        cell: (row) => (
          <IconButton
            icon={faEye}
            onClick={() => trackJob({ clusterName, id: row.id })}
          />
        ),
      },
    ],
    [trackJob]
  );

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
