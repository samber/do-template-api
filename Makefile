
BINARY=do-template-api

all: deps build

#
# Deps
#
deps-tools:
	go install github.com/cespare/reflex@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/psampaz/go-mod-outdated@latest
	go install github.com/jondot/goweight@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

	# brew install sonatype-nexus-community/nancy-tap/nancy
	# brew tap snyk/tap
	# brew install snyk
	# brew install bearer/tap/bearer

deps:
	go mod download -x

cleanup-deps:
	go mod tidy

# brew install sonatype-nexus-community/nancy-tap/nancy
audit:
	go list -json -m all | nancy sleuth

outdated:
	go list -u -m -json all | go-mod-outdated -update -direct

vulncheck:
	govulncheck ./...

#
# Build
#
build:
	go build -ldflags ${LDFLAGS} -o ${BINARY} ./cmd/main.go

debug:
	dlv --listen=:4243 --headless=true --accept-multiclient --api-version=2 debug ./cmd/main.go --continue
watch-debug:
	reflex -t 50ms -s -- sh -c 'echo \\nBUILDING && dlv --listen=:4243 --headless=true --accept-multiclient --api-version=2 debug ./cmd/main.go --continue && echo Exited'

run: 
	go run -race ./cmd/main.go
watch-run:
	reflex -t 50ms -s -- sh -c 'echo \\nBUILDING && go run -race ./cmd/main.go && echo Exited'

#
# Quality
#
lint:
	golangci-lint run --timeout 600s --max-same-issues 50 --path-prefix=./ --config=.golangci.yml ./...
lint-fix:
	golangci-lint run --fix --timeout 600s --max-same-issues 50 --path-prefix=./ --config=.golangci.yml ./...

test:
	gotest -v -race ./...
watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v -race ./...'

weight:
	goweight

coverage:
	go test -v -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html

# brew tap snyk/tap
# brew install snyk
snyk-test:
	snyk test
snyk-code-test:
	snyk code test

# brew install bearer/tap/bearer
bearer:
	bearer scan .

#
# Cleaning
#
clean:
	rm -f ${BINARY}
clean-env:
	# don't execute it every day ;) it will remove all your dependencies and build cache
	go clean -cache -testcache -modcache -x

re: clean all

.PHONY: all deps deps-toolsaudit outdated vulncheck build debug watch-debug run watch-run lint lint-fix test watch-test weight coverage clean re