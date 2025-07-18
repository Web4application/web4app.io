import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import './index.css';

// Import components
import Sidebar from './components/Sidebar';
import ChatWindow from './components/ChatWindow';
import ChatroomList from './components/ChatroomList';
import UserProfile from './components/UserProfile';

// Home Page (Displays Chatroom List and Chat Window)
const Home = () => {
  return (
    <div className="app-container">
      <Sidebar />
      <main className="main-content">
        <ChatroomList />
        <ChatWindow />
      </main>
    </div>
  );
};

// Individual Chatroom Page
const Chatroom = ({ match }) => {
  const chatroomId = match.params.id;  // Access the chatroom id from the URL
  return (
    <div className="chatroom-container">
      <h2>Chat Room #{chatroomId}</h2>
      <ChatWindow chatroomId={chatroomId} />
    </div>
  );
};

// User Profile Page
const Profile = () => {
  return (
    <div className="profile-container">
      <UserProfile />
    </div>
  );
};

// App Component (Routing Setup)
const App = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/profile" component={Profile} />
        <Route path="/chatroom/:id" component={Chatroom} /> {/* Dynamic route for chatrooms */}
      </Switch>
    </Router>
  );
};

// Render App
ReactDOM.render(<App />, document.getElementById('root'));
