import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import './ChatRoom.css';

const ChatRoom = () => {
  const { id } = useParams();
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');

  useEffect(() => {
    // Mock data for demonstration
    const mockMessages = {
      1: [
        { id: 1, user: 'Alice', text: 'Hello General Chat!' },
        { id: 2, user: 'Bob', text: 'Hey Alice!' },
      ],
      2: [{ id: 1, user: 'Charlie', text: 'Anyone into coding?' }],
      3: [{ id: 1, user: 'Dave', text: 'What’s the latest game everyone’s playing?' }],
      4: [{ id: 1, user: 'Emma', text: 'Any new songs to share?' }],
    };

    setMessages(mockMessages[id] || []);
  }, [id]);

  const handleSend = () => {
    if (newMessage.trim() === '') return;

    const newMsg = {
      id: messages.length + 1,
      user: 'You',
      text: newMessage,
    };

    setMessages([...messages, newMsg]);
    setNewMessage('');
  };

  return (
    <div className="chat-room">
      <h2>Chat Room #{id}</h2>
      <div className="messages">
        {messages.length ? (
          messages.map((msg) => (
            <div key={msg.id} className="message">
              <strong>{msg.user}:</strong> {msg.text}
            </div>
          ))
        ) : (
          <p>No messages in this room yet.</p>
        )}
      </div>
      <div className="input-container">
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          placeholder="Type a message..."
        />
        <button onClick={handleSend}>Send</button>
      </div>
    </div>
  );
};

export default ChatRoom;
