import React, {
  PropsWithChildren,
  ReactElement,
  useEffect,
  useState,
} from "react";
import { buildApiUrl } from "../api";
import AuthContext from "../contexts/AuthContext";

type UserResponse = {
  username: string;
};

export default function AuthProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [username, setUsername] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    fetch(buildApiUrl("user"))
      .then(async (response) => {
        if (response.status === 200) {
          const body = (await response.json()) as UserResponse;
          setUsername(body.username);
        } else if (response.status === 401) {
          setUsername(null);
        } else {
          const body = await response.json();
          setError(body.message ? body.message : "unknown error");
        }

        setIsLoading(false);
      })
      .catch((reason) => {
        setError(reason.message);
        setIsLoading(false);
      });
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <AuthContext.Provider
      value={{
        username,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
