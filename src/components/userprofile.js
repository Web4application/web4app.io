// src/components/UserProfile.js
import React from 'react';

const UserProfile = ({ user }) => {
  return (
    <div className="user-profile">
      <img src={user.avatar} alt="Profile" className="profile-picture" />
      <h3>{user.username}</h3>
      <p>Role: {user.role}</p>
      <p>Status: {user.status}</p>
    </div>
  );
};

export default UserProfile;
