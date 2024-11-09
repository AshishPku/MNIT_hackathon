import styles from "./Search.module.css";
const Search = ({ query, setQuery }) => {
  return (
    <input
      className={styles.search}
      value={query}
      type="text"
      placeholder="Search jobs..."
      onChange={(e) => setQuery(e.target.value)}
    ></input>
  );
};
export default Search;
