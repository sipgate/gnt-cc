import React, { ReactElement } from "react";
import styles from "./InstanceList.module.scss";
import { useParams } from "react-router-dom";
import { GntInstance } from "../../api/models";
import DataTable, { IDataTableColumn } from "react-data-table-component";
import Tag from "../../components/Tag/Tag";
import { useApi } from "../../api";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import PrefixLink from "../../components/PrefixLink";

interface InstancesResponse {
  cluster: string;
  number_of_instances: number;
  instances: GntInstance[];
}

const columns: IDataTableColumn<GntInstance>[] = [
  {
    name: "Name",
    sortable: true,
    // eslint-disable-next-line react/display-name
    cell: (row) => (
      <PrefixLink to={`/instances/${row.name}`}>{row.name}</PrefixLink>
    ),
    //sortFunction: (rowA, rowB) => rowA.name.localeCompare(rowB.name),
  },
  {
    name: "Primary Node",
    sortable: true,
    // eslint-disable-next-line react/display-name
    cell: (row) => (
      <PrefixLink to={`/nodes/${row.primaryNode}`}>
        {row.primaryNode}
      </PrefixLink>
    ),
  },
  {
    name: "Secondary Nodes",
    sortable: true,
    // eslint-disable-next-line react/display-name
    cell: (row) => (
      <div>
        <PrefixLink to={`/nodes/${row.secondaryNodes[0]}`}>
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
    cell: (row) => row.cpuCount,
    width: "60px",
    right: true,
  },
  {
    name: "Memory",
    cell: (row) => `${row.memoryTotal} MB`,
    width: "90px",
  },
];

function InstanceList(): ReactElement {
  const { clusterName } = useParams();

  if (!clusterName) {
    throw new Error("cluster not found");
  }

  const [{ data, isLoading, error }] = useApi<InstancesResponse>(
    `clusters/${clusterName}/instances`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  const { instances } = data;

  return (
    <div className={styles.instanceList}>
      <DataTable<GntInstance> columns={columns} data={instances}></DataTable>
    </div>
  );
}

export default InstanceList;
