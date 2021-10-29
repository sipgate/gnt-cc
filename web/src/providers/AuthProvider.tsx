import React, {
  PropsWithChildren,
  ReactElement,
  useEffect,
  useState,
} from "react";
import AuthContext from "../api/AuthContext";

const STORAGE_TOKEN_KEY = "gnt-cc-token";

type JwtPayload = {
  id: string;
  exp: number;
  orig_iat: number;
};

function parseJwtPayload(token: string | null): JwtPayload | null {
  if (!token) {
    return null;
  }

  return JSON.parse(atob(token.split(".")[1]));
}

export default function AuthProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const storedAuthToken = localStorage.getItem(STORAGE_TOKEN_KEY);
  const [authToken, setAuthToken] = useState(storedAuthToken);

  const tokenPayload = parseJwtPayload(authToken);

  useEffect(() => {
    if (authToken) {
      localStorage.setItem(STORAGE_TOKEN_KEY, authToken);
    } else {
      localStorage.removeItem(STORAGE_TOKEN_KEY);
    }
  }, [authToken]);

  return (
    <AuthContext.Provider
      value={{
        authToken,
        setToken: setAuthToken,
        username: tokenPayload ? tokenPayload.id : null,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
