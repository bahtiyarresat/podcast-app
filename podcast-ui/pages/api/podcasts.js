import axios from 'axios';

export default async (req, res) => {
    const { query } = req;
    const apiUrl = 'http://localhost:8080/podcasts';

    try {
        const response = await axios.get(apiUrl + (query.search || ''));
        res.status(200).json(response.data);
    } catch (error) {
        console.error('Error fetching podcasts:', error);
        res.status(500).json({ error: 'Internal Server Error' });
    }
};