import classNames from "classnames";
import React, { ReactElement, PropsWithChildren } from "react";
import { NavLink } from "react-router-dom";
import { useClusterName } from "../helpers/hooks";
interface Props {
  to: string;
  className?: string;
  activeClassName?: string;
  exact?: boolean;
}

const PrefixNavLink = ({
  to,
  children,
  className,
  activeClassName,
  exact,
}: PropsWithChildren<Props>): ReactElement => {
  const clusterName = useClusterName();

  return (
    <NavLink
      className={({ isActive }) =>
        isActive ? classNames(className, activeClassName) : className
      }
      to={`/${clusterName}${to}`}
      end={exact}
    >
      {children}
    </NavLink>
  );
};

export default PrefixNavLink;
