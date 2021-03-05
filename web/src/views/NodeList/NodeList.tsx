import React, { ReactElement } from "react";
import styles from "./NodeList.module.scss";
import DataTable, { IDataTableColumn } from "react-data-table-component";
import { GntNode } from "../../api/models";
import { useApi } from "../../api";
import { convertMBToGB } from "../../helpers";
import MemoryUtilisation from "../../components/MemoryUtilisation/MemoryUtilisation";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import { useClusterName } from "../../helpers/hooks";
import PrefixLink from "../../components/PrefixLink";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";

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
      <PrefixLink
        className={`${styles.link} ${styles.name}`}
        to={`/nodes/${row.name}`}
      >
        {row.name}
      </PrefixLink>
    ),
  },
  {
    name: "Memory",
    sortable: true,
    selector: (row) => row.memoryTotal - row.memoryFree,
    cell: (row) => (
      <MemoryUtilisation
        memoryInUse={convertMBToGB(row.memoryTotal - row.memoryFree)}
        memoryTotal={convertMBToGB(row.memoryTotal)}
        memoryUnit="GB"
      />
    ),
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
      <DataTable
        columns={columns}
        data={nodes}
        pagination
        paginationPerPage={20}
        noHeader
        highlightOnHover
      />
    );
  };

  return <ContentWrapper>{renderContent()}</ContentWrapper>;
}

export default NodeList;
