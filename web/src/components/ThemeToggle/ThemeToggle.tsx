import React, { ReactElement, useEffect, useState } from "react";
import styles from "./ThemeToggle.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSun, faMoon } from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";

const DARK_CLASS = "dark";

export const ThemeToggle = (): ReactElement => {
  const [isDark, setIsDark] = useState(false);

  useEffect(() => {
    if (isDark) {
      document.documentElement.classList.add(DARK_CLASS);
    } else {
      document.documentElement.classList.remove(DARK_CLASS);
    }
  }, [isDark]);

  useEffect(() => {
    const setPreferredMode = (prefersDark: boolean) => {
      setIsDark(prefersDark);
    };

    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    setPreferredMode(mediaQuery.matches);
    mediaQuery.addListener(({ matches: prefersDark }) =>
      setPreferredMode(prefersDark)
    );
  }, []);

  return (
    <div
      className={classNames(styles.toggle, { [styles.isDark]: isDark })}
      onClick={() => setIsDark(!isDark)}
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
