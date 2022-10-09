BUILD_PATH := cmd/ttg

main: dep build

all: dep build install

dep:
	go mod vendor

build:
	cd ${BUILD_PATH} && go build

install: PREFIX := /usr/local
install:
	install ${BUILD_PATH}/ttg ${PREFIX}/bin/
