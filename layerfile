FROM vm/ubuntu:18.04
RUN curl -sL https://deb.nodesource.com/setup_22 | bash
RUN apt install nodejs
COPY . .
RUN npm install
SECRET ENV CONFIGCAT_AUTH_KEY
RUN curl "https://api.configcat.com/v1/products/dcd53ddb-8104-4e48-8cc0-5df1088c6113/environments" \\
    -X POST \\
    -u $CONFIGCAT_AUTH_KEY \\
    -H "Content-Type: application/json" \\
    -d '{"name": "webapp.io-'$JOB_ID'"}'
RUN BACKGROUND REACT_CONFIGCAT_ENV="layerci-$JOB_ID" npm run start
EXPOSE WEBSITE localhost:8000

ENV CI_NAME=layerci \\
   CI_BUILD_NUMBER=$JOB_ID \\
   CI_BUILD_URL="https://webapp.io/qubuhub/commits?query=web4$JOB_ID" \\
   CI_BRANCH="$GIT_BRANCH" \\
   CI_PULL_REQUEST="$PULL_REQUEST_URL"

SECRET ENV COVERALLS_REPO_TOKEN

RUN (the test command)
