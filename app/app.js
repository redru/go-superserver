import io from 'socket.io-client';

const ioInstance = io();

ioInstance.use((socket, next) => {
    let token = socket.handshake.query.token;
    if (isValid(token)) {
      return next();
    }
    return next(new Error('authentication error'));
  });
  
  // then
  io.on('connection', (socket) => {
    let token = socket.handshake.query.token;
    // ...
  });
  