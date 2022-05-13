import { faSearch, faSpinner } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement, useEffect, useRef } from "react";
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
  const ref = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (ref.current) {
      ref.current.select();
    }
  }, [ref]);

  return (
    <div className={styles.searchInput}>
      <Icon icon={isLoading ? faSpinner : faSearch} spin={isLoading} />

      <input
        ref={ref}
        type="search"
        placeholder="Type something ..."
        value={value}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
}
