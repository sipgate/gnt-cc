import React, { ReactElement, ChangeEvent, useMemo } from "react";
import styles from "./InstanceList.module.scss";
import { GntInstance } from "../../api/models";
import DataTable, { IDataTableColumn } from "react-data-table-component";
import Tag from "../Tag/Tag";
import PrefixLink from "../PrefixLink";
import Input from "../Input/Input";
import FilterCheckbox from "../FilterCheckbox/FilterCheckbox";
import Button from "../Button/Button";
import { useFilter, filterInstances } from "./filters";
import Icon from "../Icon/Icon";
import { faTerminal } from "@fortawesome/free-solid-svg-icons";

const columns: IDataTableColumn<GntInstance>[] = [
  {
    name: "Name",
    sortable: true,
    selector: (row) => row.name,
    cell: (row) => (
      <PrefixLink
        className={`${styles.link} ${styles.name}`}
        to={`/instances/${row.name}`}
      >
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
    cell: (row) => `${row.memoryTotal} MB`,
    width: "90px",
    sortable: true,
    selector: (row) => row.memoryTotal,
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
  const [
    { filter, filterFields },
    { setFilter, setFilterFields, reset },
  ] = useFilter();

  const filteredInstances = useMemo(
    () => filterInstances(instances, filter, filterFields),
    [instances, filter, filterFields]
  );

  return (
    <div className={styles.instanceList}>
      <div className={styles.filterSettings}>
        <Input
          name="instance-filter"
          type="search"
          label="Filter"
          value={filter}
          onChange={(event: ChangeEvent<HTMLInputElement>) =>
            setFilter(event.target.value)
          }
        />
        <span className={styles.filterListLabel}>Filter by</span>
        <div className={styles.filterList}>
          <FilterCheckbox
            className={styles.filterCheckbox}
            label="Name"
            checked={filterFields.name}
            onChange={(checked) =>
              setFilterFields({
                ...filterFields,
                name: checked,
              })
            }
          />
          <FilterCheckbox
            className={styles.filterCheckbox}
            label="Primary Node"
            checked={filterFields.primaryNode}
            onChange={(checked) =>
              setFilterFields({
                ...filterFields,
                primaryNode: checked,
              })
            }
          />
          <FilterCheckbox
            className={styles.filterCheckbox}
            label="Secondary Node(s)"
            checked={filterFields.secondaryNodes}
            onChange={(checked) =>
              setFilterFields({
                ...filterFields,
                secondaryNodes: checked,
              })
            }
          />
        </div>
        <Button
          className={styles.filterResetButton}
          label="Reset filters"
          onClick={reset}
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
