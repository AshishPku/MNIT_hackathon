import { Link } from "react-router-dom";
import NavBar from "../pages/NavBar";
import styles from "./Homepage.module.css";
const Homepage = () => {
  return (
    <main className={styles.homepage}>
      <NavBar />
      <section>
        <h1>From Classroom to Career – Your Journey Starts Here.</h1>
        <h2>
          This platform connects users to diverse job opportunities,
          internships, mentorship programs, and essential career resources,
          creating a seamless pathway from academic achievements to professional
          success.
        </h2>
        <Link className={styles.btn} to="/app">
          Start Search Now
        </Link>
      </section>
    </main>
  );
};
export default Homepage;
