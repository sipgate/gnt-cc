import React from "react";
import { ReactElement } from "react";
import Button from "../components/Button/Button";
import Modal from "../components/Modal/Modal";
import styles from "./VNCCtrlAltDelConfirmModal.module.scss";

type Props = {
  isVisible: boolean;
  instanceName: string;
  onConfirm: () => void;
  onHide: () => void;
};

export default function VNCCtrlAltDelConfirmModal({
  isVisible,
  instanceName,
  onConfirm,
  onHide,
}: Props): ReactElement {
  return (
    <Modal hideModal={onHide} isVisible={isVisible}>
      <p>Are you sure, you would like to send Ctrl + Alt + Del to</p>

      <p className={styles.serverName}>{instanceName}</p>

      <p>This might trigger a reboot.</p>

      <div className={styles.buttons}>
        <Button
          onClick={() => {
            onConfirm();
            onHide();
          }}
          label="Confirm"
          danger
        />
        <Button onClick={onHide} label="Cancel" />
      </div>
    </Modal>
  );
}
