import React, { ReactElement } from "react";
import styles from "./Tag.module.scss";

interface Props {
  label: string;
}

const Tag = ({ label }: Props): ReactElement => {
  return <span className={styles.tag}>{label}</span>;
};

export default Tag;
