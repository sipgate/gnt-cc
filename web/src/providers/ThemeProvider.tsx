import React, {
  PropsWithChildren,
  ReactElement,
  useEffect,
  useState,
} from "react";
import ThemeContext from "../contexts/ThemeContext";

const LIGHT_CLASS = "light";
const DARK_THEME_ID = "theme-dark";

const savedThemeState = localStorage.getItem(DARK_THEME_ID);
const initialState = savedThemeState
  ? (JSON.parse(savedThemeState) as boolean)
  : false;

export default function ThemeProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [isDark, setIsDark] = useState(initialState);

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

  function setIsDarkPersistent(value: boolean) {
    setIsDark(value);
    localStorage.setItem(DARK_THEME_ID, JSON.stringify(value));
  }

  function toggleTheme() {
    setIsDarkPersistent(!isDark);
  }

  return (
    <ThemeContext.Provider
      value={{
        isDark,
        toggleTheme,
      }}
    >
      {children}
    </ThemeContext.Provider>
  );
}
