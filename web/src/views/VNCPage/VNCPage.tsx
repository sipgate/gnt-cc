import React, { ReactElement, useContext } from "react";
import { useParams } from "react-router-dom";
import AuthContext from "../../api/AuthContext";
import VNCConsole from "../../components/VNCConsole/VNCConsole";
import { useClusterName } from "../../helpers/hooks";

const VNCPage = (): ReactElement => {
  const clusterName = useClusterName();
  const { instanceName } = useParams<{ instanceName: string }>();
  const { authToken } = useContext(AuthContext);

  const url = `ws://localhost:8000/v1/clusters/${clusterName}/instances/${instanceName}/console?token=${authToken}`;

  return <VNCConsole url={url} />;
};

export default VNCPage;
