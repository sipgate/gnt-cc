import React, { ReactElement } from "react";
import { useOutletContext } from "react-router-dom";
import { GntInstance } from "../../api/models";
import InstanceList from "../../components/InstanceList/InstanceList";

export default function NodePrimaryInstances(): ReactElement {
  const { primaryInstances } = useOutletContext<{
    primaryInstances: GntInstance[];
  }>();

  return (
    <div>
      <InstanceList instances={primaryInstances} />
    </div>
  );
}
