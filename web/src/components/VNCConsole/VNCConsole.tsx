import { faTimes } from "@fortawesome/free-solid-svg-icons";
import RFB, {
  CredentialsRequiredCallback,
  DisconnectCallback,
  SecurityFailureCallback,
} from "@novnc/novnc/core/rfb.js";
import React, { ReactElement, useEffect, useRef, useState } from "react";
import VNCCtrlAltDelConfirmModal from "../../VNCCtrlAltDelConfirmModal/VNCCtrlAltDelConfirmModal";
import Button from "../Button/Button";
import Icon from "../Icon/Icon";
import LoadingIndicator from "../LoadingIndicator/LoadingIndicator";
import VNCControl from "../VNCControl/VNCControl";
import VNCCredentialsPrompt from "../VNCCredentialsPrompt/VNCCredentialsPrompt";
import styles from "./VNCConsole.module.scss";

const supportsPasting = () => navigator.clipboard.readText !== undefined;

type Props = {
  url: string;
  instanceName: string;
};

enum ConnectionState {
  DISCONNECTED,
  AUTHENTICATION_FAILED,
  CONNECTING,
  CONNECTED,
}

type Credentials = {
  username: string;
  password: string;
};

const emptyCredentials: Credentials = {
  username: "",
  password: "vncpassword",
};

const toVNCCredentials = ({
  username,
  password,
}: Credentials): { username?: string; password?: string } => ({
  username: username || undefined,
  password: password || undefined,
});

const createRFB = (
  element: HTMLElement,
  url: string,
  credentials: Credentials
): RFB => {
  return new RFB(element, url, {
    wsProtocols: ["binary"],
    credentials: toVNCCredentials(credentials),
  });
};

let rfb: RFB | null = null;

const VNCConsole = ({ url, instanceName }: Props): ReactElement => {
  const vncContainer = useRef<HTMLDivElement>(null);

  const [connectionState, setConnectionState] = useState(
    ConnectionState.DISCONNECTED
  );
  const [connectionError, setConnectionError] = useState<string | null>(null);
  const [credentials, setCredentials] = useState<Credentials>(emptyCredentials);
  const [confirmingCtrlAltDel, setConfirmingCtrlAltDel] = useState(false);

  const onConnect = () => setConnectionState(ConnectionState.CONNECTED);
  const onSecurityFailure: SecurityFailureCallback = () =>
    setConnectionState(ConnectionState.AUTHENTICATION_FAILED);
  const onDisconnect: DisconnectCallback = ({ detail: { clean } }) => {
    if (!clean) {
      setConnectionError("Unexpected disconnect");
    }
    setConnectionState(ConnectionState.DISCONNECTED);
  };
  const onCredentialsRequired: CredentialsRequiredCallback = () => {
    setConnectionState(ConnectionState.AUTHENTICATION_FAILED);
  };

  const disconnect = () => {
    if (!rfb) {
      return;
    }

    setCredentials(emptyCredentials);
    rfb.removeEventListener("connect", onConnect);
    rfb.removeEventListener("securityfailure", onSecurityFailure);
    rfb.removeEventListener("disconnect", onDisconnect);
    rfb.removeEventListener("credentialsrequired", onCredentialsRequired);
    rfb.disconnect();
  };

  const attemptConnection = (): void => {
    if (vncContainer.current === null) {
      throw new Error("VNC render Container is null");
    }

    setConnectionError(null);
    setConnectionState(ConnectionState.CONNECTING);

    rfb = createRFB(vncContainer.current, url, credentials);

    rfb.addEventListener("connect", onConnect);
    rfb.addEventListener("securityfailure", onSecurityFailure);
    rfb.addEventListener("disconnect", onDisconnect);
    rfb.addEventListener("credentialsrequired", onCredentialsRequired);
  };

  function sendCtrlAltDel() {
    rfb?.sendCtrlAltDel();
  }

  useEffect(() => {
    if (vncContainer.current !== null) {
      attemptConnection();
    }

    return () => disconnect();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [url]);

  useEffect(() => {
    if (rfb && connectionState === ConnectionState.AUTHENTICATION_FAILED) {
      rfb.sendCredentials(toVNCCredentials(credentials));
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [credentials]);

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
    <div className={styles.overlay}>
      <LoadingIndicator />
    </div>
  );

  const renderDisconnected = (): ReactElement => (
    <div className={styles.overlay}>
      <span>Disconnected</span>
      <Button label="Reconnect" round onClick={() => attemptConnection()} />
    </div>
  );

  const renderPasswordPrompt = (): ReactElement => (
    <div className={styles.overlay}>
      <VNCCredentialsPrompt
        initialValue={credentials}
        onConfirm={setCredentials}
      />
    </div>
  );

  return (
    <div className={styles.vncConsole}>
      <div ref={vncContainer} />
      {connectionState === ConnectionState.CONNECTING && renderConnecting()}
      {connectionError !== null && renderConnectionError()}
      {connectionState === ConnectionState.DISCONNECTED && renderDisconnected()}
      {connectionState === ConnectionState.AUTHENTICATION_FAILED &&
        renderPasswordPrompt()}

      <VNCCtrlAltDelConfirmModal
        isVisible={confirmingCtrlAltDel}
        onHide={() => setConfirmingCtrlAltDel(false)}
        onConfirm={sendCtrlAltDel}
        instanceName={instanceName}
      />

      <VNCControl
        isConnected={connectionState === ConnectionState.CONNECTED}
        enablePowerControl={rfb?.capabilities.power || false}
        enablePasting={supportsPasting()}
        onDisconnect={() => disconnect()}
        onCtrlAltDel={() => setConfirmingCtrlAltDel(true)}
        onClipboardPaste={() =>
          navigator.clipboard
            .readText()
            .then(rfb?.clipboardPasteFrom)
            .catch(console.error)
        }
        className={styles.control}
      />
    </div>
  );
};

export default VNCConsole;
