APP_NAME     := ttg
BUILD_PATH   := ./cmd/${APP_NAME}
RELEASE_PATH := ./release

main: dep build

all: dep build install

dep:
	go mod vendor

build:
	cd ${BUILD_PATH} && go build

clean:
	git clean -fdx

install: PREFIX := /usr/local
install:
	install ${BUILD_PATH}/ttg ${PREFIX}/bin/

release-binary: GOOS   ?= $(shell uname -s | tr '[A-Z]' '[a-z]')
release-binary: GOARCH ?= $(shell uname -m | sed 's/x86_64/amd64/')
release-binary:
	mkdir -p ${RELEASE_PATH}/${APP_NAME}-${GOOS}-${GOARCH} && cd ${BUILD_PATH} && \
		GOOS=${GOOS} GOARCH=${GOARCH} go build -mod=vendor -o ../../${RELEASE_PATH}/${APP_NAME}-${GOOS}-${GOARCH}/${APP_NAME}

release-binaries:
	${MAKE} --no-print-directory release-binary GOOS=linux  GOARCH=amd64
	${MAKE} --no-print-directory release-binary GOOS=darwin GOARCH=amd64
	${MAKE} --no-print-directory release-binary GOOS=darwin GOARCH=arm64

release-artifacts:
	cd ${RELEASE_PATH} && find . -mindepth 1 -maxdepth 1 -type d | xargs -i tar -C {} -zc ${APP_NAME} -f {}.tar.gz

release: release-binaries release-artifacts
