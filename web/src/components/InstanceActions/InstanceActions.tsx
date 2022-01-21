import { faTerminal } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement, useContext } from "react";
import { HttpMethod, useApi } from "../../api";
import { GntInstance, JobIdResponse } from "../../api/models";
import JobWatchContext from "../../contexts/JobWatchContext";
import Button from "../Button/Button";
import PrefixLink from "../PrefixLink";
import styles from "./InstanceActions.module.scss";

type Props = {
  clusterName: string;
  instance: GntInstance;
};

function useInstanceAction(
  clusterName: string,
  instanceName: string,
  action: "start" | "restart" | "shutdown"
) {
  const [, execute] = useApi<JobIdResponse>(
    {
      slug: `/clusters/${clusterName}/instances/${instanceName}/${action}`,
      method: HttpMethod.Post,
    },
    {
      manual: true,
    }
  );

  return execute;
}

function InstanceActions({ clusterName, instance }: Props): ReactElement {
  const { trackJob } = useContext(JobWatchContext);

  const executeStart = useInstanceAction(clusterName, instance.name, "start");
  const executeRestart = useInstanceAction(
    clusterName,
    instance.name,
    "restart"
  );
  const executeShutdown = useInstanceAction(
    clusterName,
    instance.name,
    "shutdown"
  );

  async function executeAndTrackJob(
    action: () => Promise<string | JobIdResponse>
  ) {
    const response = await action();

    if (typeof response === "string") {
      alert(`An error occured: ${response}`);
    } else {
      trackJob(response.jobId);
    }
  }

  async function onStart() {
    await executeAndTrackJob(executeStart);
  }
  async function onRestart() {
    await executeAndTrackJob(executeRestart);
  }
  async function onShutdown() {
    await executeAndTrackJob(executeShutdown);
  }

  return (
    <section className={styles.wrapper}>
      {!instance.isRunning && <Button onClick={onStart} label="Start" />}
      {instance.isRunning && (
        <>
          <Button onClick={onRestart} label="Restart" />
          <Button onClick={onShutdown} label="Shutdown" />
        </>
      )}

      {instance.offersVnc && (
        <PrefixLink to={`/instances/${instance.name}/console`}>
          <Button label="Open Console" icon={faTerminal} />
        </PrefixLink>
      )}
    </section>
  );
}

export default InstanceActions;
