import React, {
  useState,
  ReactElement,
  useEffect,
  PropsWithChildren,
} from "react";
import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import Icon from "./Icon/Icon";
import styled from "styled-components";

interface Props {
  label: string;
  icon?: IconDefinition;
}

const HoverOverlay = styled.div`
  position: relative;

  &::after {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    opacity: 0;
    background: var(--overlay);
    transition: opacity 0.1s;
    pointer-events: none;
    border-radius: inherit;
  }

  &:hover::after {
    opacity: 1;
  }
`;

const spacingOuter = "2rem";
const width = "200px";
const borderRadius = "2px";
const transitionDuration = "0.1ms";

const Root = styled.div`
  position: relative;
  width: ${width};
  height: 48px;
  cursor: pointer;
`;

const Current = styled(HoverOverlay)`
  position: relative;
  display: flex;
  align-items: center;
  z-index: 1;
  padding: 0 ${spacingOuter};
  border: 1px solid var(--colorEmphasisLow);
  border-radius: 24px;
  width: 100%;
  height: 100%;
`;

const Label = styled.span<{ extraMargin: boolean }>`
  width: 100%;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  color: inherit;
  margin: 0 ${(props) => (props.extraMargin ? spacingOuter : 0)} 0 0;
`;

const OptionsWrapper = styled.div<{ expanded: boolean }>`
  position: absolute;
  right: 0;
  visibility: ${(props) => (props.expanded ? "visible" : "hidden")};
  top: calc(100% + 1rem);
  min-width: ${width};
  border-top: 0;
  background: var(--colorElevationHigh);
  box-shadow: var(--dropShadowElevationHigh);
  opacity: ${(props) => (props.expanded ? 1 : 0)};
  border-radius: ${borderRadius};
  transition: opacity ${transitionDuration},
    visibility 0s ${(props) => (props.expanded ? 0 : transitionDuration)};
  z-index: 99;
`;

const Triangle = styled.span`
  width: 1rem;
  height: 1rem;
  background: var(--colorElevationHigh);
  position: absolute;
  top: -0.5rem;
  right: ${spacingOuter};
  transform: rotate(45deg);
`;

const Options = styled.div`
  overflow: hidden;
  border-radius: ${borderRadius};
`;

function Dropdown({
  label,
  icon,
  children,
}: PropsWithChildren<Props>): ReactElement {
  const [expanded, setExpanded] = useState(false);

  const handleOutsideClick = () => setExpanded(false);

  const toggle = () => setExpanded(!expanded);

  useEffect(() => {
    window.addEventListener("click", handleOutsideClick);

    return () => {
      window.removeEventListener("click", handleOutsideClick);
    };
  }, []);

  return (
    <Root
      onClick={(e) => {
        e.stopPropagation();
        toggle();
      }}
    >
      <Current>
        <Label extraMargin={!!icon}>{label}</Label>
        {icon && <Icon icon={icon} />}
      </Current>
      <OptionsWrapper expanded={expanded}>
        <Triangle />
        <Options>{children}</Options>
      </OptionsWrapper>
    </Root>
  );
}

export default Dropdown;
