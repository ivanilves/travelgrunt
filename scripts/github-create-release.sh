#!/usr/bin/env bash
#
if [[ ${#} -ne 2 ]]; then
  echo "Create release on GitHub"
  echo "Usage: ${0} GITHUB_REPO RELEASE_PATH" >>/dev/stderr
  exit 1
fi

if [[ -z "${GITHUB_TOKEN}" ]]; then
  echo "GITHUB_TOKEN not set" >>/dev/stderr
  exit 1
fi

set -euo pipefail

declare -r GITHUB_REPO="${1}"
declare -r RELEASE_PATH="${2}"

declare -r API_URL=https://api.github.com/repos/${GITHUB_REPO}/releases

pushd "${RELEASE_PATH}"
  declare -r TAG="$(cat TAG)"
  declare -r NAME="$(cat NAME)"
  declare -r BODY="$(sed 's/$/\\n/' CHANGELOG.md | tr -d '\n' | sed 's/\"/\\"/g')"
  declare -r DATA="{\"tag_name\":\"${TAG}\", \"target_commitish\": \"main\",\"name\": \"${NAME}\", \"body\": \"${BODY}\", \"draft\": false, \"prerelease\": false}"

  curl -f -X POST -H "Content-Type:application/json" \
    -H "Authorization: Token ${GITHUB_TOKEN}" "${API_URL}" \
    -d "${DATA}"
popd
