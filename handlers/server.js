const express = require('express');
const axios = require('axios');
const bodyParser = require('body-parser');
require('dotenv').config();

const app = express();
const port = 3000;

app.use(bodyParser.json()); // To parse JSON request bodies

// Webhook URL - replace this with your actual webhook URL from Discord
const DISCORD_WEBHOOK_URL = 'https://discord.com/oauth2/authorize?client_id=1169709827145089064&permissions=8&response_type=code&redirect_uri=https%3A%2F%2Fdiscord.com%2Foauth2%2Fauthorize%3Fclient_id%3D1208380409814188042&integration_type=0&scope=bot+webhook.incoming+identify+applications.commands.permissions.update+applications.store.update+applications.commands+applications.builds.read+applications.builds.upload+email+connections+guilds+activities.write+presences.write+openid+gateway.connect+payment_sources.country_code';

async function sendWebhookMessage(message) {
  try {
    const response = await axios.post(DISCORD_WEBHOOK_URL, {
      content: message, // The message you want to send to the Discord channel
    });
    return response.data;
  } catch (error) {
    console.error('Error sending webhook message:', error);
    throw error;
  }
}

// Route to handle the POST request from the frontend
app.post('/send-discord-message', async (req, res) => {
  const { message } = req.body; // Get the message from the request body

  if (!message) {
    return res.status(400).json({ error: 'Message is required' });
  }

  try {
    // Send the message to Discord via the webhook
    const discordResponse = await sendWebhookMessage(message);
    res.status(200).json({ success: true, data: discordResponse });
  } catch (error) {
    res.status(500).json({ error: 'Failed to send message' });
  }
});

app.listen(port, () => {
  console.log(`Server is running at http://localhost:${3000}`);
});
