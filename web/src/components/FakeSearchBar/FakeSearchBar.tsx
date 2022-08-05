import React, { PropsWithChildren, ReactElement, useContext } from "react";
import classNames from "classnames";
import styles from "./FakeSearchBar.module.scss";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import Icon from "../Icon/Icon";
import SearchBarContext from "../../contexts/SearchBarContext";
interface Props {
  className?: string;
}

const FakeSearchBar = ({
  className,
}: PropsWithChildren<Props>): ReactElement => {
  const { toggleVisibility } = useContext(SearchBarContext);

  return (
    <div
      className={classNames(styles.root, className)}
      onClick={(e) => {
        e.stopPropagation();
        toggleVisibility();
      }}
    >
      <div className={styles.current}>
        <Icon icon={faSearch} />
        <span className={styles.label}>
          Search instances, nodes and clusters
        </span>
      </div>
    </div>
  );
};

export default FakeSearchBar;
