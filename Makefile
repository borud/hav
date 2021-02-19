ifeq ($(GOPATH),)
GOPATH := $(HOME)/go
endif

all: test lint vet build

clean:
	@rm -f bin/*
	@go clean -testcache

build: hav

hav:
	@cd cmd/hav && go build -o ../../bin/hav

check: lint vet staticcheck revive

lint:
	@revive -exclude ./third_party/... ./... 

vet:
	@go vet ./...

staticcheck:
	@staticcheck ./...

revive:
	@revive ./...

test:
	@go test ./...

test_verbose:
	@go test ./... -v

test_cover:
	@go test ./... -cover -coverprofile=unittests.cover -coverpkg=github.com/lab5e/mesh/backend/...

test_race:
	@go test ./... -race

test_all: test_cover test_race
