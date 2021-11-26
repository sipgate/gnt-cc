import { faSignOutAlt } from "@fortawesome/free-solid-svg-icons";
import React, { ReactElement } from "react";
import { useHistory } from "react-router-dom";
import { buildApiUrl, HttpMethod } from "../../api";
import { GntCluster } from "../../api/models";
import Button from "../Button/Button";
import ClusterSelector from "../ClusterSelector/ClusterSelector";
import JobWatcher from "../JobWatcher/JobWatcher";
import PrefixNavLink from "../PrefixNavLink";
import { ThemeToggle } from "../ThemeToggle/ThemeToggle";
import styles from "./Navbar.module.scss";

const links = [
  {
    to: "",
    label: "Dashboard",
    exactActive: true,
  },
  {
    to: "/instances",
    label: "Instances",
    exactActive: true,
  },
  {
    to: "/nodes",
    label: "Nodes",
    exactActive: true,
  },
  {
    to: "/jobs",
    label: "Jobs",
    exactActive: true,
  },
];

interface Props {
  clusters: GntCluster[];
}

const Navbar = function ({ clusters }: Props): ReactElement {
  const history = useHistory();

  async function logout() {
    await fetch(buildApiUrl("logout"), { method: HttpMethod.Post });
    history.push("/login");
  }

  return (
    <nav className={styles.navbar}>
      <div className={styles.begin}>
        <ClusterSelector clusters={clusters} />

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
        <JobWatcher />
        <ThemeToggle />
        <Button label="Logout" icon={faSignOutAlt} onClick={logout}></Button>
      </div>
    </nav>
  );
};

export default Navbar;
