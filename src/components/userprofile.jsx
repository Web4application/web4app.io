import React, { useState } from 'react';
import './UserProfile.css';

const UserProfile = () => {
  const [username, setUsername] = useState('Kubu');
  const [email, setEmail] = useState('kubu@example.com');

  const handleSave = () => {
    alert(`Profile updated: ${username}, ${email}`);
  };

  return (
    <div className="user-profile">
      <h2>User Profile</h2>
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
      <button onClick={handleSave}>Save</button>
    </div>
  );
};

export default UserProfile;
