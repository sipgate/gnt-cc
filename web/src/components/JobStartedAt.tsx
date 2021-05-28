import React, { ReactElement } from "react";
import { unixToDate } from "../helpers/time";

type Props = {
  timestamp: number;
};

function JobStartedAt({ timestamp }: Props): ReactElement {
  if (timestamp < 0) {
    return <span>-</span>;
  }

  const date = unixToDate(timestamp);

  return <span title={date.toLocaleString()}>{date.toLocaleTimeString()}</span>;
}

export default JobStartedAt;
