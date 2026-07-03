/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express');
const app = express();
const port = 3000;

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Home Route
app.get('/', (req, res) => {
    res.status(200).send("Welcome to LearnCodeOnline server");
});

// GET Route
app.get('/get', (req, res) => {
    res.status(200).json({
        message: "Hello from LearnCodeOnline.in"
    });
});

// POST Route (JSON)
app.post('/post', (req, res) => {
    let myJson = req.body; // Your JSON
    res.status(200).send(myJson);
});

// POST Route (Form Data)
app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
});

// Start Server
app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});