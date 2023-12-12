MODULE=github.com/traggo/server

BUILD_DIR=./build

BUILD_EXECUTABLE_FILE=${BUILD_DIR}/traggo-dev
BUILD_EXECUTABLE_FILE_PROD=${BUILD_DIR}/traggo

COMMIT=$(shell git rev-parse --verify HEAD)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION=$(shell git describe --tags --abbrev=0 | cut -c 2-)

TAGS=""
TAGS_PROD="prod"
LD_FLAGS=-X ${MODULE}/environment.BuildDate=$(DATE) -X ${MODULE}/environment.BuildCommit=$(COMMIT) -X ${MODULE}/environment.BuildVersion=$(VERSION)
LD_FLAGS_PROD=${LD_FLAGS} -s -w

clean:
	rm -rf ${BUILD_DIR}

build:
	go build -ldflags "${LD_FLAGS}" -tags "${TAGS}" -o ${BUILD_EXECUTABLE_FILE} main.go

build-prod:
	go build -ldflags "${LD_FLAGS_PROD}" -tags "${TAGS_PROD}" -o ${BUILD_EXECUTABLE_FILE_PROD} main.go

run:
	go run -ldflags "${LD_FLAGS}" -tags "${TAGS}" main.go

run-prod:
	go run -ldflags "${LD_FLAGS_PROD}" -tags "${TAGS_PROD}" main.go