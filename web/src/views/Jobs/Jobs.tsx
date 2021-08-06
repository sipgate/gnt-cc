import { faCircleNotch, faRedoAlt } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React, { ReactElement, useEffect } from "react";
import { useApi } from "../../api";
import { GntJob } from "../../api/models";
import Button from "../../components/Button/Button";
import ContentWrapper from "../../components/ContentWrapper/ContentWrapper";
import JobList from "../../components/JobList/JobList";
import { useClusterName } from "../../helpers/hooks";
import styles from "./Jobs.module.scss";

const REFRESH_INTERVAL = 15000;

interface JobResponse {
  jobs: GntJob[];
}

const Jobs = (): ReactElement => {
  const clusterName = useClusterName();
  const [{ data, isLoading, error }, reload] = useApi<JobResponse>(
    `clusters/${clusterName}/jobs`
  );

  useEffect(() => {
    const intervalID = setInterval(reload, REFRESH_INTERVAL);
    return () => clearInterval(intervalID);
  }, []);

  const renderContent = (): ReactElement | null => {
    if (isLoading && !data) {
      return null;
    }

    if (!data) {
      return <div>Failed to load: {error}</div>;
    }

    return (
      <>
        <JobList jobs={data.jobs} />
      </>
    );
  };

  return (
    <ContentWrapper>
      <div className={styles.refreshControl}>
        <Button label="Refresh" icon={faRedoAlt} onClick={reload}></Button>
        {isLoading && (
          <span>
            Refreshing
            <FontAwesomeIcon
              className={styles.refreshIcon}
              spin
              icon={faCircleNotch}
            />
          </span>
        )}
      </div>
      {renderContent()}
    </ContentWrapper>
  );
};

export default Jobs;
