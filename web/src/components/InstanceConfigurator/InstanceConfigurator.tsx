import classNames from "classnames";
import React, { ReactElement, useState } from "react";
import { GntInstance } from "../../api/models";
import Button from "../Button/Button";
import ConfigurableCard from "../cards/ConfigurableCard/ConfigurableCard";
import styles from "./InstanceConfigurator.module.scss";

interface Props {
  instance: GntInstance;
}

interface ConfigurableState {
  value: number;
  isDirty: boolean;
}

interface ConfigurableActions {
  onChange: (value: number) => void;
  onIncrease: (large: boolean) => void;
  onDecrease: (large: boolean) => void;
  onReset: () => void;
}

const useConfigurableValue = (
  initial: number,
  min: number,
  max: number,
  step: number,
  largeStep: number
): [ConfigurableState, ConfigurableActions] => {
  const initialState = {
    value: initial,
    isDirty: false,
  };

  const [state, setState] = useState<ConfigurableState>(initialState);

  const onChange = (value: number): void => {
    setState({
      value,
      isDirty: value !== initial,
    });
  };

  const onReset = (): void => {
    setState(initialState);
  };

  const onIncrease = (large: boolean): void => {
    const value = Math.min(max, state.value + (large ? largeStep : step));
    setState({
      value,
      isDirty: value !== initial,
    });
  };

  const onDecrease = (large: boolean): void => {
    const value = Math.max(min, state.value - (large ? largeStep : step));
    setState({
      value,
      isDirty: value !== initial,
    });
  };

  return [state, { onChange, onReset, onIncrease, onDecrease }];
};

const InstanceConfigurator = ({ instance }: Props): ReactElement => {
  const [memory, memoryActions] = useConfigurableValue(
    instance.memoryTotal,
    512,
    8192,
    512,
    1024
  );

  const [cpuCount, cpuCountActions] = useConfigurableValue(
    instance.cpuCount,
    1,
    12,
    1,
    2
  );

  const resetAll = () => {
    memoryActions.onReset();
    cpuCountActions.onReset();
  };

  const isDirty = memory.isDirty || cpuCount.isDirty;

  return (
    <div className={styles.instanceConfigurator}>
      <div className={styles.header}>
        <div
          className={classNames(styles.actions, { [styles.hidden]: !isDirty })}
        >
          <Button
            className={styles.action}
            label="Reset all"
            onClick={resetAll}
          />
          <Button className={styles.action} label="Apply all" primary />
        </div>
      </div>
      <div className={styles.cards}>
        <ConfigurableCard
          title="Memory"
          unit="MB"
          {...memory}
          {...memoryActions}
          onAccept={() => alert("save")}
        />
        <ConfigurableCard
          title="vCPUs"
          unit=""
          {...cpuCount}
          {...cpuCountActions}
          onAccept={() => alert("save")}
        />
      </div>
    </div>
  );
};

export default InstanceConfigurator;
