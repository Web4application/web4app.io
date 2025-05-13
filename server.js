// server.js

const express = require('express');
const passport = require('passport');
const { Strategy } = require('passport-discord');
const session = require('express-session');
const axios = require('axios');
require('dotenv').config();

const app = express();
const port = 3000;

passport.use(new Strategy({
  clientID: process.env.DISCORD_CLIENT_ID,
  clientSecret: process.env.DISCORD_CLIENT_SECRET,
  callbackURL: process.env.DISCORD_REDIRECT_URI,
  scope: ['identify', 'guilds'],
}, async (accessToken, refreshToken, profile, done) => {
  // Store user profile and access token
  profile.accessToken = accessToken;
  return done(null, profile);
}));

passport.serializeUser((user, done) => done(null, user));
passport.deserializeUser((user, done) => done(null, user));

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: true,
}));

app.use(passport.initialize());
app.use(passport.session());

app.get('/auth/discord', passport.authenticate('discord'));

app.get('/auth/discord/callback',
  passport.authenticate('discord', { failureRedirect: '/' }),
  (req, res) => {
    res.redirect('/profile');
  });

app.get('/profile', (req, res) => {
  if (!req.isAuthenticated()) {
    return res.redirect('/');
  }
  res.json({
    user: req.user,
    guilds: req.user.guilds, // Show connected servers
  });
});

app.get('/', (req, res) => {
  res.send('<a href="/auth/discord">Login with Discord</a>');
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
