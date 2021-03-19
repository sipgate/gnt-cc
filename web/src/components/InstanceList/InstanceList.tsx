import { faTerminal } from "@fortawesome/free-solid-svg-icons";
import React, { ChangeEvent, ReactElement, useMemo, useState } from "react";
import { IDataTableColumn } from "react-data-table-component";
import { GntInstance } from "../../api/models";
import { prettyPrintMiB } from "../../helpers";
import Badge, { BadgeStatus } from "../Badge/Badge";
import CustomDataTable from "../CustomDataTable/CustomDataTable";
import Icon from "../Icon/Icon";
import Input from "../Input/Input";
import PrefixLink from "../PrefixLink";
import Tag from "../Tag/Tag";
import { filterInstances } from "./filters";
import styles from "./InstanceList.module.scss";

const columns: IDataTableColumn<GntInstance>[] = [
  {
    name: "Name",
    sortable: true,
    selector: (row) => row.name,
    cell: (row) => (
      <>
        <PrefixLink
          className={`${styles.link} ${styles.name}`}
          to={`/instances/${row.name}`}
        >
          {row.name}
        </PrefixLink>
        {!row.isRunning && (
          <Badge className={styles.badge} status={BadgeStatus.FAILURE}>
            Offline
          </Badge>
        )}
      </>
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
    name: "Secondary Node(s)",
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
    width: "90px",
    sortable: true,
    selector: (row) => row.memoryTotal,
    cell: (row) => prettyPrintMiB(row.memoryTotal),
    right: true,
  },
  {
    name: "",
    cell: (row) =>
      row.offersVnc && (
        <PrefixLink
          className={styles.link}
          title="Open Console"
          to={`/instances/${row.name}/console`}
        >
          <Icon icon={faTerminal} />
        </PrefixLink>
      ),
    width: "48px",
    sortable: false,
  },
];

interface Props {
  instances: GntInstance[];
}

function InstanceList({ instances }: Props): ReactElement {
  const [filter, setFilter] = useState("");

  const filteredInstances = useMemo(() => filterInstances(instances, filter), [
    instances,
    filter,
  ]);

  return (
    <div className={styles.instanceList}>
      <div className={styles.filterSettings}>
        <Input
          name="instance-filter"
          type="search"
          label="Search"
          value={filter}
          onChange={(event: ChangeEvent<HTMLInputElement>) =>
            setFilter(event.target.value)
          }
        />
      </div>

      <div className={styles.filterResults}>
        Showing
        <span className={styles.filterResultCount}>
          {filteredInstances.length}
        </span>
        of
        <span className={styles.filterResultCount}>{instances.length}</span>
        instances
      </div>

      <CustomDataTable<GntInstance>
        columns={columns}
        data={filteredInstances}
        keyField="name"
        defaultSortField="name"
      />
    </div>
  );
}

export default InstanceList;
