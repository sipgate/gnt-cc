import React, { useContext, useState, useEffect, ReactElement } from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  RouteProps,
  Redirect,
} from "react-router-dom";
import Login from "./views/Login/Login";
import AuthContext from "./api/AuthContext";
import ClusterWrapper from "./views/ClusterWrapper";
import { ThemeToggle } from "./components/ThemeToggle/ThemeToggle";

const STORAGE_TOKEN_KEY = "gnt-cc-token";

export function AuthenticatedRoute(props: RouteProps): ReactElement {
  const authContext = useContext(AuthContext);

  return authContext?.authToken ? (
    <Route {...props} />
  ) : (
    <Redirect to="/login" />
  );
}

interface JwtPayload {
  id: string;
  exp: number;
  orig_iat: number;
}

const parseJwtPayload = (token: string | null): JwtPayload | null => {
  if (!token) {
    return null;
  }

  return JSON.parse(atob(token.split(".")[1]));
};

function App(): ReactElement {
  const storedAuthToken = localStorage.getItem(STORAGE_TOKEN_KEY);
  const [authToken, setAuthToken] = useState(storedAuthToken);

  useEffect(() => {
    if (authToken) {
      localStorage.setItem(STORAGE_TOKEN_KEY, authToken);
    } else {
      localStorage.removeItem(STORAGE_TOKEN_KEY);
    }
  }, [authToken]);

  const tokenPayload = parseJwtPayload(authToken);

  return (
    <div className="App">
      <AuthContext.Provider
        value={{
          setToken: setAuthToken,
          username: tokenPayload ? tokenPayload.id : null,
          authToken,
        }}
      >
        <ThemeToggle />
        <Router>
          <Switch>
            <Route exact path="/login" component={Login} />

            <AuthenticatedRoute
              path="/:clusterName?"
              component={ClusterWrapper}
            />

            <Route render={() => <span>404 Not found</span>}></Route>
          </Switch>
        </Router>
      </AuthContext.Provider>
    </div>
  );
}

export default App;
