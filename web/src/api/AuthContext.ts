import { createContext } from "react";

export interface AuthContextProps {
  authToken: string | null;
  setToken: (token: string | null) => void;
}

export default createContext<AuthContextProps>({
  authToken: null,
  setToken: () => {
    throw new Error("Cannot set token");
  },
});
