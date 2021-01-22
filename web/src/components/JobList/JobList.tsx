import React, { ReactElement } from "react";
import styles from "./JobList.module.scss";
import { GntJob } from "../../api/models";
import DataTable, { IDataTableColumn } from "react-data-table-component";

const columns: IDataTableColumn<GntJob>[] = [
  {
    name: "ID",
    sortable: true,
    selector: (row) => row.id,
  },
  {
    name: "Status",
    sortable: true,
    selector: (row) => row.status,
  },
  {
    name: "Summary",
    sortable: true,
    selector: (row) => row.summary,
  },
  {
    name: "Start",
    sortable: true,
    selector: (row) => row.startedAt,
  },
  {
    name: "End",
    sortable: true,
    selector: (row) => row.endedAt,
  },
];

interface Props {
  jobs: GntJob[];
}

function JobList({ jobs }: Props): ReactElement {
  return (
    <div className={styles.jobList}>
      <DataTable<GntJob>
        columns={columns}
        data={jobs}
        keyField="name"
        pagination
        paginationPerPage={20}
        noHeader
      />
    </div>
  );
}

export default JobList;
