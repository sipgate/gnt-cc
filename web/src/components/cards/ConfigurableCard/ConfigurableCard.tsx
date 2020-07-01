import React, { ReactElement } from "react";
import styles from "./ConfigurableCard.module.scss";
import Card from "../Card/Card";
import Button from "../../Button/Button";
import { classNameHelper } from "../../../helpers";
import {
  faCheck,
  faUndo,
  faMinus,
  faPlus,
} from "@fortawesome/free-solid-svg-icons";

interface Props {
  title: string;
  unit: string;
  value: number;
  isDirty: boolean;
  onChange: (value: number) => void;
  onIncrease: (large: boolean) => void;
  onDecrease: (large: boolean) => void;
  onAccept: () => void;
  onReset: () => void;
}

const ConfigurableCard = ({
  onAccept,
  value,
  isDirty,
  onReset,
  onIncrease,
  onDecrease,
  unit,
  title,
}: Props): ReactElement => {
  const getAcceptButton = () => (
    <Button
      className={classNameHelper([
        styles.interaction,
        { [styles.hidden]: !isDirty },
      ])}
      icon={faCheck}
      onClick={onAccept}
    />
  );

  const getResetButton = () => (
    <Button
      className={classNameHelper([
        styles.interaction,
        { [styles.hidden]: !isDirty },
      ])}
      icon={faUndo}
      onClick={onReset}
    />
  );

  return (
    <Card
      className={styles.configurableCard}
      leftTitleSlot={getResetButton}
      rightTitleSlot={getAcceptButton}
      noHorizontalPadding
      subtitle={unit}
      title={title}
    >
      <div className={styles.content}>
        <div className={styles.buttonContainer}>
          <Button
            small
            round
            icon={faMinus}
            onClick={(event) => onDecrease(event.ctrlKey)}
          />
        </div>
        <span className={styles.value}>{value}</span>
        <div className={styles.buttonContainer}>
          <Button
            small
            round
            icon={faPlus}
            onClick={(event) => onIncrease(event.ctrlKey)}
          />
        </div>
      </div>
    </Card>
  );
};

export default ConfigurableCard;
