import React, { ReactElement } from "react";
import StatusBadge, { BadgeStatus } from "./StatusBadge/StatusBadge";

function getBadgeStatus(status: string): BadgeStatus {
  if (status === "success") {
    return BadgeStatus.SUCCESS;
  }
  if (["queued", "waiting", "canceling"].includes(status)) {
    return BadgeStatus.WARNING;
  }
  if (status === "error") {
    return BadgeStatus.FAILURE;
  }

  return BadgeStatus.PRIMARY;
}

type Props = {
  status: string;
};

function JobStatus({ status }: Props): ReactElement {
  const badgeStatus = getBadgeStatus(status);

  return <StatusBadge status={badgeStatus}>{status}</StatusBadge>;
}

export default JobStatus;
