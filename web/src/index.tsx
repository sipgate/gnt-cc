import React from "react";
import ReactDOM from "react-dom";
import { createTheme } from "react-data-table-component";
import "./index.scss";
import App from "./App";

createTheme("default", {
  text: {
    primary: "var(--colorEmphasisHigh)",
    secondary: "var(--colorEmphasisMedium)",
  },
  background: {
    default: "var(--colorElevationLow)",
  },
  context: {
    background: "var(--colorElevationMedium)",
    text: "var(--colorEmphasisHigh)",
  },
  divider: {
    default: "var(--colorSeparator)",
  },
  action: {
    button: "var(--colorEmphasisMedium)",
    hover: "var(--colorEmphasisHigh)",
    disabled: "var(--colorSeparator)",
  },
});

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
