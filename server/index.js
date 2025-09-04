// server/index.js
const express = require('express');
const http = require('http');
const socketIo = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketIo(server);

io.on('connection', (socket) => {
  console.log('a user connected');
  
  socket.on('message', (msg) => {
    io.emit('message', msg); // broadcast message to all clients
  });
  
  socket.on('disconnect', () => {
    console.log('user disconnected');
  });
});

server.listen(8080, () => {
  console.log('Server is running on http://localhost:8080');
});
