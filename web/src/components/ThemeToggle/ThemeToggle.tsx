import React, { ReactElement, useEffect, useState } from "react";
import styles from "./ThemeToggle.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSun, faMoon } from "@fortawesome/free-solid-svg-icons";
import { classNameHelper } from "../../helpers";

const DARK_CLASS = "dark";

export const ThemeToggle = (): ReactElement => {
  /*const systemPrefersDark = useMediaQuery(
    {
      query: "(prefers-color-scheme: dark)",
    },
    undefined,
    (prefersDark) => {
      setIsDark(prefersDark);
    }
  );*/

  const [isDark, setIsDark] = useState(false);

  useEffect(() => {
    if (isDark) {
      document.documentElement.classList.add(DARK_CLASS);
    } else {
      document.documentElement.classList.remove(DARK_CLASS);
    }
  }, [isDark]);

  return (
    <div
      className={classNameHelper([styles.toggle, { [styles.isDark]: isDark }])}
      onClick={() => setIsDark(!isDark)}
    >
      <span className={styles.knob} />
      <span className={styles.light}>
        <FontAwesomeIcon icon={faSun} />
      </span>
      <span className={styles.dark}>
        <FontAwesomeIcon icon={faMoon} />
      </span>
    </div>
  );
};
