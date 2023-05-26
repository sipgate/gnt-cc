import React, { ReactElement } from "react";
import { useOutletContext } from "react-router-dom";
import { GntInstance } from "../../api/models";
import InstanceList from "../../components/InstanceList/InstanceList";

export default function NodeSecondaryInstances(): ReactElement {
  const { secondaryInstances } = useOutletContext<{
    secondaryInstances: GntInstance[];
  }>();

  return (
    <div>
      <InstanceList instances={secondaryInstances} />
    </div>
  );
}
