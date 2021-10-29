import { faMoon, faSun } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import classNames from "classnames";
import React, { ReactElement, useContext } from "react";
import ThemeContext from "../../contexts/ThemeContext";
import styles from "./ThemeToggle.module.scss";

export const ThemeToggle = (): ReactElement => {
  const { isDark, toggleTheme } = useContext(ThemeContext);

  return (
    <div
      className={classNames(styles.toggle, { [styles.isDark]: isDark })}
      onClick={toggleTheme}
    >
      <span className={styles.knob} />
      <span>
        <FontAwesomeIcon icon={faSun} />
      </span>
      <span>
        <FontAwesomeIcon icon={faMoon} />
      </span>
    </div>
  );
};
