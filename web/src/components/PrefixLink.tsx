import React, { ReactElement, PropsWithChildren } from "react";
import { Link } from "react-router-dom";
import { useClusterName } from "../helpers/hooks";

interface Props {
  to: string;
  title?: string;
  className?: string;
}

const PrefixLink = ({
  to,
  title,
  className,
  children,
}: PropsWithChildren<Props>): ReactElement => {
  const clusterName = useClusterName();

  return (
    <Link to={`/${clusterName}${to}`} title={title} className={className}>
      {children}
    </Link>
  );
};

export default PrefixLink;
