// src/App.js
import React from 'react';
import Sidebar from './components/Sidebar';
import ChatWindow from './components/ChatWindow';
import MessageInput from './components/MessageInput';

const App = () => {
  return (
    <div className="app">
      <Sidebar />
      <div className="main-content">
        <ChatWindow />
        <MessageInput />
      </div>
    </div>
  );
};

export default App;
