import { useState, useEffect } from 'react';
import axios from 'axios';

const Home = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [podcasts, setPodcasts] = useState([]);
  const [loading, setLoading] = useState(true); // Set loading to true initially
  const [noResults, setNoResults] = useState(false);

  const fetchPodcasts = async (search = '') => {
    setLoading(true);
    try {
      const response = await axios.get(`/api/podcasts${search}`);
      setPodcasts(response.data);
      setNoResults(response.data.length === 0);
    } catch (error) {
      console.error('Error fetching podcasts:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    // Fetch all podcasts on initial load
    fetchPodcasts();
  }, []);

  const handleInputChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const handleSearch = () => {
    // Fetch podcasts based on search term
    fetchPodcasts(searchTerm ? `?search=${searchTerm}` : ''); // Include search term only if it exists
  };

  return (
      <div>
        <h1>Podcasts</h1>
        <div>
          <input
              type="text"
              value={searchTerm}
              onChange={handleInputChange}
              placeholder="Search podcasts..."
          />
          <button onClick={handleSearch}>Search</button>
        </div>
        {loading && <p>Loading podcasts...</p>}
        {noResults && <p>No podcasts found for the given search term.</p>}
        {podcasts.length > 0 && (
            <table>
              <thead>
              <tr>
                <th>Title</th>
                <th>Description</th>
                <th>Category</th>
              </tr>
              </thead>
              <tbody>
              {podcasts.map((podcast) => (
                  <tr key={podcast.id}>
                    <td>{podcast.title}</td>
                    <td>{podcast.description}</td>
                    <td>{podcast.categoryName}</td>
                  </tr>
              ))}
              </tbody>
            </table>
        )}
      </div>
  );
};

export default Home;