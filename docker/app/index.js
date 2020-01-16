const express = require('express');
const app = express();

app.get('/', (req, res) => {
    res.json({ message: 'Hi everyone' });
});

app.listen(3000, () => {
    console.log('Server is running.');
});
