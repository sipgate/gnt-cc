import React, { PropsWithChildren, ReactElement } from "react";
import styles from "./CardGrid.module.scss";

function CardGrid({ children }: PropsWithChildren<unknown>): ReactElement {
  return <div className={styles.root}>{children}</div>;
}

type SectionProps = {
  headline: string;
};

function Section({
  children,
  headline,
}: PropsWithChildren<SectionProps>): ReactElement {
  return (
    <section>
      <h2>{headline}</h2>
      {children}
    </section>
  );
}

CardGrid.Section = Section;

export default CardGrid;
