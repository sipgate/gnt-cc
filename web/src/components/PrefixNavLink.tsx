import React, { ReactElement, PropsWithChildren } from "react";
import { NavLink, useParams } from "react-router-dom";
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
  const { clusterName } = useParams();

  if (!clusterName) {
    throw new Error("Could not get clusterName from router params");
  }

  return (
    <NavLink {...rest} to={`/${clusterName}${to}`}>
      {children}
    </NavLink>
  );
};

export default PrefixNavLink;
