export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)
export SHELL := bash
export OSTYPE := $(shell uname -s)

BINARY=shipment-service-shipping-label
BINARY_MIGRATION_TOOLS=jb-shipment-migrate

# Include external Makefiles (adjust paths if necessary)
include ./misc/make/tools.Makefile
include ./misc/make/help.Makefile

.PHONY: module-generate up down destroy install-deps deps dev-env dev-env-test dev docker-stop docker-teardown lint build build-race go-generate run-tests tests image-build tenants-migration system-migration-up system-migration-down migrate-create clean-all clean-artifacts clean-docker clean start build-migration-tools lint-prepare

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
module-generate:  ## Generate modules
	@read -p "Enter Module name: " name; \
	mkdir -p modules/$${name}/{repository,usecase,delivery/http}; \
	echo "generate module $${name} done."

up: dev-env dev-air             ## Startup / Spinup Docker Compose and air
down: docker-stop               ## Stop Docker
destroy: docker-teardown clean  ## Teardown (removes volumes, tmp files, etc...)

install-deps: air gotestsum tparse mockery ## Install Development Dependencies (locally).
deps:
	@echo "Required Tools Are Available"

dev-env: ## Bootstrap Environment (with Docker-Compose).
	@ docker-compose up -d

dev-env-test: dev-env ## Run application (within Docker-Compose).
	@ $(MAKE) image-build
	docker-compose up web

dev: $(AIR) ## Starts AIR (Continuous Development app).
	air

docker-stop:
	@ docker-compose down

docker-teardown:
	@ docker-compose down --remove-orphans -v

# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
lint: $(GOLANGCI) ## Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

build: ## Builds binary
	@ printf "Building application... "
	@ go build \
		-trimpath  \
		-o $(BINARY) \
		.
	@ echo "done"

build-race: ## Builds binary (with -race flag)
	@ printf "Building application with race flag... "
	@ go build \
		-trimpath  \
		-race      \
		-o $(BINARY) \
		.
	@ echo "done"

go-generate: $(MOCKERY) ## Runs go generate ./...
	go generate ./...

run-tests: $(GOTESTSUM)
	@ gotestsum $(TESTS_ARGS) -short

tests: run-tests $(TPARSE) ## Run Tests & parse details
	@cat gotestsum.json.out | $(TPARSE) -all -top -notests

# ~~~ Docker Build ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
image-build:
	@ echo "Docker Build"
	@ DOCKER_BUILDKIT=0 docker build \
		--file Dockerfile \
		--tag go-clean-arch .

# ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
DB_SYSTEM_DSN := "postgres://$(DB_NAME):$(DB_PASSWORD)@tcp($(DB_HOST))/$(DB_NAME)"

tenants-migration: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@read -p "Enter Tenant Schema name: " name; \
	./${BINARY_MIGRATION_TOOLS} tenants $${name}

system-migration-up: $(MIGRATE) ## Apply system migrations up.
	sql-migrate up --env="systemdb"

system-migration-down: $(MIGRATE) ## Apply system migrations down.
	sql-migrate down --env="systemdb"

migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@read -p "Enter migration file name: " name; \
	read -p "Enter migration environment: " envname; \
	sql-migrate new --env="$${envname}" $${name}

# ~~~ Cleans ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
clean-all: clean-artifacts clean-docker

clean-artifacts: ## Removes Artifacts (*.out)
	@printf "Cleaning artifacts... "
	@rm -f *.out
	@echo "done."

clean-docker: ## Removes dangling docker images
	@ docker image prune -f

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

start:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	go build -o ${BINARY} *.go
	./${BINARY}

build-migration-tools:
	@if [ -f ${BINARY_MIGRATION_TOOLS} ] ; then rm ${BINARY_MIGRATION_TOOLS} ; fi
	go build -o ${BINARY_MIGRATION_TOOLS} ./misc/*.go

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
