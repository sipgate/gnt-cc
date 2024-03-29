import classNames from "classnames";
import React, { ReactElement } from "react";
import { Link } from "react-router-dom";
import styles from "./SearchResult.module.scss";

type Props = {
  name: string;
  url: string;
  selected?: boolean;
  onClick?: () => void;
};

export default function SearchResult({
  name,
  url,
  selected,
  onClick,
}: Props): ReactElement {
  return (
    <Link
      to={url}
      className={classNames(styles.searchResult, {
        [styles.selected]: selected,
      })}
      onClick={onClick}
      data-selected={selected}
    >
      {name}
    </Link>
  );
}
