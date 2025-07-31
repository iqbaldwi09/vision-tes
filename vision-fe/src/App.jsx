import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import ArticleList from './pages/ArticleList';
import EditArticle from './components/EditModal';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<ArticleList />} />
        <Route path="/edit/:id" element={<EditArticle />} />
      </Routes>
    </Router>
  );
}

export default App;
