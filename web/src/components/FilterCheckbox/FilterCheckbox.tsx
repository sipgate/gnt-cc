import React, { ChangeEvent, ReactElement } from "react";
import styles from "./FilterCheckbox.module.scss";

interface FilterCheckboxProps {
  label: string;
  checked: boolean;
  onChange: (newValue: boolean) => void;
  className?: string;
}

function FilterCheckbox({
  label,
  checked,
  onChange,
  className,
}: FilterCheckboxProps): ReactElement {
  return (
    <label className={[styles.filterCheckbox, className].join(" ")}>
      <span className={styles.label}>{label}</span>
      <input
        type="checkbox"
        checked={checked}
        onChange={(event: ChangeEvent<HTMLInputElement>) =>
          onChange(event.target.checked)
        }
      />
    </label>
  );
}

export default FilterCheckbox;