.PHONY: build
build:
	@goreleaser --skip-publish --rm-dist

.PHONY: snap
snap:
	@goreleaser --snapshot --skip-publish --rm-dist

.PHONY: fmt
fmt:
	@go fmt ./...
