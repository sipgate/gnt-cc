import React, { ReactElement } from "react";
import InstanceList from "../../components/InstanceList/InstanceList";
import { useClusterName } from "../../helpers/hooks";
import { useApi } from "../../api";
import { GntInstance } from "../../api/models";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import Hero from "../../components/Hero/Hero";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";

interface InstancesResponse {
  cluster: string;
  number_of_instances: number;
  instances: GntInstance[];
}

const Instances = (): ReactElement => {
  const clusterName = useClusterName();

  const [{ data, isLoading, error }] = useApi<InstancesResponse>(
    `clusters/${clusterName}/instances`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  return (
    <div>
      <Hero title={`Instances on ${clusterName} cluster`} />
      <ContentWrapper>
        <InstanceList instances={data.instances} />
      </ContentWrapper>
    </div>
  );
};

export default Instances;
