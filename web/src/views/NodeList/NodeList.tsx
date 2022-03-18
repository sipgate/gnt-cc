import React, { ReactElement } from "react";
import { TableColumn } from "react-data-table-component";
import { useApi } from "../../api";
import { GntNode } from "../../api/models";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import CustomDataTable from "../../components/CustomDataTable/CustomDataTable";
import MemoryUtilisation from "../../components/MemoryUtilisation/MemoryUtilisation";
import PrefixLink from "../../components/PrefixLink";
import StatusBadge, {
  BadgeStatus,
} from "../../components/StatusBadge/StatusBadge";
import CustomColorBadge from "../../CustomColorBadge/CustomColorBadge";
import { prettyPrintMiB } from "../../helpers";
import { getColorForString } from "../../helpers/colors";
import { useClusterName } from "../../helpers/hooks";
import styles from "./NodeList.module.scss";

interface NodesResponse {
  nodes: GntNode[];
  numberOfNodes: number;
  cluster: string;
}

const columns: TableColumn<GntNode>[] = [
  {
    id: "name",
    name: "Name",
    sortable: true,
    selector: (row) => row.name,
    width: "30%",
    cell: (row) => (
      <>
        <PrefixLink className={styles.name} to={`/nodes/${row.name}`}>
          {row.name}
        </PrefixLink>
        {row.isMaster && (
          <StatusBadge status={BadgeStatus.PRIMARY} className={styles.badge}>
            Master
          </StatusBadge>
        )}
        {row.isOffline && (
          <StatusBadge status={BadgeStatus.FAILURE} className={styles.badge}>
            Offline
          </StatusBadge>
        )}
      </>
    ),
  },
  {
    id: "group",
    name: "Group",
    sortable: true,
    selector: (row) => row.groupName,
    cell: (row) => (
      <CustomColorBadge color={getColorForString(row.groupName)}>
        {row.groupName}
      </CustomColorBadge>
    ),
  },
  {
    id: "primary_instances",
    name: "Primary Instances",
    sortable: true,
    selector: (row) => row.primaryInstancesCount,
    cell: (row) => (
      <PrefixLink to={`/nodes/${row.name}/primary-instances`}>
        {row.primaryInstancesCount}
      </PrefixLink>
    ),
  },
  {
    id: "secondary_instances",
    name: "Secondary Instances",
    sortable: true,
    selector: (row) => row.secondaryInstancesCount,
    cell: (row) => (
      <PrefixLink to={`/nodes/${row.name}/secondary-instances`}>
        {row.secondaryInstancesCount}
      </PrefixLink>
    ),
  },
  {
    id: "cpu_cores",
    name: "CPU Cores",
    sortable: true,
    selector: (row) => row.cpuCount,
  },
  {
    id: "memory_utilisation",
    name: "Memory Utilisation",
    sortable: true,
    selector: (row) => row.memoryTotal - row.memoryFree,
    cell: (row) => {
      const memoryUsed = row.memoryTotal - row.memoryFree;

      return (
        <MemoryUtilisation
          memoryInUse={prettyPrintMiB(memoryUsed)}
          memoryTotal={prettyPrintMiB(row.memoryTotal)}
          usagePercent={(memoryUsed / row.memoryTotal) * 100}
        />
      );
    },
  },
  {
    id: "disk_utilisation",
    name: "Disk Utilisation",
    sortable: true,
    selector: (row) => row.diskTotal - row.diskFree,
    cell: (row) => {
      const diskUsed = row.diskTotal - row.diskFree;

      return (
        <MemoryUtilisation
          memoryInUse={prettyPrintMiB(diskUsed)}
          memoryTotal={prettyPrintMiB(row.diskTotal)}
          usagePercent={(diskUsed / row.diskTotal) * 100}
        />
      );
    },
  },
];

function NodeList(): ReactElement {
  const clusterName = useClusterName();

  const [apiProps] = useApi<NodesResponse>(`clusters/${clusterName}/nodes`);

  return (
    <ContentWrapper>
      <ApiDataRenderer<NodesResponse>
        {...apiProps}
        render={({ nodes }) => (
          <CustomDataTable
            columns={columns}
            data={nodes}
            defaultSortFieldId="name"
          />
        )}
      />
    </ContentWrapper>
  );
}

export default NodeList;
