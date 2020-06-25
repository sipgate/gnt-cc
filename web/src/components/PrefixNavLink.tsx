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
  ...rest
}: PropsWithChildren<Props>): ReactElement => {
  const clusterName = useClusterName();

  return (
    <NavLink {...rest} to={`/${clusterName}${to}`}>
      {children}
    </NavLink>
  );
};

export default PrefixNavLink;
