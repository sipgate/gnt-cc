import React, { ReactElement, PropsWithChildren } from "react";
import { Link, useParams } from "react-router-dom";

interface Props {
  to: string;
  className?: string;
}

const PrefixLink = ({
  to,
  className,
  children,
}: PropsWithChildren<Props>): ReactElement => {
  const { clusterName } = useParams();

  if (!clusterName) {
    throw new Error("Could not get clusterName from router params");
  }

  return (
    <Link to={`/${clusterName}${to}`} className={className}>
      {children}
    </Link>
  );
};

export default PrefixLink;
