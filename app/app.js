import io from 'socket.io-client';

const socket = io({
  upgrade: false,
  transports: ['websocket']
});
  
// then
socket.on('connection', conn => {
  let token = conn.handshake.query.token;
  console.log(token);
  // ...
});
  