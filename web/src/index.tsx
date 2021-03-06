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
    hover: "var(--colorEmphasisHigh)",
  },
  divider: {
    default: "var(--colorSeparator)",
  },
  button: {
    default: "var(--colorEmphasisMedium)",
    focus: "var(--colorInteractionBackground)",
    hover: "var(--colorInteractionBackground)",
    disabled: "var(--colorEmphasisLow)",
  },
  action: {
    button: "var(--colorEmphasisMedium)",
    hover: "var(--colorEmphasisHigh)",
    disabled: "var(--colorSeparator)",
  },
  highlightOnHover: {
    default: "var(--colorInteractionBackground)",
    text: "var(--colorEmphasisHigh)",
  },
  sortFocus: {
    default: "var(--colorEmphasisHigh)",
  },
});

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
