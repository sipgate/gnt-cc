import RFB from "@novnc/novnc/core/rfb.js";
import React, { ReactElement, useEffect, useRef, useState } from "react";
import LoadingIndicator from "../LoadingIndicator/LoadingIndicator";
import styles from "./VNCConsole.module.scss";

type Props = {
  url: string;
};

enum ConnectionState {
  DISCONNECTED,
  CONNECTING,
  CONNECTED,
}

const createRFB = (element: HTMLElement, url: string): RFB => {
  return new RFB(element, url, {
    wsProtocols: ["binary", "base64"],
  });
};

let rfb: RFB | null = null;

const VNCConsole = ({ url }: Props): ReactElement => {
  const vncContainer = useRef<HTMLDivElement>(null);

  const [connectionState, setConnectionState] = useState(
    ConnectionState.DISCONNECTED
  );
  const [connectionError, setConnectionError] = useState<string | null>(null);

  const attemptConnection = (): void => {
    if (vncContainer.current === null) {
      throw new Error("VNC render Container is null");
    }

    setConnectionError("");
    setConnectionState(ConnectionState.CONNECTING);

    rfb = createRFB(vncContainer.current, url);

    rfb.addEventListener("connect", () =>
      setConnectionState(ConnectionState.CONNECTED)
    );
    rfb.addEventListener("securityfailure", ({ detail: { reason } }) => {
      setConnectionError(`Error while connecting: ${reason}`);
      setConnectionState(ConnectionState.DISCONNECTED);
    });
    rfb.addEventListener("disconnect", ({ detail: { clean } }) => {
      if (!clean) {
        setConnectionError("Unexpected disconnect");
      }
      setConnectionState(ConnectionState.DISCONNECTED);
    });
  };

  useEffect(() => {
    if (rfb) {
      console.warn("RFB already initialized");
      return;
    }

    if (vncContainer.current !== null) {
      attemptConnection();

      return () => {
        rfb?.disconnect();
      };
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [url]);

  const renderConnectionError = (): ReactElement => (
    <div className={styles.error}>
      <span>{connectionError}</span>
    </div>
  );

  const renderLoadingIndicator = (): ReactElement => (
    <div className={styles.loading}>
      <LoadingIndicator />
    </div>
  );

  return (
    <div className={styles.vncConsole}>
      <div style={{ height: "100vh" }} ref={vncContainer} />
      {connectionState === ConnectionState.CONNECTING &&
        renderLoadingIndicator()}
      {connectionError && renderConnectionError()}
    </div>
  );
};

export default VNCConsole;
