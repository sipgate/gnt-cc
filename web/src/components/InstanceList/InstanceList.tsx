import React, { ReactElement } from "react";
import styles from "./InstanceList.module.scss";
import { GntInstance } from "../../api/models";
import DataTable, { IDataTableColumn } from "react-data-table-component";
import Tag from "../Tag/Tag";
import PrefixLink from "../PrefixLink";

const columns: IDataTableColumn<GntInstance>[] = [
  {
    name: "Name",
    sortable: true,
    selector: (row) => row.name,
    cell: (row) => (
      <PrefixLink className={styles.link} to={`/instances/${row.name}`}>
        {row.name}
      </PrefixLink>
    ),
  },
  {
    name: "Primary Node",
    sortable: true,
    selector: (row) => row.primaryNode,
    cell: (row) => (
      <PrefixLink className={styles.link} to={`/nodes/${row.primaryNode}`}>
        {row.primaryNode}
      </PrefixLink>
    ),
  },
  {
    name: "Secondary Nodes",
    cell: (row) => (
      <div>
        <PrefixLink
          className={styles.link}
          to={`/nodes/${row.secondaryNodes[0]}`}
        >
          {row.secondaryNodes[0]}
        </PrefixLink>
        {row.secondaryNodes.length > 1 && (
          <Tag label={`+${row.secondaryNodes.length - 1}`} />
        )}
      </div>
    ),
  },
  {
    name: "vCPUs",
    width: "60px",
    selector: (row) => row.cpuCount,
    sortable: true,
    right: true,
  },
  {
    name: "Memory",
    cell: (row) => `${row.memoryTotal} MB`,
    width: "90px",
    sortable: true,
    selector: (row) => row.memoryTotal,
  },
];

interface Props {
  instances: GntInstance[];
}

function InstanceList({ instances }: Props): ReactElement {
  return (
    <div className={styles.instanceList}>
      <DataTable<GntInstance>
        columns={columns}
        data={instances}
        keyField="name"
        noHeader
      />
    </div>
  );
}

export default InstanceList;
