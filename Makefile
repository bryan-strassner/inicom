TESTDIR ?= build/test

.PHONY: tests
tests:
	go vet ./...
	mkdir -p $(TESTDIR)
	go test ./... -timeout=10s -coverprofile=build/test/cover.out
	go tool cover -func=build/test/cover.out

.PHONY: images
images: tests
	docker build -f build/package/Dockerfile .