import React, { ReactElement, PropsWithChildren, useState } from "react";
import JobWatchContext from "../contexts/JobWatchContext";

export default function JobWatchProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [trackedJobs, setTrackedJobs] = useState<number[]>([]);

  function trackJob(jobID: number) {
    if (trackedJobs.includes(jobID)) {
      return;
    }

    setTrackedJobs([...trackedJobs, jobID]);
  }

  function untrackJob(jobID: number) {
    setTrackedJobs(trackedJobs.filter((id) => id !== jobID));
  }

  return (
    <JobWatchContext.Provider
      value={{
        trackedJobs,
        trackJob,
        untrackJob,
      }}
    >
      {children}
    </JobWatchContext.Provider>
  );
}
