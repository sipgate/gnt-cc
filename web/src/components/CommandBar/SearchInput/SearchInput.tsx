import { faSearch, faSpinner } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import Icon from "../../Icon/Icon";
import styles from "./SearchInput.module.scss";

type Props = {
  value: string;
  isLoading: boolean;
  onChange: (value: string) => void;
};

export default function SearchInput({
  value,
  isLoading,
  onChange,
}: Props): ReactElement {
  return (
    <div className={styles.searchInput}>
      <Icon icon={isLoading ? faSpinner : faSearch} spin={isLoading} />

      <input
        type="search"
        autoFocus
        placeholder="Type something ..."
        value={value}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
}
