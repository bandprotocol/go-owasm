.PHONY: docker-images docker-image-linux docker-image-osx

DOCKER_TAG := 0.0.1
USER_ID := $(shell id -u)
USER_GROUP = $(shell id -g)

docker-image-linux:
	DOCKER_BUILDKIT=1 docker build .. -t owasm/go-ext-builder:$(DOCKER_TAG)-linux -f build/Dockerfile.linux

docker-image-osx:
	DOCKER_BUILDKIT=1 docker build .. -t owasm/go-ext-builder:$(DOCKER_TAG)-osx -f build/Dockerfile.osx

docker-images: docker-image-linux docker-image-osx

release-linux:
	rm -rf libgo_owasm/target/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-linux

release-osx:
	rm -rf libgo_owasm/target/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-osx

# and use them to compile release builds
release:
	rm -rf libgo_owasm/target/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-linux
	rm -rf libgo_owasm/target/release
	docker run --rm -u $(USER_ID):$(USER_GROUP) -v $(shell pwd):/code/go-owasm owasm/go-ext-builder:$(DOCKER_TAG)-osx
