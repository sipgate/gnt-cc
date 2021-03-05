import {
  faBars,
  faCircle,
  faClipboard,
  faPowerOff,
  faTimes,
  IconDefinition,
} from "@fortawesome/free-solid-svg-icons";
import classNames from "classnames";
import React, { ReactElement, useState } from "react";
import Icon from "../Icon/Icon";
import styles from "./VNCControl.module.scss";

type Props = {
  onShutdown?: () => void;
  onCtrlAltDel?: () => void;
  onClipboardPaste?: () => void;
  onDisconnect?: () => void;
  isConnected: boolean;
  enablePowerControl?: boolean;
  enablePasting?: boolean;
  className?: string;
};

const VNCControl = ({
  onClipboardPaste,
  onCtrlAltDel,
  onDisconnect,
  onShutdown,
  isConnected,
  enablePowerControl,
  enablePasting,
  className,
}: Props): ReactElement => {
  const [expanded, setExpanded] = useState(true);

  const renderAction = ({
    icon,
    label,
    onClick,
  }: {
    icon?: IconDefinition;
    label: string;
    onClick?: () => void;
  }): ReactElement => (
    <div className={styles.action} onClick={onClick}>
      <span className={styles.actionAside}>
        {icon ? <Icon icon={icon} /> : <Icon icon={faCircle} />}
      </span>
      <span className={styles.actionLabel}>{label}</span>
    </div>
  );

  return (
    <div
      className={classNames(className, styles.wrapper, {
        [styles.visible]: isConnected,
        [styles.expanded]: expanded,
      })}
    >
      <span
        className={styles.toggleButton}
        onClick={() => setExpanded(!expanded)}
      >
        <Icon icon={expanded ? faTimes : faBars} />
      </span>

      {expanded && (
        <div className={styles.container} onClick={() => setExpanded(false)}>
          <div className={styles.actions}>
            {enablePasting &&
              renderAction({
                onClick: onClipboardPaste,
                label: "Send clipboard",
                icon: faClipboard,
              })}
            {renderAction({
              onClick: onCtrlAltDel,
              label: "Send ctrl+alt+del",
            })}
            {renderAction({
              onClick: onDisconnect,
              label: "Disconnect",
            })}
            {enablePowerControl &&
              renderAction({
                onClick: onShutdown,
                label: "Shutdown options",
                icon: faPowerOff,
              })}
          </div>
        </div>
      )}
    </div>
  );
};

export default VNCControl;
