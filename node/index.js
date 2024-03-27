import express from 'express';
const app = express();
const port = process.env.PORT ?? 1111;

app.use(morgan('dev'));

app.get('/', (req, res) => {
  res.send(`Add ID: ${port}, Home Page`);
});

app.get('/app1', (req, res) => {
  res.send(`Add ID: ${port}, App1 Page`);
});

app.get('/app2', (req, res) => {
  res.send(`Add ID: ${port}, App2 Page`);
});


app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
