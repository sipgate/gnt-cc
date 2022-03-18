import React, { ReactElement, useContext } from "react";
import {
  BrowserRouter as Router,
  Navigate,
  Outlet,
  Route,
  Routes,
} from "react-router-dom";
import CommandBar from "./components/CommandBar/CommandBar";
import AuthContext from "./contexts/AuthContext";
import AuthProvider from "./providers/AuthProvider";
import JobWatchProvider from "./providers/JobWatchProvider";
import ThemeProvider from "./providers/ThemeProvider";
import ClusterWrapper from "./views/ClusterWrapper";
import Dashboard from "./views/Dashboard/Dashboard";
import InstanceConsole from "./views/InstanceConsole/InstanceConsole";
import InstanceDetail from "./views/InstanceDetail/InstanceDetail";
import Instances from "./views/Instances/Instances";
import JobDetail from "./views/JobDetail/JobDetail";
import Jobs from "./views/Jobs/Jobs";
import Login from "./views/Login/Login";
import NodeDetail from "./views/NodeDetail/NodeDetail";
import NodeList from "./views/NodeList/NodeList";
import NodePrimaryInstances from "./views/NodePrimaryInstances/NodePrimaryInstances";
import NodeSecondaryInstances from "./views/NodeSecondaryInstances/NodeSecondaryInstances";

function AuthenticatedWrapper() {
  const authContext = useContext(AuthContext);
  const isAuthenticated = !!authContext.username;

  return isAuthenticated ? <Outlet /> : <Navigate to="/login" />;
}

function UnauthenticatedWrapper() {
  const authContext = useContext(AuthContext);
  const isAuthenticated = !!authContext.username;

  return isAuthenticated ? <Navigate to="/" /> : <Outlet />;
}

function App(): ReactElement {
  return (
    <div className="App">
      <ThemeProvider>
        <AuthProvider>
          <JobWatchProvider>
            <CommandBar />
            <Router>
              <Routes>
                <Route element={<UnauthenticatedWrapper />}>
                  <Route path="/login" element={<Login />} />
                </Route>

                <Route element={<AuthenticatedWrapper />}>
                  <Route path="/" element={<ClusterWrapper />} />
                  <Route path="/:clusterName" element={<ClusterWrapper />}>
                    <Route index element={<Dashboard />} />
                    <Route
                      path="instances/:instanceName/console"
                      element={<InstanceConsole />}
                    />
                    <Route
                      path="instances/:instanceName"
                      element={<InstanceDetail />}
                    />
                    <Route path="instances" element={<Instances />} />
                    <Route path="nodes/:nodeName/*" element={<NodeDetail />}>
                      <Route
                        index
                        element={<Navigate to={`primary-instances`} />}
                      />

                      <Route
                        path="primary-instances"
                        element={<NodePrimaryInstances />}
                      />

                      <Route
                        path="secondary-instances"
                        element={<NodeSecondaryInstances />}
                      />
                    </Route>
                    <Route path="nodes" element={<NodeList />} />
                    <Route path="jobs/:jobID" element={<JobDetail />} />
                    <Route path="jobs" element={<Jobs />} />
                  </Route>
                </Route>

                <Route path="*" element={<span>404 Not found</span>} />
              </Routes>
            </Router>
          </JobWatchProvider>
        </AuthProvider>
      </ThemeProvider>
    </div>
  );
}

export default App;
