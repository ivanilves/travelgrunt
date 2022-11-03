APP_NAME    := travelgrunt
API_VERSION := 0.2
BUILD_PATH  := ./cmd/${APP_NAME}

CURRENT_PATCH = $(shell (git fetch --tags && git tag --sort=creatordate | grep -F "v${API_VERSION}." || echo -1) | tail -n1 | sed -r "s/^v([0-9]+\.){2}//")
NEXT_PATCH    = $(shell expr ${CURRENT_PATCH} + 1)
NEXT_VERSION  = v${API_VERSION}.${NEXT_PATCH}

default: dep build

all: dep build install

deploy: build install

dep:
	go mod tidy
	go mod vendor

build:
	cd ${BUILD_PATH} && go build

clean:
	git clean -fdx

install: PREFIX ?= /usr/local
install:
	install ${BUILD_PATH}/travelgrunt ${PREFIX}/bin/

check-git-branch: GIT_BRANCH ?= main
check-git-branch:
	@if [ `git rev-parse --abbrev-ref HEAD` != ${GIT_BRANCH} ]; \
		then echo "ERR: Need to be on the \"${GIT_BRANCH}\" branch" >>/dev/stderr; \
		exit 1; fi

pull-git-branch:
	git pull

next-version-tag:
	git tag ${NEXT_VERSION} && git push --tags

release: check-git-branch pull-git-branch next-version-tag
