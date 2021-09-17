import React, { PropsWithChildren, ReactElement } from "react";
import styles from "./CustomColorBadge.module.scss";

type Props = {
  color: string;
};

function CustomColorBadge({
  children,
  color,
}: PropsWithChildren<Props>): ReactElement {
  return (
    <span
      className={styles.badge}
      style={{
        backgroundColor: color,
        color: getContrastingTextColor(color), // stylelint-disable-line function-name-case
      }}
    >
      {children}
    </span>
  );
}

/* credit goes to David Halford */
function getContrastingTextColor(hex: string) {
  const threshold = 130;

  const hRed = hexToR(hex);
  const hGreen = hexToG(hex);
  const hBlue = hexToB(hex);

  function hexToR(h: string) {
    return parseInt(cutHex(h).substring(0, 2), 16);
  }
  function hexToG(h: string) {
    return parseInt(cutHex(h).substring(2, 4), 16);
  }
  function hexToB(h: string) {
    return parseInt(cutHex(h).substring(4, 6), 16);
  }
  function cutHex(h: string) {
    return h.charAt(0) == "#" ? h.substring(1, 7) : h;
  }

  const cBrightness = (hRed * 299 + hGreen * 587 + hBlue * 114) / 1000;
  if (cBrightness > threshold) {
    return "#000000";
  } else {
    return "#ffffff";
  }
}

export default CustomColorBadge;
