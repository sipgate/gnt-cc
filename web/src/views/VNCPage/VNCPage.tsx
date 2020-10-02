import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import VNCConsole from "../../components/VNCConsole/VNCConsole";

const VNCPage = (): ReactElement => {
  const { host } = useParams();
  const port = 6901;

  return (
    <div>
      <VNCConsole host={host} port={port} />
    </div>
  );
};

export default VNCPage;
