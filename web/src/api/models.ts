export interface GntInstance {
  name: string;
  primaryNode: string;
  secondaryNodes: string[];
  cpuCount: number;
  memoryTotal: number;
  isRunning: boolean;
  offersVnc: boolean;
}

export interface GntNode {
  name: string;
  memoryFree: number;
  memoryTotal: number;
  primaryInstances: string[];
  secondaryInstances: string[];
}

export interface GntCluster {
  name: string;
  hostname: string;
  description: string;
  port: number;
}
