// src/components/ChatWindow.js
import React, { useState, useEffect } from 'react';
import { socket } from '../socket';

const ChatWindow = () => {
  const [messages, setMessages] = useState([]);
  
  useEffect(() => {
    socket.on('message', (message) => {
      setMessages((prevMessages) => [...prevMessages, message]);
    });

    return () => socket.off('message');
  }, []);
  
  return (
    <div className="chat-window">
      {messages.map((msg, index) => (
        <div key={index}>
          <strong>{msg.username}: </strong>{msg.text}
        </div>
      ))}
    </div>
  );
};

export default ChatWindow;
