import { Trash2, Pencil } from 'lucide-react';
import { useState } from 'react';

const ArticleTable = ({ articles, onDelete, onEdit = () => {} }) => {
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 5;

  const totalPages = Math.ceil(articles.length / itemsPerPage);
  const startIndex = (currentPage - 1) * itemsPerPage;
  const currentArticles = articles.slice(startIndex, startIndex + itemsPerPage);

  const handlePrevious = () => {
    if (currentPage > 1) setCurrentPage(currentPage - 1);
  };

  const handleNext = () => {
    if (currentPage < totalPages) setCurrentPage(currentPage + 1);
  };

  return (
    <div className="w-full flex flex-col items-center">
      <div className="overflow-x-auto w-full">
        <table className="min-w-full text-sm text-center table-auto border border-gray-300 rounded-lg shadow">
          <thead className="bg-purple-100 text-purple-800">
            <tr>
              <th className="px-4 py-2 border">Title</th>
              <th className="px-4 py-2 border">Category</th>
              <th className="px-4 py-2 border">Action</th>
            </tr>
          </thead>
          <tbody className="bg-white">
            {currentArticles.length > 0 ? (
              currentArticles.map((article, index) => (
                <tr key={index} className="hover:bg-purple-50 transition">
                  <td className="px-4 py-2 border">{article.title}</td>
                  <td className="px-4 py-2 border">{article.category}</td>
                  <td className="px-4 py-2 border flex justify-center gap-2">
                    <button
                      className="text-blue-500 hover:text-blue-700"
                      onClick={() => onEdit(article.id)}
                      title="Edit"
                    >
                      <Pencil size={18} />
                    </button>
                    <button
                      className="text-red-500 hover:text-red-700"
                      onClick={() => onDelete(article.id)}
                      title="Delete"
                    >
                      <Trash2 size={18} />
                    </button>
                  </td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan="3" className="px-4 py-6 text-gray-500">
                  No articles to display.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination Controls */}
      {totalPages > 1 && (
        <div className="mt-4 flex items-center gap-3">
          <button
            onClick={handlePrevious}
            disabled={currentPage === 1}
            className="px-3 py-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50"
          >
            Previous
          </button>

          <span className="text-sm text-gray-700">
            Page <strong>{currentPage}</strong> of <strong>{totalPages}</strong>
          </span>

          <button
            onClick={handleNext}
            disabled={currentPage === totalPages}
            className="px-3 py-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50"
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
};

export default ArticleTable;
