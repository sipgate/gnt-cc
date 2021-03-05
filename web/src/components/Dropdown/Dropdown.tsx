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

interface Props {
  label: string;
  icon?: IconDefinition;
}

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
    <div
      className={classNames(styles.dropdown, {
        [styles.expanded]: expanded,
      })}
      onClick={(e) => {
        e.stopPropagation();
        toggle();
      }}
    >
      <div className={styles.current}>
        {icon && <Icon icon={icon} />}
        <span className={styles.label}>{label}</span>
      </div>
      <div className={styles.optionsWrapper}>
        <span className={styles.triangle} />
        <span className={styles.options}>{children}</span>
      </div>
    </div>
  );
}

export default Dropdown;
