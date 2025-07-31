const Tabs = ({ activeTab, setActiveTab }) => {
  const tabs = ['Published', 'Drafts', 'Trashed'];

  return (
    <div className="flex justify-center space-x-2 bg-white p-2 rounded shadow">
      {tabs.map((tab) => (
        <button
          key={tab}
          onClick={() => setActiveTab(tab)}
          className={`px-4 py-2 rounded ${
            activeTab === tab
              ? 'bg-purple-700 text-white'
              : 'bg-gray-100 hover:bg-gray-200'
          }`}
        >
          {tab}
        </button>
      ))}
    </div>
  );
};

export default Tabs;
