import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import VNCConsole from "../../components/VNCConsole/VNCConsole";

const VNCPage = (): ReactElement => {
  return (
      <>
      <VNCConsole />
      </>
  );
};

export default VNCPage;
