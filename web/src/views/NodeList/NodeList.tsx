import React, { ReactElement } from "react";
import { IDataTableColumn } from "react-data-table-component";
import { useApi } from "../../api";
import { GntNode } from "../../api/models";
import Badge, { BadgeStatus } from "../../components/Badge/Badge";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import CustomDataTable from "../../components/CustomDataTable/CustomDataTable";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import MemoryUtilisation from "../../components/MemoryUtilisation/MemoryUtilisation";
import PrefixLink from "../../components/PrefixLink";
import { prettyPrintMiB } from "../../helpers";
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
    cell: (row) => (
      <>
        <PrefixLink
          className={`${styles.link} ${styles.name}`}
          to={`/nodes/${row.name}`}
        >
          {row.name}
        </PrefixLink>
        {row.isMaster && (
          <Badge status={BadgeStatus.PRIMARY} className={styles.badge}>
            Master
          </Badge>
        )}
      </>
    ),
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
