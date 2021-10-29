import { createContext } from "react";

type ThemeContextProps = {
  isDark: boolean;
  toggleTheme: () => void;
};

export default createContext<ThemeContextProps>({
  isDark: false,
  toggleTheme: () => {
    throw new Error("Not implemented");
  },
});
