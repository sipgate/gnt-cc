import React, { ReactElement, PropsWithChildren, useState } from "react";
import JobWatchContext, { TrackedJob } from "../contexts/JobWatchContext";

function compareTo(a: TrackedJob): (b: TrackedJob) => boolean {
  return (b) => a.clusterName === b.clusterName && a.id === b.id;
}

export default function JobWatchProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [trackedJobs, setTrackedJobs] = useState<TrackedJob[]>([]);

  function trackJob(job: TrackedJob) {
    if (trackedJobs.find(compareTo(job))) {
      return;
    }

    setTrackedJobs([...trackedJobs, job]);
  }

  function untrackJob(job: TrackedJob) {
    setTrackedJobs(trackedJobs.filter((j) => !compareTo(job)(j)));
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
