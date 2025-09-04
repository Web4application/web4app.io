# ~/.nodenv/hooks/exec/npx-path-fix.bash
# If npx is being executed, remove nodenv's shims from PATH
# so npx won’t pick up shims for non-active node versions.

[ "$NODENV_COMMAND" = npx ] || return 0

remove_from_path() {
  local path_to_remove="$1"
  local path_before
  local result=":${PATH//\~/$HOME}:"
  while [ "$path_before" != "$result" ]; do
    path_before="$result"
    result="${result//:$path_to_remove:/:}"
  done
  result="${result%:}"
  echo "${result#:}"
}

PATH="$(remove_from_path "${NODENV_ROOT}/shims")"
export PATH
