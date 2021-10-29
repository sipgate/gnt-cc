import React, { ReactElement, useContext, useEffect, useState } from "react";
import {
  BrowserRouter as Router,
  Redirect,
  Route,
  RouteProps,
  Switch,
} from "react-router-dom";
import AuthContext from "./api/AuthContext";
import JobWatchContext from "./contexts/JobWatchContext";
import ThemeContext from "./contexts/ThemeContext";
import ClusterWrapper from "./views/ClusterWrapper";
import Login from "./views/Login/Login";

const STORAGE_TOKEN_KEY = "gnt-cc-token";
const LIGHT_CLASS = "light";
const DARK_THEME_ID = "theme-dark";

const savedThemeState = localStorage.getItem(DARK_THEME_ID);
const initialState = savedThemeState
  ? (JSON.parse(savedThemeState) as boolean)
  : false;

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
  const [trackedJobs, setTrackedJobs] = useState<number[]>([]);
  const [isDark, setIsDark] = useState(initialState);

  useEffect(() => {
    if (isDark) {
      document.documentElement.classList.remove(LIGHT_CLASS);
    } else {
      document.documentElement.classList.add(LIGHT_CLASS);
    }
  }, [isDark]);

  useEffect(() => {
    const setPreferredMode = (prefersDark: boolean) => {
      if (savedThemeState === null) {
        setIsDark(prefersDark);
      }
    };

    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    setPreferredMode(mediaQuery.matches);
    mediaQuery.addListener(({ matches: prefersDark }) =>
      setPreferredMode(prefersDark)
    );
  }, []);

  useEffect(() => {
    if (authToken) {
      localStorage.setItem(STORAGE_TOKEN_KEY, authToken);
    } else {
      localStorage.removeItem(STORAGE_TOKEN_KEY);
    }
  }, [authToken]);

  function setIsDarkPersistent(value: boolean) {
    setIsDark(value);
    localStorage.setItem(DARK_THEME_ID, JSON.stringify(value));
  }

  function toggleTheme() {
    setIsDarkPersistent(!isDark);
  }

  const tokenPayload = parseJwtPayload(authToken);

  return (
    <div className="App">
      <ThemeContext.Provider
        value={{
          isDark,
          toggleTheme,
        }}
      >
        <AuthContext.Provider
          value={{
            setToken: setAuthToken,
            username: tokenPayload ? tokenPayload.id : null,
            authToken,
          }}
        >
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
        </AuthContext.Provider>
      </ThemeContext.Provider>
    </div>
  );
}

export default App;
