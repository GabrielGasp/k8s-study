const express = require('express');
const os = require('os');

const app = express();

app.get('/', (req, res) => {
  const podName = os.hostname().split('-')

  res.send(`Hello from ${podName[podName.length - 1]}`);
});

app.get('/load', (req, res) => {
  const start = Date.now();

  for (let i = 0; i < 1000000000; i++) {
    // Do nothing
  }

  const end = Date.now();

  res.send(`Took ${end - start}ms to complete`);
});

app.get('/health', (req, res) => {
  res.sendStatus(200);
})

app.listen(8080, () => {
  console.log('Listening on port 8080');
});