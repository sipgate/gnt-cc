import React, { ReactElement } from "react";
import InstanceList from "../../components/InstanceList/InstanceList";
import { useClusterName } from "../../helpers/hooks";
import { useApi } from "../../api";
import { GntInstance } from "../../api/models";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";

interface InstancesResponse {
  cluster: string;
  numberOfInstances: number;
  instances: GntInstance[];
}

const Instances = (): ReactElement => {
  const clusterName = useClusterName();

  const [{ data, isLoading, error }] = useApi<InstancesResponse>(
    `clusters/${clusterName}/instances`
  );

  const renderContent = (): ReactElement => {
    if (isLoading) {
      return <LoadingIndicator />;
    }

    if (!data) {
      return <div>Failed to load: {error}</div>;
    }

    return <InstanceList instances={data.instances} />;
  };

  return <ContentWrapper>{renderContent()}</ContentWrapper>;
};

export default Instances;
