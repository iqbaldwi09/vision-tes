import { useState, useEffect } from 'react';

const EditModal = ({ article, onClose, onSave }) => {
  const [title, setTitle] = useState(article.title);
  const [category, setCategory] = useState(article.category);
  const [content, setContent] = useState(article.content || '');

  useEffect(() => {
    setTitle(article.title);
    setCategory(article.category);
    setContent(article.content || '');
  }, [article]);

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave({ ...article, title, category, content });
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
      <div className="bg-white rounded-xl shadow-xl p-6 w-full max-w-md mx-4">
        <h2 className="text-2xl font-bold mb-5 text-purple-700 flex items-center gap-2">
          📝 Edit Article
        </h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block mb-1 font-semibold text-sm">Title</label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Article Title"
            />
          </div>
          <div>
            <label className="block mb-1 font-semibold text-sm">Category</label>
            <input
              type="text"
              value={category}
              onChange={(e) => setCategory(e.target.value)}
              className="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="e.g. Teknologi"
            />
          </div>
          <div>
            <label className="block mb-1 font-semibold text-sm">Content</label>
            <textarea
              value={content}
              onChange={(e) => setContent(e.target.value)}
              rows={5}
              className="w-full border border-gray-300 rounded px-3 py-2 resize-none focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Write the article content..."
            ></textarea>
          </div>
          <div className="flex justify-end gap-3 pt-2">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 border border-gray-400 rounded hover:bg-gray-200 transition"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="px-4 py-2 bg-purple-600 text-white rounded hover:bg-purple-700 transition"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default EditModal;
