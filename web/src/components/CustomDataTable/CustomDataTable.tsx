import { faCaretDown } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import DataTable, { IDataTableProps } from "react-data-table-component";
import Icon from "../Icon/Icon";
import styles from "./CustomDataTable.module.scss";
export default function CustomDataTable<T>(
  props: IDataTableProps<T>
): ReactElement {
  return (
    <DataTable<T>
      pagination
      paginationPerPage={20}
      noHeader
      highlightOnHover
      sortIcon={<Icon className={styles.sortIcon} icon={faCaretDown} />}
      {...props}
    />
  );
}
