APP_NAME    := travelgrunt
API_VERSION := 0.2
BUILD_PATH  := ./cmd/${APP_NAME}

CURRENT_PATCH = $(shell (git fetch --tags && git tag --sort=creatordate | grep -F "v${API_VERSION}." || echo -1) | tail -n1 | sed -r "s/^v([0-9]+\.){2}//")
NEXT_PATCH    = $(shell expr ${CURRENT_PATCH} + 1)
NEXT_VERSION  = v${API_VERSION}.${NEXT_PATCH}

SHELL := /bin/bash

default: dep build

all: dep build test install

deploy: build install

dep:
	go mod tidy
	go mod vendor

build:
	cd ${BUILD_PATH} && go build -tags netgo,osusergo

test:
	go test -v ./...

lint:
	golangci-lint run -v ./...

staticcheck:
	staticcheck ./...

clean:
	git clean -fdx

install: PREFIX ?= /usr/local
install:
	install ${BUILD_PATH}/travelgrunt ${PREFIX}/bin/

check-git-branch: GIT_BRANCH ?= main
check-git-branch:
	@if [ `git rev-parse --abbrev-ref HEAD` != ${GIT_BRANCH} ]; \
		then echo -e "\e[33mERROR: Need to be on the \"${GIT_BRANCH}\" branch!\e[0m" >>/dev/stderr; \
		exit 1; fi

pull-git-branch:
	git pull

next-version-tag:
	git tag ${NEXT_VERSION} && git push --tags

release: check-git-branch pull-git-branch next-version-tag

semantic-commit-check: RANGE ?= main..HEAD
semantic-commit-check: REGEX := "^(feat|fix|refactor|chore|test|style|docs)(\([a-zA-Z0-9\/_-]+\))?: [a-zA-Z]"
semantic-commit-check:
	@git log --pretty="format:%s" ${RANGE} >/dev/null
	@git log --pretty="format:%s" ${RANGE} | egrep -v "Merge " \
		| egrep -v ${REGEX} | awk '{print "NON-SEMANTIC: "$$0}' | grep . \
		&& echo -e "\e[1m\e[31mFATAL: Non-semantic commit messages found (${RANGE})!\e[0m" && exit 1 \
		|| echo -e "\e[1m\e[32mOK\e[0m"
