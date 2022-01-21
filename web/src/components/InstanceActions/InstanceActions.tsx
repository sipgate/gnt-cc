import { faTerminal } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement, useContext, useState } from "react";
import { HttpMethod, useApi } from "../../api";
import { GntInstance, JobIdResponse } from "../../api/models";
import JobWatchContext from "../../contexts/JobWatchContext";
import Button from "../Button/Button";
import InstanceActionConfirmationModal from "../InstanceActionConfirmationModal/InstanceActionConfirmationModal";
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

type ConfirmationState = {
  actionName: string;
  action: () => Promise<void>;
};

function InstanceActions({ clusterName, instance }: Props): ReactElement {
  const [
    confirmationState,
    setConfirmationState,
  ] = useState<ConfirmationState | null>(null);

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

  function createExecutor(action: () => Promise<string | JobIdResponse>) {
    return async () => {
      const response = await action();

      if (typeof response === "string") {
        alert(`An error occured: ${response}`);
      } else {
        trackJob(response.jobId);
      }
    };
  }

  function onStart() {
    setConfirmationState({
      actionName: "start",
      action: createExecutor(executeStart),
    });
  }

  function onRestart() {
    setConfirmationState({
      actionName: "restart",
      action: createExecutor(executeRestart),
    });
  }

  function onShutdown() {
    setConfirmationState({
      actionName: "shutdown",
      action: createExecutor(executeShutdown),
    });
  }

  async function onActionConfirmed() {
    if (confirmationState !== null) {
      await confirmationState.action();
    }
  }

  return (
    <>
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
      <InstanceActionConfirmationModal
        isVisible={confirmationState !== null}
        onHide={() => setConfirmationState(null)}
        onConfirm={onActionConfirmed}
        actionName={confirmationState?.actionName || ""}
        instanceName={instance.name}
      />
    </>
  );
}

export default InstanceActions;
