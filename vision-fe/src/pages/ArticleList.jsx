import { useState, useEffect } from 'react';
import { fetchArticles, deleteArticle, updateArticle } from '../api/articleAPI';
import Tabs from '../components/Tabs';
import ArticleTable from '../components/ArticleTable';
import EditModal from '../components/EditModal';

const ArticleList = () => {
  const [activeTab, setActiveTab] = useState('Published');
  const [allArticles, setAllArticles] = useState([]);
  const [selectedArticle, setSelectedArticle] = useState(null);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const normalizeStatus = (status = '') => {
    const s = status.toLowerCase();
    if (s === 'publish' || s === 'published') return 'Published';
    if (s === 'draft' || s === 'drafts') return 'Drafts';
    if (s === 'trash' || s === 'trashed') return 'Trashed';
    return 'Unknown';
  };

  const loadArticles = async () => {
    try {
      const res = await fetchArticles();
      const raw = Array.isArray(res) ? res : res?.data || [];
      const normalized = raw.map(a => ({
        ...a,
        uiStatus: normalizeStatus(a.status)
      }));
      setAllArticles(normalized);
    } catch (err) {
      console.error('Gagal memuat artikel:', err);
      setAllArticles([]);
    }
  };

  useEffect(() => {
    loadArticles();
  }, []);

  const filteredArticles = allArticles.filter(
    article => article.uiStatus === activeTab
  );

  const handleDelete = async (id) => {
    try {
      await deleteArticle(id);
      await loadArticles();
    } catch (err) {
      console.error('Gagal menghapus artikel:', err);
    }
  };

  const handleEdit = (article) => {
    setSelectedArticle(article);
    setIsEditModalOpen(true);
  };

  const handleUpdate = async (updatedArticle) => {
    try {
      await updateArticle(updatedArticle.id, updatedArticle);
      await loadArticles();
      setIsEditModalOpen(false);
      setSelectedArticle(null);
    } catch (err) {
      console.error('Gagal mengupdate artikel:', err);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 px-4 py-10">
      <h1 className="text-3xl font-bold mb-8 flex items-center gap-2 text-purple-800">
        <span role="img" aria-label="doc">📄</span> All Posts
      </h1>

      <Tabs activeTab={activeTab} setActiveTab={setActiveTab} />

      <div className="w-full max-w-4xl mt-6 bg-white shadow-md rounded-lg p-4">
        {filteredArticles.length === 0 ? (
          <div className="text-center text-gray-500 py-10">No articles found.</div>
        ) : (
          <>
            <ArticleTable
              articles={filteredArticles}
              onDelete={handleDelete}
              onEdit={handleEdit}
            />
          </>
        )}
      </div>

      {isEditModalOpen && selectedArticle && (
        <EditModal
          article={selectedArticle}
          onClose={() => setIsEditModalOpen(false)}
          onSave={handleUpdate}
        />
      )}
    </div>
  );
};

export default ArticleList;
