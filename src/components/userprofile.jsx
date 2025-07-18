import React, { useState } from "react";
import Avatar from "react-avatar";
import './UserProfile.css';

const UserProfile = () => {
  const [username, setUsername] = useState("Kubu");
  const [email, setEmail] = useState("kubu@example.com");

  return (
    <div className="user-profile">
      <h2>User Profile</h2>
      <Avatar name={username} size="100" round={true} />
      <div className="profile-field">
        <label>Username:</label>
        <input 
          type="text" 
          value={username} 
          onChange={(e) => setUsername(e.target.value)} 
        />
      </div>
      <div className="profile-field">
        <label>Email:</label>
        <input 
          type="email" 
          value={email} 
          onChange={(e) => setEmail(e.target.value)} 
        />
      </div>
    </div>
  );
};

export default UserProfile;
