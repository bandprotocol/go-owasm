.PHONY: docker-images docker-image-linux docker-image-osx

DOCKER_TAG := 0.0.1
USER_ID := $(shell id -u)
USER_GROUP = $(shell id -g)

docker-image-alpine:
	DOCKER_BUILDKIT=1 docker build .. -t owasm/go-ext-builder:$(DOCKER_TAG)-alpine -f build/Dockerfile.alpine

docker-image-linux:
	DOCKER_BUILDKIT=1 docker build .. -t owasm/go-ext-builder:$(DOCKER_TAG)-linux -f build/Dockerfile.linux

docker-image-osx:
	DOCKER_BUILDKIT=1 docker build .. -t owasm/go-ext-builder:$(DOCKER_TAG)-osx -f build/Dockerfile.osx

docker-images: docker-image-osx docker-image-linux docker-image-alpine

# Creates a release build in a containerized build environment of the static library for Alpine Linux (.a)
release-alpine:
	rm -rf libgo_owasm/target/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-alpine

# Creates a release build in a containerized build environment of the shared library for glibc Linux (.so)
release-linux:
	rm -rf libgo_owasm/target/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-linux

# Creates a release build in a containerized build environment of the shared library for macOS (.dylib)
release-osx:
	rm -rf libgo_owasm/target/release
	rm -rf libgo_owasm/target/x86_64-apple-darwin/release
	rm -rf libgo_owasm/target/aarch64-apple-darwin/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-osx

releases:
	make release-alpine
	make release-linux
	make release-osx
