import { faChevronRight } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React, { ReactElement } from "react";
import { Link, useLocation } from "react-router-dom";
import styles from "./Breadcrumbs.module.scss";

const Breadcrumbs = (): ReactElement => {
  const location = useLocation();

  const crumbs = location.pathname.split("/").filter((crumb) => !!crumb);

  const crumbElements = [];
  let link = "";

  for (let i = 0; i < crumbs.length; i++) {
    link += `/${crumbs[i]}`;

    crumbElements.push(
      <Link to={link} key={i} className={styles.crumb}>
        {crumbs[i]}
      </Link>
    );

    if (i < crumbs.length - 1) {
      crumbElements.push(
        <FontAwesomeIcon
          key={`chevron-${i}`}
          className={styles.chevron}
          icon={faChevronRight}
        />
      );
    }
  }

  return (
    <div className={styles.root}>
      <span className={styles.crumbs}>{crumbElements}</span>
    </div>
  );
};

export default Breadcrumbs;
