import React, { ReactElement, useContext, useState } from "react";
import {
  BrowserRouter as Router,
  Redirect,
  Route,
  RouteProps,
  Switch,
} from "react-router-dom";
import AuthContext from "./api/AuthContext";
import JobWatchContext from "./contexts/JobWatchContext";
import AuthProvider from "./providers/AuthProvider";
import JobWatchProvider from "./providers/JobWatchProvider";
import ThemeProvider from "./providers/ThemeProvider";
import ClusterWrapper from "./views/ClusterWrapper";
import Login from "./views/Login/Login";

export function AuthenticatedRoute(props: RouteProps): ReactElement {
  const authContext = useContext(AuthContext);

  return authContext?.authToken ? (
    <Route {...props} />
  ) : (
    <Redirect to="/login" />
  );
}

function App(): ReactElement {
  return (
    <div className="App">
      <ThemeProvider>
        <AuthProvider>
          <JobWatchProvider>
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
          </JobWatchProvider>
        </AuthProvider>
      </ThemeProvider>
    </div>
  );
}

export default App;
