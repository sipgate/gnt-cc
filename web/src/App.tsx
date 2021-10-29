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
  const [trackedJobs, setTrackedJobs] = useState<number[]>([]);

  return (
    <div className="App">
      <ThemeProvider>
        <AuthProvider>
          <JobWatchContext.Provider
            value={{
              trackJob(jobID) {
                if (trackedJobs.includes(jobID)) {
                  return;
                }

                setTrackedJobs([...trackedJobs, jobID]);
              },
              untrackJob(jobID) {
                setTrackedJobs(trackedJobs.filter((id) => id !== jobID));
              },
              trackedJobs,
            }}
          >
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
          </JobWatchContext.Provider>
        </AuthProvider>
      </ThemeProvider>
    </div>
  );
}

export default App;
