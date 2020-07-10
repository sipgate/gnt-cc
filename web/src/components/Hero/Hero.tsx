import React, { PropsWithChildren, ReactElement } from "react";
import styles from "./Hero.module.scss";

interface Props {
  title: string;
}

const Hero = ({ title, children }: PropsWithChildren<Props>): ReactElement => {
  return (
    <div className={styles.hero}>
      <h1 className={styles.title}>{title}</h1>
      {children}
    </div>
  );
};

export default Hero;
