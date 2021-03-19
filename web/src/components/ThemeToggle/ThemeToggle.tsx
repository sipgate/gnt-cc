import React, { ReactElement, useEffect, useState } from "react";
import styles from "./ThemeToggle.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSun, faMoon } from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";

const LIGHT_CLASS = "light";
const DARK_THEME_ID = "theme-dark";

const savedThemeState = localStorage.getItem(DARK_THEME_ID);
const initialState = savedThemeState
  ? (JSON.parse(savedThemeState) as boolean)
  : false;

export const ThemeToggle = (): ReactElement => {
  const [isDark, setIsDark] = useState(initialState);

  function setIsDarkPersistent(value: boolean) {
    setIsDark(value);
    localStorage.setItem(DARK_THEME_ID, JSON.stringify(value));
  }

  useEffect(() => {
    if (isDark) {
      document.documentElement.classList.remove(LIGHT_CLASS);
    } else {
      document.documentElement.classList.add(LIGHT_CLASS);
    }
  }, [isDark]);

  useEffect(() => {
    const setPreferredMode = (prefersDark: boolean) => {
      if (savedThemeState === null) {
        setIsDark(prefersDark);
      }
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
      onClick={() => setIsDarkPersistent(!isDark)}
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
