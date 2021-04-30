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
};

export interface GntInstance {
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
}

export interface GntNode {
  name: string;
  memoryFree: number;
  memoryTotal: number;
  diskTotal: number;
  diskFree: number;
  cpuCount: number;
  primaryInstances: string[];
  secondaryInstances: string[];
  isMaster: boolean;
  isMasterCandidate: boolean;
  isMasterCapable: boolean;
  isDrained: boolean;
  isOffline: boolean;
  isVMCapable: boolean;
}

export interface GntCluster {
  name: string;
  hostname: string;
  description: string;
  port: number;
}

export interface GntJob {
  id: number;
  summary: string;
  receivedAt: number;
  startedAt: number;
  endedAt: number;
  status: string;
}
