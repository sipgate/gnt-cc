import classNames from "classnames";
import React, { PropsWithChildren, ReactElement } from "react";
import { Link } from "react-router-dom";
import styles from "./TabBar.module.scss";

const TabBar = ({ children }: PropsWithChildren<unknown>): ReactElement => {
  return <div className={styles.tabBar}>{children}</div>;
};

type TabProps = {
  label: string;
  to: string;
  isActive?: boolean;
};

const Tab = ({ label, to, isActive }: TabProps) => (
  <Link
    to={to}
    className={classNames(styles.tab, {
      [styles.active]: isActive,
    })}
    key={label}
  >
    {label}
  </Link>
);

TabBar.Tab = Tab;

export default TabBar;
