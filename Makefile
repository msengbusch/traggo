MODULE=github.com/traggo/server

BUILD_DIR=./build
BUILD_DIR_WEB=./web/dist
BUILD_EXEC_CLI_DEV=${BUILD_DIR}/traggo-cli-dev
BUILD_EXEC_CLI_PROD=${BUILD_DIR}/traggo-cli
BUILD_EXEC_SERVER_DEV=${BUILD_DIR}/traggo-server-dev
BUILD_EXEC_SERVER_PROD=${BUILD_DIR}/traggo-server

FLAGS_DATA=-X ${MODULE}/environment.BuildDate=$(DATE) -X ${MODULE}/environment.BuildCommit=$(COMMIT) -X ${MODULE}/environment.BuildVersion=$(VERSION)
FlAGS=${FLAGS_DATA}
FLAGS_DEV=${FlAGS}
FLAGS_PROD=${FlAGS} -s -w

TAGS=
TAGS_DEV=${TAGS}
TAGS_PROD=${TAGS} prod

COMMIT=$(shell git rev-parse --verify HEAD)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION=$(shell git describe --tags --abbrev=0 | cut -c 2-)

.PHONY: clean
clean:
	rm -rf ${BUILD_DIR}

build-cli:
	go build -ldflags "${FLAGS_DEV}" -tags "${TAGS_DEV}" -o "${BUILD_EXEC_CLI_DEV}" cmd/cli/main.go

build-cli-prod:
	go build -ldflags "${FLAGS_PROD}" -tags "${TAGS_PROD}" -o "${BUILD_EXEC_CLI_PROD}" cmd/cli/main.go

build-server:
	go build -ldflags "${FLAGS_DEV}" -tags "${TAGS_DEV}" -o "${BUILD_EXEC_SERVER_DEV}" cmd/server/main.go

build-server-prod:
	go build -ldflags "${FLAGS_PROD}" -tags "${TAGS_PROD}" -o "${BUILD_EXEC_SERVER_PROD}" cmd/server/main.go

build-web:
	cd web && bun run build

.PHONY: run-cli
run-cli:
	go run -ldflags "${FLAGS_DEV}" -tags "${TAGS_DEV}"

.PHONY: run-cli-prod
run-cli-prod:
	go run -ldflags "${FLAGS_PROD}" -tags "${TAGS_PROD}"

.PHONY: run-server
run-server:
	go run -ldflags "${FLAGS_DEV}" -tags "${TAGS_DEV}"

.PHONY: run-server-prod
run-server-prod:
	go run -ldflags "${FLAGS_PROD}" -tags "${TAGS_PROD}"