// src/components/MessageInput.js
import React, { useState } from 'react';
import { socket } from '../socket';

const MessageInput = () => {
  const [message, setMessage] = useState('');

  const sendMessage = (e) => {
    e.preventDefault();
    socket.emit('message', { username: 'User', text: message });
    setMessage('');
  };

  return (
    <form onSubmit={sendMessage}>
      <input 
        type="text" 
        value={message} 
        onChange={(e) => setMessage(e.target.value)} 
        placeholder="Type a message" 
      />
      <button type="submit">Send</button>
    </form>
  );
};

export default MessageInput;
