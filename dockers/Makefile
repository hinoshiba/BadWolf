GOLANG=dev-go
DC=docker-compose

SRCS := $(shell find . -type f)

all: go ## test & build

build: $(SRCS) ## build to all container
	$(DC) build

run: $(SRCS) ## start up to all container
	$(DC) up -d

.PHONY: clean
clean: ## stop container and cleanup data
	$(DC) down --rmi all

go: ## run golang
	$(DC) exec $(GOLANG) make
gobuild: ## build golang
	$(DC) exec $(GOLANG) make build
gotest: ## test golang
	$(DC) exec $(GOLANG) make test
goclean: ## clean golang
	$(DC) exec $(GOLANG) make clean

.PHONY: help
	all: help
help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)