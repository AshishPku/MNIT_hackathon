import styles from "./Logo.module.css";
import logo from "../assets/image/image/logo.png";
import { Link } from "react-router-dom";
const Logo = () => {
  return (
    <Link to="/">
      <img className={styles.logoimg} src={logo} alt="logo" />
    </Link>
  );
};
export default Logo;
