import React, { ReactElement, ChangeEvent } from "react";
import Icon from "./Icon/Icon";
import { faSquare, faCheckSquare } from "@fortawesome/free-solid-svg-icons";

interface Props {
  checked: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
}

const Checkbox = ({ checked, onChange }: Props): ReactElement => {
  return (
    <>
      <input
        style={{ display: "none" }}
        type="checkbox"
        checked={checked}
        onChange={onChange}
      />
      {!checked && <Icon icon={faSquare} />}
      {checked && <Icon icon={faCheckSquare} />}
    </>
  );
};

export default Checkbox;
