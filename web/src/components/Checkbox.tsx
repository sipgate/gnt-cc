import React, { ReactElement, ChangeEvent } from "react";
import Icon from "./Icon/Icon";
import { faSquare, faCheckSquare } from "@fortawesome/free-solid-svg-icons";
import styled from "styled-components";

interface Props {
  checked: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
}

const StyledInput = styled.input`
  display: none;
`;

const Checkbox = ({ checked, onChange }: Props): ReactElement => {
  return (
    <>
      <StyledInput type="checkbox" checked={checked} onChange={onChange} />
      {!checked && <Icon icon={faSquare} />}
      {checked && <Icon icon={faCheckSquare} />}
    </>
  );
};

export default Checkbox;
