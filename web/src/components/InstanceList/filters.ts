import { GntInstance } from "../../api/models";

export const filterInstances = (
  instances: GntInstance[],
  filter: string
): GntInstance[] => {
  if (filter === "") {
    return instances;
  }

  return instances.filter((instance) => {
    if (instance.name.includes(filter)) {
      return true;
    }

    if (instance.primaryNode.includes(filter)) {
      return true;
    }

    for (const node of instance.secondaryNodes) {
      if (node.includes(filter)) {
        return true;
      }
    }

    return false;
  });
};
