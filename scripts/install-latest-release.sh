#!/bin/sh
#
if [ -z "${PREFIX}" ]; then
  PREFIX=/usr/local
fi

set -euo pipefail

LATEST_TAG=$(curl -s -f https://api.github.com/repos/ivanilves/travelgrunt/releases/latest | grep "^  \"tag_name\"" | sed 's/.*: "//;s/",$//')

OS=$(uname -s | tr '[A-Z]' '[a-z]')
ARCH=$(uname -m | sed 's/x86_64/amd64/')

curl -s -f -L -o - https://github.com/ivanilves/travelgrunt/releases/download/${LATEST_TAG}/travelgrunt-${OS}-${ARCH}.tar.gz | tar -xz -C ${PREFIX}/bin
