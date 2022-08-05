import { createContext } from "react";

type SearchBarContextProps = {
  isVisible: boolean;
  toggleVisibility: () => void;
  setVisible: (visible: boolean) => void;
};

export default createContext<SearchBarContextProps>({
  isVisible: false,
  toggleVisibility: () => {
    throw new Error("Not implemented");
  },
  setVisible: (boolean) => {
    throw new Error("Not implemented");
  },
});
