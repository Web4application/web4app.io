# Modernized Layerfile for React + Cypress
FROM vm/ubuntu:22.04

# Install Node.js 18 LTS + build essentials
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash && \
    apt-get install -y nodejs python3 make gcc build-essential

# Cypress OS dependencies
RUN apt-get install -y \
    libgtk2.0-0 libgtk-3-0 libgbm-dev \
    libnotify-dev libgconf-2-4 libnss3 libxss1 \
    libasound2 libxtst6 xauth xvfb wait-for-it

# VM memory + Node heap settings
MEMORY 2G
ENV NODE_OPTIONS=--max-old-space-size=2048

RUN node --version
RUN npm --version

# Copy package files first (better caching)
COPY package.json ./
COPY package-lock.json ./
CACHE ~/.npm ~/.cache/Cypress

# Install dependencies
RUN npm ci

# Copy rest of the app
COPY . .

# Build frontend
RUN npm run build

# Show Cypress install info
RUN npx cypress info

# Print CI environment variables
RUN npx @bahmutov/print-env GIT CI RETRY USER SPLIT

# Serve production build in background
RUN BACKGROUND npx serve -s dist -l 8080

# Wait until app is ready
RUN npx wait-on http://localhost:8080

# Unique staging link
EXPOSE WEBSITE http://localhost:8080

# Cypress Dashboard secret
SECRET ENV CYPRESS_RECORD_KEY

# Run Cypress tests in parallel across 3 workers
SPLIT UNORDERED 3
RUN npx cypress run --record --parallel --browser chrome --headless \
  --ci-build-id $JOB_ID-$RETRY_INDEX
