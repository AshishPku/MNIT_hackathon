import styles from "./JobCard.module.css";
const JobCard = ({ job }) => {
  return (
    <div className={styles.job_card} key={job.id}>
      <div className={styles.job_title}>{job.title}</div>
      <div className={styles.location}>{job.location}</div>
      <div className={styles.job_type}>{job.type}</div>
      <div className={styles.description}>{job.description}</div>
      <div className={styles.section_title}>Skills</div>
      <div className={styles.skills}>
        <span className={styles.badge}>Risk Management</span>
        <span className={styles.badge}>Budgeting</span>
        <span className={styles.badge}>Financial Reporting</span>
      </div>
      <div className={styles.section_title}>Perks</div>
      <div className={styles.perks}>
        <span className={styles.badge}>Health Insurance</span>
        <span className={styles.badge}>401(k) Matching</span>
      </div>
      <div className={styles.salary}>Salary Range: £59,569 - £103,946 GBP</div>
      <div className={styles.equity}>Equity: 0.1% - 3.1%</div>
      <div className={styles.posted_date}>{job.posted}</div>
      <a href={job.apply} className={styles.apply_btn}>
        Apply Now
      </a>
    </div>
  );
};

export default JobCard;
