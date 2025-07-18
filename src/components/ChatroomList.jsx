import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const ChatroomList = () => {
  const [chatrooms, setChatrooms] = useState([
    { id: 1, name: 'General Chat', description: 'Chat with everyone!' },
    { id: 2, name: 'Tech Talk', description: 'Discuss technology and coding' },
    { id: 3, name: 'Gaming', description: 'Talk about your favorite games' },
    { id: 4, name: 'Music', description: 'Share your favorite tunes' },
    // Add more rooms as needed
  ]);

  return (
    <div className="chatroom-list">
      <h3>Chat Rooms</h3>
      <ul>
        {chatrooms.map((chatroom) => (
          <li key={chatroom.id}>
            <Link to={`/chatroom/${chatroom.id}`}>
              <div className="chatroom-item">
                <h4>{chatroom.name}</h4>
                <p>{chatroom.description}</p>
              </div>
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ChatroomList;
