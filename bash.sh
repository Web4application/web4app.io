#!/usr/bin/env bash

set -e

if [ "$1" = "--debug" ]; then
  export NODENV_DEBUG=1
  shift
fi

if [ -n "$NODENV_DEBUG" ]; then
  # https://wiki-dev.bash-hackers.org/scripting/debuggingtips#making_xtrace_more_useful
  export PS4='+(${BASH_SOURCE}:${LINENO}): ${FUNCNAME[0]:+${FUNCNAME[0]}(): }'
  set -x
fi

abort() {
  { if [ "$#" -eq 0 ]; then cat -
    else echo "nodenv: $*"
    fi
  } >&2
  exit 1
}

if enable -f "${BASH_SOURCE%/*}"/../libexec/nodenv-realpath.dylib realpath 2>/dev/null; then
  abs_dirname() {
    local path
    path="$(realpath "$1")"
    echo "${path%/*}"
  }
else
  [ -z "$NODENV_NATIVE_EXT" ] || abort "failed to load \`realpath' builtin"

  READLINK=$(type -p greadlink readlink 2>/dev/null | head -n1)
  [ -n "$READLINK" ] || abort "cannot find readlink - are you missing GNU coreutils?"

  resolve_link() {
    $READLINK "$1"
  }

  abs_dirname() {
    local cwd="$PWD"
    local path="$1"

    while [ -n "$path" ]; do
      cd "${path%/*}"
      local name="${path##*/}"
      path="$(resolve_link "$name" || true)"
    done

    pwd
    cd "$cwd"
  }
fi

if [ -z "${NODENV_ROOT}" ]; then
  NODENV_ROOT="${HOME}/.nodenv"
else
  NODENV_ROOT="${NODENV_ROOT%/}"
fi
export NODENV_ROOT

if [ -z "${NODENV_DIR}" ]; then
  NODENV_DIR="$PWD"
else
  [[ $NODENV_DIR == /* ]] || NODENV_DIR="$PWD/$NODENV_DIR"
  cd "$NODENV_DIR" 2>/dev/null || abort "cannot change working directory to \`$NODENV_DIR'"
  NODENV_DIR="$PWD"
  cd "$OLDPWD"
fi
export NODENV_DIR

[ -n "$NODENV_ORIG_PATH" ] || export NODENV_ORIG_PATH="$PATH"

shopt -s nullglob

bin_path="$(abs_dirname "$0")"
for plugin_bin in "${NODENV_ROOT}/plugins/"*/bin; do
  PATH="${plugin_bin}:${PATH}"
done
export PATH="${bin_path}:${PATH}"

NODENV_HOOK_PATH="${NODENV_HOOK_PATH}:${NODENV_ROOT}/nodenv.d"
if [ "${bin_path%/*}" != "$NODENV_ROOT" ]; then
  # Add nodenv's own `nodenv.d` unless nodenv was cloned to NODENV_ROOT
  NODENV_HOOK_PATH="${NODENV_HOOK_PATH}:${bin_path%/*}/nodenv.d"
fi
NODENV_HOOK_PATH="${NODENV_HOOK_PATH}:/usr/local/etc/nodenv.d:/etc/nodenv.d:/usr/lib/nodenv/hooks"
for plugin_hook in "${NODENV_ROOT}/plugins/"*/etc/nodenv.d; do
  NODENV_HOOK_PATH="${NODENV_HOOK_PATH}:${plugin_hook}"
done
NODENV_HOOK_PATH="${NODENV_HOOK_PATH#:}"
export NODENV_HOOK_PATH

shopt -u nullglob


command="$1"
case "$command" in
"" )
  { nodenv---version
    nodenv-help
  } | abort
  ;;
-v | --version )
  exec nodenv---version
  ;;
-h | --help )
  exec nodenv-help
  ;;
* )
  command_path="$(command -v "nodenv-$command" || true)"
  if [ -z "$command_path" ]; then
    if [ "$command" == "shell" ]; then
      abort "shell integration not enabled. Run \`nodenv init' for instructions."
    else
      abort "no such command \`$command'"
    fi
  fi

  shift 1
  if [ "$1" = --help ]; then
    if [[ "$command" == "sh-"* ]]; then
      echo "nodenv help \"$command\""
    else
      exec nodenv-help "$command"
    fi
  else
    exec "$command_path" "$@"
  fi
  ;;
esac

# Initialize Go module if needed
go mod init web4app
go mod tidy

# Build and run locally
go build -o web4app main.go
./web4app
docker build -t web4app .
docker run -p 8080:8080 --env-file .env web4app
npm install -g netlify-cli
netlify init
netlify deploy --prod
npm install lola-ai

cd ./web4app.io
yarn create semantic-module

curl -SL https://github.com/docker/compose/releases/download/v2.37.3/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

 Start-BitsTransfer -Source "https://github.com/docker/compose/releases/download/v2.37.3/docker-compose-windows-x86_64.exe" -Destination Env:
 

 composer create-project hunwalk/yii2-basic-firestarter <project_pilot_AI> --prefer-dist

 php yii migrate-user
 php yii migrate-rbac
 php yii migrate

 cp .env.web4app .env

npm install --save-dev @commitlint/config-conventional @commitlint/cli
echo "export default {extends: ['@commitlint/config-conventional']};" > commitlint.config.js

npm install commitizen -g

npm
commitizen init cz-conventional-changelog --save-dev --save-exact

# yarn
commitizen init cz-conventional-changelog --yarn --dev --exact

# pnpm
commitizen init cz-conventional-changelog --pnpm --save-dev --save-exact

npx commitizen init cz-conventional-changelog --save-dev --save-exact

curl http://localhost:8080/api/hello
curl http://localhost:8080/api/time
curl -X POST http://localhost:8080/api/echo -d '{"text":"Web4 is here!"}' -H "Content-Type: application/json"
curl http://localhost:8080/api/status

git add .
git commit -m "Added files from phone"
git push origin main

cmake -S . -B build -DCMAKE_INSTALL_PREFIX=/usr/local -DBUILD_SHARED_LIBS=ON
cmake --build build
sudo cmake --install build
cmake -S . -B build -DCMAKE_INSTALL_PREFIX=/usr/local -DBUILD_SHARED_LIBS=OFF
cmake --build build
sudo cmake --install build


