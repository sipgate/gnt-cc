export type GntDisk = {
  uuid: string;
  name: string;
  capacity: number;
  template: string;
};

export type GntNic = {
  uuid: string;
  name: string;
  mode: string;
  bridge: string;
  mac: string;
  vlan: string;
};

export type GntInstance = {
  name: string;
  primaryNode: string;
  secondaryNodes: string[];
  cpuCount: number;
  memoryTotal: number;
  isRunning: boolean;
  offersVnc: boolean;
  disks: GntDisk[];
  nics: GntNic[];
  tags: string[];
};

export type GntNode = {
  name: string;
  memoryFree: number;
  memoryTotal: number;
  diskTotal: number;
  diskFree: number;
  cpuCount: number;
  primaryInstancesCount: number;
  secondaryInstancesCount: number;
  isMaster: boolean;
  isMasterCandidate: boolean;
  isMasterCapable: boolean;
  isDrained: boolean;
  isOffline: boolean;
  isVMCapable: boolean;
  groupName: string;
};

export type GntCluster = {
  name: string;
  hostname: string;
  description: string;
  port: number;
};

export type GntJob = {
  id: number;
  clusterName: string;
  summary: string;
  receivedAt: number;
  startedAt: number;
  endedAt: number;
  status: string;
};

export type GntJobLogEntry = {
  serial: number;
  message: string;
  startedAt: number;
};

export type GntJobWithLog = GntJob & {
  log: GntJobLogEntry[];
};

export type JobIdResponse = {
  jobId: number;
};
