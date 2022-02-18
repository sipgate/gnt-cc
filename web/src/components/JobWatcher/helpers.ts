import { TrackedJob } from "./../../contexts/JobWatchContext";
import { GntJob } from "../../api/models";

export enum WatcherStatus {
  InProgress,
  Succeeded,
  HasFailures,
}

function isJobInProgress(job: GntJob): boolean {
  return job.status !== "error" && job.status !== "success";
}

function areSameJobs(a: GntJob, b: GntJob): boolean {
  return a.clusterName === b.clusterName && a.id === b.id;
}

export function getFinishedJobs(jobs: GntJob[]): GntJob[] {
  return jobs.filter((job) => !isJobInProgress(job));
}

export function getUnfinishedJobs(jobs: GntJob[]): GntJob[] {
  return jobs.filter(isJobInProgress);
}

export function groupJobIdsByCluster(
  jobs: TrackedJob[]
): Record<string, number[]> {
  const map: Record<string, number[]> = {};

  for (const { clusterName, id } of jobs) {
    if (!map[clusterName]) {
      map[clusterName] = [];
    }

    map[clusterName].push(id);
  }

  return map;
}

export function joinJobListsUnique(
  originalList: GntJob[],
  overrideList: GntJob[]
): GntJob[] {
  return [
    ...originalList.filter(
      (aValue) => !overrideList.find((bValue) => areSameJobs(aValue, bValue))
    ),
    ...overrideList,
  ];
}

export function getWatcherStatus(jobs: GntJob[]): WatcherStatus {
  for (const job of jobs) {
    if (isJobInProgress(job)) {
      return WatcherStatus.InProgress;
    }
  }

  for (const job of jobs) {
    if (job.status === "error") {
      return WatcherStatus.HasFailures;
    }
  }

  return WatcherStatus.Succeeded;
}

export function sortJobs(jobs: GntJob[]): GntJob[] {
  return jobs.sort((a, b) => {
    if (isJobInProgress(a) && !isJobInProgress(b)) {
      return -1;
    }

    if (!isJobInProgress(a) && isJobInProgress(b)) {
      return 1;
    }

    return 0;
  });
}
