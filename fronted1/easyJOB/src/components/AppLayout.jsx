import Logo from "./Logo";
import styles from "./AppLayout.module.css";
import { useEffect, useState } from "react";
import Search from "./Search";
import JobCard from "./JobCard";
const AppLayout = () => {
  const [query, setQuery] = useState("s");
  const [jobs, setJobs] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  useEffect(() => {
    const controller = new AbortController();
    async function FetchData() {
      try {
        setIsLoading(true);
        const res = await fetch(
          `https://dqcf6trx-10000.inc1.devtunnels.ms/search?title=${query}`,
          { signal: controller.signal }
        );
        if (!res.ok) {
          throw new Error("Something went wrong in fetching the movies.");
        }
        const data = await res.json();
        if (data.Response === "False") {
          throw new Error("Movie not Found");
        }
        setJobs(data);
        setIsLoading(false);
      } catch (err) {
        if (err.name !== "AbortError") {
          console.error(err);
        }
      } finally {
        setIsLoading(false);
      }
    }
    FetchData();
    return function () {
      controller.abort();
    };
  }, [query]);
  return (
    <section>
      <Logo />
      <div className={styles.horline}></div>
      <div className={styles.container}>
        <Search query={query} setQuery={setQuery} />
      </div>
      {isLoading && (
        <div className={styles.container1}>
          <div className={styles.spinner}>
            <svg className={styles.spinner} viewBox="0 0 50 50">
              <circle className={styles.path} cx="25" cy="25" r="20"></circle>
            </svg>
          </div>
        </div>
      )}
      {!isLoading && (
        <div>
          <ul className={styles.cardbox}>
            {jobs?.map((job) => (
              <JobCard key={job.id} job={job} />
            ))}
          </ul>
        </div>
      )}
    </section>
  );
};

export default AppLayout;
