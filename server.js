const express = require('express');
const axios = require('axios');
const bodyParser = require('body-parser');
require('dotenv').config();

const app = express();
const port = 3000;

app.use(bodyParser.json()); // To parse JSON request bodies

// Webhook URL - replace this with your actual webhook URL from Discord
const DISCORD_WEBHOOK_URL = 'https://discord.com/api/webhooks/1262587975582089276/0yb7BZaAfjLplEOQSQxR23_ETlbiR5wVpLyrddSITqFsulyABp_OQ466MPhVMVYwbNUi';

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
  console.log(`Server is running at http://localhost:${port}`);
});
