import { useEffect, useState } from "react";
import JobCard from "./JobCard";
import styles from "./Jobs.module.css";
import NavBar from "../pages/NavBar";
const URL = "https://dqcf6trx-10000.inc1.devtunnels.ms/home";
const Jobs = () => {
  const [jobs, setjobs] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  useEffect(function () {
    async function fetchJobs() {
      try {
        setIsLoading(true);
        const res = await fetch(URL);
        const data = await res.json();
        console.log(data);
        setjobs(data);
      } catch (err) {
        console.log("There is some error in fetching the request.");
      } finally {
        setIsLoading(false);
      }
    }
    fetchJobs();
  }, []);
  return (
    <div>
      <NavBar />
      <div className={styles.horline}></div>
      {isLoading && (
        <div className={styles.container}>
          <div className={styles.spinner}>
            <svg className={styles.spinner} viewBox="0 0 50 50">
              <circle className={styles.path} cx="25" cy="25" r="20"></circle>
            </svg>
          </div>
        </div>
      )}
      {!isLoading && (
        <ul className={styles.cardbox}>
          {jobs?.map((job) => (
            <JobCard key={job.id} job={job} />
          ))}
        </ul>
      )}
    </div>
  );
};
export default Jobs;
