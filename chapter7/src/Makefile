VERSION?="0.0.1"

default: test

test: fmtcheck
	go test ./... -cover -race -timeout 60s

testint: fmtcheck
	go test ./... -race -cover -tags=integration -timeout 60s -count=1

cover: fmtcheck
	go test ./... -coverpkg=./... -coverprofile=cover.out -tags=integration -covermode=atomic && \
	go tool cover -func=cover.out && \
	rm cover.out

fmt:
	go fmt ./...

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint: fmtcheck
	golangci-lint run -E golint --exclude-use-default=false --build-tags integration

deeplint: fmtcheck
	golangci-lint run --enable-all --exclude-use-default=false -D dupl --build-tags integration

ci: fmtcheck lint testint

modsync: fmtcheck
	go mod tidy && \
	go mod vendor


# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: default test testint cover fmt fmtcheck lint deeplint ci modsync
