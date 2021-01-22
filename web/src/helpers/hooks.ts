import { useParams } from "react-router-dom";

export const useClusterName = (): string => {
  const { clusterName } = useParams<{ clusterName: string }>();

  if (!clusterName) {
    throw new Error("Cannot get cluster name from router params.");
  }

  return clusterName;
};
