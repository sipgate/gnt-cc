import React, { ReactElement } from "react";
import styles from "./InstanceList.module.scss";
import { GntInstance } from "../../api/models";
import DataTable, { IDataTableColumn } from "react-data-table-component";
import Tag from "../../components/Tag/Tag";
import { useApi } from "../../api";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import PrefixLink from "../../components/PrefixLink";
import { useClusterName } from "../../helpers/hooks";

interface InstancesResponse {
  cluster: string;
  number_of_instances: number;
  instances: GntInstance[];
}

const columns: IDataTableColumn<GntInstance>[] = [
  {
    name: "Name",
    sortable: true,
    selector: (row) => row.name,
    // eslint-disable-next-line react/display-name
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
    // eslint-disable-next-line react/display-name
    cell: (row) => (
      <PrefixLink className={styles.link} to={`/nodes/${row.primaryNode}`}>
        {row.primaryNode}
      </PrefixLink>
    ),
  },
  {
    name: "Secondary Nodes",
    // eslint-disable-next-line react/display-name
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

function InstanceList(): ReactElement {
  const clusterName = useClusterName();

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
