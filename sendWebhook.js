// sendWebhook.js

const axios = require('axios');

async function sendWebhookMessage(webhookURL, message) {
  try {
    const response = await axios.post(webhookURL, {
      content: message,
    });
    console.log('Message sent:', response.data);
  } catch (error) {
    console.error('Error sending webhook message:', error);
  }
}

// Usage example
const webhookURL = 'https://discord.com/api/webhooks/your_webhook_id';
const message = 'Hello, this is a message from the web4app.io platform!';
sendWebhookMessage(webhookURL, message);
