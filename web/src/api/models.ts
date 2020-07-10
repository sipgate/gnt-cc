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
  primaryInstances: GntInstance[];
  secondaryInstances: GntInstance[];
}

export interface GntCluster {
  name: string;
  hostname: string;
  description: string;
  port: number;
}
