// server.js (inside the /profile route)

app.get('/profile', async (req, res) => {
  if (!req.isAuthenticated()) {
    return res.redirect('/');
  }

  try {
    const { accessToken } = req.user;
    const guildsResponse = await axios.get('https://discord.com/api/v10/users/@me/guilds', {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    res.json({
      user: req.user,
      guilds: guildsResponse.data, // Display connected guilds
    });
  } catch (error) {
    console.error('Error fetching Discord data:', error);
    res.status(500).send('Error fetching Discord data');
  }
});
