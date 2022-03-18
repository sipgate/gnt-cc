import React from "react";
import ReactDOM from "react-dom";
import { createTheme } from "react-data-table-component";
import "./index.scss";
import App from "./App";

createTheme("default", {
  text: {
    primary: "var(--color-emphasis-high)",
    secondary: "var(--color-emphasis-medium)",
  },
  background: {
    default: "var(--color-elevation-low)",
  },
  context: {
    background: "var(--color-elevation-medium)",
    text: "var(--color-emphasis-high)",
    hover: "var(--color-emphasis-high)",
  },
  divider: {
    default: "var(--color-separator)",
  },
  button: {
    default: "var(--color-emphasis-medium)",
    focus: "var(--color-interaction-background)",
    hover: "var(--color-interaction-background)",
    disabled: "var(--color-emphasis-low)",
  },
  action: {
    button: "var(--color-emphasis-medium)",
    hover: "var(--color-emphasis-high)",
    disabled: "var(--color-separator)",
  },
  highlightOnHover: {
    default: "var(--color-interaction-background)",
    text: "var(--color-emphasis-high)",
  },
  sortFocus: {
    default: "var(--color-emphasis-high)",
  },
});

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
