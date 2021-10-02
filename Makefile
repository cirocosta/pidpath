.PHONY: build
build:
	@goreleaser --snapshot --skip-publish --rm-dist

.PHONY: fmt
fmt:
	@go fmt ./...
