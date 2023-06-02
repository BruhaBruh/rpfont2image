help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ## run application
	go mod tidy && go mod download && \
    go run ./cmd/app

build: ## build application for platforms
	bash ./build.sh github.com/bruhabruh/rpfont2image