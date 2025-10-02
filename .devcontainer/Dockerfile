FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu-22.04

# Install Node.js + Python
RUN apt-get update && apt-get install -y \
    python3 python3-pip python3-venv \
    curl git build-essential \
    && curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs \
    && npm install -g npm@latest

# Set user
USER vscode
