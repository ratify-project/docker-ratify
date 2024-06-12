DOCKER_PLUGINS = docker-ratify
GO_BUILD_FLAGS =

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: build

.PHONY: FORCE
FORCE:

bin/%: cmd/% FORCE
	go build $(GO_BUILD_FLAGS) -o $@ ./$<

.PHONY: build
build: $(addprefix bin/,$(DOCKER_PLUGINS)) ## builds binaries

.PHONY: install
install: $(addprefix install-,$(DOCKER_PLUGINS)) ## installs the docker plugins

.PHONY: install-docker-%
install-docker-%: bin/docker-%
	cp $< ~/.docker/cli-plugins/

.PHONY: check-line-endings
check-line-endings: ## check line endings
	! find . -name "*.go" -type f -exec file "{}" ";" | grep CRLF

.PHONY: fix-line-endings
fix-line-endings: ## fix line endings
	find . -type f -name "*.go" -exec sed -i -e "s/\r//g" {} +
