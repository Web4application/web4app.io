import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Sidebar from './components/Sidebar';
import ChatroomList from './components/ChatroomList';
import ChatRoom from './components/ChatRoom';
import UserProfile from './components/UserProfile';
import './App.css';

const App = () => {
  return (
    <Router>
      <div className="app-container">
        <Sidebar />
        <main className="main-content">
          <Routes>
            <Route path="/" element={<ChatroomList />} />
            <Route path="/profile" element={<UserProfile />} />
            <Route path="/chatroom/:id" element={<ChatRoom />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
};

export default App;
