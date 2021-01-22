import React, { ReactElement } from "react";
import { useParams } from "react-router-dom";
import { useClusterName } from "../../helpers/hooks";
import { GntInstance, GntNode } from "../../api/models";
import { useApi } from "../../api";
import LoadingIndicator from "../../components/LoadingIndicator/LoadingIndicator";
import Hero from "../../components/Hero/Hero";
import InstanceList from "../../components/InstanceList/InstanceList";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import Button from "../../components/Button/Button";

interface NodeResponse {
  node: GntNode;
  primaryInstances: GntInstance[];
  secondaryInstances: GntInstance[];
}

const NodeDetail = (): ReactElement => {
  const { nodeName } = useParams<{ nodeName: string }>();
  const clusterName = useClusterName();

  const [{ data, isLoading, error }] = useApi<NodeResponse>(
    `clusters/${clusterName}/nodes/${nodeName}`
  );

  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (!data) {
    return <div>Failed to load: {error}</div>;
  }

  const { node } = data;

  return (
    <>
      <Hero title={node.name}>
        <Button label="Empty Node" danger />
      </Hero>

      <ContentWrapper>
        <div>
          <p>Primary Instances on {nodeName}</p>
          <InstanceList instances={data.primaryInstances} />
        </div>

        <div>
          <p>Secondary Instances on {nodeName}</p>
          <InstanceList instances={data.secondaryInstances} />
        </div>
      </ContentWrapper>
    </>
  );
};

export default NodeDetail;
