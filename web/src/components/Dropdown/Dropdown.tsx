import React, {
  useState,
  ReactElement,
  useEffect,
  PropsWithChildren,
} from "react";
import styles from "./Dropdown.module.scss";
import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
import classNames from "classnames";
import Icon from "../Icon/Icon";

export enum Alignment {
  LEFT = "left",
  CENTER = "center",
  RIGHT = "right",
}

type Props = {
  label?: string;
  icon?: IconDefinition;
  align?: Alignment;
};

function Dropdown({
  label,
  icon,
  align = Alignment.LEFT,
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
    <div
      className={classNames(
        styles.root,
        {
          [styles.expanded]: expanded,
          [styles.hasLabel]: !!label,
        },
        styles[align]
      )}
      onClick={(e) => {
        e.stopPropagation();
        toggle();
      }}
    >
      <div className={styles.current}>
        {icon && <Icon icon={icon} />}
        {label && <span className={styles.label}>{label}</span>}
      </div>
      <div className={styles.optionsWrapper}>
        <span className={styles.triangle} />
        <div className={styles.options}>{children}</div>
      </div>
    </div>
  );
}

export default Dropdown;
