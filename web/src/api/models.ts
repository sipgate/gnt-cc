export interface GntInstance {
  name: string;
  primaryNode: string;
  secondaryNodes: string[];
  cpuCount: number;
  memoryTotal: number;
}

export interface GntNode {
  name: string;
  memoryFree: number;
  memoryTotal: number;
}

export interface GntCluster {
  name: string;
  hostname: string;
  description: string;
  port: number;
}
