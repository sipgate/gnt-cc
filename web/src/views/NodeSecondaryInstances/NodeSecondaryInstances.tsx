import React, { ReactElement } from "react";
import { GntInstance } from "../../api/models";
import InstanceList from "../../components/InstanceList/InstanceList";

type Props = {
  instances: GntInstance[];
};

export default function NodeSecondaryInstances({
  instances,
}: Props): ReactElement {
  return (
    <div>
      <InstanceList instances={instances} />
    </div>
  );
}
