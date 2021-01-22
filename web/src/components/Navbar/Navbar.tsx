import React, { useContext, ReactElement } from "react";
import styles from "./Navbar.module.scss";
import logo from "../../assets/ganeti_logo.svg";
import Button from "../Button/Button";
import { faSignOutAlt } from "@fortawesome/free-solid-svg-icons";
import AuthContext from "../../api/AuthContext";
import ClusterSelector from "../ClusterSelector/ClusterSelector";
import PrefixNavLink from "../PrefixNavLink";
import PrefixLink from "../PrefixLink";
import { GntCluster } from "../../api/models";

const links = [
  {
    to: "/",
    label: "Dashboard",
    exactActive: true,
  },
  {
    to: "/instances",
    label: "Instances",
  },
  {
    to: "/nodes",
    label: "Nodes",
  },
  {
    to: "/jobs",
    label: "Jobs",
  },
];

interface Props {
  clusters: GntCluster[];
}

const Navbar = function ({ clusters }: Props): ReactElement {
  const authContext = useContext(AuthContext);

  const logout = () => authContext.setToken(null);

  return (
    <nav className={styles.navbar}>
      <div className={styles.begin}>
        <PrefixLink className={styles.logoContainer} to="/">
          <img className={styles.logo} src={logo} alt="Gnt-CC Logo" />
        </PrefixLink>

        <div className={styles.items}>
          {links.map((link) => (
            <PrefixNavLink
              key={link.to}
              to={link.to}
              activeClassName={styles.active}
              exact={link.exactActive}
              className={styles.item}
            >
              {link.label}
            </PrefixNavLink>
          ))}
        </div>
      </div>
      <div className={styles.end}>
        <ClusterSelector clusters={clusters} />
        <Button
          className={styles.logoutButton}
          label="Logout"
          icon={faSignOutAlt}
          onClick={logout}
        ></Button>
      </div>
    </nav>
  );
};

export default Navbar;
