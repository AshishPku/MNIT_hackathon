import { NavLink } from "react-router-dom";
import Logo from "../components/Logo";
import styles from "./NavBar.module.css";
const NavBar = () => {
  return (
    <nav className={styles.container}>
      <Logo />
      <ul className={styles.navContainer}>
        <li>
          <NavLink className={styles.item} to="/jobs">
            Jobs
          </NavLink>
        </li>
        <li>
          <NavLink className={styles.item} to="/counceling">
            Counceling
          </NavLink>
        </li>
        <li>
          <NavLink className={styles.item} to="/savedjob">
            Savedjob
          </NavLink>
        </li>
      </ul>
    </nav>
  );
};

export default NavBar;
