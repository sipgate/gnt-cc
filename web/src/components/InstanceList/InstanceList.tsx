import React, { ReactElement, useState, ChangeEvent, useEffect } from "react";
import styles from "./InstanceList.module.scss";
import { GntInstance } from "../../api/models";
import DataTable, { IDataTableColumn } from "react-data-table-component";
import Tag from "../Tag/Tag";
import PrefixLink from "../PrefixLink";
import Input from "../Input/Input";

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

const filterInstances = (
  instances: GntInstance[],
  filter: string
): GntInstance[] => {
  if (filter === "") {
    return instances;
  }

  return instances.filter((instance) => {
    if (instance.name.includes(filter)) {
      return true;
    }

    if (instance.primaryNode.includes(filter)) {
      return true;
    }

    for (const node of instance.secondaryNodes) {
      if (node.includes(filter)) {
        return true;
      }
    }

    return false;
  });
};

function InstanceList({ instances }: Props): ReactElement {
  const [filter, setFilter] = useState("");
  const [filteredInstances, setFilteredInstances] = useState(instances);

  useEffect(() => setFilteredInstances(filterInstances(instances, filter)), [
    filter,
    instances,
  ]);

  return (
    <div className={styles.instanceList}>
      <Input
        name="instance-filter"
        type="search"
        label="Filter"
        value={filter}
        onChange={(event: ChangeEvent<HTMLInputElement>) =>
          setFilter(event.target.value)
        }
      />
      <DataTable<GntInstance>
        columns={columns}
        data={filteredInstances}
        keyField="name"
        pagination
        paginationPerPage={20}
        noHeader
      />
    </div>
  );
}

export default InstanceList;
