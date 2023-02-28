# Makefile to build the project

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint scan-gosec

test:
	go test `go list ./...`

test-cov:
	go test `go list ./...` ${COVERAGE}

test-int:
	go test `go list ./...` -tags=integration

test-int-cov:
	go test `go list ./...` -tags=integration ${COVERAGE}

lint:
	golangci-lint run

scan-gosec:
	gosec ./...

tidy:
	go mod tidy
