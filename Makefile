LDFLAGS := -X github.com/msugakov/release-stackrox/pkg/version.version=$$(git describe --tags --abbrev=10 --dirty)

.PHONY: run
run:
	go run -ldflags "$(LDFLAGS)" cmd/release-stackrox/main.go
