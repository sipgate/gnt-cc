import React, { ReactElement } from "react";
import { Link } from "react-router-dom";
import styles from "./SearchResult.module.scss";

type Props = {
  name: string;
  url: string;
  onClick?: () => void;
};

export default function SearchResult({
  name,
  url,
  onClick,
}: Props): ReactElement {
  return (
    <Link to={url} className={styles.searchResult} onClick={onClick}>
      {name}
    </Link>
  );
}
