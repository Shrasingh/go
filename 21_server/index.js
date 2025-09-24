const express = require('express'); // loading express module
const app = express();// creating an express app
const port = 8000; // server will run on this 

app.use(express.json()); // middleware 
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
  res.status(200).send('Hello World!');
});

app.post('/data', (req, res) => {
  let myjson = req.body;
  res.status(200).json(myjson);
});


app.post('/json', (req, res) => {
  const myjson = req.body;
  res.status(200).send(JSON.stringify(myjson));
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
