import { createContext } from "react";
type JobWatchContextProps = {
  trackedJobs: number[];
  trackJob: (jobID: number) => void;
};
export default createContext<JobWatchContextProps>({
  trackedJobs: [],
  trackJob: () => {
    throw new Error("Not implemented");
  },
});
