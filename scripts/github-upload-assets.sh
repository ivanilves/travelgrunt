#!/usr/bin/env bash
#
if [[ ${#} -ne 2 ]]; then
  echo "Upload assets to GitHub release"
  echo "Usage: ${0} GITHUB_REPO ASSETS_PATH" >>/dev/stderr
  exit 1
fi

if [[ -z "${GITHUB_TOKEN}" ]]; then
  echo "GITHUB_TOKEN not set" >>/dev/stderr
  exit 1
fi

set -euo pipefail

declare -r GITHUB_REPO="${1}"
declare -r ASSETS_PATH="${2}"

declare -r API_URL=https://api.github.com/repos/${GITHUB_REPO}/releases
declare -r UPLOAD_URL=https://uploads.github.com/repos/${GITHUB_REPO}/releases

pushd "${ASSETS_PATH}"
  declare -r TAG="$(cat TAG)"

  ID=$(curl -s -H "Authorization: Token ${GITHUB_TOKEN}" "${API_URL}/tags/${TAG}" | jq ".id")

  for FILE in $(find ./ -mindepth 1 -maxdepth 1 -type f -name "*.tar.gz"); do
    curl -f -H "Content-Type: $(file -b --mime-type ${FILE})" \
      -H "Authorization: Token ${GITHUB_TOKEN}" \
      --data-binary @${FILE} \
      "${UPLOAD_URL}/${ID}/assets?name=$(basename ${FILE})"
    echo
  done
popd
