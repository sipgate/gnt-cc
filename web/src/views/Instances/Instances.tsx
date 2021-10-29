import React, { ReactElement } from "react";
import { useApi } from "../../api";
import { GntInstance } from "../../api/models";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import InstanceList from "../../components/InstanceList/InstanceList";
import { useClusterName } from "../../helpers/hooks";

interface InstancesResponse {
  cluster: string;
  numberOfInstances: number;
  instances: GntInstance[];
}

const Instances = (): ReactElement => {
  const clusterName = useClusterName();

  const [apiProps] = useApi<InstancesResponse>(
    `clusters/${clusterName}/instances`
  );

  return (
    <ContentWrapper>
      <ApiDataRenderer<InstancesResponse>
        {...apiProps}
        render={({ instances }) => <InstanceList instances={instances} />}
      />
    </ContentWrapper>
  );
};

export default Instances;
