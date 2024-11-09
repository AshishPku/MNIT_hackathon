import { BrowserRouter, Route, Routes } from "react-router-dom";
import Homepage from "./components/Homepage";
import Jobs from "./components/Jobs";
import SavedJob from "./components/SavedJob";
import PostJob from "./components/PostJob";
import AppLayout from "./components/AppLayout";
import PageNotfound from "./components/PageNotfound";
const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Homepage />} />
        <Route path="/jobs" element={<Jobs />} />
        <Route path="/counceling" element={<PostJob />} />
        <Route path="/savedjob" element={<SavedJob />} />
        <Route path="/app" element={<AppLayout />} />
        <Route path="*" element={<PageNotfound />} />
      </Routes>
    </BrowserRouter>
  );
};
export default App;
