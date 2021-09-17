import React, { ReactElement } from "react";
import { IDataTableColumn } from "react-data-table-component";
import { useApi } from "../../api";
import { GntNode } from "../../api/models";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import CustomDataTable from "../../components/CustomDataTable/CustomDataTable";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
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

const columns: IDataTableColumn<GntNode>[] = [
  {
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
    name: "CPU Cores",
    sortable: true,
    selector: (row) => row.cpuCount,
  },
  {
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

  const [{ data, isLoading, error }] = useApi<NodesResponse>(
    `clusters/${clusterName}/nodes`
  );

  const renderContent = (): ReactElement => {
    if (isLoading) {
      return <LoadingIndicator />;
    }

    if (!data) {
      return <div>Failed to load: {error}</div>;
    }

    const { nodes } = data;

    return (
      <CustomDataTable columns={columns} data={nodes} defaultSortField="name" />
    );
  };

  return <ContentWrapper>{renderContent()}</ContentWrapper>;
}

export default NodeList;
