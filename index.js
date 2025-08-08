const express = require('express');
const app = express();
app.use(express.json());
app.use(express.static('public'));

const urlMap = new Map();

app.post('/shorten', async (req, res) => {
    const { original } = req.body;
    if (!original || !isValidUrl(original)) {
        return res.status(400).json({ error: 'Valid URL required' });
    }
    const shortId = Math.random().toString(36).substring(2, 8);
    const short = `https://${req.get('host')}/${shortId}`;
    urlMap.set(shortId, original);
    res.json({ short });
});

app.get('/:shortId', (req, res) => {
    const original = urlMap.get(req.params.shortId);
    if (original) return res.redirect(original);
    res.status(404).send('Not Found');
});

function isValidUrl(url) {
    try {
        new URL(url);
        return true;
    } catch {
        return false;
    }
}

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));