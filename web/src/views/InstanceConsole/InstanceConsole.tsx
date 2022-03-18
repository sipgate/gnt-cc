import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { buildWSURL } from "../../api";
import VNCConsole from "../../components/VNCConsole/VNCConsole";
import { useClusterName } from "../../helpers/hooks";

const InstanceConsole = (): ReactElement => {
  const clusterName = useClusterName();
  const { instanceName } = useParams<{ instanceName: string }>();

  const url = buildWSURL(
    `/clusters/${clusterName}/instances/${instanceName}/console`
  );

  return <VNCConsole url={url} instanceName={instanceName || ""} />;
};

export default InstanceConsole;
