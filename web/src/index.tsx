import React from "react";
import ReactDOM from "react-dom";
import { createTheme } from "react-data-table-component";
import "./index.scss";
import App from "./App";

createTheme("default", {
  text: {
    primary: "var(--foregroundPrimary)",
    secondary: "var(--foregroundSecondary)",
  },
  background: {
    default: "var(--backgroundPrimary)",
  },
  context: {
    background: "var(--backgroundSecondary)",
    text: "var(--foregroundPrimary)",
  },
  divider: {
    default: "var(--backgroundSecondary)",
  },
  action: {
    button: "rgba(0,0,0,.54)",
    hover: "rgba(0,0,0,.08)",
    disabled: "rgba(0,0,0,.12)",
  },
});

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
