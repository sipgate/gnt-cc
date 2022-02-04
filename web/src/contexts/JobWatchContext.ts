import { createContext } from "react";

export type TrackedJob = {
  clusterName: string;
  id: number;
};

type JobWatchContextProps = {
  trackedJobs: TrackedJob[];
  trackJob: (job: TrackedJob) => void;
  untrackJob: (job: TrackedJob) => void;
};

export default createContext<JobWatchContextProps>({
  trackedJobs: [],
  trackJob: () => {
    throw new Error("Not implemented");
  },
  untrackJob: () => {
    throw new Error("Not implemented");
  },
});
