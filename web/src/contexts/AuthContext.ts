import { createContext } from "react";

export type AuthContextProps = {
  username: string | null;
};

export default createContext<AuthContextProps>({
  username: null,
});
