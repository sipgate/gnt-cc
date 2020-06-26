import React, { ReactElement, useEffect, useState } from "react";
import Toggle from "react-toggle";
import "react-toggle/style.css";
import styles from "./ThemeToggle.module.scss";

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
    <Toggle
      className={styles.toggle}
      checked={isDark}
      onChange={(toggleEvent) => setIsDark(toggleEvent.target.checked)}
      icons={{ checked: "🌙", unchecked: "🔆" }}
      aria-label="Dark mode"
    />
  );
};
