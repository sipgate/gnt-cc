import { createContext } from "react";

export type AuthContextProps = {
  username: string | null;
  setUsername: (username: string | null) => void;
};

export default createContext<AuthContextProps>({
  username: null,
  setUsername: () => {
    throw new Error("not implemented");
  },
});
