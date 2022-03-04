import { faRedoAlt } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useApi } from "../../api";
import { GntJob } from "../../api/models";
import ApiDataRenderer from "../../components/ApiDataRenderer/ApiDataRenderer";
import Button from "../../components/Button/Button";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import JobList from "../../components/JobList/JobList";
import { useClusterName } from "../../helpers/hooks";

interface JobResponse {
  jobs: GntJob[];
}

const Jobs = (): ReactElement => {
  const clusterName = useClusterName();
  const [apiProps, reload] = useApi<JobResponse>(
    `clusters/${clusterName}/jobs`
  );

  return (
    <ContentWrapper>
      <Button label="Refresh" icon={faRedoAlt} onClick={reload}></Button>
      <ApiDataRenderer
        {...apiProps}
        render={({ jobs }) => <JobList jobs={jobs} />}
      />
    </ContentWrapper>
  );
};

export default Jobs;
