import { faBolt, faRobot } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import Icon from "../Icon/Icon";
import LoadingIndicator from "../LoadingIndicator/LoadingIndicator";
import styles from "./ApiDataRenderer.module.scss";

type Props<T> = {
  data: T | null;
  isLoading: boolean;
  error: string | null;
  render: (data: T) => ReactElement;
};

function renderError(error: string): ReactElement {
  return (
    <div className={styles.error}>
      <header>
        <Icon className={styles.robot} icon={faRobot} />
        <Icon className={styles.bolt} icon={faBolt} />
      </header>
      <div>
        <h1>An error occured</h1>
        <p>{error}</p>
      </div>
    </div>
  );
}

function ApiDataRenderer<T>({
  data,
  isLoading,
  error,
  render,
}: Props<T>): ReactElement {
  if (isLoading) {
    return <LoadingIndicator />;
  }

  if (error) {
    return renderError(error);
  }

  if (!data) {
    return renderError("No data returned from server");
  }

  return render(data);
}

export default ApiDataRenderer;
