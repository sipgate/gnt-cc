import React, {
    useState,
    ReactElement,
    useEffect,
    PropsWithChildren,
  } from "react";
  import styles from "./Dropdown.module.scss";
  import { IconDefinition } from "@fortawesome/fontawesome-svg-core";
  import { classNameHelper } from "../../helpers";
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
  
    const classNames = classNameHelper([
      styles.dropdown,
      {
        [styles.expanded]: expanded,
        [styles.hasIcon]: !!icon,
      },
    ]);
  
    return (
      <div
        className={classNames}
        onClick={(e) => {
          e.stopPropagation();
          toggle();
        }}
      >
        <div className={styles.current}>
          <span className={styles.label}>{label}</span>
          {icon && <Icon icon={icon} />}
        </div>
        <div className={styles.optionsWrapper}>
          <span className={styles.triangle} />
          <span className={styles.options}>{children}</span>
        </div>
      </div>
    );
  }
  
  export default Dropdown;
  