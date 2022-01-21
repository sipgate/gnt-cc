import React, { ReactElement } from "react";
import Button from "../Button/Button";
import Modal from "../Modal/Modal";
import Name from "../Name/Name";
import styles from "./InstanceActionConfirmationModal.module.scss";

type Props = {
  isVisible: boolean;
  actionName: string;
  instanceName: string;
  onConfirm: () => void;
  onHide: () => void;
};

function InstanceActionConfirmationModal({
  isVisible,
  actionName,
  instanceName,
  onConfirm,
  onHide,
}: Props): ReactElement {
  return (
    <Modal hideModal={onHide} isVisible={isVisible}>
      <p>Are you sure you would like to {actionName}</p>
      <Name>{instanceName}</Name>

      <div className={styles.buttons}>
        <Button
          onClick={() => {
            onConfirm();
            onHide();
          }}
          label={actionName}
          primary
        />
        <Button onClick={onHide} label="Cancel" />
      </div>
    </Modal>
  );
}

export default InstanceActionConfirmationModal;
