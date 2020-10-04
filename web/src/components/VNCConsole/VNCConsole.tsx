import { faTimes } from "@fortawesome/free-solid-svg-icons";
import RFB, {
  DisconnectCallback,
  SecurityFailureCallback,
} from "@novnc/novnc/core/rfb.js";
import React, { ReactElement, useEffect, useRef, useState } from "react";
import Button from "../Button/Button";
import Icon from "../Icon/Icon";
import LoadingIndicator from "../LoadingIndicator/LoadingIndicator";
import VNCControl from "../VNCControl/VNCControl";
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

  const onConnect = () => setConnectionState(ConnectionState.CONNECTED);
  const onSecurityFailure: SecurityFailureCallback = ({
    detail: { reason },
  }) => {
    setConnectionError(`Error while connecting: ${reason}`);
    setConnectionState(ConnectionState.DISCONNECTED);
  };
  const onDisconnect: DisconnectCallback = ({ detail: { clean } }) => {
    if (!clean) {
      setConnectionError(
        connectionState === ConnectionState.CONNECTED
          ? "Unexpected disconnect"
          : "Cannot establish connection"
      );
    }
    setConnectionState(ConnectionState.DISCONNECTED);
  };

  const attemptConnection = (): void => {
    if (vncContainer.current === null) {
      throw new Error("VNC render Container is null");
    }

    setConnectionError("");
    setConnectionState(ConnectionState.CONNECTING);

    rfb = createRFB(vncContainer.current, url);

    rfb.addEventListener("connect", onConnect);
    rfb.addEventListener("securityfailure", onSecurityFailure);
    rfb.addEventListener("disconnect", onDisconnect);
  };

  useEffect(() => {
    if (vncContainer.current !== null) {
      attemptConnection();

      return () => {
        rfb?.removeEventListener("connect", onConnect);
        rfb?.removeEventListener("securityfailure", onSecurityFailure);
        rfb?.removeEventListener("disconnect", onDisconnect);
        rfb?.disconnect();
      };
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [url]);

  const renderConnectionError = (): ReactElement => (
    <div className={styles.error}>
      <span>{connectionError}</span>
      <span
        className={styles.errorClose}
        onClick={() => setConnectionError(null)}
      >
        <Icon icon={faTimes} />
      </span>
    </div>
  );

  const renderConnecting = (): ReactElement => (
    <div className={styles.connecting}>
      <LoadingIndicator />
    </div>
  );

  const renderDisconnected = (): ReactElement => (
    <div className={styles.disconnected}>
      <Button label="Reconnect" round onClick={() => attemptConnection()} />
    </div>
  );

  return (
    <div className={styles.vncConsole}>
      <VNCControl isConnected={connectionState === ConnectionState.CONNECTED} />
      <div className={styles.viewer}>
        <div ref={vncContainer} />
        {connectionState === ConnectionState.CONNECTING && renderConnecting()}
        {connectionError && renderConnectionError()}
        {(connectionError ||
          connectionState === ConnectionState.DISCONNECTED) &&
          renderDisconnected()}
      </div>
    </div>
  );
};

export default VNCConsole;
