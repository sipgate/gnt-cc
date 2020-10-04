import {
  faClipboard,
  faPowerOff,
  faTerminal,
  faUnlink,
} from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { classNameHelper } from "../../helpers";
import Icon from "../Icon/Icon";
import styles from "./VNCControl.module.scss";

type Props = {
  onShutdown?: () => void;
  onCtrlAltDel?: () => void;
  onClipboardPaste?: () => void;
  onDisconnect?: () => void;
  isConnected: boolean;
};

const VNCControl = ({
  onClipboardPaste,
  onCtrlAltDel,
  onDisconnect,
  onShutdown,
  isConnected,
}: Props): ReactElement => {
  const className = classNameHelper([
    styles.wrapper,
    {
      [styles.connected]: isConnected,
    },
  ]);

  return (
    <div className={className}>
      <button
        disabled={!isConnected}
        className={styles.button}
        onClick={onClipboardPaste}
        title="Send clipboard"
      >
        <Icon icon={faClipboard} />
      </button>
      <button
        disabled={!isConnected}
        className={styles.button}
        onClick={onCtrlAltDel}
        title="Send ctrl+alt+del"
      >
        <Icon icon={faTerminal} />
      </button>
      <button
        disabled={!isConnected}
        className={styles.button}
        onClick={onDisconnect}
        title="Disconnect"
      >
        <Icon icon={faUnlink} />
      </button>
      <button
        disabled={!isConnected}
        className={styles.button}
        onClick={onShutdown}
        title="Power off"
      >
        <Icon icon={faPowerOff} />
      </button>
    </div>
  );
};

export default VNCControl;
