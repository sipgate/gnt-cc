import React, { PropsWithChildren, ReactElement } from "react";
import styles from "./SearchResults.module.scss";

type Props = {
  headline: string;
};

export default function SearchResults({
  headline,
  children,
}: PropsWithChildren<Props>): ReactElement {
  return (
    <section className={styles.searchResults}>
      <h1>{headline}</h1>
      {children}
    </section>
  );
}
